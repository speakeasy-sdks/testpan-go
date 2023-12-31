// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OapIVersion string

const (
	OapIVersionOpenAPISpecScoreV2 OapIVersion = "OpenApiSpecScoreV2"
	OapIVersionOpenAPISpecScoreV3 OapIVersion = "OpenApiSpecScoreV3"
)

func (e OapIVersion) ToPointer() *OapIVersion {
	return &e
}

func (e *OapIVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OpenApiSpecScoreV2":
		fallthrough
	case "OpenApiSpecScoreV3":
		*e = OapIVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OapIVersion: %v", v)
	}
}

type OpenAPISpecScoreBase struct {
	General     *OpenAPISpecScoreElementsList `json:"general,omitempty"`
	OapIVersion *OapIVersion                  `json:"oapIVersion,omitempty"`
	Security    *OpenAPISpecScoreElementsList `json:"security,omitempty"`
	Tags        *OpenAPISpecTags              `json:"tags,omitempty"`
}

func (o *OpenAPISpecScoreBase) GetGeneral() *OpenAPISpecScoreElementsList {
	if o == nil {
		return nil
	}
	return o.General
}

func (o *OpenAPISpecScoreBase) GetOapIVersion() *OapIVersion {
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
