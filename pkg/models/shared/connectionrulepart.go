// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionRulePartConnectionRulePartType string

const (
	ConnectionRulePartConnectionRulePartTypeAppNameConnectionRulePart         ConnectionRulePartConnectionRulePartType = "AppNameConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeAppTypeConnectionRulePart         ConnectionRulePartConnectionRulePartType = "AppTypeConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeAppLabelConnectionRulePart        ConnectionRulePartConnectionRulePartType = "AppLabelConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeAppAnyConnectionRulePart          ConnectionRulePartConnectionRulePartType = "AppAnyConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypePodNameConnectionRulePart         ConnectionRulePartConnectionRulePartType = "PodNameConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypePodLablesConnectionRulePart       ConnectionRulePartConnectionRulePartType = "PodLablesConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypePodAnyConnectionRulePart          ConnectionRulePartConnectionRulePartType = "PodAnyConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeEnvironmentNameConnectionRulePart ConnectionRulePartConnectionRulePartType = "EnvironmentNameConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeEnvironmentAnyConnectionRulePart  ConnectionRulePartConnectionRulePartType = "EnvironmentAnyConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeIPRangeConnectionRulePart         ConnectionRulePartConnectionRulePartType = "IpRangeConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeExternalConnectionRulePart        ConnectionRulePartConnectionRulePartType = "ExternalConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeFqdnConnectionRulePart            ConnectionRulePartConnectionRulePartType = "FqdnConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeServiceNameConnectionRulePart     ConnectionRulePartConnectionRulePartType = "ServiceNameConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeAnyConnectionRulePart             ConnectionRulePartConnectionRulePartType = "AnyConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeExpansionNameConnectionRulePart   ConnectionRulePartConnectionRulePartType = "ExpansionNameConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeExpansionLabelsConnectionRulePart ConnectionRulePartConnectionRulePartType = "ExpansionLabelsConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeExpansionAnyConnectionRulePart    ConnectionRulePartConnectionRulePartType = "ExpansionAnyConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeKafkaConnectionRulePart           ConnectionRulePartConnectionRulePartType = "KafkaConnectionRulePart"
	ConnectionRulePartConnectionRulePartTypeAPIServiceConnectionRulePart      ConnectionRulePartConnectionRulePartType = "ApiServiceConnectionRulePart"
)

func (e ConnectionRulePartConnectionRulePartType) ToPointer() *ConnectionRulePartConnectionRulePartType {
	return &e
}

func (e *ConnectionRulePartConnectionRulePartType) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionRulePartConnectionRulePartType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRulePartConnectionRulePartType: %v", v)
	}
}

type ConnectionRulePart struct {
	ConnectionRulePartType ConnectionRulePartConnectionRulePartType `json:"connectionRulePartType"`
}

func (o *ConnectionRulePart) GetConnectionRulePartType() ConnectionRulePartConnectionRulePartType {
	if o == nil {
		return ConnectionRulePartConnectionRulePartType("")
	}
	return o.ConnectionRulePartType
}