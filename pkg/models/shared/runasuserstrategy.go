// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RunAsUserStrategy string

const (
	RunAsUserStrategyMustRunAs        RunAsUserStrategy = "MustRunAs"
	RunAsUserStrategyMustRunAsNonRoot RunAsUserStrategy = "MustRunAsNonRoot"
	RunAsUserStrategyRunAsAny         RunAsUserStrategy = "RunAsAny"
)

func (e RunAsUserStrategy) ToPointer() *RunAsUserStrategy {
	return &e
}

func (e *RunAsUserStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MustRunAs":
		fallthrough
	case "MustRunAsNonRoot":
		fallthrough
	case "RunAsAny":
		*e = RunAsUserStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunAsUserStrategy: %v", v)
	}
}
