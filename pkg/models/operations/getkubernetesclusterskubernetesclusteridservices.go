// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

type GetKubernetesClustersKubernetesClusterIDServicesRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
	// if true, return only services deployed on namespace with label istio-injection=enabled
	ShowIstioOnly *bool `default:"false" queryParam:"style=form,explode=true,name=showIstioOnly"`
}

func (g GetKubernetesClustersKubernetesClusterIDServicesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetKubernetesClustersKubernetesClusterIDServicesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesRequest) GetShowIstioOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ShowIstioOnly
}

type GetKubernetesClustersKubernetesClusterIDServicesResponse struct {
	// unknown error
	APIResponse *sdkerrors.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// success
	Classes []shared.KubernetesService
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKubernetesClustersKubernetesClusterIDServicesResponse) GetClasses() []shared.KubernetesService {
	if o == nil {
		return nil
	}
	return o.Classes
}
