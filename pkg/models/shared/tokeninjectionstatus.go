// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TokenInjectionStatus string

const (
	TokenInjectionStatusNoInjection TokenInjectionStatus = "NO_INJECTION"
	TokenInjectionStatusSuccess     TokenInjectionStatus = "SUCCESS"
	TokenInjectionStatusError       TokenInjectionStatus = "ERROR"
)

func (e TokenInjectionStatus) ToPointer() *TokenInjectionStatus {
	return &e
}

func (e *TokenInjectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_INJECTION":
		fallthrough
	case "SUCCESS":
		fallthrough
	case "ERROR":
		*e = TokenInjectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TokenInjectionStatus: %v", v)
	}
}
