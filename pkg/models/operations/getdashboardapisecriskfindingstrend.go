// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetDashboardApisecRiskFindingsTrendAPISecSource - source filter. an enum representing the source of the APIs service in scope
type GetDashboardApisecRiskFindingsTrendAPISecSource string

const (
	GetDashboardApisecRiskFindingsTrendAPISecSourceInternal GetDashboardApisecRiskFindingsTrendAPISecSource = "INTERNAL"
	GetDashboardApisecRiskFindingsTrendAPISecSourceExternal GetDashboardApisecRiskFindingsTrendAPISecSource = "EXTERNAL"
)

func (e GetDashboardApisecRiskFindingsTrendAPISecSource) ToPointer() *GetDashboardApisecRiskFindingsTrendAPISecSource {
	return &e
}

func (e *GetDashboardApisecRiskFindingsTrendAPISecSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = GetDashboardApisecRiskFindingsTrendAPISecSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDashboardApisecRiskFindingsTrendAPISecSource: %v", v)
	}
}

type GetDashboardApisecRiskFindingsTrendRequest struct {
	// source filter. an enum representing the source of the APIs service in scope
	APISecSource GetDashboardApisecRiskFindingsTrendAPISecSource `default:"INTERNAL" queryParam:"style=form,explode=true,name=apiSecSource"`
	// the desired number of days in graph
	NumOfDays *int64 `default:"30" queryParam:"style=form,explode=true,name=numOfDays"`
}

func (g GetDashboardApisecRiskFindingsTrendRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDashboardApisecRiskFindingsTrendRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetDashboardApisecRiskFindingsTrendRequest) GetAPISecSource() GetDashboardApisecRiskFindingsTrendAPISecSource {
	if o == nil {
		return GetDashboardApisecRiskFindingsTrendAPISecSource("")
	}
	return o.APISecSource
}

func (o *GetDashboardApisecRiskFindingsTrendRequest) GetNumOfDays() *int64 {
	if o == nil {
		return nil
	}
	return o.NumOfDays
}

type GetDashboardApisecRiskFindingsTrendResponse struct {
	// OK
	APISecRiskFindingsTrendWidget *shared.APISecRiskFindingsTrendWidget
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDashboardApisecRiskFindingsTrendResponse) GetAPISecRiskFindingsTrendWidget() *shared.APISecRiskFindingsTrendWidget {
	if o == nil {
		return nil
	}
	return o.APISecRiskFindingsTrendWidget
}

func (o *GetDashboardApisecRiskFindingsTrendResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDashboardApisecRiskFindingsTrendResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDashboardApisecRiskFindingsTrendResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
