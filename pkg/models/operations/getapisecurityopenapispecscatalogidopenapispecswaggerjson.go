// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON200ApplicationJSONString *string
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse) GetGetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON200ApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON200ApplicationJSONString
}