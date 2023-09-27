// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetServerlessPolicyHistoryResponse struct {
	// Success
	ConnectionPolicyHistories []shared.ConnectionPolicyHistory
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetServerlessPolicyHistoryResponse) GetConnectionPolicyHistories() []shared.ConnectionPolicyHistory {
	if o == nil {
		return nil
	}
	return o.ConnectionPolicyHistories
}

func (o *GetServerlessPolicyHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerlessPolicyHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerlessPolicyHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
