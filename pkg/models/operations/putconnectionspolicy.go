// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutConnectionsPolicyResponse struct {
	// Success
	ConnectionsPolicy *shared.ConnectionsPolicy
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutConnectionsPolicyResponse) GetConnectionsPolicy() *shared.ConnectionsPolicy {
	if o == nil {
		return nil
	}
	return o.ConnectionsPolicy
}

func (o *PutConnectionsPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutConnectionsPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutConnectionsPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
