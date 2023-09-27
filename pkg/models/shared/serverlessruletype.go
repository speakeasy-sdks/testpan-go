// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ServerlessRuleTypeServerlessRuleType string

const (
	ServerlessRuleTypeServerlessRuleTypeFunctionAnyServerlessRuleType  ServerlessRuleTypeServerlessRuleType = "FunctionAnyServerlessRuleType"
	ServerlessRuleTypeServerlessRuleTypeFunctionNameServerlessRuleType ServerlessRuleTypeServerlessRuleType = "FunctionNameServerlessRuleType"
	ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType  ServerlessRuleTypeServerlessRuleType = "FunctionArnServerlessRuleType"
)

func (e ServerlessRuleTypeServerlessRuleType) ToPointer() *ServerlessRuleTypeServerlessRuleType {
	return &e
}

func (e *ServerlessRuleTypeServerlessRuleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FunctionAnyServerlessRuleType":
		fallthrough
	case "FunctionNameServerlessRuleType":
		fallthrough
	case "FunctionArnServerlessRuleType":
		*e = ServerlessRuleTypeServerlessRuleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerlessRuleTypeServerlessRuleType: %v", v)
	}
}

// ServerlessRuleType - identify the serverless functions matching type. Only one of the below should be not null, and  used.
type ServerlessRuleType struct {
	ServerlessFunctionValidation *ServerlessFunctionValidation        `json:"serverlessFunctionValidation,omitempty"`
	ServerlessRuleType           ServerlessRuleTypeServerlessRuleType `json:"serverlessRuleType"`
}

func (o *ServerlessRuleType) GetServerlessFunctionValidation() *ServerlessFunctionValidation {
	if o == nil {
		return nil
	}
	return o.ServerlessFunctionValidation
}

func (o *ServerlessRuleType) GetServerlessRuleType() ServerlessRuleTypeServerlessRuleType {
	if o == nil {
		return ServerlessRuleTypeServerlessRuleType("")
	}
	return o.ServerlessRuleType
}
