// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteDependencyElementUser struct {
	Deployers *DeleteDependencyDeployerElement `json:"deployers,omitempty"`
}

func (o *DeleteDependencyElementUser) GetDeployers() *DeleteDependencyDeployerElement {
	if o == nil {
		return nil
	}
	return o.Deployers
}
