// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AwsEnvironmentInput struct {
	ID   *string `json:"id,omitempty"`
	Tags []Tag   `json:"tags,omitempty"`
	// Like VPC but also includes the name
	Vpc VPCDescriptionInput `json:"vpc"`
}

func (o *AwsEnvironmentInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AwsEnvironmentInput) GetTags() []Tag {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AwsEnvironmentInput) GetVpc() VPCDescriptionInput {
	if o == nil {
		return VPCDescriptionInput{}
	}
	return o.Vpc
}