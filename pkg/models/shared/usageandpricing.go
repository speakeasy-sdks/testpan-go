// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type UsageAndPricing struct {
	Date            *time.Time         `json:"date,omitempty"`
	PodsUsageGraph  []GraphNumberPoint `json:"podsUsageGraph,omitempty"`
	PricingDetails  *PricingDetails    `json:"pricingDetails,omitempty"`
	VcpusUsageGraph []GraphNumberPoint `json:"vcpusUsageGraph,omitempty"`
}

func (u UsageAndPricing) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UsageAndPricing) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UsageAndPricing) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *UsageAndPricing) GetPodsUsageGraph() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.PodsUsageGraph
}

func (o *UsageAndPricing) GetPricingDetails() *PricingDetails {
	if o == nil {
		return nil
	}
	return o.PricingDetails
}

func (o *UsageAndPricing) GetVcpusUsageGraph() []GraphNumberPoint {
	if o == nil {
		return nil
	}
	return o.VcpusUsageGraph
}
