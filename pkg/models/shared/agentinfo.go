// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type AgentInfo struct {
	Enabled *bool `json:"enabled,omitempty"`
	// The last time that the agent sent telemetries
	LastActive *time.Time `json:"lastActive,omitempty"`
	// The current controller state.
	Status  *ControllerStatus `json:"status,omitempty"`
	Version *string           `json:"version,omitempty"`
}

func (a AgentInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AgentInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AgentInfo) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AgentInfo) GetLastActive() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastActive
}

func (o *AgentInfo) GetStatus() *ControllerStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AgentInfo) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
