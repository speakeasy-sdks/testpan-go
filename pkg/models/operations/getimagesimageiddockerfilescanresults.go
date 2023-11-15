// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetImagesImageIDDockerfileScanResultsQueryParamSortDir - sorting direction
type GetImagesImageIDDockerfileScanResultsQueryParamSortDir string

const (
	GetImagesImageIDDockerfileScanResultsQueryParamSortDirAsc  GetImagesImageIDDockerfileScanResultsQueryParamSortDir = "ASC"
	GetImagesImageIDDockerfileScanResultsQueryParamSortDirDesc GetImagesImageIDDockerfileScanResultsQueryParamSortDir = "DESC"
)

func (e GetImagesImageIDDockerfileScanResultsQueryParamSortDir) ToPointer() *GetImagesImageIDDockerfileScanResultsQueryParamSortDir {
	return &e
}

func (e *GetImagesImageIDDockerfileScanResultsQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetImagesImageIDDockerfileScanResultsQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetImagesImageIDDockerfileScanResultsQueryParamSortDir: %v", v)
	}
}

type GetImagesImageIDDockerfileScanResultsRequest struct {
	ImageID string `pathParam:"style=simple,explode=false,name=imageId"`
	// Return ignored / not ignored entries
	IsIgnored *bool `default:"false" queryParam:"style=form,explode=true,name=isIgnored"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// sorting direction
	SortDir *GetImagesImageIDDockerfileScanResultsQueryParamSortDir `default:"DESC" queryParam:"style=form,explode=true,name=sortDir"`
}

func (g GetImagesImageIDDockerfileScanResultsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetImagesImageIDDockerfileScanResultsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetImagesImageIDDockerfileScanResultsRequest) GetImageID() string {
	if o == nil {
		return ""
	}
	return o.ImageID
}

func (o *GetImagesImageIDDockerfileScanResultsRequest) GetIsIgnored() *bool {
	if o == nil {
		return nil
	}
	return o.IsIgnored
}

func (o *GetImagesImageIDDockerfileScanResultsRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetImagesImageIDDockerfileScanResultsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetImagesImageIDDockerfileScanResultsRequest) GetSortDir() *GetImagesImageIDDockerfileScanResultsQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

type GetImagesImageIDDockerfileScanResultsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.DockerfileScanResult
}

func (o *GetImagesImageIDDockerfileScanResultsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetImagesImageIDDockerfileScanResultsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetImagesImageIDDockerfileScanResultsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetImagesImageIDDockerfileScanResultsResponse) GetClasses() []shared.DockerfileScanResult {
	if o == nil {
		return nil
	}
	return o.Classes
}
