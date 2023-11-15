// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ChallengeName string

const (
	ChallengeNameGoogleLogin ChallengeName = "GOOGLE_LOGIN"
)

func (e ChallengeName) ToPointer() *ChallengeName {
	return &e
}

func (e *ChallengeName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_LOGIN":
		*e = ChallengeName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChallengeName: %v", v)
	}
}

type ChallengeRequest struct {
	ChallengeName *ChallengeName `json:"ChallengeName,omitempty"`
}

func (o *ChallengeRequest) GetChallengeName() *ChallengeName {
	if o == nil {
		return nil
	}
	return o.ChallengeName
}
