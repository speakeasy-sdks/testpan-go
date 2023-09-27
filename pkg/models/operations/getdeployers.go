// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
)

// GetDeployersSortDir - sorting direction
type GetDeployersSortDir string

const (
	GetDeployersSortDirAsc  GetDeployersSortDir = "ASC"
	GetDeployersSortDirDesc GetDeployersSortDir = "DESC"
)

func (e GetDeployersSortDir) ToPointer() *GetDeployersSortDir {
	return &e
}

func (e *GetDeployersSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetDeployersSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployersSortDir: %v", v)
	}
}

// GetDeployersSortKey - sort key
type GetDeployersSortKey string

const (
	GetDeployersSortKeyDeployer GetDeployersSortKey = "deployer"
	GetDeployersSortKeyType     GetDeployersSortKey = "type"
)

func (e GetDeployersSortKey) ToPointer() *GetDeployersSortKey {
	return &e
}

func (e *GetDeployersSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deployer":
		fallthrough
	case "type":
		*e = GetDeployersSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDeployersSortKey: %v", v)
	}
}

type GetDeployersRequest struct {
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// Filter deployers by name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// Filter deployers by rule creation
	RuleCreation *bool `queryParam:"style=form,explode=true,name=ruleCreation"`
	// Filter deployers by security checks
	SecurityCheck *bool `queryParam:"style=form,explode=true,name=securityCheck"`
	// sorting direction
	SortDir *GetDeployersSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey GetDeployersSortKey `queryParam:"style=form,explode=true,name=sortKey"`
}

func (g GetDeployersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDeployersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetDeployersRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetDeployersRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetDeployersRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetDeployersRequest) GetRuleCreation() *bool {
	if o == nil {
		return nil
	}
	return o.RuleCreation
}

func (o *GetDeployersRequest) GetSecurityCheck() *bool {
	if o == nil {
		return nil
	}
	return o.SecurityCheck
}

func (o *GetDeployersRequest) GetSortDir() *GetDeployersSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetDeployersRequest) GetSortKey() GetDeployersSortKey {
	if o == nil {
		return GetDeployersSortKey("")
	}
	return o.SortKey
}

type GetDeployersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Deployers []shared.Deployer
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDeployersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeployersResponse) GetDeployers() []shared.Deployer {
	if o == nil {
		return nil
	}
	return o.Deployers
}

func (o *GetDeployersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeployersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
