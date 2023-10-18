// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperationsDiffsDonutPieChart struct {
	GeneralDiffs           *int64 `json:"generalDiffs,omitempty"`
	OperationsWithoutDiffs *int64 `json:"operationsWithoutDiffs,omitempty"`
	ShadowDiffs            *int64 `json:"shadowDiffs,omitempty"`
	TotalOperations        *int64 `json:"totalOperations,omitempty"`
	ZombieDiffs            *int64 `json:"zombieDiffs,omitempty"`
}

func (o *OperationsDiffsDonutPieChart) GetGeneralDiffs() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralDiffs
}

func (o *OperationsDiffsDonutPieChart) GetOperationsWithoutDiffs() *int64 {
	if o == nil {
		return nil
	}
	return o.OperationsWithoutDiffs
}

func (o *OperationsDiffsDonutPieChart) GetShadowDiffs() *int64 {
	if o == nil {
		return nil
	}
	return o.ShadowDiffs
}

func (o *OperationsDiffsDonutPieChart) GetTotalOperations() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalOperations
}

func (o *OperationsDiffsDonutPieChart) GetZombieDiffs() *int64 {
	if o == nil {
		return nil
	}
	return o.ZombieDiffs
}