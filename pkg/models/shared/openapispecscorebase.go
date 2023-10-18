// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OpenAPISpecScoreBaseOapIVersion string

const (
	OpenAPISpecScoreBaseOapIVersionOpenAPISpecScoreV2 OpenAPISpecScoreBaseOapIVersion = "OpenApiSpecScoreV2"
	OpenAPISpecScoreBaseOapIVersionOpenAPISpecScoreV3 OpenAPISpecScoreBaseOapIVersion = "OpenApiSpecScoreV3"
)

func (e OpenAPISpecScoreBaseOapIVersion) ToPointer() *OpenAPISpecScoreBaseOapIVersion {
	return &e
}

func (e *OpenAPISpecScoreBaseOapIVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OpenApiSpecScoreV2":
		fallthrough
	case "OpenApiSpecScoreV3":
		*e = OpenAPISpecScoreBaseOapIVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpenAPISpecScoreBaseOapIVersion: %v", v)
	}
}

type OpenAPISpecScoreBase struct {
	General     *OpenAPISpecScoreElementsList    `json:"general,omitempty"`
	OapIVersion *OpenAPISpecScoreBaseOapIVersion `json:"oapIVersion,omitempty"`
	Security    *OpenAPISpecScoreElementsList    `json:"security,omitempty"`
	Tags        *OpenAPISpecTags                 `json:"tags,omitempty"`
}

func (o *OpenAPISpecScoreBase) GetGeneral() *OpenAPISpecScoreElementsList {
	if o == nil {
		return nil
	}
	return o.General
}

func (o *OpenAPISpecScoreBase) GetOapIVersion() *OpenAPISpecScoreBaseOapIVersion {
	if o == nil {
		return nil
	}
	return o.OapIVersion
}

func (o *OpenAPISpecScoreBase) GetSecurity() *OpenAPISpecScoreElementsList {
	if o == nil {
		return nil
	}
	return o.Security
}

func (o *OpenAPISpecScoreBase) GetTags() *OpenAPISpecTags {
	if o == nil {
		return nil
	}
	return o.Tags
}