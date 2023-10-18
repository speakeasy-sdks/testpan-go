// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteDependencyDeployerElement struct {
	CdPolicies []DeleteDependencyElement `json:"cdPolicies,omitempty"`
	Deployers  []DeleteDependencyElement `json:"deployers,omitempty"`
}

func (o *DeleteDependencyDeployerElement) GetCdPolicies() []DeleteDependencyElement {
	if o == nil {
		return nil
	}
	return o.CdPolicies
}

func (o *DeleteDependencyDeployerElement) GetDeployers() []DeleteDependencyElement {
	if o == nil {
		return nil
	}
	return o.Deployers
}