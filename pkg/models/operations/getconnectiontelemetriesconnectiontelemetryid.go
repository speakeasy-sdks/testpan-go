// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
	"time"
)

type GetConnectionTelemetriesConnectionTelemetryIDRequest struct {
	// connection telemetry ID
	ConnectionTelemetryID string `pathParam:"style=simple,explode=false,name=connectionTelemetryId"`
	// End date of the query
	EndTime time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Start date of the query
	StartTime time.Time `queryParam:"style=form,explode=true,name=startTime"`
}

func (g GetConnectionTelemetriesConnectionTelemetryIDRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetConnectionTelemetriesConnectionTelemetryIDRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDRequest) GetConnectionTelemetryID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionTelemetryID
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDRequest) GetEndTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndTime
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDRequest) GetStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartTime
}

type GetConnectionTelemetriesConnectionTelemetryIDResponse struct {
	// OK
	ConnectionTelemetry *shared.ConnectionTelemetry
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDResponse) GetConnectionTelemetry() *shared.ConnectionTelemetry {
	if o == nil {
		return nil
	}
	return o.ConnectionTelemetry
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionTelemetriesConnectionTelemetryIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}