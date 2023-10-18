// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APISecRiskFindingsTrendWidget struct {
	Critical    []GraphNumberPoint `json:"critical,omitempty"`
	High        []GraphNumberPoint `json:"high,omitempty"`
	Low         []GraphNumberPoint `json:"low,omitempty"`
	Medium      []GraphNumberPoint `json:"medium,omitempty"`
	NoKnownRisk []GraphNumberPoint `json:"noKnownRisk,omitempty"`
}

func (o *APISecRiskFindingsTrendWidget) GetCritical() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.Critical
}

func (o *APISecRiskFindingsTrendWidget) GetHigh() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.High
}

func (o *APISecRiskFindingsTrendWidget) GetLow() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.Low
}

func (o *APISecRiskFindingsTrendWidget) GetMedium() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.Medium
}

func (o *APISecRiskFindingsTrendWidget) GetNoKnownRisk() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.NoKnownRisk
}