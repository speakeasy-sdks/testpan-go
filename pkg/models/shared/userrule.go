// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UserRule struct {
	// App rule action
	Action    *AppRuleAction `json:"action,omitempty"`
	ID        *string        `json:"id,omitempty"`
	IsDeleted *bool          `json:"isDeleted,omitempty"`
	Name      *string        `json:"name,omitempty"`
}

func (o *UserRule) GetAction() *AppRuleAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *UserRule) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UserRule) GetIsDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeleted
}

func (o *UserRule) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
