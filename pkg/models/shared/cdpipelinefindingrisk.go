// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CDPipelineFindingRisk string

const (
	CDPipelineFindingRiskNoRisk CDPipelineFindingRisk = "NO_RISK"
	CDPipelineFindingRiskMedium CDPipelineFindingRisk = "MEDIUM"
	CDPipelineFindingRiskHigh   CDPipelineFindingRisk = "HIGH"
)

func (e CDPipelineFindingRisk) ToPointer() *CDPipelineFindingRisk {
	return &e
}

func (e *CDPipelineFindingRisk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_RISK":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		*e = CDPipelineFindingRisk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CDPipelineFindingRisk: %v", v)
	}
}
