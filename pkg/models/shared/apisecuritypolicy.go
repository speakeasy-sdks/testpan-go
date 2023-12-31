// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APISecurityPolicy struct {
	CategoryConditions *APISecurityPolicyCategoryConditions `json:"categoryConditions,omitempty"`
	Description        *string                              `json:"description,omitempty"`
	GlobalCondition    *APISecurityPolicyGlobalCondition    `json:"globalCondition,omitempty"`
	ID                 *string                              `json:"id,omitempty"`
	Name               string                               `json:"name"`
}

func (o *APISecurityPolicy) GetCategoryConditions() *APISecurityPolicyCategoryConditions {
	if o == nil {
		return nil
	}
	return o.CategoryConditions
}

func (o *APISecurityPolicy) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *APISecurityPolicy) GetGlobalCondition() *APISecurityPolicyGlobalCondition {
	if o == nil {
		return nil
	}
	return o.GlobalCondition
}

func (o *APISecurityPolicy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *APISecurityPolicy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
