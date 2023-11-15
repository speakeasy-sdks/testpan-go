// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionRulePartType string

const (
	ConnectionRulePartTypeAppNameConnectionRulePart         ConnectionRulePartType = "AppNameConnectionRulePart"
	ConnectionRulePartTypeAppTypeConnectionRulePart         ConnectionRulePartType = "AppTypeConnectionRulePart"
	ConnectionRulePartTypeAppLabelConnectionRulePart        ConnectionRulePartType = "AppLabelConnectionRulePart"
	ConnectionRulePartTypeAppAnyConnectionRulePart          ConnectionRulePartType = "AppAnyConnectionRulePart"
	ConnectionRulePartTypePodNameConnectionRulePart         ConnectionRulePartType = "PodNameConnectionRulePart"
	ConnectionRulePartTypePodLablesConnectionRulePart       ConnectionRulePartType = "PodLablesConnectionRulePart"
	ConnectionRulePartTypePodAnyConnectionRulePart          ConnectionRulePartType = "PodAnyConnectionRulePart"
	ConnectionRulePartTypeEnvironmentNameConnectionRulePart ConnectionRulePartType = "EnvironmentNameConnectionRulePart"
	ConnectionRulePartTypeEnvironmentAnyConnectionRulePart  ConnectionRulePartType = "EnvironmentAnyConnectionRulePart"
	ConnectionRulePartTypeIPRangeConnectionRulePart         ConnectionRulePartType = "IpRangeConnectionRulePart"
	ConnectionRulePartTypeExternalConnectionRulePart        ConnectionRulePartType = "ExternalConnectionRulePart"
	ConnectionRulePartTypeFqdnConnectionRulePart            ConnectionRulePartType = "FqdnConnectionRulePart"
	ConnectionRulePartTypeServiceNameConnectionRulePart     ConnectionRulePartType = "ServiceNameConnectionRulePart"
	ConnectionRulePartTypeAnyConnectionRulePart             ConnectionRulePartType = "AnyConnectionRulePart"
	ConnectionRulePartTypeExpansionNameConnectionRulePart   ConnectionRulePartType = "ExpansionNameConnectionRulePart"
	ConnectionRulePartTypeExpansionLabelsConnectionRulePart ConnectionRulePartType = "ExpansionLabelsConnectionRulePart"
	ConnectionRulePartTypeExpansionAnyConnectionRulePart    ConnectionRulePartType = "ExpansionAnyConnectionRulePart"
	ConnectionRulePartTypeKafkaConnectionRulePart           ConnectionRulePartType = "KafkaConnectionRulePart"
	ConnectionRulePartTypeAPIServiceConnectionRulePart      ConnectionRulePartType = "ApiServiceConnectionRulePart"
)

func (e ConnectionRulePartType) ToPointer() *ConnectionRulePartType {
	return &e
}

func (e *ConnectionRulePartType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AppNameConnectionRulePart":
		fallthrough
	case "AppTypeConnectionRulePart":
		fallthrough
	case "AppLabelConnectionRulePart":
		fallthrough
	case "AppAnyConnectionRulePart":
		fallthrough
	case "PodNameConnectionRulePart":
		fallthrough
	case "PodLablesConnectionRulePart":
		fallthrough
	case "PodAnyConnectionRulePart":
		fallthrough
	case "EnvironmentNameConnectionRulePart":
		fallthrough
	case "EnvironmentAnyConnectionRulePart":
		fallthrough
	case "IpRangeConnectionRulePart":
		fallthrough
	case "ExternalConnectionRulePart":
		fallthrough
	case "FqdnConnectionRulePart":
		fallthrough
	case "ServiceNameConnectionRulePart":
		fallthrough
	case "AnyConnectionRulePart":
		fallthrough
	case "ExpansionNameConnectionRulePart":
		fallthrough
	case "ExpansionLabelsConnectionRulePart":
		fallthrough
	case "ExpansionAnyConnectionRulePart":
		fallthrough
	case "KafkaConnectionRulePart":
		fallthrough
	case "ApiServiceConnectionRulePart":
		*e = ConnectionRulePartType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRulePartType: %v", v)
	}
}

type ConnectionRulePart struct {
	ConnectionRulePartType ConnectionRulePartType `json:"connectionRulePartType"`
}

func (o *ConnectionRulePart) GetConnectionRulePartType() ConnectionRulePartType {
	if o == nil {
		return ConnectionRulePartType("")
	}
	return o.ConnectionRulePartType
}
