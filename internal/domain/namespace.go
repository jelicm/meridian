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
	return fmt.Sprintf("%s/%s", n.orgId, n.name)
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

type SeccompProfile struct {
	Namespace    string
	Application  string
	Name         string
	Version      string
	Architecture string
}
