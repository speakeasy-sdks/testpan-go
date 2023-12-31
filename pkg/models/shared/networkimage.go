// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NetworkImage struct {
	JfrogProperties            []Property             `json:"jfrogProperties,omitempty"`
	Repository                 *string                `json:"repository,omitempty"`
	Tag                        *string                `json:"tag,omitempty"`
	VulnerabilitySeverityLevel *VulnerabilitySeverity `json:"vulnerabilitySeverityLevel,omitempty"`
}

func (o *NetworkImage) GetJfrogProperties() []Property {
	if o == nil {
		return nil
	}
	return o.JfrogProperties
}

func (o *NetworkImage) GetRepository() *string {
	if o == nil {
		return nil
	}
	return o.Repository
}

func (o *NetworkImage) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

func (o *NetworkImage) GetVulnerabilitySeverityLevel() *VulnerabilitySeverity {
	if o == nil {
		return nil
	}
	return o.VulnerabilitySeverityLevel
}
