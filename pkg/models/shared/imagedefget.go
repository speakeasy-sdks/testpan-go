// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

// ImageDefGet - Authorized image hash
type ImageDefGet struct {
	// dockerfile scan results summary by severity
	DockerfileScanResultsSummary *DockerfileScanResultsSummary `json:"dockerfileScanResultsSummary,omitempty"`
	ID                           *string                       `json:"id,omitempty"`
	// Valid hash for the image. * will authorize image name without validating hash
	ImageHash       *string          `json:"imageHash,omitempty"`
	ImageName       *string          `json:"imageName,omitempty"`
	ImageSourceType *ImageSourceType `default:"DOCKER_PLUGIN_CI" json:"imageSourceType"`
	ImageTags       []string         `json:"imageTags,omitempty"`
	// Specify if the image is identified
	IsIdentified *bool `json:"isIdentified,omitempty"`
	// Specify if the image has been scanned during CI
	IsScanned *bool      `json:"isScanned,omitempty"`
	Licenses  []string   `json:"licenses,omitempty"`
	TimeAdded *time.Time `json:"timeAdded,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (i ImageDefGet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImageDefGet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImageDefGet) GetDockerfileScanResultsSummary() *DockerfileScanResultsSummary {
	if o == nil {
		return nil
	}
	return o.DockerfileScanResultsSummary
}

func (o *ImageDefGet) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ImageDefGet) GetImageHash() *string {
	if o == nil {
		return nil
	}
	return o.ImageHash
}

func (o *ImageDefGet) GetImageName() *string {
	if o == nil {
		return nil
	}
	return o.ImageName
}

func (o *ImageDefGet) GetImageSourceType() *ImageSourceType {
	if o == nil {
		return nil
	}
	return o.ImageSourceType
}

func (o *ImageDefGet) GetImageTags() []string {
	if o == nil {
		return nil
	}
	return o.ImageTags
}

func (o *ImageDefGet) GetIsIdentified() *bool {
	if o == nil {
		return nil
	}
	return o.IsIdentified
}

func (o *ImageDefGet) GetIsScanned() *bool {
	if o == nil {
		return nil
	}
	return o.IsScanned
}

func (o *ImageDefGet) GetLicenses() []string {
	if o == nil {
		return nil
	}
	return o.Licenses
}

func (o *ImageDefGet) GetTimeAdded() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimeAdded
}

func (o *ImageDefGet) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}
