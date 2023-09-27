// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UnidentifiedServerlessRule struct {
	Action UnidentifiedServerlessRuleAction `json:"action"`
	Name   *string                          `json:"name,omitempty"`
}

func (o *UnidentifiedServerlessRule) GetAction() UnidentifiedServerlessRuleAction {
	if o == nil {
		return UnidentifiedServerlessRuleAction("")
	}
	return o.Action
}

func (o *UnidentifiedServerlessRule) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
