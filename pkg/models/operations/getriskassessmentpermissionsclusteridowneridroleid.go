// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest struct {
	ClusterID string `pathParam:"style=simple,explode=false,name=clusterId"`
	OwnerID   string `pathParam:"style=simple,explode=false,name=ownerId"`
	RoleID    string `pathParam:"style=simple,explode=false,name=roleId"`
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	PermissionRoleResponse *shared.PermissionRoleResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse) GetPermissionRoleResponse() *shared.PermissionRoleResponse {
	if o == nil {
		return nil
	}
	return o.PermissionRoleResponse
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}