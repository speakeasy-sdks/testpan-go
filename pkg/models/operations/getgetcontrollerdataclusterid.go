// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetGetControllerDataClusterIDRequest struct {
	// Portshift cluster ID
	ClusterID string `pathParam:"style=simple,explode=false,name=clusterId"`
}

func (o *GetGetControllerDataClusterIDRequest) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

type GetGetControllerDataClusterIDResponse struct {
	// unknown error
	APIResponse *shared.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// Success
	ControllerDataResponse *shared.ControllerDataResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGetControllerDataClusterIDResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetGetControllerDataClusterIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGetControllerDataClusterIDResponse) GetControllerDataResponse() *shared.ControllerDataResponse {
	if o == nil {
		return nil
	}
	return o.ControllerDataResponse
}

func (o *GetGetControllerDataClusterIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGetControllerDataClusterIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}