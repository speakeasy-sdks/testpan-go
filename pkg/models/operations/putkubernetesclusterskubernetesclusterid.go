// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutKubernetesClustersKubernetesClusterIDRequest struct {
	KubernetesCluster shared.KubernetesCluster `request:"mediaType=application/json"`
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *PutKubernetesClustersKubernetesClusterIDRequest) GetKubernetesCluster() shared.KubernetesCluster {
	if o == nil {
		return shared.KubernetesCluster{}
	}
	return o.KubernetesCluster
}

func (o *PutKubernetesClustersKubernetesClusterIDRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type PutKubernetesClustersKubernetesClusterIDResponse struct {
	// unknown error
	APIResponse *sdkerrors.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// OK
	KubernetesCluster *shared.KubernetesCluster
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutKubernetesClustersKubernetesClusterIDResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *PutKubernetesClustersKubernetesClusterIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutKubernetesClustersKubernetesClusterIDResponse) GetKubernetesCluster() *shared.KubernetesCluster {
	if o == nil {
		return nil
	}
	return o.KubernetesCluster
}

func (o *PutKubernetesClustersKubernetesClusterIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutKubernetesClustersKubernetesClusterIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
