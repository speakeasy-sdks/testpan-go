// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TruncationStatus struct {
	// is truncation enabled.
	IsTruncationEnabled bool `json:"isTruncationEnabled"`
	// truncation interval, in days.
	TruncateTimeInDays int64 `json:"truncateTimeInDays"`
}

func (o *TruncationStatus) GetIsTruncationEnabled() bool {
	if o == nil {
		return false
	}
	return o.IsTruncationEnabled
}

func (o *TruncationStatus) GetTruncateTimeInDays() int64 {
	if o == nil {
		return 0
	}
	return o.TruncateTimeInDays
}