// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIEndpoint struct {
	// API service this endpoint belongs to. Empty if still undetermined.
	APIID      *string                     `json:"api_id,omitempty"`
	Compliance *APIServiceComplianceSimple `json:"compliance,omitempty"`
	// IP v4/v6 address of the API endpoint
	Host string `json:"host"`
	// Hostname of the API endpoint if known
	Hostname *string `json:"hostname,omitempty"`
	// Unique id of the Endpoint
	Identifier string  `json:"identifier"`
	Location   *string `json:"location,omitempty"`
	// Port of the API endpoint
	Port int64 `json:"port"`
	// An enumeration.
	Proto IPProtoEnum `json:"proto"`
	// An enumeration.
	Scheme *URLSchemeEnum `json:"scheme,omitempty"`
}

func (o *APIEndpoint) GetAPIID() *string {
	if o == nil {
		return nil
	}
	return o.APIID
}

func (o *APIEndpoint) GetCompliance() *APIServiceComplianceSimple {
	if o == nil {
		return nil
	}
	return o.Compliance
}

func (o *APIEndpoint) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *APIEndpoint) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *APIEndpoint) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *APIEndpoint) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *APIEndpoint) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *APIEndpoint) GetProto() IPProtoEnum {
	if o == nil {
		return IPProtoEnum("")
	}
	return o.Proto
}

func (o *APIEndpoint) GetScheme() *URLSchemeEnum {
	if o == nil {
		return nil
	}
	return o.Scheme
}
