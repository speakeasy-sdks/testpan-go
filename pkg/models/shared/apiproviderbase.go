// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIProviderBase struct {
	// Unique id of the subject API as assigned by Crankshaft
	Identifier string  `json:"identifier"`
	Industry   *string `json:"industry,omitempty"`
	Location   *string `json:"location,omitempty"`
	// Name of the provider, typically an FQDN
	Name string `json:"name"`
}

func (o *APIProviderBase) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *APIProviderBase) GetIndustry() *string {
	if o == nil {
		return nil
	}
	return o.Industry
}

func (o *APIProviderBase) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *APIProviderBase) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
