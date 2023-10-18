// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutSeccompProfilesProfileIDRequest struct {
	SeccompProfileInput shared.SeccompProfileInput `request:"mediaType=application/json"`
	ProfileID           string                     `pathParam:"style=simple,explode=false,name=profileId"`
}

func (o *PutSeccompProfilesProfileIDRequest) GetSeccompProfileInput() shared.SeccompProfileInput {
	if o == nil {
		return shared.SeccompProfileInput{}
	}
	return o.SeccompProfileInput
}

func (o *PutSeccompProfilesProfileIDRequest) GetProfileID() string {
	if o == nil {
		return ""
	}
	return o.ProfileID
}

type PutSeccompProfilesProfileIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	SeccompProfile *shared.SeccompProfile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSeccompProfilesProfileIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSeccompProfilesProfileIDResponse) GetSeccompProfile() *shared.SeccompProfile {
	if o == nil {
		return nil
	}
	return o.SeccompProfile
}

func (o *PutSeccompProfilesProfileIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSeccompProfilesProfileIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}