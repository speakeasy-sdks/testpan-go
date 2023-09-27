// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ClusterEventRuleRecommendationPeriod struct {
	ClusterEventRuleRecommendations []ClusterEventRuleRecommendation `json:"clusterEventRuleRecommendations,omitempty"`
	TimePeriod                      *RecommendationTimePeriod        `json:"timePeriod,omitempty"`
	TotalEvents                     *int64                           `json:"totalEvents,omitempty"`
}

func (o *ClusterEventRuleRecommendationPeriod) GetClusterEventRuleRecommendations() []ClusterEventRuleRecommendation {
	if o == nil {
		return nil
	}
	return o.ClusterEventRuleRecommendations
}

func (o *ClusterEventRuleRecommendationPeriod) GetTimePeriod() *RecommendationTimePeriod {
	if o == nil {
		return nil
	}
	return o.TimePeriod
}

func (o *ClusterEventRuleRecommendationPeriod) GetTotalEvents() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalEvents
}
