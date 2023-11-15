// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type APIServiceSpecPathHitCountGraphPoint struct {
	Date  *time.Time `json:"date,omitempty"`
	Graph *float64   `json:"graph,omitempty"`
}

func (a APIServiceSpecPathHitCountGraphPoint) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIServiceSpecPathHitCountGraphPoint) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIServiceSpecPathHitCountGraphPoint) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *APIServiceSpecPathHitCountGraphPoint) GetGraph() *float64 {
	if o == nil {
		return nil
	}
	return o.Graph
}
