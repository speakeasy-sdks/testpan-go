// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ControllerStatus - The current controller state.
type ControllerStatus string

const (
	ControllerStatusPendingInstall        ControllerStatus = "PENDING_INSTALL"
	ControllerStatusActive                ControllerStatus = "ACTIVE"
	ControllerStatusInactive              ControllerStatus = "INACTIVE"
	ControllerStatusStopped               ControllerStatus = "STOPPED"
	ControllerStatusTerminated            ControllerStatus = "TERMINATED"
	ControllerStatusUnknown               ControllerStatus = "UNKNOWN"
	ControllerStatusWaitingForUserUpdate  ControllerStatus = "WAITING_FOR_USER_UPDATE"
	ControllerStatusAutoUpgradeInProgress ControllerStatus = "AUTO_UPGRADE_IN_PROGRESS"
)

func (e ControllerStatus) ToPointer() *ControllerStatus {
	return &e
}

func (e *ControllerStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING_INSTALL":
		fallthrough
	case "ACTIVE":
		fallthrough
	case "INACTIVE":
		fallthrough
	case "STOPPED":
		fallthrough
	case "TERMINATED":
		fallthrough
	case "UNKNOWN":
		fallthrough
	case "WAITING_FOR_USER_UPDATE":
		fallthrough
	case "AUTO_UPGRADE_IN_PROGRESS":
		*e = ControllerStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControllerStatus: %v", v)
	}
}
