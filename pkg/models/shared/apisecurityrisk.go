// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type APISecurityRisk string

const (
	APISecurityRiskUnknown  APISecurityRisk = "UNKNOWN"
	APISecurityRiskNeutral  APISecurityRisk = "NEUTRAL"
	APISecurityRiskLow      APISecurityRisk = "LOW"
	APISecurityRiskMedium   APISecurityRisk = "MEDIUM"
	APISecurityRiskHigh     APISecurityRisk = "HIGH"
	APISecurityRiskCritical APISecurityRisk = "CRITICAL"
	APISecurityRiskNoRisk   APISecurityRisk = "NO_RISK"
)

func (e APISecurityRisk) ToPointer() *APISecurityRisk {
	return &e
}

func (e *APISecurityRisk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "NEUTRAL":
		fallthrough
	case "LOW":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		fallthrough
	case "CRITICAL":
		fallthrough
	case "NO_RISK":
		*e = APISecurityRisk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APISecurityRisk: %v", v)
	}
}
