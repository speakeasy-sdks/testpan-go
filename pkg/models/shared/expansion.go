// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

// Expansion - represent expansion object
type Expansion struct {
	AccountName               *string `json:"accountName,omitempty"`
	ClusterID                 string  `json:"clusterId"`
	ClusterName               *string `json:"clusterName,omitempty"`
	ControllerEnabled         *bool   `json:"controllerEnabled,omitempty"`
	ControllerID              *string `json:"controllerId,omitempty"`
	ControllerIsUpdateEnabled *bool   `json:"controllerIsUpdateEnabled,omitempty"`
	// The last time that the agent sent telemetries
	ControllerLastActive *time.Time `json:"controllerLastActive,omitempty"`
	ControllerStatus     *string    `json:"controllerStatus,omitempty"`
	ControllerVersion    *string    `json:"controllerVersion,omitempty"`
	// unique Id
	ID                *string           `json:"id,omitempty"`
	InstanceID        *string           `json:"instanceId,omitempty"`
	Labels            []Label           `json:"labels,omitempty"`
	Name              string            `json:"name"`
	NamespaceID       string            `json:"namespaceId"`
	ShouldSendMetrics *bool             `json:"shouldSendMetrics,omitempty"`
	WorkloadAddresses []WorkloadAddress `json:"workloadAddresses"`
}

func (e Expansion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Expansion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Expansion) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *Expansion) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

func (o *Expansion) GetClusterName() *string {
	if o == nil {
		return nil
	}
	return o.ClusterName
}

func (o *Expansion) GetControllerEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.ControllerEnabled
}

func (o *Expansion) GetControllerID() *string {
	if o == nil {
		return nil
	}
	return o.ControllerID
}

func (o *Expansion) GetControllerIsUpdateEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.ControllerIsUpdateEnabled
}

func (o *Expansion) GetControllerLastActive() *time.Time {
	if o == nil {
		return nil
	}
	return o.ControllerLastActive
}

func (o *Expansion) GetControllerStatus() *string {
	if o == nil {
		return nil
	}
	return o.ControllerStatus
}

func (o *Expansion) GetControllerVersion() *string {
	if o == nil {
		return nil
	}
	return o.ControllerVersion
}

func (o *Expansion) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Expansion) GetInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.InstanceID
}

func (o *Expansion) GetLabels() []Label {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *Expansion) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Expansion) GetNamespaceID() string {
	if o == nil {
		return ""
	}
	return o.NamespaceID
}

func (o *Expansion) GetShouldSendMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.ShouldSendMetrics
}

func (o *Expansion) GetWorkloadAddresses() []WorkloadAddress {
	if o == nil {
		return []WorkloadAddress{}
	}
	return o.WorkloadAddresses
}
