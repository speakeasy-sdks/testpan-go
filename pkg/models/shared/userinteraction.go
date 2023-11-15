// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UserInteraction string

const (
	UserInteractionNone     UserInteraction = "NONE"
	UserInteractionRequired UserInteraction = "REQUIRED"
)

func (e UserInteraction) ToPointer() *UserInteraction {
	return &e
}

func (e *UserInteraction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NONE":
		fallthrough
	case "REQUIRED":
		*e = UserInteraction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserInteraction: %v", v)
	}
}
