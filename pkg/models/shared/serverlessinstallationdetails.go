// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServerlessInstallationDetails struct {
	ExternalID      *string `json:"externalId,omitempty"`
	InstallationURL *string `json:"installationUrl,omitempty"`
}

func (o *ServerlessInstallationDetails) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *ServerlessInstallationDetails) GetInstallationURL() *string {
	if o == nil {
		return nil
	}
	return o.InstallationURL
}
