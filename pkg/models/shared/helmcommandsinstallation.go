// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type HelmCommandsInstallation struct {
	// Cmd of istio values
	IstioHelmCommand *string `json:"istioHelmCommand,omitempty"`
	// Cmd of panoptica values
	PanopticaHelmCommand *string `json:"panopticaHelmCommand,omitempty"`
	// Cmd of vault values
	VaultHelmCommand *string `json:"vaultHelmCommand,omitempty"`
}

func (o *HelmCommandsInstallation) GetIstioHelmCommand() *string {
	if o == nil {
		return nil
	}
	return o.IstioHelmCommand
}

func (o *HelmCommandsInstallation) GetPanopticaHelmCommand() *string {
	if o == nil {
		return nil
	}
	return o.PanopticaHelmCommand
}

func (o *HelmCommandsInstallation) GetVaultHelmCommand() *string {
	if o == nil {
		return nil
	}
	return o.VaultHelmCommand
}
