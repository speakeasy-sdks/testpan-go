// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type KubernetesAPIRuleOrigin string

const (
	KubernetesAPIRuleOriginUser            KubernetesAPIRuleOrigin = "USER"
	KubernetesAPIRuleOriginAutomatedPolicy KubernetesAPIRuleOrigin = "AUTOMATED_POLICY"
	KubernetesAPIRuleOriginSystem          KubernetesAPIRuleOrigin = "SYSTEM"
)

func (e KubernetesAPIRuleOrigin) ToPointer() *KubernetesAPIRuleOrigin {
	return &e
}

func (e *KubernetesAPIRuleOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER":
		fallthrough
	case "AUTOMATED_POLICY":
		fallthrough
	case "SYSTEM":
		*e = KubernetesAPIRuleOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KubernetesAPIRuleOrigin: %v", v)
	}
}
