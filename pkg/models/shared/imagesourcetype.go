// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ImageSourceType string

const (
	ImageSourceTypeJenkinsPluginCi ImageSourceType = "JENKINS_PLUGIN_CI"
	ImageSourceTypeDockerPluginCi  ImageSourceType = "DOCKER_PLUGIN_CI"
	ImageSourceTypeRiskAssessment  ImageSourceType = "RISK_ASSESSMENT"
	ImageSourceTypeJfrogXray       ImageSourceType = "JFROG_XRAY"
	ImageSourceTypeRuntime         ImageSourceType = "RUNTIME"
)

func (e ImageSourceType) ToPointer() *ImageSourceType {
	return &e
}

func (e *ImageSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JENKINS_PLUGIN_CI":
		fallthrough
	case "DOCKER_PLUGIN_CI":
		fallthrough
	case "RISK_ASSESSMENT":
		fallthrough
	case "JFROG_XRAY":
		fallthrough
	case "RUNTIME":
		*e = ImageSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ImageSourceType: %v", v)
	}
}
