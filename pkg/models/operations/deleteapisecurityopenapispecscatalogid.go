// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type DeleteAPISecurityOpenAPISpecsCatalogIDRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type DeleteAPISecurityOpenAPISpecsCatalogIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	OpenAPISpecScoreStatus *shared.OpenAPISpecScoreStatus
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDResponse) GetOpenAPISpecScoreStatus() *shared.OpenAPISpecScoreStatus {
	if o == nil {
		return nil
	}
	return o.OpenAPISpecScoreStatus
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
