// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

// PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType - The approve action type (ADD/REMOVE)
type PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType string

const (
	PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionTypeAdd    PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType = "ADD"
	PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionTypeRemove PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType = "REMOVE"
)

func (e PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType) ToPointer() *PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType {
	return &e
}

func (e *PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADD":
		fallthrough
	case "REMOVE":
		*e = PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType: %v", v)
	}
}

type PostRiskAssessmentPermissionsOwnerIDApproveRequest struct {
	UUIDList shared.UUIDList `request:"mediaType=application/json"`
	// The approve action type (ADD/REMOVE)
	ActionType PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType `queryParam:"style=form,explode=true,name=actionType"`
	OwnerID    string                                                          `pathParam:"style=simple,explode=false,name=ownerId"`
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveRequest) GetUUIDList() shared.UUIDList {
	if o == nil {
		return shared.UUIDList{}
	}
	return o.UUIDList
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveRequest) GetActionType() PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType {
	if o == nil {
		return PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionType("")
	}
	return o.ActionType
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveRequest) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

type PostRiskAssessmentPermissionsOwnerIDApproveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostRiskAssessmentPermissionsOwnerIDApproveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
