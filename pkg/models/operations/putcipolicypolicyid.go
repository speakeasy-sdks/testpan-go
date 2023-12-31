// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutCiPolicyPolicyIDRequest struct {
	CiPolicy shared.CiPolicyInput `request:"mediaType=application/json"`
	PolicyID string               `pathParam:"style=simple,explode=false,name=policyId"`
}

func (o *PutCiPolicyPolicyIDRequest) GetCiPolicy() shared.CiPolicyInput {
	if o == nil {
		return shared.CiPolicyInput{}
	}
	return o.CiPolicy
}

func (o *PutCiPolicyPolicyIDRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type PutCiPolicyPolicyIDResponse struct {
	// Success
	CiPolicy *shared.CiPolicy
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutCiPolicyPolicyIDResponse) GetCiPolicy() *shared.CiPolicy {
	if o == nil {
		return nil
	}
	return o.CiPolicy
}

func (o *PutCiPolicyPolicyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCiPolicyPolicyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCiPolicyPolicyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
