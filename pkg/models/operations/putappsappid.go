// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"net/http"
)

type PutAppsAppIDRequest struct {
	App   shared.App `request:"mediaType=application/json"`
	AppID string     `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *PutAppsAppIDRequest) GetApp() shared.App {
	if o == nil {
		return shared.App{}
	}
	return o.App
}

func (o *PutAppsAppIDRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type PutAppsAppIDResponse struct {
	// App was edited.
	App *shared.App
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAppsAppIDResponse) GetApp() *shared.App {
	if o == nil {
		return nil
	}
	return o.App
}

func (o *PutAppsAppIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAppsAppIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAppsAppIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
