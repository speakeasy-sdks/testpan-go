// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// APISecurity - Mode of the API security
type APISecurity string

const (
	APISecurityEnabled  APISecurity = "ENABLED"
	APISecurityDisabled APISecurity = "DISABLED"
)

func (e APISecurity) ToPointer() *APISecurity {
	return &e
}

func (e *APISecurity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ENABLED":
		fallthrough
	case "DISABLED":
		*e = APISecurity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APISecurity: %v", v)
	}
}
