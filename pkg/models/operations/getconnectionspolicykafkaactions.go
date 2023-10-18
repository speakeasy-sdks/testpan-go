// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetConnectionsPolicyKafkaActionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	GetConnectionsPolicyKafkaActions200ApplicationJSONStrings []string
}

func (o *GetConnectionsPolicyKafkaActionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionsPolicyKafkaActionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionsPolicyKafkaActionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConnectionsPolicyKafkaActionsResponse) GetGetConnectionsPolicyKafkaActions200ApplicationJSONStrings() []string {
	if o == nil {
		return nil
	}
	return o.GetConnectionsPolicyKafkaActions200ApplicationJSONStrings
}