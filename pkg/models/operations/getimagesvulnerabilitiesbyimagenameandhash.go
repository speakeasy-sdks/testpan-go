// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetImagesVulnerabilitiesByImageNameAndHashSortDir - sorting direction
type GetImagesVulnerabilitiesByImageNameAndHashSortDir string

const (
	GetImagesVulnerabilitiesByImageNameAndHashSortDirAsc  GetImagesVulnerabilitiesByImageNameAndHashSortDir = "ASC"
	GetImagesVulnerabilitiesByImageNameAndHashSortDirDesc GetImagesVulnerabilitiesByImageNameAndHashSortDir = "DESC"
)

func (e GetImagesVulnerabilitiesByImageNameAndHashSortDir) ToPointer() *GetImagesVulnerabilitiesByImageNameAndHashSortDir {
	return &e
}

func (e *GetImagesVulnerabilitiesByImageNameAndHashSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetImagesVulnerabilitiesByImageNameAndHashSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetImagesVulnerabilitiesByImageNameAndHashSortDir: %v", v)
	}
}

type GetImagesVulnerabilitiesByImageNameAndHashRequest struct {
	// the image sha256
	ImageHash string `queryParam:"style=form,explode=true,name=imageHash"`
	// the image name without tag
	ImageName string `queryParam:"style=form,explode=true,name=imageName"`
	// Return ignored / not ignored entries
	IsIgnored *bool   `default:"false" queryParam:"style=form,explode=true,name=isIgnored"`
	LayerID   *string `queryParam:"style=form,explode=true,name=layerId"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Return entries from this offset (pagination)
	Offset                         *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	ShowOnlyVulnerabilitiesWithFix *bool    `default:"false" queryParam:"style=form,explode=true,name=showOnlyVulnerabilitiesWithFix"`
	// sorting direction
	SortDir *GetImagesVulnerabilitiesByImageNameAndHashSortDir `default:"DESC" queryParam:"style=form,explode=true,name=sortDir"`
}

func (g GetImagesVulnerabilitiesByImageNameAndHashRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetImagesVulnerabilitiesByImageNameAndHashRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetImageHash() string {
	if o == nil {
		return ""
	}
	return o.ImageHash
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetImageName() string {
	if o == nil {
		return ""
	}
	return o.ImageName
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetIsIgnored() *bool {
	if o == nil {
		return nil
	}
	return o.IsIgnored
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetLayerID() *string {
	if o == nil {
		return nil
	}
	return o.LayerID
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetShowOnlyVulnerabilitiesWithFix() *bool {
	if o == nil {
		return nil
	}
	return o.ShowOnlyVulnerabilitiesWithFix
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashRequest) GetSortDir() *GetImagesVulnerabilitiesByImageNameAndHashSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

type GetImagesVulnerabilitiesByImageNameAndHashResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Vulnerabilities []shared.Vulnerability
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetImagesVulnerabilitiesByImageNameAndHashResponse) GetVulnerabilities() []shared.Vulnerability {
	if o == nil {
		return nil
	}
	return o.Vulnerabilities
}
