// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CiPolicyInput struct {
	Description            *string                 `json:"description,omitempty"`
	DockerfileScanCiPolicy *DockerfileScanCiPolicy `json:"dockerfileScanCiPolicy,omitempty"`
	Name                   string                  `json:"name"`
	VulnerabilityCiPolicy  *VulnerabilityCiPolicy  `json:"vulnerabilityCiPolicy,omitempty"`
}

func (o *CiPolicyInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CiPolicyInput) GetDockerfileScanCiPolicy() *DockerfileScanCiPolicy {
	if o == nil {
		return nil
	}
	return o.DockerfileScanCiPolicy
}

func (o *CiPolicyInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CiPolicyInput) GetVulnerabilityCiPolicy() *VulnerabilityCiPolicy {
	if o == nil {
		return nil
	}
	return o.VulnerabilityCiPolicy
}
