// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"io"
	"net/http"
)

type GetCloudAccountsCloudAccountIDDownloadBundleRequest struct {
	// cloud account ID
	CloudAccountID string `pathParam:"style=simple,explode=false,name=cloudAccountId"`
}

func (o *GetCloudAccountsCloudAccountIDDownloadBundleRequest) GetCloudAccountID() string {
	if o == nil {
		return ""
	}
	return o.CloudAccountID
}

type GetCloudAccountsCloudAccountIDDownloadBundleResponse struct {
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

func (o *GetCloudAccountsCloudAccountIDDownloadBundleResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetCloudAccountsCloudAccountIDDownloadBundleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCloudAccountsCloudAccountIDDownloadBundleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCloudAccountsCloudAccountIDDownloadBundleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCloudAccountsCloudAccountIDDownloadBundleResponse) GetStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Stream
}
