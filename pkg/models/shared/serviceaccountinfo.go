// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServiceAccountInfo struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *ServiceAccountInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceAccountInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
