// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AgentInstance struct {
	// The id of the instance on which the agent is running
	ID *string `json:"id,omitempty"`
	// The name of the instance on which the agent is running
	Name *string `json:"name,omitempty"`
}

func (o *AgentInstance) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AgentInstance) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}