// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteDependencyElement struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *DeleteDependencyElement) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteDependencyElement) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
