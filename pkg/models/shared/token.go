// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type TokenAttributeType string

const (
	TokenAttributeTypeRequestHeader TokenAttributeType = "REQUEST_HEADER"
	TokenAttributeTypeQueryParam    TokenAttributeType = "QUERY_PARAM"
)

func (e TokenAttributeType) ToPointer() *TokenAttributeType {
	return &e
}

func (e *TokenAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REQUEST_HEADER":
		fallthrough
	case "QUERY_PARAM":
		*e = TokenAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TokenAttributeType: %v", v)
	}
}

type Token struct {
	Apis            []string            `json:"apis,omitempty"`
	AttributeName   *string             `json:"attributeName,omitempty"`
	AttributeType   *TokenAttributeType `json:"attributeType,omitempty"`
	ExpirationDate  *time.Time          `json:"expirationDate,omitempty"`
	HTTPPath        *string             `json:"httpPath,omitempty"`
	ID              *string             `json:"id,omitempty"`
	Name            string              `json:"name"`
	VaultSecretPath string              `json:"vaultSecretPath"`
}

func (t Token) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Token) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Token) GetApis() []string {
	if o == nil {
		return nil
	}
	return o.Apis
}

func (o *Token) GetAttributeName() *string {
	if o == nil {
		return nil
	}
	return o.AttributeName
}

func (o *Token) GetAttributeType() *TokenAttributeType {
	if o == nil {
		return nil
	}
	return o.AttributeType
}

func (o *Token) GetExpirationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

func (o *Token) GetHTTPPath() *string {
	if o == nil {
		return nil
	}
	return o.HTTPPath
}

func (o *Token) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Token) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Token) GetVaultSecretPath() string {
	if o == nil {
		return ""
	}
	return o.VaultSecretPath
}
