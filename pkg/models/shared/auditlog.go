// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

// AuditLog - Single telemetry entry
type AuditLog struct {
	Action     *string    `json:"action,omitempty"`
	Module     *string    `json:"module,omitempty"`
	ObjectName *string    `json:"objectName,omitempty"`
	ObjectType *string    `json:"objectType,omitempty"`
	Time       *time.Time `json:"time,omitempty"`
	User       *string    `json:"user,omitempty"`
}

func (a AuditLog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditLog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuditLog) GetAction() *string {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *AuditLog) GetModule() *string {
	if o == nil {
		return nil
	}
	return o.Module
}

func (o *AuditLog) GetObjectName() *string {
	if o == nil {
		return nil
	}
	return o.ObjectName
}

func (o *AuditLog) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *AuditLog) GetTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *AuditLog) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
