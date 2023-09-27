// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetAppsSortDir - sorting direction
type GetAppsSortDir string

const (
	GetAppsSortDirAsc  GetAppsSortDir = "ASC"
	GetAppsSortDirDesc GetAppsSortDir = "DESC"
)

func (e GetAppsSortDir) ToPointer() *GetAppsSortDir {
	return &e
}

func (e *GetAppsSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetAppsSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppsSortDir: %v", v)
	}
}

// GetAppsSortKey - App sort key
type GetAppsSortKey string

const (
	GetAppsSortKeyName GetAppsSortKey = "name"
	GetAppsSortKeyType GetAppsSortKey = "type"
)

func (e GetAppsSortKey) ToPointer() *GetAppsSortKey {
	return &e
}

func (e *GetAppsSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "name":
		fallthrough
	case "type":
		*e = GetAppsSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppsSortKey: %v", v)
	}
}

type GetAppsRequest struct {
	// When true, the API will return an xlsx file, and pagination will be ignored
	DownloadAsXlsx *bool `queryParam:"style=form,explode=true,name=downloadAsXlsx"`
	// Filter Apps by name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// When true, the pagination params will be ignored
	NoPagination *bool `queryParam:"style=form,explode=true,name=noPagination"`
	// sorting direction
	SortDir *GetAppsSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// App sort key
	SortKey *GetAppsSortKey `queryParam:"style=form,explode=true,name=sortKey"`
	// Filter Apps by type
	Type []string `queryParam:"style=form,explode=false,name=type"`
}

func (g GetAppsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAppsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAppsRequest) GetDownloadAsXlsx() *bool {
	if o == nil {
		return nil
	}
	return o.DownloadAsXlsx
}

func (o *GetAppsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetAppsRequest) GetNoPagination() *bool {
	if o == nil {
		return nil
	}
	return o.NoPagination
}

func (o *GetAppsRequest) GetSortDir() *GetAppsSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetAppsRequest) GetSortKey() *GetAppsSortKey {
	if o == nil {
		return nil
	}
	return o.SortKey
}

func (o *GetAppsRequest) GetType() []string {
	if o == nil {
		return nil
	}
	return o.Type
}

type GetAppsResponse struct {
	// Created
	Apps []shared.App
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAppsResponse) GetApps() []shared.App {
	if o == nil {
		return nil
	}
	return o.Apps
}

func (o *GetAppsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
