// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutCloudAccountsCloudAccountIDRequest struct {
	CloudAccount shared.CloudAccountInput `request:"mediaType=application/json"`
	// cloud account ID
	CloudAccountID string `pathParam:"style=simple,explode=false,name=cloudAccountId"`
}

func (o *PutCloudAccountsCloudAccountIDRequest) GetCloudAccount() shared.CloudAccountInput {
	if o == nil {
		return shared.CloudAccountInput{}
	}
	return o.CloudAccount
}

func (o *PutCloudAccountsCloudAccountIDRequest) GetCloudAccountID() string {
	if o == nil {
		return ""
	}
	return o.CloudAccountID
}

type PutCloudAccountsCloudAccountIDResponse struct {
	// Success
	CloudAccount *shared.CloudAccount
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutCloudAccountsCloudAccountIDResponse) GetCloudAccount() *shared.CloudAccount {
	if o == nil {
		return nil
	}
	return o.CloudAccount
}

func (o *PutCloudAccountsCloudAccountIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutCloudAccountsCloudAccountIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutCloudAccountsCloudAccountIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
