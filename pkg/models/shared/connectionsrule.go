// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
)

type ConnectionsRuleStatus string

const (
	ConnectionsRuleStatusEnabled  ConnectionsRuleStatus = "ENABLED"
	ConnectionsRuleStatusDisabled ConnectionsRuleStatus = "DISABLED"
	ConnectionsRuleStatusDeleted  ConnectionsRuleStatus = "DELETED"
)

func (e ConnectionsRuleStatus) ToPointer() *ConnectionsRuleStatus {
	return &e
}

func (e *ConnectionsRuleStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = ConnectionsRuleStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionsRuleStatus: %v", v)
	}
}

// ConnectionsRule - A rule that states what Apps are allowed to communicate with each other.
type ConnectionsRule struct {
	// ENCRYPT is not allowed in default rule
	Action         ConnectionRuleAction       `json:"action"`
	Destination    *ConnectionRulePart        `json:"destination,omitempty"`
	GroupName      *string                    `json:"groupName,omitempty"`
	ID             *string                    `json:"id,omitempty"`
	IsRuleActive   *bool                      `json:"isRuleActive,omitempty"`
	Layer7Settings *Layer7SettingsPart        `json:"layer7Settings,omitempty"`
	Name           string                     `json:"name"`
	RuleOrigin     *ConnectionRuleOrigin      `default:"USER" json:"ruleOrigin"`
	RuleType       *NetworkConnectionRuleType `json:"ruleType,omitempty"`
	Source         *ConnectionRulePart        `json:"source,omitempty"`
	Status         *ConnectionsRuleStatus     `json:"status,omitempty"`
}

func (c ConnectionsRule) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionsRule) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionsRule) GetAction() ConnectionRuleAction {
	if o == nil {
		return ConnectionRuleAction("")
	}
	return o.Action
}

func (o *ConnectionsRule) GetDestination() *ConnectionRulePart {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *ConnectionsRule) GetGroupName() *string {
	if o == nil {
		return nil
	}
	return o.GroupName
}

func (o *ConnectionsRule) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConnectionsRule) GetIsRuleActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsRuleActive
}

func (o *ConnectionsRule) GetLayer7Settings() *Layer7SettingsPart {
	if o == nil {
		return nil
	}
	return o.Layer7Settings
}

func (o *ConnectionsRule) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionsRule) GetRuleOrigin() *ConnectionRuleOrigin {
	if o == nil {
		return nil
	}
	return o.RuleOrigin
}

func (o *ConnectionsRule) GetRuleType() *NetworkConnectionRuleType {
	if o == nil {
		return nil
	}
	return o.RuleType
}

func (o *ConnectionsRule) GetSource() *ConnectionRulePart {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *ConnectionsRule) GetStatus() *ConnectionsRuleStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
