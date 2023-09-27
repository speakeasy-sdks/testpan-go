// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PostPodSecurityPolicyProfilesBatchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Added
	PodSecurityPolicies []shared.PodSecurityPolicy
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostPodSecurityPolicyProfilesBatchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostPodSecurityPolicyProfilesBatchResponse) GetPodSecurityPolicies() []shared.PodSecurityPolicy {
	if o == nil {
		return nil
	}
	return o.PodSecurityPolicies
}

func (o *PostPodSecurityPolicyProfilesBatchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostPodSecurityPolicyProfilesBatchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
