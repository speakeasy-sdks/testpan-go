// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Property struct {
	Key    *string  `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

func (o *Property) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Property) GetValues() []string {
	if o == nil {
		return nil
	}
	return o.Values
}