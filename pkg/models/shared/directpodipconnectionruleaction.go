// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DirectPodIPConnectionRuleAction string

const (
	DirectPodIPConnectionRuleActionDetect DirectPodIPConnectionRuleAction = "DETECT"
	DirectPodIPConnectionRuleActionBlock  DirectPodIPConnectionRuleAction = "BLOCK"
)

func (e DirectPodIPConnectionRuleAction) ToPointer() *DirectPodIPConnectionRuleAction {
	return &e
}

func (e *DirectPodIPConnectionRuleAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DETECT":
		fallthrough
	case "BLOCK":
		*e = DirectPodIPConnectionRuleAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectPodIPConnectionRuleAction: %v", v)
	}
}
