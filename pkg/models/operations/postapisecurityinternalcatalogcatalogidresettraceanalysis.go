// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
