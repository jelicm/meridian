package domain

import (
	"fmt"
	"slices"
)

const (
	SECCOMP_APP_WILDCARD = "*"
	SECCOMP_DEFAULT_ARCH = "x86"
)

var (
	supportedResourceQuotas = []string{"mem", "cpu", "disk"}
)

type ResourceQuota struct {
	ResourceName string
	Quota        float64
}

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

func MakeAppId(orgId, namespaceName, appName string) string {
	return fmt.Sprintf("%s/%s", MakeNamespaceId(orgId, namespaceName), appName)
}

type Namespace struct {
	orgId          string
	name           string
	resourceQuotas []ResourceQuota
	profileVersion string
}

func (n Namespace) GetOrgId() string {
	return n.orgId
}

func (n Namespace) GetName() string {
	return n.name
}

func (n Namespace) GetResourceQuotas() []ResourceQuota {
	quotas := make([]ResourceQuota, len(n.resourceQuotas))
	copy(quotas, n.resourceQuotas)
	return quotas
}

func (n *Namespace) AddResourceQuota(quota ResourceQuota) error {
	if !slices.Contains(supportedResourceQuotas, quota.ResourceName) {
		return fmt.Errorf("quotas for a resource with name %s are not supported", quota.ResourceName)
	}
	n.resourceQuotas = append(n.resourceQuotas, quota)
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

type NamespaceTreeNode struct {
	Namespace Namespace
	Children  []Namespace
}

type NamespaceTree struct {
	Root *NamespaceTreeNode
}

type App struct {
	namespace      Namespace
	name           string
	resourceQuotas []ResourceQuota
}

func (a App) GetNamespace() Namespace {
	return a.namespace
}

func (a App) GetName() string {
	return a.name
}

func (a App) GetId() string {
	return MakeAppId(a.namespace.orgId, a.namespace.name, a.name)
}

func (a App) GetResourceQuotas() []ResourceQuota {
	quotas := make([]ResourceQuota, len(a.resourceQuotas))
	copy(quotas, a.resourceQuotas)
	return quotas
}

func (a *App) AddResourceQuota(quota ResourceQuota) error {
	if !slices.Contains(supportedResourceQuotas, quota.ResourceName) {
		return fmt.Errorf("quotas for a resource with name %s are not supported", quota.ResourceName)
	}
	a.resourceQuotas = append(a.resourceQuotas, quota)
	return nil
}

type NamespaceStore interface {
	Add(namespace, parent Namespace) error
	Get(id string) (Namespace, error)
	GetHierarchy(rootId string) (NamespaceTree, error)
	Remove(id string) error
}

type AppStore interface {
	Add(app App) error
	FindByNamespace(namespaceId string) ([]App, error)
	Remove(id string) error
}
