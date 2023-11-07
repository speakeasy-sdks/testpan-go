// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PostRiskAssessmentKubernetesClusterIDSettingsRequest struct {
	RiskAssessmentClusterScanConfig shared.RiskAssessmentClusterScanConfig `request:"mediaType=application/json"`
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsRequest) GetRiskAssessmentClusterScanConfig() shared.RiskAssessmentClusterScanConfig {
	if o == nil {
		return shared.RiskAssessmentClusterScanConfig{}
	}
	return o.RiskAssessmentClusterScanConfig
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type PostRiskAssessmentKubernetesClusterIDSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Scan with the returned scanId was added to execution queue.
	Res *string
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostRiskAssessmentKubernetesClusterIDSettingsResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
