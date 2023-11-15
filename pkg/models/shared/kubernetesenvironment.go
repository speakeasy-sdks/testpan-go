// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type KubernetesEnvironment struct {
	ID                       *string  `json:"id,omitempty"`
	KubernetesCluster        string   `json:"kubernetesCluster"`
	KubernetesClusterName    *string  `json:"kubernetesClusterName,omitempty"`
	NamespaceLabels          []Label  `json:"namespaceLabels,omitempty"`
	Namespaces               []string `json:"namespaces,omitempty"`
	PreserveOriginalSourceIP *bool    `json:"preserveOriginalSourceIp,omitempty"`
}

func (o *KubernetesEnvironment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KubernetesEnvironment) GetKubernetesCluster() string {
	if o == nil {
		return ""
	}
	return o.KubernetesCluster
}

func (o *KubernetesEnvironment) GetKubernetesClusterName() *string {
	if o == nil {
		return nil
	}
	return o.KubernetesClusterName
}

func (o *KubernetesEnvironment) GetNamespaceLabels() []Label {
	if o == nil {
		return nil
	}
	return o.NamespaceLabels
}

func (o *KubernetesEnvironment) GetNamespaces() []string {
	if o == nil {
		return nil
	}
	return o.Namespaces
}

func (o *KubernetesEnvironment) GetPreserveOriginalSourceIP() *bool {
	if o == nil {
		return nil
	}
	return o.PreserveOriginalSourceIP
}
