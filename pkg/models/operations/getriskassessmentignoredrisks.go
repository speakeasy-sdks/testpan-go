// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetRiskAssessmentIgnoredRisksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	IgnoredRisks []shared.IgnoredRisk
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRiskAssessmentIgnoredRisksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRiskAssessmentIgnoredRisksResponse) GetIgnoredRisks() []shared.IgnoredRisk {
	if o == nil {
		return nil
	}
	return o.IgnoredRisks
}

func (o *GetRiskAssessmentIgnoredRisksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRiskAssessmentIgnoredRisksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
