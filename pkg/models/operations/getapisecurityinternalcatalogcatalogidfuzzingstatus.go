// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetAPISecurityInternalCatalogCatalogIDFuzzingStatusRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingStatusRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse struct {
	// Success
	APIServiceFuzzingProgress *shared.APIServiceFuzzingProgress
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse) GetAPIServiceFuzzingProgress() *shared.APIServiceFuzzingProgress {
	if o == nil {
		return nil
	}
	return o.APIServiceFuzzingProgress
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
