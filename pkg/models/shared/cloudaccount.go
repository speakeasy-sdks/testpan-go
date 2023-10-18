// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type CloudAccountValidateFunction string

const (
	CloudAccountValidateFunctionHashValidation      CloudAccountValidateFunction = "HASH_VALIDATION"
	CloudAccountValidateFunctionSignatureValidation CloudAccountValidateFunction = "SIGNATURE_VALIDATION"
	CloudAccountValidateFunctionNone                CloudAccountValidateFunction = "NONE"
)

func (e CloudAccountValidateFunction) ToPointer() *CloudAccountValidateFunction {
	return &e
}

func (e *CloudAccountValidateFunction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HASH_VALIDATION":
		fallthrough
	case "SIGNATURE_VALIDATION":
		fallthrough
	case "NONE":
		*e = CloudAccountValidateFunction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CloudAccountValidateFunction: %v", v)
	}
}

// CloudAccountInput - represent cloud account object
type CloudAccountInput struct {
	CloudProvider               *CloudProviderType              `json:"cloudProvider,omitempty"`
	InstallVulnerabilityScanner *bool                           `default:"false" json:"installVulnerabilityScanner"`
	Name                        *string                         `json:"name,omitempty"`
	PeriodicJobExpression       ServerlessPeriodicJobExpression `json:"periodicJobExpression"`
	Regions                     []string                        `json:"regions,omitempty"`
	SecurityThreats             *CloudAccountSecurityThreats    `json:"securityThreats,omitempty"`
	ValidateFunction            *CloudAccountValidateFunction   `json:"validateFunction,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (c CloudAccountInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CloudAccountInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CloudAccountInput) GetCloudProvider() *CloudProviderType {
	if o == nil {
		return nil
	}
	return o.CloudProvider
}

func (o *CloudAccountInput) GetInstallVulnerabilityScanner() *bool {
	if o == nil {
		return nil
	}
	return o.InstallVulnerabilityScanner
}

func (o *CloudAccountInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CloudAccountInput) GetPeriodicJobExpression() ServerlessPeriodicJobExpression {
	if o == nil {
		return ServerlessPeriodicJobExpression{}
	}
	return o.PeriodicJobExpression
}

func (o *CloudAccountInput) GetRegions() []string {
	if o == nil {
		return nil
	}
	return o.Regions
}

func (o *CloudAccountInput) GetSecurityThreats() *CloudAccountSecurityThreats {
	if o == nil {
		return nil
	}
	return o.SecurityThreats
}

func (o *CloudAccountInput) GetValidateFunction() *CloudAccountValidateFunction {
	if o == nil {
		return nil
	}
	return o.ValidateFunction
}

func (o *CloudAccountInput) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}

type CloudAccountInstallationStatus string

const (
	CloudAccountInstallationStatusInstalled           CloudAccountInstallationStatus = "INSTALLED"
	CloudAccountInstallationStatusPendingInstallation CloudAccountInstallationStatus = "PENDING_INSTALLATION"
	CloudAccountInstallationStatusFailed              CloudAccountInstallationStatus = "FAILED"
)

func (e CloudAccountInstallationStatus) ToPointer() *CloudAccountInstallationStatus {
	return &e
}

func (e *CloudAccountInstallationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INSTALLED":
		fallthrough
	case "PENDING_INSTALLATION":
		fallthrough
	case "FAILED":
		*e = CloudAccountInstallationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CloudAccountInstallationStatus: %v", v)
	}
}

// CloudAccount - represent cloud account object
type CloudAccount struct {
	// the identifier id from the cloud account provider. account ID for AWS and subscription ID in Azure
	CloudAccountID              *string                         `json:"cloudAccountId,omitempty"`
	CloudProvider               *CloudProviderType              `json:"cloudProvider,omitempty"`
	ID                          *string                         `json:"id,omitempty"`
	InstallVulnerabilityScanner *bool                           `default:"false" json:"installVulnerabilityScanner"`
	InstallationStatus          *CloudAccountInstallationStatus `json:"installationStatus,omitempty"`
	LastScanned                 *time.Time                      `json:"lastScanned,omitempty"`
	Name                        *string                         `json:"name,omitempty"`
	PeriodicJobExpression       ServerlessPeriodicJobExpression `json:"periodicJobExpression"`
	Regions                     []string                        `json:"regions,omitempty"`
	SecurityThreats             *CloudAccountSecurityThreats    `json:"securityThreats,omitempty"`
	ValidateFunction            *CloudAccountValidateFunction   `json:"validateFunction,omitempty"`
	// Vulnerabilities summary by severity
	VulnerabilitiesSummary *VulnerabilitiesSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (c CloudAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CloudAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CloudAccount) GetCloudAccountID() *string {
	if o == nil {
		return nil
	}
	return o.CloudAccountID
}

func (o *CloudAccount) GetCloudProvider() *CloudProviderType {
	if o == nil {
		return nil
	}
	return o.CloudProvider
}

func (o *CloudAccount) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CloudAccount) GetInstallVulnerabilityScanner() *bool {
	if o == nil {
		return nil
	}
	return o.InstallVulnerabilityScanner
}

func (o *CloudAccount) GetInstallationStatus() *CloudAccountInstallationStatus {
	if o == nil {
		return nil
	}
	return o.InstallationStatus
}

func (o *CloudAccount) GetLastScanned() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastScanned
}

func (o *CloudAccount) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CloudAccount) GetPeriodicJobExpression() ServerlessPeriodicJobExpression {
	if o == nil {
		return ServerlessPeriodicJobExpression{}
	}
	return o.PeriodicJobExpression
}

func (o *CloudAccount) GetRegions() []string {
	if o == nil {
		return nil
	}
	return o.Regions
}

func (o *CloudAccount) GetSecurityThreats() *CloudAccountSecurityThreats {
	if o == nil {
		return nil
	}
	return o.SecurityThreats
}

func (o *CloudAccount) GetValidateFunction() *CloudAccountValidateFunction {
	if o == nil {
		return nil
	}
	return o.ValidateFunction
}

func (o *CloudAccount) GetVulnerabilitiesSummary() *VulnerabilitiesSummary {
	if o == nil {
		return nil
	}
	return o.VulnerabilitiesSummary
}