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

// QueryParamSortDir - sorting direction
type QueryParamSortDir string

const (
	QueryParamSortDirAsc  QueryParamSortDir = "ASC"
	QueryParamSortDirDesc QueryParamSortDir = "DESC"
)

func (e QueryParamSortDir) ToPointer() *QueryParamSortDir {
	return &e
}

func (e *QueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = QueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSortDir: %v", v)
	}
}

// QueryParamSortKey - sort key
type QueryParamSortKey string

const (
	QueryParamSortKeyTime   QueryParamSortKey = "time"
	QueryParamSortKeyStatus QueryParamSortKey = "status"
)

func (e QueryParamSortKey) ToPointer() *QueryParamSortKey {
	return &e
}

func (e *QueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "time":
		fallthrough
	case "status":
		*e = QueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSortKey: %v", v)
	}
}

type GetCdRequest struct {
	// End date of the query
	EndTime time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// Resource name
	ResourceName *string `queryParam:"style=form,explode=true,name=resourceName"`
	// sorting direction
	SortDir *QueryParamSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey *QueryParamSortKey `default:"time" queryParam:"style=form,explode=true,name=sortKey"`
	// Start date of the query
	StartTime time.Time `queryParam:"style=form,explode=true,name=startTime"`
}

func (g GetCdRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCdRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetCdRequest) GetEndTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndTime
}

func (o *GetCdRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetCdRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetCdRequest) GetResourceName() *string {
	if o == nil {
		return nil
	}
	return o.ResourceName
}

func (o *GetCdRequest) GetSortDir() *QueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetCdRequest) GetSortKey() *QueryParamSortKey {
	if o == nil {
		return nil
	}
	return o.SortKey
}

func (o *GetCdRequest) GetStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartTime
}

type GetCdResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.CDPipelineResult
}

func (o *GetCdResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCdResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCdResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCdResponse) GetClasses() []shared.CDPipelineResult {
	if o == nil {
		return nil
	}
	return o.Classes
}
