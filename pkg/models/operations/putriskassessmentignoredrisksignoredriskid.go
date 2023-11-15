// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest struct {
	CiPolicy      shared.CiPolicyInput `request:"mediaType=application/json"`
	IgnoredRiskID string               `pathParam:"style=simple,explode=false,name=ignoredRiskId"`
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest) GetCiPolicy() shared.CiPolicyInput {
	if o == nil {
		return shared.CiPolicyInput{}
	}
	return o.CiPolicy
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest) GetIgnoredRiskID() string {
	if o == nil {
		return ""
	}
	return o.IgnoredRiskID
}

type PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	IgnoredRisk *shared.IgnoredRisk
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse) GetIgnoredRisk() *shared.IgnoredRisk {
	if o == nil {
		return nil
	}
	return o.IgnoredRisk
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
