// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PodDefinitionInput struct {
	ClusterID      string           `json:"clusterId"`
	Containers     []Container      `json:"containers"`
	InitContainers []Container      `json:"initContainers,omitempty"`
	Kind           *PodTemplateKind `json:"kind,omitempty"`
	Labels         []Label          `json:"labels,omitempty"`
	// in pod template, this is the normalized name (for example, get it from pod -> replicaset -> deployment).
	//
	Name string `json:"name"`
	// The source type of the pod definition
	PodDefinitionSource *PodDefinitionSource `json:"podDefinitionSource,omitempty"`
}

func (o *PodDefinitionInput) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

func (o *PodDefinitionInput) GetContainers() []Container {
	if o == nil {
		return []Container{}
	}
	return o.Containers
}

func (o *PodDefinitionInput) GetInitContainers() []Container {
	if o == nil {
		return nil
	}
	return o.InitContainers
}

func (o *PodDefinitionInput) GetKind() *PodTemplateKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *PodDefinitionInput) GetLabels() []Label {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PodDefinitionInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PodDefinitionInput) GetPodDefinitionSource() *PodDefinitionSource {
	if o == nil {
		return nil
	}
	return o.PodDefinitionSource
}
