// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ScoreFindingGroup struct {
	Count    int64          `json:"count"`
	Findings []ScoreFinding `json:"findings"`
}

func (o *ScoreFindingGroup) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *ScoreFindingGroup) GetFindings() []ScoreFinding {
	if o == nil {
		return []ScoreFinding{}
	}
	return o.Findings
}
