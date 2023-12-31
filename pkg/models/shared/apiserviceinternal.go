// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type APIServiceInternal struct {
	ClientWorkloads   []APIServiceClientWorkload  `json:"clientWorkloads,omitempty"`
	Cluster           *string                     `json:"cluster,omitempty"`
	Compliance        *APIServiceComplianceSimple `json:"compliance,omitempty"`
	CreationTimestamp *time.Time                  `json:"creation_timestamp,omitempty"`
	// Textual description of the Service
	Description      *string    `json:"description,omitempty"`
	FuzzingEndTime   *time.Time `json:"fuzzingEndTime,omitempty"`
	FuzzingStartTime *time.Time `json:"fuzzingStartTime,omitempty"`
	// An enumeration.
	FuzzingStatus        *FuzzingStatus `json:"fuzzingStatus,omitempty"`
	FuzzingStatusMessage *string        `json:"fuzzingStatusMessage,omitempty"`
	GatewayInfo          *Gateway       `json:"gatewayInfo,omitempty"`
	// Unique identifier of the subject API as assigned by Crankshaft
	Identifier string     `json:"identifier"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	// Location of the documentation. This can be an URL for example
	LinkDoc *string `json:"link_doc,omitempty"`
	// API name (for external) or destination workload (for internal)
	Name                   string                   `json:"name"`
	Namespace              *string                  `json:"namespace,omitempty"`
	OpenapiSpecAvailablity *OpenAPISpecAvailability `json:"openapi_spec_availablity,omitempty"`
	Port                   *int64                   `json:"port,omitempty"`
	// An `enum`eration.
	Risk  *APISecurityRiskSeverity `json:"risk,omitempty"`
	Score *APIServiceScore         `json:"score,omitempty"`
	// Api status enumeration.
	Status            *APISecurityAPIStatus `json:"status,omitempty"`
	StatusDescription *string               `json:"status_description,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (a APIServiceInternal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIServiceInternal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIServiceInternal) GetClientWorkloads() []APIServiceClientWorkload {
	if o == nil {
		return nil
	}
	return o.ClientWorkloads
}

func (o *APIServiceInternal) GetCluster() *string {
	if o == nil {
		return nil
	}
	return o.Cluster
}

func (o *APIServiceInternal) GetCompliance() *APIServiceComplianceSimple {
	if o == nil {
		return nil
	}
	return o.Compliance
}

func (o *APIServiceInternal) GetCreationTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTimestamp
}

func (o *APIServiceInternal) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *APIServiceInternal) GetFuzzingEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.FuzzingEndTime
}

func (o *APIServiceInternal) GetFuzzingStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.FuzzingStartTime
}

func (o *APIServiceInternal) GetFuzzingStatus() *FuzzingStatus {
	if o == nil {
		return nil
	}
	return o.FuzzingStatus
}

func (o *APIServiceInternal) GetFuzzingStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.FuzzingStatusMessage
}

func (o *APIServiceInternal) GetGatewayInfo() *Gateway {
	if o == nil {
		return nil
	}
	return o.GatewayInfo
}

func (o *APIServiceInternal) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *APIServiceInternal) GetLastUpdate() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUpdate
}

func (o *APIServiceInternal) GetLinkDoc() *string {
	if o == nil {
		return nil
	}
	return o.LinkDoc
}

func (o *APIServiceInternal) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *APIServiceInternal) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *APIServiceInternal) GetOpenapiSpecAvailablity() *OpenAPISpecAvailability {
	if o == nil {
		return nil
	}
	return o.OpenapiSpecAvailablity
}

func (o *APIServiceInternal) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *APIServiceInternal) GetRisk() *APISecurityRiskSeverity {
	if o == nil {
		return nil
	}
	return o.Risk
}

func (o *APIServiceInternal) GetScore() *APIServiceScore {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *APIServiceInternal) GetStatus() *APISecurityAPIStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *APIServiceInternal) GetStatusDescription() *string {
	if o == nil {
		return nil
	}
	return o.StatusDescription
}

func (o *APIServiceInternal) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}
