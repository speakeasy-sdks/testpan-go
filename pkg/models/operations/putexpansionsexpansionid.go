// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutExpansionsExpansionIDRequest struct {
	ExpansionPut shared.ExpansionPut `request:"mediaType=application/json"`
	ExpansionID  string              `pathParam:"style=simple,explode=false,name=expansionId"`
}

func (o *PutExpansionsExpansionIDRequest) GetExpansionPut() shared.ExpansionPut {
	if o == nil {
		return shared.ExpansionPut{}
	}
	return o.ExpansionPut
}

func (o *PutExpansionsExpansionIDRequest) GetExpansionID() string {
	if o == nil {
		return ""
	}
	return o.ExpansionID
}

type PutExpansionsExpansionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Expansion *shared.Expansion
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutExpansionsExpansionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutExpansionsExpansionIDResponse) GetExpansion() *shared.Expansion {
	if o == nil {
		return nil
	}
	return o.Expansion
}

func (o *PutExpansionsExpansionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutExpansionsExpansionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
