// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TrustedSignerCluster struct {
	ID         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Status     *TrustedSignerClusterStatus     `json:"status,omitempty"`
	Validation *TrustedSignerClusterValidation `json:"validation,omitempty"`
}

func (o *TrustedSignerCluster) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrustedSignerCluster) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrustedSignerCluster) GetStatus() *TrustedSignerClusterStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TrustedSignerCluster) GetValidation() *TrustedSignerClusterValidation {
	if o == nil {
		return nil
	}
	return o.Validation
}
