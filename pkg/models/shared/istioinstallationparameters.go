// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
)

// IstioInstallationParameters - istio related information
type IstioInstallationParameters struct {
	// indicates whether Istio is already installed on this cluster (which means Secure Application should not install it)
	IsIstioAlreadyInstalled *bool `default:"false" json:"isIstioAlreadyInstalled"`
	// when istio already installed, choose the version from supported istio versions list: /istio/supportedVersions
	IstioVersion *string `json:"istioVersion,omitempty"`
}

func (i IstioInstallationParameters) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IstioInstallationParameters) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IstioInstallationParameters) GetIsIstioAlreadyInstalled() *bool {
	if o == nil {
		return nil
	}
	return o.IsIstioAlreadyInstalled
}

func (o *IstioInstallationParameters) GetIstioVersion() *string {
	if o == nil {
		return nil
	}
	return o.IstioVersion
}
