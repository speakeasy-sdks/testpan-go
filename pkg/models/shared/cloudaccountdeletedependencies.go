// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CloudAccountDeleteDependencies struct {
	Rules []DeleteDependencyElement `json:"rules,omitempty"`
}

func (o *CloudAccountDeleteDependencies) GetRules() []DeleteDependencyElement {
	if o == nil {
		return nil
	}
	return o.Rules
}
