// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PostTruncationWorkloadsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TruncationStatus *shared.TruncationStatus
}

func (o *PostTruncationWorkloadsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostTruncationWorkloadsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostTruncationWorkloadsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostTruncationWorkloadsResponse) GetTruncationStatus() *shared.TruncationStatus {
	if o == nil {
		return nil
	}
	return o.TruncationStatus
}
