// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetAPISecurityInternalCatalogCatalogIDFuzzingTestsRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingTestsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.APIServiceFuzzingTest
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse) GetClasses() []shared.APIServiceFuzzingTest {
	if o == nil {
		return nil
	}
	return o.Classes
}
