// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type GetMitreReportDownloadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The Mitre report
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	GetMitreReportDownload200ApplicationJSONBinaryString io.ReadCloser
}

func (o *GetMitreReportDownloadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMitreReportDownloadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMitreReportDownloadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMitreReportDownloadResponse) GetGetMitreReportDownload200ApplicationJSONBinaryString() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.GetMitreReportDownload200ApplicationJSONBinaryString
}