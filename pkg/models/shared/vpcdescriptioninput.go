// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// VPCDescriptionInput - Describes an AWS VPC.
type VPCDescriptionInput struct {
	AwsAccountID string `json:"awsAccountId"`
	// The "Name" tag of the VPC.
	//
	Name     *string `json:"name,omitempty"`
	RegionID string  `json:"regionId"`
	// AWS VPC ID
	VpcID string `json:"vpcId"`
}

func (o *VPCDescriptionInput) GetAwsAccountID() string {
	if o == nil {
		return ""
	}
	return o.AwsAccountID
}

func (o *VPCDescriptionInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *VPCDescriptionInput) GetRegionID() string {
	if o == nil {
		return ""
	}
	return o.RegionID
}

func (o *VPCDescriptionInput) GetVpcID() string {
	if o == nil {
		return ""
	}
	return o.VpcID
}
