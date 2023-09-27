// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetRegistriesSortDir - sorting direction
type GetRegistriesSortDir string

const (
	GetRegistriesSortDirAsc  GetRegistriesSortDir = "ASC"
	GetRegistriesSortDirDesc GetRegistriesSortDir = "DESC"
)

func (e GetRegistriesSortDir) ToPointer() *GetRegistriesSortDir {
	return &e
}

func (e *GetRegistriesSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetRegistriesSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRegistriesSortDir: %v", v)
	}
}

type GetRegistriesRequest struct {
	// sorting direction
	SortDir *GetRegistriesSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	sortKey *string `const:"url" queryParam:"style=form,explode=true,name=sortKey"`
}

func (g GetRegistriesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRegistriesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRegistriesRequest) GetSortDir() *GetRegistriesSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetRegistriesRequest) GetSortKey() *string {
	return types.String("url")
}

type GetRegistriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Registries []shared.Registry
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRegistriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRegistriesResponse) GetRegistries() []shared.Registry {
	if o == nil {
		return nil
	}
	return o.Registries
}

func (o *GetRegistriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRegistriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
