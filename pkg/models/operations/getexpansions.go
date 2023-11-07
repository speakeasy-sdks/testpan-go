// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetExpansionsQueryParamSortDir - sorting direction
type GetExpansionsQueryParamSortDir string

const (
	GetExpansionsQueryParamSortDirAsc  GetExpansionsQueryParamSortDir = "ASC"
	GetExpansionsQueryParamSortDirDesc GetExpansionsQueryParamSortDir = "DESC"
)

func (e GetExpansionsQueryParamSortDir) ToPointer() *GetExpansionsQueryParamSortDir {
	return &e
}

func (e *GetExpansionsQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetExpansionsQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetExpansionsQueryParamSortDir: %v", v)
	}
}

// GetExpansionsQueryParamSortKey - sort key
type GetExpansionsQueryParamSortKey string

const (
	GetExpansionsQueryParamSortKeyName GetExpansionsQueryParamSortKey = "name"
)

func (e GetExpansionsQueryParamSortKey) ToPointer() *GetExpansionsQueryParamSortKey {
	return &e
}

func (e *GetExpansionsQueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "name":
		*e = GetExpansionsQueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetExpansionsQueryParamSortKey: %v", v)
	}
}

type GetExpansionsRequest struct {
	// Filter expansions by cluster name
	ClusterName *string `queryParam:"style=form,explode=true,name=clusterName"`
	// Filter the clusters by controller status
	ControllerStatus *string `queryParam:"style=form,explode=true,name=controllerStatus"`
	// Filter the clusters by controller version
	ControllerVersion *string `queryParam:"style=form,explode=true,name=controllerVersion"`
	// When true, the API will return an xlsx file, and pagination will be ignored
	DownloadAsXlsx *bool `queryParam:"style=form,explode=true,name=downloadAsXlsx"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Filter expansions by name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// Filter expansions by namespace
	NamespaceName *string `queryParam:"style=form,explode=true,name=namespaceName"`
	// When true, the pagination params will be ignored
	NoPagination *bool `queryParam:"style=form,explode=true,name=noPagination"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// sorting direction
	SortDir *GetExpansionsQueryParamSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey GetExpansionsQueryParamSortKey `queryParam:"style=form,explode=true,name=sortKey"`
}

func (g GetExpansionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetExpansionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetExpansionsRequest) GetClusterName() *string {
	if o == nil {
		return nil
	}
	return o.ClusterName
}

func (o *GetExpansionsRequest) GetControllerStatus() *string {
	if o == nil {
		return nil
	}
	return o.ControllerStatus
}

func (o *GetExpansionsRequest) GetControllerVersion() *string {
	if o == nil {
		return nil
	}
	return o.ControllerVersion
}

func (o *GetExpansionsRequest) GetDownloadAsXlsx() *bool {
	if o == nil {
		return nil
	}
	return o.DownloadAsXlsx
}

func (o *GetExpansionsRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetExpansionsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetExpansionsRequest) GetNamespaceName() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceName
}

func (o *GetExpansionsRequest) GetNoPagination() *bool {
	if o == nil {
		return nil
	}
	return o.NoPagination
}

func (o *GetExpansionsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetExpansionsRequest) GetSortDir() *GetExpansionsQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetExpansionsRequest) GetSortKey() GetExpansionsQueryParamSortKey {
	if o == nil {
		return GetExpansionsQueryParamSortKey("")
	}
	return o.SortKey
}

type GetExpansionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.Expansion
}

func (o *GetExpansionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetExpansionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetExpansionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetExpansionsResponse) GetClasses() []shared.Expansion {
	if o == nil {
		return nil
	}
	return o.Classes
}
