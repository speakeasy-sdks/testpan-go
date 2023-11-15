// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
	"time"
)

// QueryParamProtectionStatus - When true, the API will return only protected pods
type QueryParamProtectionStatus string

const (
	QueryParamProtectionStatusFull           QueryParamProtectionStatus = "FULL"
	QueryParamProtectionStatusDeploymentOnly QueryParamProtectionStatus = "DEPLOYMENT_ONLY"
	QueryParamProtectionStatusConnectionOnly QueryParamProtectionStatus = "CONNECTION_ONLY"
	QueryParamProtectionStatusDisabled       QueryParamProtectionStatus = "DISABLED"
	QueryParamProtectionStatusAll            QueryParamProtectionStatus = "ALL"
)

func (e QueryParamProtectionStatus) ToPointer() *QueryParamProtectionStatus {
	return &e
}

func (e *QueryParamProtectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FULL":
		fallthrough
	case "DEPLOYMENT_ONLY":
		fallthrough
	case "CONNECTION_ONLY":
		fallthrough
	case "DISABLED":
		fallthrough
	case "ALL":
		*e = QueryParamProtectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamProtectionStatus: %v", v)
	}
}

type GetAppTelemetriesQueryParamResult string

const (
	GetAppTelemetriesQueryParamResultAllow  GetAppTelemetriesQueryParamResult = "ALLOW"
	GetAppTelemetriesQueryParamResultDetect GetAppTelemetriesQueryParamResult = "DETECT"
	GetAppTelemetriesQueryParamResultBlock  GetAppTelemetriesQueryParamResult = "BLOCK"
)

func (e GetAppTelemetriesQueryParamResult) ToPointer() *GetAppTelemetriesQueryParamResult {
	return &e
}

func (e *GetAppTelemetriesQueryParamResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALLOW":
		fallthrough
	case "DETECT":
		fallthrough
	case "BLOCK":
		*e = GetAppTelemetriesQueryParamResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppTelemetriesQueryParamResult: %v", v)
	}
}

// GetAppTelemetriesQueryParamSortDir - sorting direction
type GetAppTelemetriesQueryParamSortDir string

const (
	GetAppTelemetriesQueryParamSortDirAsc  GetAppTelemetriesQueryParamSortDir = "ASC"
	GetAppTelemetriesQueryParamSortDirDesc GetAppTelemetriesQueryParamSortDir = "DESC"
)

func (e GetAppTelemetriesQueryParamSortDir) ToPointer() *GetAppTelemetriesQueryParamSortDir {
	return &e
}

func (e *GetAppTelemetriesQueryParamSortDir) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = GetAppTelemetriesQueryParamSortDir(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppTelemetriesQueryParamSortDir: %v", v)
	}
}

// GetAppTelemetriesQueryParamSortKey - sort key
type GetAppTelemetriesQueryParamSortKey string

const (
	GetAppTelemetriesQueryParamSortKeyAppName         GetAppTelemetriesQueryParamSortKey = "appName"
	GetAppTelemetriesQueryParamSortKeyAppType         GetAppTelemetriesQueryParamSortKey = "appType"
	GetAppTelemetriesQueryParamSortKeyExecutable      GetAppTelemetriesQueryParamSortKey = "executable"
	GetAppTelemetriesQueryParamSortKeyEnvironmentName GetAppTelemetriesQueryParamSortKey = "environmentName"
	GetAppTelemetriesQueryParamSortKeyRisk            GetAppTelemetriesQueryParamSortKey = "risk"
	GetAppTelemetriesQueryParamSortKeyStatus          GetAppTelemetriesQueryParamSortKey = "status"
	GetAppTelemetriesQueryParamSortKeyStartTime       GetAppTelemetriesQueryParamSortKey = "startTime"
	GetAppTelemetriesQueryParamSortKeyFinishTime      GetAppTelemetriesQueryParamSortKey = "finishTime"
	GetAppTelemetriesQueryParamSortKeyWorkloadRisk    GetAppTelemetriesQueryParamSortKey = "workloadRisk"
)

func (e GetAppTelemetriesQueryParamSortKey) ToPointer() *GetAppTelemetriesQueryParamSortKey {
	return &e
}

