// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetRiskAssessmentPermissionsPermissionRisk - the risk to filter by
type GetRiskAssessmentPermissionsPermissionRisk string

const (
	GetRiskAssessmentPermissionsPermissionRiskNoRisk   GetRiskAssessmentPermissionsPermissionRisk = "NO_RISK"
	GetRiskAssessmentPermissionsPermissionRiskMedium   GetRiskAssessmentPermissionsPermissionRisk = "MEDIUM"
	GetRiskAssessmentPermissionsPermissionRiskHigh     GetRiskAssessmentPermissionsPermissionRisk = "HIGH"
	GetRiskAssessmentPermissionsPermissionRiskApproved GetRiskAssessmentPermissionsPermissionRisk = "APPROVED"
)

func (e GetRiskAssessmentPermissionsPermissionRisk) ToPointer() *GetRiskAssessmentPermissionsPermissionRisk {
	return &e
}

func (e *GetRiskAssessmentPermissionsPermissionRisk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_RISK":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		fallthrough
	case "APPROVED":
		*e = GetRiskAssessmentPermissionsPermissionRisk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRiskAssessmentPermissionsPermissionRisk: %v", v)
	}
}

// GetRiskAssessmentPermissionsSortDir - sorting direction
type GetRiskAssessmentPermissionsSortDir string

const (
	GetRiskAssessmentPermissionsSortDirAsc  GetRiskAssessmentPermissionsSortDir = "ASC"
	GetRiskAssessmentPermissionsSortDirDesc GetRiskAssessmentPermissionsSortDir = "DESC"
)

func (e GetRiskAssessmentPermissionsSortDir) ToPointer() *GetRiskAssessmentPermissionsSortDir {
	return &e
}

func (e *GetRiskAssessmentPermissionsSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetRiskAssessmentPermissionsSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRiskAssessmentPermissionsSortDir: %v", v)
	}
}

// GetRiskAssessmentPermissionsSortKey - sort key
type GetRiskAssessmentPermissionsSortKey string

const (
	GetRiskAssessmentPermissionsSortKeyPermissionRisk GetRiskAssessmentPermissionsSortKey = "permissionRisk"
)

func (e GetRiskAssessmentPermissionsSortKey) ToPointer() *GetRiskAssessmentPermissionsSortKey {
	return &e
}

func (e *GetRiskAssessmentPermissionsSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "permissionRisk":
		*e = GetRiskAssessmentPermissionsSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRiskAssessmentPermissionsSortKey: %v", v)
	}
}

type GetRiskAssessmentPermissionsRequest struct {
	// the clusters ids to filter by
	ClustersIds []string `queryParam:"style=form,explode=true,name=clustersIds"`
	// include systems default owners
	IncludeSystemOwners *bool `default:"false" queryParam:"style=form,explode=true,name=includeSystemOwners"`
	// the risk to filter by
	PermissionRisk *GetRiskAssessmentPermissionsPermissionRisk `queryParam:"style=form,explode=true,name=permissionRisk"`
	// sorting direction
	SortDir *GetRiskAssessmentPermissionsSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey *GetRiskAssessmentPermissionsSortKey `default:"permissionRisk" queryParam:"style=form,explode=true,name=sortKey"`
}

func (g GetRiskAssessmentPermissionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRiskAssessmentPermissionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRiskAssessmentPermissionsRequest) GetClustersIds() []string {
	if o == nil {
		return nil
	}
	return o.ClustersIds
}

func (o *GetRiskAssessmentPermissionsRequest) GetIncludeSystemOwners() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeSystemOwners
}

func (o *GetRiskAssessmentPermissionsRequest) GetPermissionRisk() *GetRiskAssessmentPermissionsPermissionRisk {
	if o == nil {
		return nil
	}
	return o.PermissionRisk
}

func (o *GetRiskAssessmentPermissionsRequest) GetSortDir() *GetRiskAssessmentPermissionsSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetRiskAssessmentPermissionsRequest) GetSortKey() *GetRiskAssessmentPermissionsSortKey {
	if o == nil {
		return nil
	}
	return o.SortKey
}

type GetRiskAssessmentPermissionsResponse struct {
	// Success
	ClusterPermissions []shared.ClusterPermission
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRiskAssessmentPermissionsResponse) GetClusterPermissions() []shared.ClusterPermission {
	if o == nil {
		return nil
	}
	return o.ClusterPermissions
}

func (o *GetRiskAssessmentPermissionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRiskAssessmentPermissionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRiskAssessmentPermissionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}