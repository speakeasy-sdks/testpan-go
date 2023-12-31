// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionProtocol string

const (
	ConnectionProtocolTCPPerformanceMetrics  ConnectionProtocol = "TcpPerformanceMetrics"
	ConnectionProtocolHTTPPerformanceMetrics ConnectionProtocol = "HttpPerformanceMetrics"
)

func (e ConnectionProtocol) ToPointer() *ConnectionProtocol {
	return &e
}

func (e *ConnectionProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TcpPerformanceMetrics":
		fallthrough
	case "HttpPerformanceMetrics":
		*e = ConnectionProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionProtocol: %v", v)
	}
}

type PerformanceMetrics struct {
	ConnectionProtocol *ConnectionProtocol `json:"connectionProtocol,omitempty"`
	// Return a list of total received bytes per connection
	TotalReceivedBytes []PerformanceMetricsGraphPoint `json:"totalReceivedBytes,omitempty"`
	// Return a list of total sent bytes per connection
	TotalSentBytes []PerformanceMetricsGraphPoint `json:"totalSentBytes,omitempty"`
}

func (o *PerformanceMetrics) GetConnectionProtocol() *ConnectionProtocol {
	if o == nil {
		return nil
	}
	return o.ConnectionProtocol
}

func (o *PerformanceMetrics) GetTotalReceivedBytes() []PerformanceMetricsGraphPoint {
	if o == nil {
		return nil
	}
	return o.TotalReceivedBytes
}

func (o *PerformanceMetrics) GetTotalSentBytes() []PerformanceMetricsGraphPoint {
	if o == nil {
		return nil
	}
	return o.TotalSentBytes
}
