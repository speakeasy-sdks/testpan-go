// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AwsEnvironment struct {
	ID   *string `json:"id,omitempty"`
	Tags []Tag   `json:"tags,omitempty"`
	// Like VPC but also includes the name
	Vpc VPCDescription `json:"vpc"`
}

func (o *AwsEnvironment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AwsEnvironment) GetTags() []Tag {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AwsEnvironment) GetVpc() VPCDescription {
	if o == nil {
		return VPCDescription{}
	}
	return o.Vpc
}
