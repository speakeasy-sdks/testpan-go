// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type DeleteSettingsIntegrationsCaIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteSettingsIntegrationsCaIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteSettingsIntegrationsCaIDResponse struct {
	// unknown error
	APIResponse *shared.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSettingsIntegrationsCaIDResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *DeleteSettingsIntegrationsCaIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSettingsIntegrationsCaIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSettingsIntegrationsCaIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
