// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
	"time"
)

// GetAuditLogsQueryParamSortDir - sorting direction
type GetAuditLogsQueryParamSortDir string

const (
	GetAuditLogsQueryParamSortDirAsc  GetAuditLogsQueryParamSortDir = "ASC"
	GetAuditLogsQueryParamSortDirDesc GetAuditLogsQueryParamSortDir = "DESC"
)

func (e GetAuditLogsQueryParamSortDir) ToPointer() *GetAuditLogsQueryParamSortDir {
	return &e
}

func (e *GetAuditLogsQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetAuditLogsQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAuditLogsQueryParamSortDir: %v", v)
	}
}

// GetAuditLogsQueryParamSortKey - sort key
type GetAuditLogsQueryParamSortKey string

const (
	GetAuditLogsQueryParamSortKeyTime       GetAuditLogsQueryParamSortKey = "time"
	GetAuditLogsQueryParamSortKeyAction     GetAuditLogsQueryParamSortKey = "action"
	GetAuditLogsQueryParamSortKeyObjectType GetAuditLogsQueryParamSortKey = "objectType"
)

func (e GetAuditLogsQueryParamSortKey) ToPointer() *GetAuditLogsQueryParamSortKey {
	return &e
}

func (e *GetAuditLogsQueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "time":
		fallthrough
	case "action":
		fallthrough
	case "objectType":
		*e = GetAuditLogsQueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAuditLogsQueryParamSortKey: %v", v)
	}
}

type GetAuditLogsRequest struct {
	// Actions
	Actions []string `queryParam:"style=form,explode=false,name=actions"`
	// When true, the API will return an xlsx file, and pagination will be ignored
	DownloadAsXlsx *bool `queryParam:"style=form,explode=true,name=downloadAsXlsx"`
	// End date of the query
	EndTime time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Object Type
	ObjectType *string `queryParam:"style=form,explode=true,name=objectType"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// sorting direction
	SortDir *GetAuditLogsQueryParamSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey *GetAuditLogsQueryParamSortKey `default:"time" queryParam:"style=form,explode=true,name=sortKey"`
	// Start date of the query
	StartTime time.Time `queryParam:"style=form,explode=true,name=startTime"`
	// User name
	User *string `queryParam:"style=form,explode=true,name=user"`
}

func (g GetAuditLogsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAuditLogsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAuditLogsRequest) GetActions() []string {
	if o == nil {
		return nil
	}
	return o.Actions
}

func (o *GetAuditLogsRequest) GetDownloadAsXlsx() *bool {
	if o == nil {
		return nil
	}
	return o.DownloadAsXlsx
}

func (o *GetAuditLogsRequest) GetEndTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndTime
}

func (o *GetAuditLogsRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetAuditLogsRequest) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *GetAuditLogsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetAuditLogsRequest) GetSortDir() *GetAuditLogsQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetAuditLogsRequest) GetSortKey() *GetAuditLogsQueryParamSortKey {
	if o == nil {
		return nil
	}
	return o.SortKey
}

func (o *GetAuditLogsRequest) GetStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartTime
}

func (o *GetAuditLogsRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

type GetAuditLogsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.AuditLog
}

func (o *GetAuditLogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuditLogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuditLogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAuditLogsResponse) GetClasses() []shared.AuditLog {
	if o == nil {
		return nil
	}
	return o.Classes
}