func (e *GetAppTelemetriesQueryParamSortKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "appName":
		fallthrough
	case "appType":
		fallthrough
	case "executable":
		fallthrough
	case "environmentName":
		fallthrough
	case "risk":
		fallthrough
	case "status":
		fallthrough
	case "startTime":
		fallthrough
	case "finishTime":
		fallthrough
	case "workloadRisk":
		*e = GetAppTelemetriesQueryParamSortKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppTelemetriesQueryParamSortKey: %v", v)
	}
}

type WorkloadRisks string

const (
	WorkloadRisksLow      WorkloadRisks = "LOW"
	WorkloadRisksMedium   WorkloadRisks = "MEDIUM"
	WorkloadRisksHigh     WorkloadRisks = "HIGH"
	WorkloadRisksCritical WorkloadRisks = "CRITICAL"
)

func (e WorkloadRisks) ToPointer() *WorkloadRisks {
	return &e
}

func (e *WorkloadRisks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOW":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		fallthrough
	case "CRITICAL":
		*e = WorkloadRisks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkloadRisks: %v", v)
	}
}

type GetAppTelemetriesRequest struct {
	// Defined App name
	AppName []string `queryParam:"style=form,explode=false,name=appName"`
	// Empty string means no filtering. "UNDEFINED" means telemetries with no App type
	AppType []string `queryParam:"style=form,explode=false,name=appType"`
	// When true, the API will return an xlsx file, and pagination will be ignored
	DownloadAsXlsx *bool `queryParam:"style=form,explode=true,name=downloadAsXlsx"`
	// End date of the query
	EndTime time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Empty string means no filtering. "UNDEFINED" means telemetries with no App type
	EnvironmentName []string `queryParam:"style=form,explode=false,name=environmentName"`
	Executable      []string `queryParam:"style=form,explode=false,name=executable"`
	// When true, the API will filter out "OS Internal" and "User OS Internal" App types
	HideInternals *bool `default:"false" queryParam:"style=form,explode=true,name=hideInternals"`
	// Highest DockerfileScan Result
	HighestDockerfileScanResult []string `queryParam:"style=form,explode=false,name=highestDockerfileScanResult"`
	// Defined host name
	Host []string `queryParam:"style=form,explode=false,name=host"`
	// Array of images id
	ImagesID []string `queryParam:"style=form,explode=false,name=imagesId"`
	// app is identified filter
	IsIdentified *bool `queryParam:"style=form,explode=true,name=isIdentified"`
	// The number of entries to return (pagination)
	MaxResults *float64 `default:"100" queryParam:"style=form,explode=true,name=maxResults"`
	// namespace filter. a base 64 representation of a list of NamespacesFilter definition object
	NamespacesFilter *string `queryParam:"style=form,explode=true,name=namespacesFilter"`
	// Return entries from this offset (pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// When true, the API will return only protected pods
	ProtectionStatus *QueryParamProtectionStatus `default:"ALL" queryParam:"style=form,explode=true,name=protectionStatus"`
	// app result filter
	Result []GetAppTelemetriesQueryParamResult `queryParam:"style=form,explode=false,name=result"`
	// When true, the telemetries API will only return entries with the App name
	ShowOnlyEntriesWithAppName *bool `default:"false" queryParam:"style=form,explode=true,name=showOnlyEntriesWithAppName"`
	// When true, the API will only return entries that violate the active policy
	ShowOnlyViolations *bool `queryParam:"style=form,explode=true,name=showOnlyViolations"`
	// When true, the telemetries API will also return workloads that are part of the Kubernetes system
	ShowSystemPods *bool `default:"false" queryParam:"style=form,explode=true,name=showSystemPods"`
	// sorting direction
	SortDir *GetAppTelemetriesQueryParamSortDir `default:"ASC" queryParam:"style=form,explode=true,name=sortDir"`
	// sort key
	SortKey GetAppTelemetriesQueryParamSortKey `queryParam:"style=form,explode=true,name=sortKey"`
	// Start date of the query
	StartTime time.Time `queryParam:"style=form,explode=true,name=startTime"`
	// App status
	Status []string `queryParam:"style=form,explode=false,name=status"`
	// Highest vulnerability
	VulnerabilityLevelsFilter []string `queryParam:"style=form,explode=false,name=vulnerabilityLevelsFilter"`
	// workloadRisk filter
	WorkloadRisks []WorkloadRisks `queryParam:"style=form,explode=false,name=workloadRisks"`
}

