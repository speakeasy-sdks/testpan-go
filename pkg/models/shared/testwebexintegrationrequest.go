// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TestWebexIntegrationRequest struct {
	URL string `json:"url"`
}

func (o *TestWebexIntegrationRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
