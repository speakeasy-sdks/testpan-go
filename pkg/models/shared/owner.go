// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Owner struct {
	ID                     *string                 `json:"id,omitempty"`
	Namespace              *string                 `json:"namespace,omitempty"`
	Owner                  *string                 `json:"owner,omitempty"`
	OwnerType              *PermissionOwnerType    `json:"ownerType,omitempty"`
	PermissionInfo         []PermissionInfo        `json:"permissionInfo,omitempty"`
	SystemDefaultOwnerType *SystemDefaultOwnerType `json:"systemDefaultOwnerType,omitempty"`
}

func (o *Owner) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Owner) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *Owner) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *Owner) GetOwnerType() *PermissionOwnerType {
	if o == nil {
		return nil
	}
	return o.OwnerType
}

func (o *Owner) GetPermissionInfo() []PermissionInfo {
	if o == nil {
		return nil
	}
	return o.PermissionInfo
}

func (o *Owner) GetSystemDefaultOwnerType() *SystemDefaultOwnerType {
	if o == nil {
		return nil
	}
	return o.SystemDefaultOwnerType
}
