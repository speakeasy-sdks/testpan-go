// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionRoleResources struct {
	APIGroups     []string        `json:"apiGroups,omitempty"`
	ResourceNames []string        `json:"resourceNames,omitempty"`
	Resources     []string        `json:"resources,omitempty"`
	Risk          *PermissionRisk `json:"risk,omitempty"`
	RiskReason    []string        `json:"riskReason,omitempty"`
	Verbs         []string        `json:"verbs,omitempty"`
}

func (o *PermissionRoleResources) GetAPIGroups() []string {
	if o == nil {
		return nil
	}
	return o.APIGroups
}

func (o *PermissionRoleResources) GetResourceNames() []string {
	if o == nil {
		return nil
	}
	return o.ResourceNames
}

func (o *PermissionRoleResources) GetResources() []string {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *PermissionRoleResources) GetRisk() *PermissionRisk {
	if o == nil {
		return nil
	}
	return o.Risk
}

func (o *PermissionRoleResources) GetRiskReason() []string {
	if o == nil {
		return nil
	}
	return o.RiskReason
}

func (o *PermissionRoleResources) GetVerbs() []string {
	if o == nil {
		return nil
	}
	return o.Verbs
}
