// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CategoryScoreGrade struct {
	AdditionalInfo []AdditionalInfo `json:"additional_info,omitempty"`
	// An enumeration.
	Confidence      *RiskConfidenceEnum `json:"confidence,omitempty"`
	CountersHistory *CountersHistory    `json:"counters_history,omitempty"`
	Critical        ScoreFindingGroup   `json:"critical"`
	High            ScoreFindingGroup   `json:"high"`
	Low             ScoreFindingGroup   `json:"low"`
	Medium          ScoreFindingGroup   `json:"medium"`
	Name            string              `json:"name"`
	// An `enum`eration.
	Risk          APISecurityRiskSeverity `json:"risk"`
	ScorerVersion int64                   `json:"scorer_version"`
	// An enumeration.
	Trend        *RiskTrendEnum    `json:"trend,omitempty"`
	Unclassified ScoreFindingGroup `json:"unclassified"`
}

func (o *CategoryScoreGrade) GetAdditionalInfo() []AdditionalInfo {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *CategoryScoreGrade) GetConfidence() *RiskConfidenceEnum {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *CategoryScoreGrade) GetCountersHistory() *CountersHistory {
	if o == nil {
		return nil
	}
	return o.CountersHistory
}

func (o *CategoryScoreGrade) GetCritical() ScoreFindingGroup {
	if o == nil {
		return ScoreFindingGroup{}
	}
	return o.Critical
}

func (o *CategoryScoreGrade) GetHigh() ScoreFindingGroup {
	if o == nil {
		return ScoreFindingGroup{}
	}
	return o.High
}

func (o *CategoryScoreGrade) GetLow() ScoreFindingGroup {
	if o == nil {
		return ScoreFindingGroup{}
	}
	return o.Low
}

func (o *CategoryScoreGrade) GetMedium() ScoreFindingGroup {
	if o == nil {
		return ScoreFindingGroup{}
	}
	return o.Medium
}

func (o *CategoryScoreGrade) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CategoryScoreGrade) GetRisk() APISecurityRiskSeverity {
	if o == nil {
		return APISecurityRiskSeverity("")
	}
	return o.Risk
}

func (o *CategoryScoreGrade) GetScorerVersion() int64 {
	if o == nil {
		return 0
	}
	return o.ScorerVersion
}

func (o *CategoryScoreGrade) GetTrend() *RiskTrendEnum {
	if o == nil {
		return nil
	}
	return o.Trend
}

func (o *CategoryScoreGrade) GetUnclassified() ScoreFindingGroup {
	if o == nil {
		return ScoreFindingGroup{}
	}
	return o.Unclassified
}
