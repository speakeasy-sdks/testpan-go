// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PodSecurityPolicy struct {
	AllowPrivilegeEscalation        *bool                      `json:"allowPrivilegeEscalation,omitempty"`
	AllowedCapabilities             []string                   `json:"allowedCapabilities,omitempty"`
	AllowedHostPaths                []AllowedHostPath          `json:"allowedHostPaths,omitempty"`
	AllowedProcMountTypes           []AllowedProcMountType     `json:"allowedProcMountTypes,omitempty"`
	AllowedUnsafeSysctls            []string                   `json:"allowedUnsafeSysctls,omitempty"`
	DefaultAllowPrivilegeEscalation *bool                      `json:"defaultAllowPrivilegeEscalation,omitempty"`
	Description                     *string                    `json:"description,omitempty"`
	ForbiddenSysctls                []string                   `json:"forbiddenSysctls,omitempty"`
	FsGroup                         RunAsGroupStrategyOptions  `json:"fsGroup"`
	HostIPC                         *bool                      `json:"hostIPC,omitempty"`
	HostNetwork                     *bool                      `json:"hostNetwork,omitempty"`
	HostPID                         *bool                      `json:"hostPID,omitempty"`
	HostPorts                       []HostPortRange            `json:"hostPorts,omitempty"`
	ID                              *string                    `json:"id,omitempty"`
	IsSecurecnDefaultProfile        *bool                      `json:"isSecurecnDefaultProfile,omitempty"`
	Name                            string                     `json:"name"`
	Privileged                      *bool                      `json:"privileged,omitempty"`
	ReadOnlyRootFileSystem          *bool                      `json:"readOnlyRootFileSystem,omitempty"`
	RequiredDropCapabilities        []string                   `json:"requiredDropCapabilities,omitempty"`
	RunAsGroup                      *RunAsGroupStrategyOptions `json:"runAsGroup,omitempty"`
	RunAsUser                       RunAsUserStrategyOptions   `json:"runAsUser"`
	SeccompProfile                  *string                    `json:"seccompProfile,omitempty"`
	SupplementalGroups              RunAsGroupStrategyOptions  `json:"supplementalGroups"`
	Volumes                         []PSPVolumeTypes           `json:"volumes,omitempty"`
}

func (o *PodSecurityPolicy) GetAllowPrivilegeEscalation() *bool {
	if o == nil {
		return nil
	}
	return o.AllowPrivilegeEscalation
}

func (o *PodSecurityPolicy) GetAllowedCapabilities() []string {
	if o == nil {
		return nil
	}
	return o.AllowedCapabilities
}

func (o *PodSecurityPolicy) GetAllowedHostPaths() []AllowedHostPath {
	if o == nil {
		return nil
	}
	return o.AllowedHostPaths
}

func (o *PodSecurityPolicy) GetAllowedProcMountTypes() []AllowedProcMountType {
	if o == nil {
		return nil
	}
	return o.AllowedProcMountTypes
}

func (o *PodSecurityPolicy) GetAllowedUnsafeSysctls() []string {
	if o == nil {
		return nil
	}
	return o.AllowedUnsafeSysctls
}

func (o *PodSecurityPolicy) GetDefaultAllowPrivilegeEscalation() *bool {
	if o == nil {
		return nil
	}
	return o.DefaultAllowPrivilegeEscalation
}

func (o *PodSecurityPolicy) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PodSecurityPolicy) GetForbiddenSysctls() []string {
	if o == nil {
		return nil
	}
	return o.ForbiddenSysctls
}

func (o *PodSecurityPolicy) GetFsGroup() RunAsGroupStrategyOptions {
	if o == nil {
		return RunAsGroupStrategyOptions{}
	}
	return o.FsGroup
}

func (o *PodSecurityPolicy) GetHostIPC() *bool {
	if o == nil {
		return nil
	}
	return o.HostIPC
}

func (o *PodSecurityPolicy) GetHostNetwork() *bool {
	if o == nil {
		return nil
	}
	return o.HostNetwork
}

func (o *PodSecurityPolicy) GetHostPID() *bool {
	if o == nil {
		return nil
	}
	return o.HostPID
}

func (o *PodSecurityPolicy) GetHostPorts() []HostPortRange {
	if o == nil {
		return nil
	}
	return o.HostPorts
}

func (o *PodSecurityPolicy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PodSecurityPolicy) GetIsSecurecnDefaultProfile() *bool {
	if o == nil {
		return nil
	}
	return o.IsSecurecnDefaultProfile
}

func (o *PodSecurityPolicy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PodSecurityPolicy) GetPrivileged() *bool {
	if o == nil {
		return nil
	}
	return o.Privileged
}

func (o *PodSecurityPolicy) GetReadOnlyRootFileSystem() *bool {
	if o == nil {
		return nil
	}
	return o.ReadOnlyRootFileSystem
}

func (o *PodSecurityPolicy) GetRequiredDropCapabilities() []string {
	if o == nil {
		return nil
	}
	return o.RequiredDropCapabilities
}

func (o *PodSecurityPolicy) GetRunAsGroup() *RunAsGroupStrategyOptions {
	if o == nil {
		return nil
	}
	return o.RunAsGroup
}

func (o *PodSecurityPolicy) GetRunAsUser() RunAsUserStrategyOptions {
	if o == nil {
		return RunAsUserStrategyOptions{}
	}
	return o.RunAsUser
}

func (o *PodSecurityPolicy) GetSeccompProfile() *string {
	if o == nil {
		return nil
	}
	return o.SeccompProfile
}

func (o *PodSecurityPolicy) GetSupplementalGroups() RunAsGroupStrategyOptions {
	if o == nil {
		return RunAsGroupStrategyOptions{}
	}
	return o.SupplementalGroups
}

func (o *PodSecurityPolicy) GetVolumes() []PSPVolumeTypes {
	if o == nil {
		return nil
	}
	return o.Volumes
}
