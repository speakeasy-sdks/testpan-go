// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ScanConfiguration - scan configuration information
type ScanConfiguration struct {
	// Number of available scanners in cluster
	NumberOfScanners *int64 `json:"numberOfScanners,omitempty"`
	// Cluster scan types
	ScanTypes []ScanType `json:"scanTypes,omitempty"`
}

func (o *ScanConfiguration) GetNumberOfScanners() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfScanners
}

func (o *ScanConfiguration) GetScanTypes() []ScanType {
	if o == nil {
		return nil
	}
	return o.ScanTypes
}
