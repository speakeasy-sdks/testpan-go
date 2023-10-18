// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type APIRiskInfoServiceRisk string

const (
	APIRiskInfoServiceRiskCritical APIRiskInfoServiceRisk = "CRITICAL"
	APIRiskInfoServiceRiskHigh     APIRiskInfoServiceRisk = "HIGH"
	APIRiskInfoServiceRiskMedium   APIRiskInfoServiceRisk = "MEDIUM"
	APIRiskInfoServiceRiskLow      APIRiskInfoServiceRisk = "LOW"
)

func (e APIRiskInfoServiceRisk) ToPointer() *APIRiskInfoServiceRisk {
	return &e
}

func (e *APIRiskInfoServiceRisk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CRITICAL":
		fallthrough
	case "HIGH":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "LOW":
		*e = APIRiskInfoServiceRisk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIRiskInfoServiceRisk: %v", v)
	}
}

type APIRiskInfo struct {
	ServiceID   *string                 `json:"serviceId,omitempty"`
	ServiceName *string                 `json:"serviceName,omitempty"`
	ServiceRisk *APIRiskInfoServiceRisk `json:"serviceRisk,omitempty"`
	// An `enum`eration.
	ServiceType *APIServiceType `json:"serviceType,omitempty"`
}

func (o *APIRiskInfo) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *APIRiskInfo) GetServiceName() *string {
	if o == nil {
		return nil
	}
	return o.ServiceName
}

func (o *APIRiskInfo) GetServiceRisk() *APIRiskInfoServiceRisk {
	if o == nil {
		return nil
	}
	return o.ServiceRisk
}

func (o *APIRiskInfo) GetServiceType() *APIServiceType {
	if o == nil {
		return nil
	}
	return o.ServiceType
}