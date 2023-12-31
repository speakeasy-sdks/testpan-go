// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir - sorting direction
type GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir string

const (
	GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDirAsc  GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir = "ASC"
	GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDirDesc GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir = "DESC"
)

func (e GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir) ToPointer() *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir {
	return &e
}

func (e *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir: %v", v)
	}
}

// GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey - sort key
type GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey string

const (
	GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKeyRisk GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey = "risk"
)

func (e GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey) ToPointer() *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey {
	return &e
}

func (e *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "risk":
		*e = GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey: %v", v)
	}
}

type GetRiskAssessmentPermissionsClusterIDOwnerIDRequest struct {
	ClusterID string `pathParam:"style=simple,explode=false,name=clusterId"`
	// Return approved / not approved entries
	IsApproved *bool  `default:"false" queryParam:"style=form,explode=true,name=isApproved"`
	OwnerID    string `pathParam:"style=simple,explode=false,name=ownerId"`
	// sorting direction
	SortDir *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey `default:"risk" queryParam:"style=form,explode=true,name=sortKey"`
}

func (g GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) GetIsApproved() *bool {
	if o == nil {
		return nil
	}
	return o.IsApproved
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) GetSortDir() *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDRequest) GetSortKey() *GetRiskAssessmentPermissionsClusterIDOwnerIDQueryParamSortKey {
	if o == nil {
		return nil
	}
	return o.SortKey
}

type GetRiskAssessmentPermissionsClusterIDOwnerIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	PermissionResponse *shared.PermissionResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDResponse) GetPermissionResponse() *shared.PermissionResponse {
	if o == nil {
		return nil
	}
	return o.PermissionResponse
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRiskAssessmentPermissionsClusterIDOwnerIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
