// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type GetGatewaysClustersGatewayType string

const (
	GetGatewaysClustersGatewayTypeApigeeX      GetGatewaysClustersGatewayType = "APIGEE_X"
	GetGatewaysClustersGatewayTypeKongInternal GetGatewaysClustersGatewayType = "KONG_INTERNAL"
	GetGatewaysClustersGatewayTypeTykInternal  GetGatewaysClustersGatewayType = "TYK_INTERNAL"
	GetGatewaysClustersGatewayTypeF5BigIP      GetGatewaysClustersGatewayType = "F5_BIG_IP"
)

func (e GetGatewaysClustersGatewayType) ToPointer() *GetGatewaysClustersGatewayType {
	return &e
}

func (e *GetGatewaysClustersGatewayType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APIGEE_X":
		fallthrough
	case "KONG_INTERNAL":
		fallthrough
	case "TYK_INTERNAL":
		fallthrough
	case "F5_BIG_IP":
		*e = GetGatewaysClustersGatewayType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGatewaysClustersGatewayType: %v", v)
	}
}

type GetGatewaysClustersRequest struct {
	GatewayType GetGatewaysClustersGatewayType `queryParam:"style=form,explode=true,name=gatewayType"`
}

func (o *GetGatewaysClustersRequest) GetGatewayType() GetGatewaysClustersGatewayType {
	if o == nil {
		return GetGatewaysClustersGatewayType("")
	}
	return o.GatewayType
}

type GetGatewaysClustersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	GatewayClusterInfos []shared.GatewayClusterInfo
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGatewaysClustersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGatewaysClustersResponse) GetGatewayClusterInfos() []shared.GatewayClusterInfo {
	if o == nil {
		return nil
	}
	return o.GatewayClusterInfos
}

func (o *GetGatewaysClustersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGatewaysClustersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
