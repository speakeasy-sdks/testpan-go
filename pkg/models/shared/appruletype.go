// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RuleType string

const (
	RuleTypeInjectionRuleType RuleType = "InjectionRuleType"
	RuleTypeViolationRuleType RuleType = "ViolationRuleType"
)

func (e RuleType) ToPointer() *RuleType {
	return &e
}

func (e *RuleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "InjectionRuleType":
		fallthrough
	case "ViolationRuleType":
		*e = RuleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RuleType: %v", v)
	}
}

// AppRuleType - identify the app rule type. Only one of the below should be not null, and  used.
type AppRuleType struct {
	RuleType RuleType `json:"ruleType"`
}

func (o *AppRuleType) GetRuleType() RuleType {
	if o == nil {
		return RuleType("")
	}
	return o.RuleType
}
