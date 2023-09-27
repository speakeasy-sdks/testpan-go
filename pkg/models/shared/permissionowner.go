// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionOwner struct {
	ID                 *string         `json:"id,omitempty"`
	IsApproved         *bool           `json:"isApproved,omitempty"`
	IsSystemPermission *bool           `json:"isSystemPermission,omitempty"`
	Name               *string         `json:"name,omitempty"`
	Namespace          *string         `json:"namespace,omitempty"`
	Risk               *PermissionRisk `json:"risk,omitempty"`
	Scope              *string         `json:"scope,omitempty"`
}

func (o *PermissionOwner) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PermissionOwner) GetIsApproved() *bool {
	if o == nil {
		return nil
	}
	return o.IsApproved
}

func (o *PermissionOwner) GetIsSystemPermission() *bool {
	if o == nil {
		return nil
	}
	return o.IsSystemPermission
}

func (o *PermissionOwner) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PermissionOwner) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PermissionOwner) GetRisk() *PermissionRisk {
	if o == nil {
		return nil
	}
	return o.Risk
}

func (o *PermissionOwner) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}
