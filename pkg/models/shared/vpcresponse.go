// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type VpcResponse struct {
	// AWS VPC ID
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *VpcResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *VpcResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
