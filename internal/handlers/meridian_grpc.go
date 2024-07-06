package handlers

import (
	"context"
	"log"
	"strings"

	"github.com/c12s/meridian/internal/domain"
	"github.com/c12s/meridian/pkg/api"
	pulsar_api "github.com/c12s/pulsar/model/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MeridianGrpcHandler struct {
	api.UnimplementedMeridianServer
	namespaces domain.NamespaceStore
	apps       domain.AppStore
	pulsar     pulsar_api.SeccompServiceClient
}

func NewMeridianGrpcHandler(namespaces domain.NamespaceStore, apps domain.AppStore, pulsar pulsar_api.SeccompServiceClient) api.MeridianServer {
	return MeridianGrpcHandler{
		namespaces: namespaces,
		apps:       apps,
		pulsar:     pulsar,
	}
}

func (m MeridianGrpcHandler) AddNamespace(ctx context.Context, req *api.AddNamespaceReq) (*api.AddNamespaceResp, error) {
	namespace, err := m.namespaces.Get(domain.MakeNamespaceId(req.OrgId, req.Name))
	if err == nil {
		err = status.Error(codes.AlreadyExists, "namespace already exists")
		return nil, err
	}
	var parent *domain.Namespace
	if req.ParentName != "" {
		p, err := m.namespaces.Get(domain.MakeNamespaceId(req.OrgId, req.ParentName))
		if err != nil {
			log.Println(err)
			err = status.Error(codes.NotFound, "parent namespace not found")
			return nil, err
		}
		parent = &p
	}
	namespace = domain.NewNamespace(req.OrgId, req.Name, req.Profile.Version, req.Labels)
	for resource, quota := range req.Quotas {
		err := namespace.AddResourceQuota(resource, quota)
		if err != nil {
			log.Println(err)
			err = status.Error(codes.InvalidArgument, err.Error())
			return nil, err
		}
	}
	err = m.sendSeccompProfile(ctx,
		req.SeccompDefinitionStrategy,
		namespace.GetSeccompProfile(),
		req.Profile, parent)
	if err != nil {
		return nil, err
	}
	err = m.namespaces.Add(namespace, parent)
	if err != nil {
		log.Println(err)
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &api.AddNamespaceResp{}, nil
}

func (m MeridianGrpcHandler) RemoveNamespace(ctx context.Context, req *api.RemoveNamespaceReq) (*api.RemoveNamespaceResp, error) {
	tree, err := m.namespaces.GetHierarchy(domain.MakeNamespaceId(req.OrgId, req.Name))
	if err == nil && (len(tree.Root.Children) > 0 || len(tree.Root.Apps) > 0) {
		err = status.Error(codes.InvalidArgument, "namespace must not have applications or child namespaces")
		return nil, err
	}
	err = m.namespaces.Remove(domain.MakeNamespaceId(req.OrgId, req.Name))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &api.RemoveNamespaceResp{}, nil
}

func (m MeridianGrpcHandler) AddApp(ctx context.Context, req *api.AddAppReq) (*api.AddAppResp, error) {
	namespace, err := m.namespaces.Get(domain.MakeNamespaceId(req.OrgId, req.Namespace))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.NotFound, "namespace not found")
		return nil, err
	}
	app := domain.NewApp(namespace, req.Name, req.Profile.Version)
	for resource, quota := range req.Quotas {
		err := app.AddResourceQuota(resource, quota)
		if err != nil {
			log.Println(err)
			err = status.Error(codes.InvalidArgument, err.Error())
			return nil, err
		}
	}
	err = m.sendSeccompProfile(ctx,
		req.SeccompDefinitionStrategy,
		app.GetSeccompProfile(),
		req.Profile, &namespace)
	if err != nil {
		return nil, err
	}
	err = m.apps.Add(app)
	if err != nil {
		log.Println(err)
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &api.AddAppResp{}, nil
}

func (m MeridianGrpcHandler) RemoveApp(ctx context.Context, req *api.RemoveAppReq) (*api.RemoveAppResp, error) {
	err := m.apps.Remove(domain.MakeAppId(req.OrgId, req.Namespace, req.Name))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &api.RemoveAppResp{}, nil
}

func (m MeridianGrpcHandler) GetNamespace(ctx context.Context, req *api.GetNamespaceReq) (*api.GetNamespaceResp, error) {
	namespace, err := m.namespaces.Get(domain.MakeNamespaceId(req.OrgId, req.Name))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.NotFound, "namespace not found")
		return nil, err
	}
	return &api.GetNamespaceResp{
		Name:      namespace.GetName(),
		Labels:    namespace.GetLabels(),
		Total:     namespace.GetResourceQuotas(),
		Available: namespace.GetAvailable(),
		Utilized:  namespace.GetUtilized(),
		Profile:   m.getSeccompProfile(ctx, namespace.GetSeccompProfile()),
	}, nil
}

func (m MeridianGrpcHandler) GetNamespaceHierarchy(ctx context.Context, req *api.GetNamespaceHierarchyReq) (*api.GetNamespaceHierarchyResp, error) {
	tree, err := m.namespaces.GetHierarchy(domain.MakeNamespaceId(req.OrgId, "default"))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.NotFound, "namespace hierarchy not found")
		return nil, err
	}
	return m.mapNamespaceTreeNode(ctx, &tree.Root), nil
}

