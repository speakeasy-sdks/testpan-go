// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServerlessFunctionNames struct {
	Names []string `json:"names,omitempty"`
}

func (o *ServerlessFunctionNames) GetNames() []string {
	if o == nil {
		return nil
	}
	return o.Names
}
