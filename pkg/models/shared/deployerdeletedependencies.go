// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeployerDeleteDependencies struct {
	CdPolicies []DeleteDependencyElement `json:"cdPolicies,omitempty"`
}

func (o *DeployerDeleteDependencies) GetCdPolicies() []DeleteDependencyElement {
	if o == nil {
		return nil
	}
	return o.CdPolicies
}