// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type KubernetesAuditLogAction string

const (
	KubernetesAuditLogActionCreate  KubernetesAuditLogAction = "CREATE"
	KubernetesAuditLogActionUpdate  KubernetesAuditLogAction = "UPDATE"
	KubernetesAuditLogActionDelete  KubernetesAuditLogAction = "DELETE"
	KubernetesAuditLogActionConnect KubernetesAuditLogAction = "CONNECT"
)

func (e KubernetesAuditLogAction) ToPointer() *KubernetesAuditLogAction {
	return &e
}

func (e *KubernetesAuditLogAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREATE":
		fallthrough
	case "UPDATE":
		fallthrough
	case "DELETE":
		fallthrough
	case "CONNECT":
		*e = KubernetesAuditLogAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KubernetesAuditLogAction: %v", v)
	}
}