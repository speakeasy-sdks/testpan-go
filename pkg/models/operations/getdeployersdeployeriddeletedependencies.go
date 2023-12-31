// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetDeployersDeployerIDDeleteDependenciesRequest struct {
	DeployerID string `pathParam:"style=simple,explode=false,name=deployerId"`
}

func (o *GetDeployersDeployerIDDeleteDependenciesRequest) GetDeployerID() string {
	if o == nil {
		return ""
	}
	return o.DeployerID
}

type GetDeployersDeployerIDDeleteDependenciesResponse struct {
	// unknown error
	APIResponse *sdkerrors.APIResponse
	// HTTP response content type for this operation
	ContentType string
	// Success
	DeployerDeleteDependencies *shared.DeployerDeleteDependencies
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDeployersDeployerIDDeleteDependenciesResponse) GetAPIResponse() *sdkerrors.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *GetDeployersDeployerIDDeleteDependenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeployersDeployerIDDeleteDependenciesResponse) GetDeployerDeleteDependencies() *shared.DeployerDeleteDependencies {
	if o == nil {
		return nil
	}
	return o.DeployerDeleteDependencies
}

func (o *GetDeployersDeployerIDDeleteDependenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeployersDeployerIDDeleteDependenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
