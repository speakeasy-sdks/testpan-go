// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
)

type ProxyConfiguration struct {
	// Specifies if the proxy configuration should be used
	EnableProxy *bool   `default:"false" json:"enableProxy"`
	HTTPSProxy  *string `json:"httpsProxy,omitempty"`
}

func (p ProxyConfiguration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProxyConfiguration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProxyConfiguration) GetEnableProxy() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxy
}

func (o *ProxyConfiguration) GetHTTPSProxy() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxy
}
