package handlers

import (
	"context"
	"log"

	"github.com/c12s/meridian/internal/domain"
	"github.com/c12s/meridian/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MeridianGrpcHandler struct {
	api.UnimplementedMeridianServer
	namespaces domain.NamespaceStore
	apps       domain.AppStore
}

func NewMeridianGrpcHandler(namespaces domain.NamespaceStore, apps domain.AppStore) api.MeridianServer {
	return MeridianGrpcHandler{
		namespaces: namespaces,
		apps:       apps,
	}
}

func (m MeridianGrpcHandler) AddNamespace(ctx context.Context, req *api.AddNamespaceReq) (*api.AddNamespaceResp, error) {
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
	namespace := domain.NewNamespace(req.OrgId, req.Name, req.Profile.Version, req.Labels)
	for resource, quota := range req.Quotas {
		err := namespace.AddResourceQuota(resource, quota)
		if err != nil {
			log.Println(err)
			err = status.Error(codes.InvalidArgument, err.Error())
			return nil, err
		}
	}
	// todo: send seccomp profile
	err := m.namespaces.Add(namespace, parent)
	if err != nil {
		log.Println(err)
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &api.AddNamespaceResp{}, nil
}

func (m MeridianGrpcHandler) RemoveNamespace(ctx context.Context, req *api.RemoveNamespaceReq) (*api.RemoveNamespaceResp, error) {
	err := m.namespaces.Remove(domain.MakeNamespaceId(req.OrgId, req.Name))
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
	// todo: send seccomp profile
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
		Name:   namespace.GetName(),
		Labels: namespace.GetLabels(),
		Total:  namespace.GetResourceQuotas(),
		// todo: calculate available and utilized
		// todo: get seccomp profile
		Profile: &api.SeccompProfile{},
	}, nil
}

func (m MeridianGrpcHandler) GetNamespaceHierarchy(ctx context.Context, req *api.GetNamespaceHierarchyReq) (*api.GetNamespaceHierarchyResp, error) {
	tree, err := m.namespaces.GetHierarchy(domain.MakeNamespaceId(req.OrgId, "default"))
	if err != nil {
		log.Println(err)
		err = status.Error(codes.NotFound, "namespace hierarchy not found")
		return nil, err
	}
	return mapNamespaceTreeNode(&tree.Root), nil
}

func mapNamespaceTreeNode(node *domain.NamespaceTreeNode) *api.GetNamespaceHierarchyResp {
	resp := &api.GetNamespaceHierarchyResp{
		Namespace: &api.GetNamespaceHierarchyResp_Namespace{
			Name:   node.Namespace.GetName(),
			Labels: node.Namespace.GetLabels(),
			Total:  node.Namespace.GetResourceQuotas(),
			// todo: calculate available and utilized
			// todo: get seccomp profile
			Profile: &api.SeccompProfile{},
		},
	}
	for _, app := range node.Apps {
		resp.Apps = append(resp.Apps, &api.GetNamespaceHierarchyResp_App{
			Name:  app.GetName(),
			Total: app.GetResourceQuotas(),
			// todo: calculate available and utilized
			// todo: get seccomp profile
			Profile: &api.SeccompProfile{},
		})
	}
	for _, child := range node.Children {
		resp.Namespaces = append(resp.Namespaces, mapNamespaceTreeNode(child))
	}
	return resp
}
