// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OpenAPISpecScoreElementsList struct {
	Elements []OpenAPISpecScoreElement `json:"elements,omitempty"`
	// An `enum`eration.
	Severity *APISecurityRiskSeverity `json:"severity,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (o *OpenAPISpecScoreElementsList) GetElements() []OpenAPISpecScoreElement {
	if o == nil {
		return nil
	}
	return o.Elements
}

func (o *OpenAPISpecScoreElementsList) GetSeverity() *APISecurityRiskSeverity {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *OpenAPISpecScoreElementsList) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}