func (g GetAppTelemetriesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAppTelemetriesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAppTelemetriesRequest) GetAppName() []string {
	if o == nil {
		return nil
	}
	return o.AppName
}

func (o *GetAppTelemetriesRequest) GetAppType() []string {
	if o == nil {
		return nil
	}
	return o.AppType
}

func (o *GetAppTelemetriesRequest) GetDownloadAsXlsx() *bool {
	if o == nil {
		return nil
	}
	return o.DownloadAsXlsx
}

func (o *GetAppTelemetriesRequest) GetEndTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndTime
}

func (o *GetAppTelemetriesRequest) GetEnvironmentName() []string {
	if o == nil {
		return nil
	}
	return o.EnvironmentName
}

func (o *GetAppTelemetriesRequest) GetExecutable() []string {
	if o == nil {
		return nil
	}
	return o.Executable
}

func (o *GetAppTelemetriesRequest) GetHideInternals() *bool {
	if o == nil {
		return nil
	}
	return o.HideInternals
}

func (o *GetAppTelemetriesRequest) GetHighestDockerfileScanResult() []string {
	if o == nil {
		return nil
	}
	return o.HighestDockerfileScanResult
}

func (o *GetAppTelemetriesRequest) GetHost() []string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *GetAppTelemetriesRequest) GetImagesID() []string {
	if o == nil {
		return nil
	}
	return o.ImagesID
}

func (o *GetAppTelemetriesRequest) GetIsIdentified() *bool {
	if o == nil {
		return nil
	}
	return o.IsIdentified
}

func (o *GetAppTelemetriesRequest) GetMaxResults() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxResults
}

func (o *GetAppTelemetriesRequest) GetNamespacesFilter() *string {
	if o == nil {
		return nil
	}
	return o.NamespacesFilter
}

func (o *GetAppTelemetriesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetAppTelemetriesRequest) GetProtectionStatus() *QueryParamProtectionStatus {
	if o == nil {
		return nil
	}
	return o.ProtectionStatus
}

func (o *GetAppTelemetriesRequest) GetResult() []GetAppTelemetriesQueryParamResult {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *GetAppTelemetriesRequest) GetShowOnlyEntriesWithAppName() *bool {
	if o == nil {
		return nil
	}
	return o.ShowOnlyEntriesWithAppName
}

func (o *GetAppTelemetriesRequest) GetShowOnlyViolations() *bool {
	if o == nil {
		return nil
	}
	return o.ShowOnlyViolations
}

func (o *GetAppTelemetriesRequest) GetShowSystemPods() *bool {
	if o == nil {
		return nil
	}
	return o.ShowSystemPods
}

func (o *GetAppTelemetriesRequest) GetSortDir() *GetAppTelemetriesQueryParamSortDir {
	if o == nil {
		return nil
	}
	return o.SortDir
}

func (o *GetAppTelemetriesRequest) GetSortKey() GetAppTelemetriesQueryParamSortKey {
	if o == nil {
		return GetAppTelemetriesQueryParamSortKey("")
	}
	return o.SortKey
}

func (o *GetAppTelemetriesRequest) GetStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartTime
}

func (o *GetAppTelemetriesRequest) GetStatus() []string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetAppTelemetriesRequest) GetVulnerabilityLevelsFilter() []string {
	if o == nil {
		return nil
	}
	return o.VulnerabilityLevelsFilter
}

func (o *GetAppTelemetriesRequest) GetWorkloadRisks() []WorkloadRisks {
	if o == nil {
		return nil
	}
	return o.WorkloadRisks
}

type GetAppTelemetriesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.AppTelemetry
}

func (o *GetAppTelemetriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppTelemetriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppTelemetriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAppTelemetriesResponse) GetClasses() []shared.AppTelemetry {
	if o == nil {
		return nil
	}
	return o.Classes
}
