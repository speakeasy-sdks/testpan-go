// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PortshiftAwsSubnet - the interface that represents the node's entity type. Possible types are: PortshiftAwsInstance, PortshiftAwsSecurityGroupViolationInfo, PortshiftAwsSubnet
type PortshiftAwsSubnet struct {
	CidrBlock             *string               `json:"cidrBlock,omitempty"`
	State                 *string               `json:"state,omitempty"`
	SubnetID              *string               `json:"subnetId,omitempty"`
	ViolationsMapNodeType ViolationsMapNodeType `json:"violationsMapNodeType"`
	VpcID                 *string               `json:"vpcId,omitempty"`
}

func (o *PortshiftAwsSubnet) GetCidrBlock() *string {
	if o == nil {
		return nil
	}
	return o.CidrBlock
}

func (o *PortshiftAwsSubnet) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *PortshiftAwsSubnet) GetSubnetID() *string {
	if o == nil {
		return nil
	}
	return o.SubnetID
}

func (o *PortshiftAwsSubnet) GetViolationsMapNodeType() ViolationsMapNodeType {
	if o == nil {
		return ViolationsMapNodeType("")
	}
	return o.ViolationsMapNodeType
}

func (o *PortshiftAwsSubnet) GetVpcID() *string {
	if o == nil {
		return nil
	}
	return o.VpcID
}
