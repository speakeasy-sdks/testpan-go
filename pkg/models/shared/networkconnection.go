// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"time"
)

type NetworkConnectionViolation struct {
	EncryptRule      *ConnectionRuleBasic `json:"encryptRule,omitempty"`
	EncryptionReason *EncryptionReason    `json:"encryptionReason,omitempty"`
	LastViolation    *time.Time           `json:"lastViolation,omitempty"`
	Rule             *ConnectionRuleBasic `json:"rule,omitempty"`
}

func (n NetworkConnectionViolation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NetworkConnectionViolation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NetworkConnectionViolation) GetEncryptRule() *ConnectionRuleBasic {
	if o == nil {
		return nil
	}
	return o.EncryptRule
}

func (o *NetworkConnectionViolation) GetEncryptionReason() *EncryptionReason {
	if o == nil {
		return nil
	}
	return o.EncryptionReason
}

func (o *NetworkConnectionViolation) GetLastViolation() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastViolation
}

func (o *NetworkConnectionViolation) GetRule() *ConnectionRuleBasic {
	if o == nil {
		return nil
	}
	return o.Rule
}

type NetworkConnection struct {
	// Destination App id
	DestinationID         *string `json:"destinationId,omitempty"`
	DestinationPortNumber *int64  `json:"destinationPortNumber,omitempty"`
	ID                    *string `json:"id,omitempty"`
	NumberOfConnections   *int64  `json:"numberOfConnections,omitempty"`
	Protocol              *string `json:"protocol,omitempty"`
	// Source App id
	SourceID  *string                     `json:"sourceId,omitempty"`
	StartTime *time.Time                  `json:"startTime,omitempty"`
	Violation *NetworkConnectionViolation `json:"violation,omitempty"`
}

func (n NetworkConnection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NetworkConnection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NetworkConnection) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *NetworkConnection) GetDestinationPortNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.DestinationPortNumber
}

func (o *NetworkConnection) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NetworkConnection) GetNumberOfConnections() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfConnections
}

func (o *NetworkConnection) GetProtocol() *string {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *NetworkConnection) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

func (o *NetworkConnection) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *NetworkConnection) GetViolation() *NetworkConnectionViolation {
	if o == nil {
		return nil
	}
	return o.Violation
}