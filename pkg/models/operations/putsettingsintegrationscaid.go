// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutSettingsIntegrationsCaIDRequest struct {
	CaIntegrationRequestInput shared.CaIntegrationRequestInput `request:"mediaType=application/json"`
	ID                        string                           `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutSettingsIntegrationsCaIDRequest) GetCaIntegrationRequestInput() shared.CaIntegrationRequestInput {
	if o == nil {
		return shared.CaIntegrationRequestInput{}
	}
	return o.CaIntegrationRequestInput
}

func (o *PutSettingsIntegrationsCaIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutSettingsIntegrationsCaIDResponse struct {
	// Success
	CaIntegrationResponse *shared.CaIntegrationResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSettingsIntegrationsCaIDResponse) GetCaIntegrationResponse() *shared.CaIntegrationResponse {
	if o == nil {
		return nil
	}
	return o.CaIntegrationResponse
}

func (o *PutSettingsIntegrationsCaIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSettingsIntegrationsCaIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSettingsIntegrationsCaIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}