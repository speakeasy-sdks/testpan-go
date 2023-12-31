// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type APIServiceDrillDownExternal struct {
	// API classification label as determined by Crankshaft, e.g. ['meetings', 'messaging']
	Classification    []string                   `json:"classification,omitempty"`
	ClientWorkloads   []APIServiceClientWorkload `json:"clientWorkloads,omitempty"`
	Compliance        *APIServiceCompliance      `json:"compliance,omitempty"`
	CreationTimestamp *time.Time                 `json:"creation_timestamp,omitempty"`
	// Textual description of the Service
	Description *string `json:"description,omitempty"`
	// Unique identifier of the subject API as assigned by Crankshaft
	Identifier          string `json:"identifier"`
	IsAPITracingEnabled *bool  `json:"isApiTracingEnabled,omitempty"`
	// Location of the documentation. This can be an URL for example
	LinkDoc *string `json:"link_doc,omitempty"`
	// API name, usually an FQDN as determined by crankshaft, it can be logical or can correspond to one of the endpoints where the API is reachable, i.e. api.webex.com
	Name                   string                   `json:"name"`
	OpenapiSpecAvailablity *OpenAPISpecAvailability `json:"openapi_spec_availablity,omitempty"`
	Provider               *APIProviderBase         `json:"provider,omitempty"`
	// An `enum`eration.
	Risk  *APISecurityRiskSeverity `json:"risk,omitempty"`
	Score *APIServiceScore         `json:"score,omitempty"`
	// Api status enumeration.
	Status            *APISecurityAPIStatus `json:"status,omitempty"`
	StatusDescription *string               `json:"status_description,omitempty"`
}

func (a APIServiceDrillDownExternal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIServiceDrillDownExternal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIServiceDrillDownExternal) GetClassification() []string {
	if o == nil {
		return nil
	}
	return o.Classification
}

func (o *APIServiceDrillDownExternal) GetClientWorkloads() []APIServiceClientWorkload {
	if o == nil {
		return nil
	}
	return o.ClientWorkloads
}

func (o *APIServiceDrillDownExternal) GetCompliance() *APIServiceCompliance {
	if o == nil {
		return nil
	}
	return o.Compliance
}

func (o *APIServiceDrillDownExternal) GetCreationTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTimestamp
}

func (o *APIServiceDrillDownExternal) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *APIServiceDrillDownExternal) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *APIServiceDrillDownExternal) GetIsAPITracingEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IsAPITracingEnabled
}

func (o *APIServiceDrillDownExternal) GetLinkDoc() *string {
	if o == nil {
		return nil
	}
	return o.LinkDoc
}

func (o *APIServiceDrillDownExternal) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *APIServiceDrillDownExternal) GetOpenapiSpecAvailablity() *OpenAPISpecAvailability {
	if o == nil {
		return nil
	}
	return o.OpenapiSpecAvailablity
}

func (o *APIServiceDrillDownExternal) GetProvider() *APIProviderBase {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *APIServiceDrillDownExternal) GetRisk() *APISecurityRiskSeverity {
	if o == nil {
		return nil
	}
	return o.Risk
}

func (o *APIServiceDrillDownExternal) GetScore() *APIServiceScore {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *APIServiceDrillDownExternal) GetStatus() *APISecurityAPIStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *APIServiceDrillDownExternal) GetStatusDescription() *string {
	if o == nil {
		return nil
	}
	return o.StatusDescription
}
