// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type K8sCISBenchmarkClustersSummary struct {
	AgentActive          *bool                     `json:"agentActive,omitempty"`
	Compliance           *int                      `default:"0" json:"compliance"`
	ID                   *string                   `json:"id,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	NumberOfActionItems  *int                      `default:"0" json:"numberOfActionItems"`
	NumberOfInfoItems    *int                      `default:"0" json:"numberOfInfoItems"`
	NumberOfItemsFailed  *int                      `default:"0" json:"numberOfItemsFailed"`
	NumberOfItemsPassed  *int                      `default:"0" json:"numberOfItemsPassed"`
	NumberOfNodes        *int                      `default:"0" json:"numberOfNodes"`
	NumberOfScannedNodes *int                      `default:"0" json:"numberOfScannedNodes"`
	Orchestration        *string                   `json:"orchestration,omitempty"`
	ScanState            *K8sCisBenchmarkScanState `json:"scanState,omitempty"`
	ScanTime             *time.Time                `json:"scanTime,omitempty"`
}

func (k K8sCISBenchmarkClustersSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *K8sCISBenchmarkClustersSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *K8sCISBenchmarkClustersSummary) GetAgentActive() *bool {
	if o == nil {
		return nil
	}
	return o.AgentActive
}

func (o *K8sCISBenchmarkClustersSummary) GetCompliance() *int {
	if o == nil {
		return nil
	}
	return o.Compliance
}

func (o *K8sCISBenchmarkClustersSummary) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *K8sCISBenchmarkClustersSummary) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfActionItems() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfActionItems
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfInfoItems() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfInfoItems
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfItemsFailed() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfItemsFailed
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfItemsPassed() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfItemsPassed
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfNodes() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfNodes
}

func (o *K8sCISBenchmarkClustersSummary) GetNumberOfScannedNodes() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfScannedNodes
}

func (o *K8sCISBenchmarkClustersSummary) GetOrchestration() *string {
	if o == nil {
		return nil
	}
	return o.Orchestration
}

func (o *K8sCISBenchmarkClustersSummary) GetScanState() *K8sCisBenchmarkScanState {
	if o == nil {
		return nil
	}
	return o.ScanState
}

func (o *K8sCISBenchmarkClustersSummary) GetScanTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ScanTime
}