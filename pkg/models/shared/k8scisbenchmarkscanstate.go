// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type K8sCisBenchmarkScanState string

const (
	K8sCisBenchmarkScanStateReady      K8sCisBenchmarkScanState = "READY"
	K8sCisBenchmarkScanStateInProgress K8sCisBenchmarkScanState = "IN_PROGRESS"
	K8sCisBenchmarkScanStateStarting   K8sCisBenchmarkScanState = "STARTING"
	K8sCisBenchmarkScanStateError      K8sCisBenchmarkScanState = "ERROR"
)

func (e K8sCisBenchmarkScanState) ToPointer() *K8sCisBenchmarkScanState {
	return &e
}

func (e *K8sCisBenchmarkScanState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "READY":
		fallthrough
	case "IN_PROGRESS":
		fallthrough
	case "STARTING":
		fallthrough
	case "ERROR":
		*e = K8sCisBenchmarkScanState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for K8sCisBenchmarkScanState: %v", v)
	}
}
