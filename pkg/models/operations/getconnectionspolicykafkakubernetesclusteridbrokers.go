// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Strings []string
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
