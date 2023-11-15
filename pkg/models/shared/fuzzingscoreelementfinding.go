// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type FuzzingScoreElementFinding struct {
	AdditionalInfo map[string]string `json:"additionalInfo,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Source         *string           `json:"source,omitempty"`
	SpecPath       *string           `json:"specPath,omitempty"`
}

func (o *FuzzingScoreElementFinding) GetAdditionalInfo() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *FuzzingScoreElementFinding) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *FuzzingScoreElementFinding) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FuzzingScoreElementFinding) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *FuzzingScoreElementFinding) GetSpecPath() *string {
	if o == nil {
		return nil
	}
	return o.SpecPath
}
