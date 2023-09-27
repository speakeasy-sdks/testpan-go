// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RunAsUserStrategyOptions struct {
	Ranges []IDRange          `json:"ranges,omitempty"`
	Rule   *RunAsUserStrategy `json:"rule,omitempty"`
}

func (o *RunAsUserStrategyOptions) GetRanges() []IDRange {
	if o == nil {
		return nil
	}
	return o.Ranges
}

func (o *RunAsUserStrategyOptions) GetRule() *RunAsUserStrategy {
	if o == nil {
		return nil
	}
	return o.Rule
}
