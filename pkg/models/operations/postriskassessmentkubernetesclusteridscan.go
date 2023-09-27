// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostRiskAssessmentKubernetesClusterIDScanRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *PostRiskAssessmentKubernetesClusterIDScanRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type PostRiskAssessmentKubernetesClusterIDScanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Scan with the returned scanId was added to execution queue.
	PostRiskAssessmentKubernetesClusterIDScan201ApplicationJSONUUIDString *string
}

func (o *PostRiskAssessmentKubernetesClusterIDScanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostRiskAssessmentKubernetesClusterIDScanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostRiskAssessmentKubernetesClusterIDScanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostRiskAssessmentKubernetesClusterIDScanResponse) GetPostRiskAssessmentKubernetesClusterIDScan201ApplicationJSONUUIDString() *string {
	if o == nil {
		return nil
	}
	return o.PostRiskAssessmentKubernetesClusterIDScan201ApplicationJSONUUIDString
}
