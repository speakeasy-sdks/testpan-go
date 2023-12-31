// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type State struct {
	Count            *int64            `json:"count,omitempty"`
	IsOpen           *bool             `json:"isOpen,omitempty"`
	LastSeen         *time.Time        `json:"lastSeen,omitempty"`
	Layer7Attributes []Layer7Attribute `json:"layer7Attributes,omitempty"`
	Protocol         *string           `json:"protocol,omitempty"`
	StartTime        *time.Time        `json:"startTime,omitempty"`
}

func (s State) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *State) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *State) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *State) GetIsOpen() *bool {
	if o == nil {
		return nil
	}
	return o.IsOpen
}

func (o *State) GetLastSeen() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastSeen
}

func (o *State) GetLayer7Attributes() []Layer7Attribute {
	if o == nil {
		return nil
	}
	return o.Layer7Attributes
}

func (o *State) GetProtocol() *string {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *State) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

type ConnectionTelemetry struct {
	APITokens []string    `json:"apiTokens,omitempty"`
	ID        *string     `json:"id,omitempty"`
	Source    *AppEnvInfo `json:"source,omitempty"`
	State     *State      `json:"state,omitempty"`
	Target    *AppEnvInfo `json:"target,omitempty"`
	// If there is a connection violation according to the policy - this object will hold the violation info
	Violation *ConnectionViolation `json:"violation,omitempty"`
}

func (o *ConnectionTelemetry) GetAPITokens() []string {
	if o == nil {
		return nil
	}
	return o.APITokens
}

func (o *ConnectionTelemetry) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConnectionTelemetry) GetSource() *AppEnvInfo {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *ConnectionTelemetry) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *ConnectionTelemetry) GetTarget() *AppEnvInfo {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *ConnectionTelemetry) GetViolation() *ConnectionViolation {
	if o == nil {
		return nil
	}
	return o.Violation
}
