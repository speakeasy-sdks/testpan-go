// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UsageAndPricingSection struct {
	Periods []UsageAndPricingPeriod `json:"periods,omitempty"`
}

func (o *UsageAndPricingSection) GetPeriods() []UsageAndPricingPeriod {
	if o == nil {
		return nil
	}
	return o.Periods
}