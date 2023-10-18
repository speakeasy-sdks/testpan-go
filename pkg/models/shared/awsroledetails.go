// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AWSRoleDetails - A given name for the AWS role that Secure Application can connect to.
type AWSRoleDetails struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

func (o *AWSRoleDetails) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AWSRoleDetails) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}