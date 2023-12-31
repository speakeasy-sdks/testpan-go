// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AgentsUpdateSettingsInput struct {
	IsManualUpdate     *bool `json:"isManualUpdate,omitempty"`
	IsUpdateNowEnabled *bool `json:"isUpdateNowEnabled,omitempty"`
}

func (o *AgentsUpdateSettingsInput) GetIsManualUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.IsManualUpdate
}

func (o *AgentsUpdateSettingsInput) GetIsUpdateNowEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IsUpdateNowEnabled
}
