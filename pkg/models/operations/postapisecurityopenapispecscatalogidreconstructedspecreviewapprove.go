// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveRequest struct {
	APIReconstructedSpec shared.APIReconstructedSpecInput `request:"mediaType=application/json"`
	CatalogID            string                           `pathParam:"style=simple,explode=false,name=catalogId"`
}

func (o *PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveRequest) GetAPIReconstructedSpec() shared.APIReconstructedSpecInput {
	if o == nil {
		return shared.APIReconstructedSpecInput{}
	}
	return o.APIReconstructedSpec
}

func (o *PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
