// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// APIResponse - An object that is return in all cases of failures.
type APIResponse struct {
	Message *string `json:"message,omitempty"`
}

func (o *APIResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
