// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetKubernetesAPIPolicyHistoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	KubernetesAPIPolicyHistories []shared.KubernetesAPIPolicyHistory
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetKubernetesAPIPolicyHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKubernetesAPIPolicyHistoryResponse) GetKubernetesAPIPolicyHistories() []shared.KubernetesAPIPolicyHistory {
	if o == nil {
		return nil
	}
	return o.KubernetesAPIPolicyHistories
}

func (o *GetKubernetesAPIPolicyHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKubernetesAPIPolicyHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
