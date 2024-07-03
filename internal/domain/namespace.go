package domain

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
)

const (
	SECCOMP_APP_WILDCARD = "*"
	SECCOMP_DEFAULT_ARCH = "x86"
)

type SeccompProfile struct {
	Namespace    string
	Application  string
	Name         string
	Version      string
	Architecture string
}

func MakeNamespaceId(orgId, namespaceName string) string {
	return fmt.Sprintf("%s/%s", orgId, namespaceName)
}

type Namespace struct {
	orgId          string
	name           string
	resourceQuotas ResourceQuotas
	profileVersion string
}

func NewNamespace(orgId, name, profileVersion string) Namespace {
	return Namespace{
		orgId:          orgId,
		name:           name,
		profileVersion: profileVersion,
		resourceQuotas: make(ResourceQuotas, 0),
	}
}

func (n Namespace) GetOrgId() string {
	return n.orgId
}

func (n Namespace) GetName() string {
	return n.name
}

func (n Namespace) GetResourceQuotas() ResourceQuotas {
	quotas := make(ResourceQuotas)
	maps.Copy(quotas, n.resourceQuotas)
	return quotas
}

func (n *Namespace) AddResourceQuota(resource string, quota float64) error {
	if !slices.Contains(SupportedResourceQuotas, resource) {
		return fmt.Errorf("quotas for a resource with name %s are not supported", resource)
	}
	n.resourceQuotas[resource] = quota
	return nil
}

func (n Namespace) GetProfileVersion() string {
	return n.profileVersion
}

func (n Namespace) GetId() string {
	return MakeNamespaceId(n.orgId, n.name)
}

func (n Namespace) GetSeccompProfile() SeccompProfile {
	return SeccompProfile{
		Namespace:    n.GetId(),
		Application:  SECCOMP_APP_WILDCARD,
		Name:         fmt.Sprintf("%s profile", n.GetId()),
		Version:      n.profileVersion,
		Architecture: SECCOMP_DEFAULT_ARCH,
	}
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		OrgId          string         `json:"org_id"`
		Name           string         `json:"name"`
		SeccompProfile SeccompProfile `json:"seccomp_profile"`
		ResourceQuotas ResourceQuotas `json:"resource_quotas"`
	}{
		OrgId:          n.orgId,
		Name:           n.name,
		SeccompProfile: n.GetSeccompProfile(),
		ResourceQuotas: n.resourceQuotas,
	})
}

type NamespaceTreeNode struct {
	Namespace Namespace           `json:"namespace"`
	Apps      []App               `json:"applications"`
	Children  []NamespaceTreeNode `json:"child_namespaces"`
}

type NamespaceTree struct {
	Root NamespaceTreeNode
}

type NamespaceStore interface {
	Add(namespace Namespace, parent *Namespace) error
	Get(id string) (Namespace, error)
	GetHierarchy(rootId string) (NamespaceTree, error)
	Remove(id string) error
}
