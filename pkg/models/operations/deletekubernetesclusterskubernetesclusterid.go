// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type DeleteKubernetesClustersKubernetesClusterIDRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *DeleteKubernetesClustersKubernetesClusterIDRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type DeleteKubernetesClustersKubernetesClusterIDResponse struct {
	// unknown error
	APIResponse *shared.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteKubernetesClustersKubernetesClusterIDResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *DeleteKubernetesClustersKubernetesClusterIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteKubernetesClustersKubernetesClusterIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteKubernetesClustersKubernetesClusterIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
