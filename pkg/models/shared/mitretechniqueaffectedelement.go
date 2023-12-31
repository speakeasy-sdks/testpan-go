// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type MitreTechniqueAffectedElementType string

const (
	MitreTechniqueAffectedElementTypeMitreTechniqueAffectedOwner     MitreTechniqueAffectedElementType = "MitreTechniqueAffectedOwner"
	MitreTechniqueAffectedElementTypeMitreTechniqueAffectedWorkload  MitreTechniqueAffectedElementType = "MitreTechniqueAffectedWorkload"
	MitreTechniqueAffectedElementTypeMitreTechniqueAffectedNamespace MitreTechniqueAffectedElementType = "MitreTechniqueAffectedNamespace"
	MitreTechniqueAffectedElementTypeMitreTechniqueAffectedCluster   MitreTechniqueAffectedElementType = "MitreTechniqueAffectedCluster"
)

func (e MitreTechniqueAffectedElementType) ToPointer() *MitreTechniqueAffectedElementType {
	return &e
}

func (e *MitreTechniqueAffectedElementType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MitreTechniqueAffectedOwner":
		fallthrough
	case "MitreTechniqueAffectedWorkload":
		fallthrough
	case "MitreTechniqueAffectedNamespace":
		fallthrough
	case "MitreTechniqueAffectedCluster":
		*e = MitreTechniqueAffectedElementType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MitreTechniqueAffectedElementType: %v", v)
	}
}

type MitreTechniqueAffectedElement struct {
	MitreTechniqueAffectedElementType *MitreTechniqueAffectedElementType `json:"MitreTechniqueAffectedElementType,omitempty"`
}

func (o *MitreTechniqueAffectedElement) GetMitreTechniqueAffectedElementType() *MitreTechniqueAffectedElementType {
	if o == nil {
		return nil
	}
	return o.MitreTechniqueAffectedElementType
}
