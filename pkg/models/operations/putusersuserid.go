// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutUsersUserIDRequest struct {
	EditUser shared.EditUser `request:"mediaType=application/json"`
	UserID   string          `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *PutUsersUserIDRequest) GetEditUser() shared.EditUser {
	if o == nil {
		return shared.EditUser{}
	}
	return o.EditUser
}

func (o *PutUsersUserIDRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type PutUsersUserIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	User *shared.User
}

func (o *PutUsersUserIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutUsersUserIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutUsersUserIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutUsersUserIDResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
