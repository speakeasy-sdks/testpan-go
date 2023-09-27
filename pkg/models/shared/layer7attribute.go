// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Layer7Attribute - Collection of layer 7 attributes
type Layer7Attribute struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func (o *Layer7Attribute) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *Layer7Attribute) GetValues() []string {
	if o == nil {
		return []string{}
	}
	return o.Values
}
