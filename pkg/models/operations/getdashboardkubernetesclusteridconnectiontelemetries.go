// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetDashboardKubernetesClusterIDConnectionTelemetriesRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *GetDashboardKubernetesClusterIDConnectionTelemetriesRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type GetDashboardKubernetesClusterIDConnectionTelemetriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TimeBasedWidget *shared.TimeBasedWidget
}

func (o *GetDashboardKubernetesClusterIDConnectionTelemetriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDashboardKubernetesClusterIDConnectionTelemetriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDashboardKubernetesClusterIDConnectionTelemetriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDashboardKubernetesClusterIDConnectionTelemetriesResponse) GetTimeBasedWidget() *shared.TimeBasedWidget {
	if o == nil {
		return nil
	}
	return o.TimeBasedWidget
}
