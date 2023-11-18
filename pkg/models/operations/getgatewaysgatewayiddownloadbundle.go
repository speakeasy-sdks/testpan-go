// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"io"
	"net/http"
)

type GetGatewaysGatewayIDDownloadBundleRequest struct {
	GatewayID string `pathParam:"style=simple,explode=false,name=gatewayId"`
}

func (o *GetGatewaysGatewayIDDownloadBundleRequest) GetGatewayID() string {
	if o == nil {
		return ""
	}
	return o.GatewayID
}

type GetGatewaysGatewayIDDownloadBundleResponse struct {
	// unknown error
	APIResponse *sdkerrors.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Stream io.ReadCloser
}

func (o *GetGatewaysGatewayIDDownloadBundleResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetGatewaysGatewayIDDownloadBundleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGatewaysGatewayIDDownloadBundleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGatewaysGatewayIDDownloadBundleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGatewaysGatewayIDDownloadBundleResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
