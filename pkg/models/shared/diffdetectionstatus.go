// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type DiffDetectionStatusStatus string

const (
	DiffDetectionStatusStatusReady      DiffDetectionStatusStatus = "READY"
	DiffDetectionStatusStatusInProgress DiffDetectionStatusStatus = "IN_PROGRESS"
)

func (e DiffDetectionStatusStatus) ToPointer() *DiffDetectionStatusStatus {
	return &e
}

func (e *DiffDetectionStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "READY":
		fallthrough
	case "IN_PROGRESS":
		*e = DiffDetectionStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiffDetectionStatusStatus: %v", v)
	}
}

// DiffDetectionStatus - diff detection status. in case of in progress, will reveal the detection end time
type DiffDetectionStatus struct {
	Endtime *time.Time                 `json:"endtime,omitempty"`
	Status  *DiffDetectionStatusStatus `json:"status,omitempty"`
}

func (d DiffDetectionStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DiffDetectionStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DiffDetectionStatus) GetEndtime() *time.Time {
	if o == nil {
		return nil
	}
	return o.Endtime
}

func (o *DiffDetectionStatus) GetStatus() *DiffDetectionStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
