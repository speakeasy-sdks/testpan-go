// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TrustedSignerKey struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (o *TrustedSignerKey) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *TrustedSignerKey) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
