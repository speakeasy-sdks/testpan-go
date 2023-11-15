// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
)

type WorkloadAddress struct {
	Address         string           `json:"address"`
	NetworkProtocol *NetworkProtocol `default:"TCP" json:"networkProtocol"`
}

func (w WorkloadAddress) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkloadAddress) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WorkloadAddress) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *WorkloadAddress) GetNetworkProtocol() *NetworkProtocol {
	if o == nil {
		return nil
	}
	return o.NetworkProtocol
}
