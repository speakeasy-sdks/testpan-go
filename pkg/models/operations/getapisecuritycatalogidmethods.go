// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetAPISecurityCatalogIDMethodsRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
	// spec tags names
	Tags []string `queryParam:"style=form,explode=false,name=tags"`
}

func (o *GetAPISecurityCatalogIDMethodsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *GetAPISecurityCatalogIDMethodsRequest) GetTags() []string {
	if o == nil {
		return []string{}
	}
	return o.Tags
}

type GetAPISecurityCatalogIDMethodsResponse struct {
	// Success
	APIServiceMethods []shared.APIServiceMethod
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPISecurityCatalogIDMethodsResponse) GetAPIServiceMethods() []shared.APIServiceMethod {
	if o == nil {
		return nil
	}
	return o.APIServiceMethods
}

func (o *GetAPISecurityCatalogIDMethodsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityCatalogIDMethodsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityCatalogIDMethodsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
