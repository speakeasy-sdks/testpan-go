// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetDashboardKubernetesClusterIDKubernetesAuditLogsRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *GetDashboardKubernetesClusterIDKubernetesAuditLogsRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TimeBasedWidget *shared.TimeBasedWidget
}

func (o *GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse) GetTimeBasedWidget() *shared.TimeBasedWidget {
	if o == nil {
		return nil
	}
	return o.TimeBasedWidget
}
