// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DefaultKubernetesAPIRule struct {
	Action *KubernetesAPIRuleAction `json:"action,omitempty"`
}

func (o *DefaultKubernetesAPIRule) GetAction() *KubernetesAPIRuleAction {
	if o == nil {
		return nil
	}
	return o.Action
}