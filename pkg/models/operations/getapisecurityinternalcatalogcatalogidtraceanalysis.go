// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

type GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
}

func (g GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TraceAnalysisDetails *shared.TraceAnalysisDetails
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse) GetTraceAnalysisDetails() *shared.TraceAnalysisDetails {
	if o == nil {
		return nil
	}
	return o.TraceAnalysisDetails
}
