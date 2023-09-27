// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutKubernetesAPIPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	KubernetesAPIPolicy *shared.KubernetesAPIPolicy
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutKubernetesAPIPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutKubernetesAPIPolicyResponse) GetKubernetesAPIPolicy() *shared.KubernetesAPIPolicy {
	if o == nil {
		return nil
	}
	return o.KubernetesAPIPolicy
}

func (o *PutKubernetesAPIPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutKubernetesAPIPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
