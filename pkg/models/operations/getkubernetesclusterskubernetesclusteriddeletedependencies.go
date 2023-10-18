// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetKubernetesClustersKubernetesClusterIDDeleteDependenciesRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse struct {
	// unknown error
	APIResponse *shared.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// Success
	KubernetesClusterDeleteDependencies *shared.KubernetesClusterDeleteDependencies
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse) GetKubernetesClusterDeleteDependencies() *shared.KubernetesClusterDeleteDependencies {
	if o == nil {
		return nil
	}
	return o.KubernetesClusterDeleteDependencies
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}