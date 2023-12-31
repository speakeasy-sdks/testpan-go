// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NetworkProtocol string

const (
	NetworkProtocolTCP  NetworkProtocol = "TCP"
	NetworkProtocolHTTP NetworkProtocol = "HTTP"
)

func (e NetworkProtocol) ToPointer() *NetworkProtocol {
	return &e
}

func (e *NetworkProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TCP":
		fallthrough
	case "HTTP":
		*e = NetworkProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NetworkProtocol: %v", v)
	}
}
