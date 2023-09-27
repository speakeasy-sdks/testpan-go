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

// GetTokensSortDir - sorting direction
type GetTokensSortDir string

const (
	GetTokensSortDirAsc  GetTokensSortDir = "ASC"
	GetTokensSortDirDesc GetTokensSortDir = "DESC"
)

func (e GetTokensSortDir) ToPointer() *GetTokensSortDir {
	return &e
}

func (e *GetTokensSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetTokensSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTokensSortDir: %v", v)
	}
}

type GetTokensRequest struct {
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// When true, the pagination params will be ignored
	NoPagination *bool `queryParam:"style=form,explode=true,name=noPagination"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// sorting direction
	SortDir *GetTokensSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// the token sort key
	sortKey *string `const:"expirationDate" queryParam:"style=form,explode=true,name=sortKey"`
	// Defined token name
	TokenName *string `queryParam:"style=form,explode=true,name=tokenName"`
}

func (g GetTokensRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTokensRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTokensRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetTokensRequest) GetNoPagination() *bool {
	if o == nil {
		return nil
	}
	return o.NoPagination
}

func (o *GetTokensRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetTokensRequest) GetSortDir() *GetTokensSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetTokensRequest) GetSortKey() *string {
	return types.String("expirationDate")
}

func (o *GetTokensRequest) GetTokenName() *string {
	if o == nil {
		return nil
	}
	return o.TokenName
}

type GetTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Tokens []shared.Token
}

func (o *GetTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTokensResponse) GetTokens() []shared.Token {
	if o == nil {
		return nil
	}
	return o.Tokens
}
