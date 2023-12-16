// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AWSRole - Defines a role ARN that Secure Application can connect to.
type AWSRole struct {
	Arn         string  `json:"arn"`
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
}

func (o *AWSRole) GetArn() string {
	if o == nil {
		return ""
	}
	return o.Arn
}

func (o *AWSRole) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AWSRole) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AWSRole) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
