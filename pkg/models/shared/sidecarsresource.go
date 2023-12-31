// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SidecarsResource struct {
	ProxyInitLimitsCPU      *string `json:"proxyInitLimitsCpu,omitempty"`
	ProxyInitLimitsMemory   *string `json:"proxyInitLimitsMemory,omitempty"`
	ProxyInitRequestsCPU    *string `json:"proxyInitRequestsCpu,omitempty"`
	ProxyInitRequestsMemory *string `json:"proxyInitRequestsMemory,omitempty"`
	ProxyLimitsCPU          *string `json:"proxyLimitsCpu,omitempty"`
	ProxyLimitsMemory       *string `json:"proxyLimitsMemory,omitempty"`
	ProxyRequestCPU         *string `json:"proxyRequestCpu,omitempty"`
	ProxyRequestMemory      *string `json:"proxyRequestMemory,omitempty"`
}

func (o *SidecarsResource) GetProxyInitLimitsCPU() *string {
	if o == nil {
		return nil
	}
	return o.ProxyInitLimitsCPU
}

func (o *SidecarsResource) GetProxyInitLimitsMemory() *string {
	if o == nil {
		return nil
	}
	return o.ProxyInitLimitsMemory
}

func (o *SidecarsResource) GetProxyInitRequestsCPU() *string {
	if o == nil {
		return nil
	}
	return o.ProxyInitRequestsCPU
}

func (o *SidecarsResource) GetProxyInitRequestsMemory() *string {
	if o == nil {
		return nil
	}
	return o.ProxyInitRequestsMemory
}

func (o *SidecarsResource) GetProxyLimitsCPU() *string {
	if o == nil {
		return nil
	}
	return o.ProxyLimitsCPU
}

func (o *SidecarsResource) GetProxyLimitsMemory() *string {
	if o == nil {
		return nil
	}
	return o.ProxyLimitsMemory
}

func (o *SidecarsResource) GetProxyRequestCPU() *string {
	if o == nil {
		return nil
	}
	return o.ProxyRequestCPU
}

func (o *SidecarsResource) GetProxyRequestMemory() *string {
	if o == nil {
		return nil
	}
	return o.ProxyRequestMemory
}
