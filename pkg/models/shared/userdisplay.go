// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type UserDisplayStatus string

const (
	UserDisplayStatusEnabled  UserDisplayStatus = "ENABLED"
	UserDisplayStatusDisabled UserDisplayStatus = "DISABLED"
)

func (e UserDisplayStatus) ToPointer() *UserDisplayStatus {
	return &e
}

func (e *UserDisplayStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ENABLED":
		fallthrough
	case "DISABLED":
		*e = UserDisplayStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserDisplayStatus: %v", v)
	}
}

type UserDisplay struct {
	// The Secure Application account ID to which the user belongs
	AccountID   *string `json:"accountId,omitempty"`
	Description *string `json:"description,omitempty"`
	// The email of the user.
	Email    *string `json:"email,omitempty"`
	FullName string  `json:"fullName"`
	// ID of the user as created by Secure Application management.
	ID        *string    `json:"id,omitempty"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
	// The role of the user
	NormalizedRole *string `json:"normalizedRole,omitempty"`
	// The role of the user
	Role   *Role             `json:"role,omitempty"`
	Status UserDisplayStatus `json:"status"`
}

func (u UserDisplay) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserDisplay) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserDisplay) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UserDisplay) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UserDisplay) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UserDisplay) GetFullName() string {
	if o == nil {
		return ""
	}
	return o.FullName
}

func (o *UserDisplay) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UserDisplay) GetLastLogin() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLogin
}

func (o *UserDisplay) GetNormalizedRole() *string {
	if o == nil {
		return nil
	}
	return o.NormalizedRole
}

func (o *UserDisplay) GetRole() *Role {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *UserDisplay) GetStatus() UserDisplayStatus {
	if o == nil {
		return UserDisplayStatus("")
	}
	return o.Status
}
