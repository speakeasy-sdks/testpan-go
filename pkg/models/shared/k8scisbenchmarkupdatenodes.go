// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type K8sCISBenchmarkUpdateNodes struct {
	ClusterID string                          `json:"clusterId"`
	Index     string                          `json:"index"`
	Nodes     []K8sCISBenchmarkUpdateNode     `json:"nodes"`
	Status    K8sCISBenchmarkUpdateNodeStatus `json:"status"`
}

func (o *K8sCISBenchmarkUpdateNodes) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

func (o *K8sCISBenchmarkUpdateNodes) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *K8sCISBenchmarkUpdateNodes) GetNodes() []K8sCISBenchmarkUpdateNode {
	if o == nil {
		return []K8sCISBenchmarkUpdateNode{}
	}
	return o.Nodes
}

func (o *K8sCISBenchmarkUpdateNodes) GetStatus() K8sCISBenchmarkUpdateNodeStatus {
	if o == nil {
		return K8sCISBenchmarkUpdateNodeStatus("")
	}
	return o.Status
}
