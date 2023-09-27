// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAwsAwsAccountIDRegionsRequest struct {
	// AWS account Id
	AwsAccountID string `pathParam:"style=simple,explode=false,name=awsAccountId"`
}

func (o *GetAwsAwsAccountIDRegionsRequest) GetAwsAccountID() string {
	if o == nil {
		return ""
	}
	return o.AwsAccountID
}

type GetAwsAwsAccountIDRegionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	GetAwsAwsAccountIDRegions200ApplicationJSONStrings []string
}

func (o *GetAwsAwsAccountIDRegionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAwsAwsAccountIDRegionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAwsAwsAccountIDRegionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAwsAwsAccountIDRegionsResponse) GetGetAwsAwsAccountIDRegions200ApplicationJSONStrings() []string {
	if o == nil {
		return nil
	}
	return o.GetAwsAwsAccountIDRegions200ApplicationJSONStrings
}
