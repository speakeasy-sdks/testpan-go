// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DockerfileScanResultsSummary - dockerfile scan results summary by severity
type DockerfileScanResultsSummary struct {
	Fatal *int64 `json:"fatal,omitempty"`
	Info  *int64 `json:"info,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Warn  *int64 `json:"warn,omitempty"`
}

func (o *DockerfileScanResultsSummary) GetFatal() *int64 {
	if o == nil {
		return nil
	}
	return o.Fatal
}

func (o *DockerfileScanResultsSummary) GetInfo() *int64 {
	if o == nil {
		return nil
	}
	return o.Info
}

func (o *DockerfileScanResultsSummary) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *DockerfileScanResultsSummary) GetWarn() *int64 {
	if o == nil {
		return nil
	}
	return o.Warn
}
