// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// APIServiceType - An `enum`eration.
type APIServiceType string

const (
	APIServiceTypeInternal APIServiceType = "INTERNAL"
	APIServiceTypeExternal APIServiceType = "EXTERNAL"
)

func (e APIServiceType) ToPointer() *APIServiceType {
	return &e
}

func (e *APIServiceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = APIServiceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIServiceType: %v", v)
	}
}