// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutServerlessPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	ServerlessPolicy *shared.ServerlessPolicy
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutServerlessPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutServerlessPolicyResponse) GetServerlessPolicy() *shared.ServerlessPolicy {
	if o == nil {
		return nil
	}
	return o.ServerlessPolicy
}

func (o *PutServerlessPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutServerlessPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
