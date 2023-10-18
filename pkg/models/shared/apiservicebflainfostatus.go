// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type APIServiceBflaInfoStatus string

const (
	APIServiceBflaInfoStatusNoSpec              APIServiceBflaInfoStatus = "NO_SPEC"
	APIServiceBflaInfoStatusReady               APIServiceBflaInfoStatus = "READY"
	APIServiceBflaInfoStatusInProgressLearning  APIServiceBflaInfoStatus = "IN_PROGRESS_LEARNING"
	APIServiceBflaInfoStatusInProgressDetection APIServiceBflaInfoStatus = "IN_PROGRESS_DETECTION"
)

func (e APIServiceBflaInfoStatus) ToPointer() *APIServiceBflaInfoStatus {
	return &e
}

func (e *APIServiceBflaInfoStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_SPEC":
		fallthrough
	case "READY":
		fallthrough
	case "IN_PROGRESS_LEARNING":
		fallthrough
	case "IN_PROGRESS_DETECTION":
		*e = APIServiceBflaInfoStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIServiceBflaInfoStatus: %v", v)
	}
}