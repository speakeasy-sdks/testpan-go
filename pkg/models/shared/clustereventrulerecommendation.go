// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ClusterEventRuleRecommendation struct {
	ClusterEventRule       *KubernetesAPIRule `json:"clusterEventRule,omitempty"`
	ID                     *string            `json:"id,omitempty"`
	NumberOfAffectedEvents *int64             `json:"numberOfAffectedEvents,omitempty"`
}

func (o *ClusterEventRuleRecommendation) GetClusterEventRule() *KubernetesAPIRule {
	if o == nil {
		return nil
	}
	return o.ClusterEventRule
}

func (o *ClusterEventRuleRecommendation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ClusterEventRuleRecommendation) GetNumberOfAffectedEvents() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfAffectedEvents
}
