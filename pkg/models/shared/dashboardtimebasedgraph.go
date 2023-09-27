// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DashboardTimeBasedGraph struct {
	// the graph points
	Graph []GraphNumberPoint           `json:"graph,omitempty"`
	Info  *DashboardTimeBasedGraphInfo `json:"info,omitempty"`
}

func (o *DashboardTimeBasedGraph) GetGraph() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.Graph
}

func (o *DashboardTimeBasedGraph) GetInfo() *DashboardTimeBasedGraphInfo {
	if o == nil {
		return nil
	}
	return o.Info
}
