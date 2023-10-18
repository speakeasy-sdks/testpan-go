// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"io"
	"net/http"
)

type GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
	// The time interval for sending telemetries
	SendTelemetriesIntervalSec *int64 `default:"30" queryParam:"style=form,explode=true,name=sendTelemetriesIntervalSec"`
}

func (g GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest) GetSendTelemetriesIntervalSec() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTelemetriesIntervalSec
}

type GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse struct {
	// unknown error
	APIResponse *shared.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	GetKubernetesClustersKubernetesClusterIDDownloadBundle200ApplicationJSONBinaryString io.ReadCloser
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse) GetGetKubernetesClustersKubernetesClusterIDDownloadBundle200ApplicationJSONBinaryString() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.GetKubernetesClustersKubernetesClusterIDDownloadBundle200ApplicationJSONBinaryString
}