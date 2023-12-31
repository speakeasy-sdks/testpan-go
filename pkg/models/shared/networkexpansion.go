// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NetworkExpansion struct {
	Cluster     *string  `json:"cluster,omitempty"`
	Ids         []string `json:"ids,omitempty"`
	InstanceID  *string  `json:"instanceId,omitempty"`
	Labels      []Label  `json:"labels,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Namespace   *string  `json:"namespace,omitempty"`
	NamespaceID *string  `json:"namespaceId,omitempty"`
}

func (o *NetworkExpansion) GetCluster() *string {
	if o == nil {
		return nil
	}
	return o.Cluster
}

func (o *NetworkExpansion) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *NetworkExpansion) GetInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.InstanceID
}

func (o *NetworkExpansion) GetLabels() []Label {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *NetworkExpansion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NetworkExpansion) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *NetworkExpansion) GetNamespaceID() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceID
}
