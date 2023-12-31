// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

// ViolationInfo - If the the App is running on an environment on which it is not allowed to run, this object contains the rule it violated.
type ViolationInfo struct {
	// App rule action
	Action                 *AppRuleAction        `json:"action,omitempty"`
	Comment                *string               `json:"comment,omitempty"`
	DefaultRule            *DefaultRule          `json:"defaultRule,omitempty"`
	LastViolation          *time.Time            `json:"lastViolation,omitempty"`
	MutateRule             *UserRule             `json:"mutateRule,omitempty"`
	PspViolationReasons    []string              `json:"pspViolationReasons,omitempty"`
	UnidentifiedPodReasons []string              `json:"unidentifiedPodReasons,omitempty"`
	UnidentifiedPodsRule   *UnidentifiedPodsRule `json:"unidentifiedPodsRule,omitempty"`
	UserRule               *UserRule             `json:"userRule,omitempty"`
	ViolationReasons       []ViolationReason     `json:"violationReasons,omitempty"`
}

func (v ViolationInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *ViolationInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ViolationInfo) GetAction() *AppRuleAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *ViolationInfo) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ViolationInfo) GetDefaultRule() *DefaultRule {
	if o == nil {
		return nil
	}
	return o.DefaultRule
}

func (o *ViolationInfo) GetLastViolation() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastViolation
}

func (o *ViolationInfo) GetMutateRule() *UserRule {
	if o == nil {
		return nil
	}
	return o.MutateRule
}

func (o *ViolationInfo) GetPspViolationReasons() []string {
	if o == nil {
		return nil
	}
	return o.PspViolationReasons
}

func (o *ViolationInfo) GetUnidentifiedPodReasons() []string {
	if o == nil {
		return nil
	}
	return o.UnidentifiedPodReasons
}

func (o *ViolationInfo) GetUnidentifiedPodsRule() *UnidentifiedPodsRule {
	if o == nil {
		return nil
	}
	return o.UnidentifiedPodsRule
}

func (o *ViolationInfo) GetUserRule() *UserRule {
	if o == nil {
		return nil
	}
	return o.UserRule
}

func (o *ViolationInfo) GetViolationReasons() []ViolationReason {
	if o == nil {
		return nil
	}
	return o.ViolationReasons
}
