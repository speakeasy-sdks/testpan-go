// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPISecurityInternalCatalogCatalogIDBflaResetRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *PostAPISecurityInternalCatalogCatalogIDBflaResetRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type PostAPISecurityInternalCatalogCatalogIDBflaResetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	PostAPISecurityInternalCatalogCatalogIDBflaReset201ApplicationJSONUUIDString *string
}

func (o *PostAPISecurityInternalCatalogCatalogIDBflaResetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPISecurityInternalCatalogCatalogIDBflaResetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPISecurityInternalCatalogCatalogIDBflaResetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostAPISecurityInternalCatalogCatalogIDBflaResetResponse) GetPostAPISecurityInternalCatalogCatalogIDBflaReset201ApplicationJSONUUIDString() *string {
	if o == nil {
		return nil
	}
	return o.PostAPISecurityInternalCatalogCatalogIDBflaReset201ApplicationJSONUUIDString
}