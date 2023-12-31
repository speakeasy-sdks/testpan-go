// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UsageAndPricingPeriodEnum string

const (
	UsageAndPricingPeriodEnumYear    UsageAndPricingPeriodEnum = "YEAR"
	UsageAndPricingPeriodEnumMonths  UsageAndPricingPeriodEnum = "MONTHS"
	UsageAndPricingPeriodEnumLastDay UsageAndPricingPeriodEnum = "LAST_DAY"
)

func (e UsageAndPricingPeriodEnum) ToPointer() *UsageAndPricingPeriodEnum {
	return &e
}

func (e *UsageAndPricingPeriodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "YEAR":
		fallthrough
	case "MONTHS":
		fallthrough
	case "LAST_DAY":
		*e = UsageAndPricingPeriodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UsageAndPricingPeriodEnum: %v", v)
	}
}
