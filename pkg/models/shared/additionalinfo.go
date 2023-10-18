// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AdditionalInfo struct {
	AffectedEndpoints []string          `json:"affected_endpoints,omitempty"`
	AffectedSpecPaths []string          `json:"affected_spec_paths,omitempty"`
	Entries           map[string]string `json:"entries,omitempty"`
}

func (o *AdditionalInfo) GetAffectedEndpoints() []string {
	if o == nil {
		return nil
	}
	return o.AffectedEndpoints
}

func (o *AdditionalInfo) GetAffectedSpecPaths() []string {
	if o == nil {
		return nil
	}
	return o.AffectedSpecPaths
}

func (o *AdditionalInfo) GetEntries() map[string]string {
	if o == nil {
		return nil
	}
	return o.Entries
}