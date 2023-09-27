// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type KubernetesAPIPolicy struct {
	DefaultRule *DefaultKubernetesAPIRule `json:"defaultRule,omitempty"`
	UserRules   []KubernetesAPIRule       `json:"userRules,omitempty"`
}

func (o *KubernetesAPIPolicy) GetDefaultRule() *DefaultKubernetesAPIRule {
	if o == nil {
		return nil
	}
	return o.DefaultRule
}

func (o *KubernetesAPIPolicy) GetUserRules() []KubernetesAPIRule {
	if o == nil {
		return nil
	}
	return o.UserRules
}
