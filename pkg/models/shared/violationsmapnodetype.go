// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ViolationsMapNodeType string

const (
	ViolationsMapNodeTypeSecurityGroup ViolationsMapNodeType = "SECURITY_GROUP"
	ViolationsMapNodeTypeSubnet        ViolationsMapNodeType = "SUBNET"
	ViolationsMapNodeTypeInstance      ViolationsMapNodeType = "INSTANCE"
)

func (e ViolationsMapNodeType) ToPointer() *ViolationsMapNodeType {
	return &e
}

func (e *ViolationsMapNodeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SECURITY_GROUP":
		fallthrough
	case "SUBNET":
		fallthrough
	case "INSTANCE":
		*e = ViolationsMapNodeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ViolationsMapNodeType: %v", v)
	}
}
