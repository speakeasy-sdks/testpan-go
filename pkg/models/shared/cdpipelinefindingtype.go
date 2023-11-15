// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CDPipelineFindingType string

const (
	CDPipelineFindingTypePermissions     CDPipelineFindingType = "PERMISSIONS"
	CDPipelineFindingTypeSecurityContext CDPipelineFindingType = "SECURITY_CONTEXT"
	CDPipelineFindingTypeSecrets         CDPipelineFindingType = "SECRETS"
)

func (e CDPipelineFindingType) ToPointer() *CDPipelineFindingType {
	return &e
}

func (e *CDPipelineFindingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PERMISSIONS":
		fallthrough
	case "SECURITY_CONTEXT":
		fallthrough
	case "SECRETS":
		*e = CDPipelineFindingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CDPipelineFindingType: %v", v)
	}
}
