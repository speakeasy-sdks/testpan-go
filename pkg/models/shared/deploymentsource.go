// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DeploymentSource string

const (
	DeploymentSourceHelm     DeploymentSource = "HELM"
	DeploymentSourceDeployer DeploymentSource = "DEPLOYER"
)

func (e DeploymentSource) ToPointer() *DeploymentSource {
	return &e
}

func (e *DeploymentSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HELM":
		fallthrough
	case "DEPLOYER":
		*e = DeploymentSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeploymentSource: %v", v)
	}
}
