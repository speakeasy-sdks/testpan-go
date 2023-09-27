// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WorkloadRisk struct {
	Level   *WorkloadRiskLevel   `json:"level,omitempty"`
	Reasons []WorkloadRiskReason `json:"reasons,omitempty"`
}

func (o *WorkloadRisk) GetLevel() *WorkloadRiskLevel {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *WorkloadRisk) GetReasons() []WorkloadRiskReason {
	if o == nil {
		return nil
	}
	return o.Reasons
}
