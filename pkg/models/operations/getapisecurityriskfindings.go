// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// APISecSource - source filter. an enum representing the source of the APIs service in scope
type APISecSource string

const (
	APISecSourceInternal APISecSource = "INTERNAL"
	APISecSourceExternal APISecSource = "EXTERNAL"
)

func (e APISecSource) ToPointer() *APISecSource {
	return &e
}

func (e *APISecSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = APISecSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APISecSource: %v", v)
	}
}

type Risks string

const (
	RisksLow      Risks = "LOW"
	RisksMedium   Risks = "MEDIUM"
	RisksHigh     Risks = "HIGH"
	RisksCritical Risks = "CRITICAL"
	RisksAll      Risks = "ALL"
)

func (e Risks) ToPointer() *Risks {
	return &e
}

func (e *Risks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOW":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		fallthrough
	case "CRITICAL":
		fallthrough
	case "ALL":
		*e = Risks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Risks: %v", v)
	}
}

// GetAPISecurityRiskFindingsQueryParamSortDir - sorting direction
type GetAPISecurityRiskFindingsQueryParamSortDir string

const (
	GetAPISecurityRiskFindingsQueryParamSortDirAsc  GetAPISecurityRiskFindingsQueryParamSortDir = "ASC"
	GetAPISecurityRiskFindingsQueryParamSortDirDesc GetAPISecurityRiskFindingsQueryParamSortDir = "DESC"
)

func (e GetAPISecurityRiskFindingsQueryParamSortDir) ToPointer() *GetAPISecurityRiskFindingsQueryParamSortDir {
	return &e
}

func (e *GetAPISecurityRiskFindingsQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetAPISecurityRiskFindingsQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAPISecurityRiskFindingsQueryParamSortDir: %v", v)
	}
}

// GetAPISecurityRiskFindingsQueryParamSortKey - Risk finding sort key.
type GetAPISecurityRiskFindingsQueryParamSortKey string

const (
	GetAPISecurityRiskFindingsQueryParamSortKeyName GetAPISecurityRiskFindingsQueryParamSortKey = "NAME"
	GetAPISecurityRiskFindingsQueryParamSortKeyRisk GetAPISecurityRiskFindingsQueryParamSortKey = "RISK"
)

func (e GetAPISecurityRiskFindingsQueryParamSortKey) ToPointer() *GetAPISecurityRiskFindingsQueryParamSortKey {
	return &e
}

func (e *GetAPISecurityRiskFindingsQueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NAME":
		fallthrough
	case "RISK":
		*e = GetAPISecurityRiskFindingsQueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAPISecurityRiskFindingsQueryParamSortKey: %v", v)
	}
}

type GetAPISecurityRiskFindingsRequest struct {
	// source filter. an enum representing the source of the APIs service in scope
	APISecSource APISecSource `default:"INTERNAL" queryParam:"style=form,explode=true,name=apiSecSource"`
	// Category of the risk finding
	Category *string `queryParam:"style=form,explode=true,name=category"`
	// Show finding with detect elements only
	Detected *bool `queryParam:"style=form,explode=true,name=detected"`
	// Affected element of the risk finding
	Element *string `queryParam:"style=form,explode=true,name=element"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Name of the risk finding name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// The API risk filter
	Risks []Risks `queryParam:"style=form,explode=false,name=risks"`
	// sorting direction
	SortDir *GetAPISecurityRiskFindingsQueryParamSortDir `default:"DESC" queryParam:"style=form,explode=true,name=sortDir"`
	// Risk finding sort key.
	SortKey GetAPISecurityRiskFindingsQueryParamSortKey `default:"RISK" queryParam:"style=form,explode=true,name=sortKey"`
	// Source of the risk finding
	Source *string `queryParam:"style=form,explode=true,name=source"`
}

func (g GetAPISecurityRiskFindingsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAPISecurityRiskFindingsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAPISecurityRiskFindingsRequest) GetAPISecSource() APISecSource {
	if o == nil {
		return APISecSource("")
	}
	return o.APISecSource
}

func (o *GetAPISecurityRiskFindingsRequest) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *GetAPISecurityRiskFindingsRequest) GetDetected() *bool {
	if o == nil {
		return nil
	}
	return o.Detected
}

func (o *GetAPISecurityRiskFindingsRequest) GetElement() *string {
	if o == nil {
		return nil
	}
	return o.Element
}

func (o *GetAPISecurityRiskFindingsRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetAPISecurityRiskFindingsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetAPISecurityRiskFindingsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetAPISecurityRiskFindingsRequest) GetRisks() []Risks {
	if o == nil {
		return nil
	}
	return o.Risks
}

func (o *GetAPISecurityRiskFindingsRequest) GetSortDir() *GetAPISecurityRiskFindingsQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetAPISecurityRiskFindingsRequest) GetSortKey() GetAPISecurityRiskFindingsQueryParamSortKey {
	if o == nil {
		return GetAPISecurityRiskFindingsQueryParamSortKey("")
	}
	return o.SortKey
}

func (o *GetAPISecurityRiskFindingsRequest) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

type GetAPISecurityRiskFindingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	RiskFindings *shared.RiskFindings
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPISecurityRiskFindingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityRiskFindingsResponse) GetRiskFindings() *shared.RiskFindings {
	if o == nil {
		return nil
	}
	return o.RiskFindings
}

func (o *GetAPISecurityRiskFindingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityRiskFindingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
