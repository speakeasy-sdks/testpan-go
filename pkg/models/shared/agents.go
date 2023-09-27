// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Agents struct {
	Agents             []Agent `json:"agents,omitempty"`
	LatestAgentVersion *string `json:"latestAgentVersion,omitempty"`
	LatestIstioVersion *string `json:"latestIstioVersion,omitempty"`
}

func (o *Agents) GetAgents() []Agent {
	if o == nil {
		return nil
	}
	return o.Agents
}

func (o *Agents) GetLatestAgentVersion() *string {
	if o == nil {
		return nil
	}
	return o.LatestAgentVersion
}

func (o *Agents) GetLatestIstioVersion() *string {
	if o == nil {
		return nil
	}
	return o.LatestIstioVersion
}
