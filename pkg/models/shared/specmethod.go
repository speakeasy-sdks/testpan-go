// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SpecMethod struct {
	Description *string              `json:"description,omitempty"`
	Diffs       *SpecScoreDiffsLists `json:"diffs,omitempty"`
	Findings    *SpecScoreFindings   `json:"findings,omitempty"`
	Method      *HTTPMethod          `json:"method,omitempty"`
	Path        *string              `json:"path,omitempty"`
	// An `enum`eration.
	Severity *APISecurityRiskSeverity `json:"severity,omitempty"`
	Tag      *string                  `json:"tag,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (o *SpecMethod) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SpecMethod) GetDiffs() *SpecScoreDiffsLists {
	if o == nil {
		return nil
	}
	return o.Diffs
}

func (o *SpecMethod) GetFindings() *SpecScoreFindings {
	if o == nil {
		return nil
	}
	return o.Findings
}

func (o *SpecMethod) GetMethod() *HTTPMethod {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *SpecMethod) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *SpecMethod) GetSeverity() *APISecurityRiskSeverity {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *SpecMethod) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

func (o *SpecMethod) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}
