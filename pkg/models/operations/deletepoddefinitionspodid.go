// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeletePodDefinitionsPodIDRequest struct {
	PodID string `pathParam:"style=simple,explode=false,name=podId"`
}

func (o *DeletePodDefinitionsPodIDRequest) GetPodID() string {
	if o == nil {
		return ""
	}
	return o.PodID
}

type DeletePodDefinitionsPodIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeletePodDefinitionsPodIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePodDefinitionsPodIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePodDefinitionsPodIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
