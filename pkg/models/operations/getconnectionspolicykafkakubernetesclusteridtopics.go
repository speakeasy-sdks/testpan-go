// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest struct {
	// Secure Application Kubernetes cluster ID
	KubernetesClusterID string `pathParam:"style=simple,explode=false,name=kubernetesClusterId"`
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest) GetKubernetesClusterID() string {
	if o == nil {
		return ""
	}
	return o.KubernetesClusterID
}

type GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Strings []string
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
