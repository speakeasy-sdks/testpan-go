// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetNetworkMapQueueRequestIDRequest struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestId"`
}

func (o *GetNetworkMapQueueRequestIDRequest) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

type GetNetworkMapQueueRequestIDResponse struct {
	// Success
	BackgroundJobResponse *shared.BackgroundJobResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetNetworkMapQueueRequestIDResponse) GetBackgroundJobResponse() *shared.BackgroundJobResponse {
	if o == nil {
		return nil
	}
	return o.BackgroundJobResponse
}

func (o *GetNetworkMapQueueRequestIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNetworkMapQueueRequestIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNetworkMapQueueRequestIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
