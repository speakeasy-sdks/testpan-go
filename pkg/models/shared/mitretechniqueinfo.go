// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MitreTechniqueInfo struct {
	AffectedElements   []MitreTechniqueAffectedElement `json:"affectedElements,omitempty"`
	AffectedTechniques []string                        `json:"affectedTechniques,omitempty"`
	Explanation        *string                         `json:"explanation,omitempty"`
	FixDescription     *string                         `json:"fixDescription,omitempty"`
	HowToFix           *string                         `json:"howToFix,omitempty"`
	IsFixAvilable      *bool                           `json:"isFixAvilable,omitempty"`
}

func (o *MitreTechniqueInfo) GetAffectedElements() []MitreTechniqueAffectedElement {
	if o == nil {
		return nil
	}
	return o.AffectedElements
}

func (o *MitreTechniqueInfo) GetAffectedTechniques() []string {
	if o == nil {
		return nil
	}
	return o.AffectedTechniques
}

func (o *MitreTechniqueInfo) GetExplanation() *string {
	if o == nil {
		return nil
	}
	return o.Explanation
}

func (o *MitreTechniqueInfo) GetFixDescription() *string {
	if o == nil {
		return nil
	}
	return o.FixDescription
}

func (o *MitreTechniqueInfo) GetHowToFix() *string {
	if o == nil {
		return nil
	}
	return o.HowToFix
}

func (o *MitreTechniqueInfo) GetIsFixAvilable() *bool {
	if o == nil {
		return nil
	}
	return o.IsFixAvilable
}
