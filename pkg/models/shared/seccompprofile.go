// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SeccompProfile struct {
	Data                *string  `json:"data,omitempty"`
	ID                  *string  `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	PodSecurityPolicies []string `json:"podSecurityPolicies,omitempty"`
}

func (o *SeccompProfile) GetData() *string {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *SeccompProfile) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SeccompProfile) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SeccompProfile) GetPodSecurityPolicies() []string {
	if o == nil {
		return nil
	}
	return o.PodSecurityPolicies
}
