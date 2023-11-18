// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest struct {
	EditKubernetesClusterManagedByHelm shared.EditKubernetesClusterManagedByHelm `request:"mediaType=application/json"`
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest) GetEditKubernetesClusterManagedByHelm() shared.EditKubernetesClusterManagedByHelm {
	if o == nil {
		return shared.EditKubernetesClusterManagedByHelm{}
	}
	return o.EditKubernetesClusterManagedByHelm
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse struct {
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

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse) GetKubernetesCluster() *shared.KubernetesCluster {
	if o == nil {
		return nil
	}
	return o.KubernetesCluster
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