func (m *MeridianGrpcHandler) mapNamespaceTreeNode(ctx context.Context, node *domain.NamespaceTreeNode) *api.GetNamespaceHierarchyResp {
	resp := &api.GetNamespaceHierarchyResp{
		Namespace: &api.GetNamespaceHierarchyResp_Namespace{
			Name:      node.Namespace.GetName(),
			Labels:    node.Namespace.GetLabels(),
			Total:     node.Namespace.GetResourceQuotas(),
			Available: node.Namespace.GetAvailable(),
			Utilized:  node.Namespace.GetUtilized(),
			Profile:   m.getSeccompProfile(ctx, node.Namespace.GetSeccompProfile()),
		},
	}
	for _, app := range node.Apps {
		resp.Apps = append(resp.Apps, &api.GetNamespaceHierarchyResp_App{
			Name:    app.GetName(),
			Total:   app.GetResourceQuotas(),
			Profile: m.getSeccompProfile(ctx, app.GetSeccompProfile()),
		})
	}
	for _, child := range node.Children {
		resp.Namespaces = append(resp.Namespaces, m.mapNamespaceTreeNode(ctx, child))
	}
	return resp
}

func (m *MeridianGrpcHandler) sendSeccompProfile(ctx context.Context, strategy string, metadata domain.SeccompProfile, profileDefinition *api.SeccompProfile, parent *domain.Namespace) error {
	switch strings.ToLower(strategy) {
	case "redefine":
		profile := &pulsar_api.SeccompProfileDefinitionRequest{
			Profile: &pulsar_api.SeccompProfile{
				Namespace:    metadata.Namespace,
				Application:  metadata.Application,
				Name:         metadata.Name,
				Version:      metadata.Version,
				Architecture: metadata.Architecture,
			},
			Definition: &pulsar_api.SeccompProfileDefinition{
				DefaultAction: profileDefinition.DefaultAction,
				Architectures: []string{metadata.Architecture},
			},
		}
		for _, syscall := range profileDefinition.Syscalls {
			profile.Definition.Syscalls = append(profile.Definition.Syscalls, &pulsar_api.Syscalls{
				Names:  syscall.Names,
				Action: syscall.Action,
			})
		}
		_, err := m.pulsar.DefineSeccompProfile(ctx, profile)
		if err != nil {
			return err
		}
	case "extend":
		if parent == nil {
			return status.Error(codes.InvalidArgument, "cannot inherit or extend seccomp profiles - there is no parent")
		}
		profile := &pulsar_api.ExtendSeccompProfileRequest{
			ExtendProfile: &pulsar_api.SeccompProfile{
				Namespace:    parent.GetId(),
				Application:  parent.GetSeccompProfile().Application,
				Name:         parent.GetSeccompProfile().Name,
				Version:      parent.GetSeccompProfile().Version,
				Architecture: parent.GetSeccompProfile().Architecture,
			},
			DefineProfile: &pulsar_api.SeccompProfile{
				Namespace:    metadata.Namespace,
				Application:  metadata.Application,
				Name:         metadata.Name,
				Version:      metadata.Version,
				Architecture: metadata.Architecture,
			},
		}
		for _, syscall := range profileDefinition.Syscalls {
			profile.Syscalls = append(profile.Syscalls, &pulsar_api.Syscalls{
				Names:  syscall.Names,
				Action: syscall.Action,
			})
		}
		_, err := m.pulsar.ExtendSeccompProfile(ctx, profile)
		if err != nil {
			return err
		}
	// inherit
	default:
		if parent == nil {
			return status.Error(codes.InvalidArgument, "cannot inherit or extend seccomp profiles - there is no parent")
		}
		profile, err := m.pulsar.GetSeccompProfile(ctx, &pulsar_api.SeccompProfile{
			Namespace:    parent.GetSeccompProfile().Namespace,
			Application:  parent.GetSeccompProfile().Application,
			Name:         parent.GetSeccompProfile().Name,
			Version:      parent.GetSeccompProfile().Version,
			Architecture: parent.GetSeccompProfile().Architecture,
		})
		if err != nil {
			return err
		}
		_, err = m.pulsar.DefineSeccompProfile(ctx, &pulsar_api.SeccompProfileDefinitionRequest{
			Profile: &pulsar_api.SeccompProfile{
				Namespace:    metadata.Namespace,
				Application:  metadata.Application,
				Name:         metadata.Name,
				Version:      metadata.Version,
				Architecture: metadata.Architecture,
			},
			Definition: profile.Definition,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MeridianGrpcHandler) getSeccompProfile(ctx context.Context, metadata domain.SeccompProfile) *api.SeccompProfile {
	resp, err := m.pulsar.GetSeccompProfile(ctx, &pulsar_api.SeccompProfile{
		Namespace:    metadata.Namespace,
		Application:  metadata.Application,
		Name:         metadata.Name,
		Version:      metadata.Version,
		Architecture: metadata.Architecture,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	profile := &api.SeccompProfile{
		Version:       metadata.Version,
		DefaultAction: resp.Definition.DefaultAction,
	}
	for _, syscall := range resp.Definition.Syscalls {
		profile.Syscalls = append(profile.Syscalls, &api.SyscallRule{
			Names:  syscall.Names,
			Action: syscall.Action,
		})
	}
	return profile
}
