// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPISecurityPolicyPolicyIDRequest struct {
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
}

func (o *DeleteAPISecurityPolicyPolicyIDRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type DeleteAPISecurityPolicyPolicyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteAPISecurityPolicyPolicyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPISecurityPolicyPolicyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPISecurityPolicyPolicyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
