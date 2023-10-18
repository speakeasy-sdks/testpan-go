// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SpecScoreElementFinding struct {
	AdditionalInfo []AdditionalInfo `json:"additional_info,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Mitigation     *string          `json:"mitigation,omitempty"`
	Name           *string          `json:"name,omitempty"`
	Occurrences    *int64           `json:"occurrences,omitempty"`
	Source         *string          `json:"source,omitempty"`
	SpecPath       *string          `json:"specPath,omitempty"`
}

func (o *SpecScoreElementFinding) GetAdditionalInfo() []AdditionalInfo {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *SpecScoreElementFinding) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SpecScoreElementFinding) GetMitigation() *string {
	if o == nil {
		return nil
	}
	return o.Mitigation
}

func (o *SpecScoreElementFinding) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SpecScoreElementFinding) GetOccurrences() *int64 {
	if o == nil {
		return nil
	}
	return o.Occurrences
}

func (o *SpecScoreElementFinding) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *SpecScoreElementFinding) GetSpecPath() *string {
	if o == nil {
		return nil
	}
	return o.SpecPath
}