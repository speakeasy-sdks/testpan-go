// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TrustedSignerCloudAccountInput struct {
	ID         *string                         `json:"id,omitempty"`
	Status     *TrustedSignerClusterStatus     `json:"status,omitempty"`
	Validation *TrustedSignerClusterValidation `json:"validation,omitempty"`
}

func (o *TrustedSignerCloudAccountInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrustedSignerCloudAccountInput) GetStatus() *TrustedSignerClusterStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TrustedSignerCloudAccountInput) GetValidation() *TrustedSignerClusterValidation {
	if o == nil {
		return nil
	}
	return o.Validation
}

type TrustedSignerCloudAccount struct {
	ID         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Status     *TrustedSignerClusterStatus     `json:"status,omitempty"`
	Validation *TrustedSignerClusterValidation `json:"validation,omitempty"`
}

func (o *TrustedSignerCloudAccount) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrustedSignerCloudAccount) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrustedSignerCloudAccount) GetStatus() *TrustedSignerClusterStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TrustedSignerCloudAccount) GetValidation() *TrustedSignerClusterValidation {
	if o == nil {
		return nil
	}
	return o.Validation
}