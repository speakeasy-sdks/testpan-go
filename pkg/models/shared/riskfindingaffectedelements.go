// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RiskFindingAffectedElements struct {
	ExternalElements []RiskFindingAffectedElementObject `json:"externalElements,omitempty"`
	InternalElements []RiskFindingAffectedElementObject `json:"internalElements,omitempty"`
}

func (o *RiskFindingAffectedElements) GetExternalElements() []RiskFindingAffectedElementObject {
	if o == nil {
		return nil
	}
	return o.ExternalElements
}

func (o *RiskFindingAffectedElements) GetInternalElements() []RiskFindingAffectedElementObject {
	if o == nil {
		return nil
	}
	return o.InternalElements
}
