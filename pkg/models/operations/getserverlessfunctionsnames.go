// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetServerlessFunctionsNamesRequest struct {
	// Filter cloud accounts by name
	CloudAccountName *string `queryParam:"style=form,explode=true,name=cloudAccountName"`
	// Defined function name
	FuncName []string `queryParam:"style=form,explode=false,name=funcName"`
	// Filter cloud accounts by region
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetServerlessFunctionsNamesRequest) GetCloudAccountName() *string {
	if o == nil {
		return nil
	}
	return o.CloudAccountName
}

func (o *GetServerlessFunctionsNamesRequest) GetFuncName() []string {
	if o == nil {
		return nil
	}
	return o.FuncName
}

func (o *GetServerlessFunctionsNamesRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type GetServerlessFunctionsNamesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	ServerlessFunctionNames []shared.ServerlessFunctionNames
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetServerlessFunctionsNamesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerlessFunctionsNamesResponse) GetServerlessFunctionNames() []shared.ServerlessFunctionNames {
	if o == nil {
		return nil
	}
	return o.ServerlessFunctionNames
}

func (o *GetServerlessFunctionsNamesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerlessFunctionsNamesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}