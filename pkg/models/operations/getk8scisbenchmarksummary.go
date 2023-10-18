// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetK8sCISBenchmarkSummaryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	K8sCISBenchmarkAccountSummary *shared.K8sCISBenchmarkAccountSummary
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetK8sCISBenchmarkSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetK8sCISBenchmarkSummaryResponse) GetK8sCISBenchmarkAccountSummary() *shared.K8sCISBenchmarkAccountSummary {
	if o == nil {
		return nil
	}
	return o.K8sCISBenchmarkAccountSummary
}

func (o *GetK8sCISBenchmarkSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetK8sCISBenchmarkSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}