// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TestTeamsIntegrationRequest struct {
	URL string `json:"url"`
}

func (o *TestTeamsIntegrationRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}