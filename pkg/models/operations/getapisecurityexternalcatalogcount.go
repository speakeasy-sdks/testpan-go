// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
	"time"
)

type GetAPISecurityExternalCatalogCountRequest struct {
	// When false, only services with specs wikk be returned
	IncludeServiceWithNoSpec *bool `default:"true" queryParam:"style=form,explode=true,name=includeServiceWithNoSpec"`
	// the Api Catalog name filter
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// Only Apis updated since this date
	UpdatedAfter *time.Time `queryParam:"style=form,explode=true,name=updatedAfter"`
}

func (g GetAPISecurityExternalCatalogCountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAPISecurityExternalCatalogCountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAPISecurityExternalCatalogCountRequest) GetIncludeServiceWithNoSpec() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeServiceWithNoSpec
}

func (o *GetAPISecurityExternalCatalogCountRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetAPISecurityExternalCatalogCountRequest) GetUpdatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAfter
}

type GetAPISecurityExternalCatalogCountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Number of APIs in the inventory
	Integer *int64
}

func (o *GetAPISecurityExternalCatalogCountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPISecurityExternalCatalogCountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPISecurityExternalCatalogCountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPISecurityExternalCatalogCountResponse) GetInteger() *int64 {
	if o == nil {
		return nil
	}
	return o.Integer
}
