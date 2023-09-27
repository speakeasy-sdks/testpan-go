# Dashboard
(*Dashboard*)

## Overview

APIs to get dashboard statistics

### Available Operations

* [GetDashboardApisecRiskFindings](#getdashboardapisecriskfindings) - Get API sec risk findings widget
* [GetDashboardApisecRiskFindingsTrend](#getdashboardapisecriskfindingstrend) - Get API sec risk findings trend graph widget for the last 30 days
* [GetDashboardApisecSpecsAndOperationsDiffs](#getdashboardapisecspecsandoperationsdiffs) - Get API sec specs and operations diffs widget
* [GetDashboardApisecTopRiskyApis](#getdashboardapisectopriskyapis) - Get API sec top risky APIs widget
* [GetDashboardApisecTopRiskyFindings](#getdashboardapisectopriskyfindings) - Get API sec top risky findings widget
* [GetDashboardClusters](#getdashboardclusters) - Get the active clusters
* [GetDashboardConnectionTelemetries](#getdashboardconnectiontelemetries) - Get pod connection dashboard data for all clusters
* [GetDashboardKubernetesAuditLogs](#getdashboardkubernetesauditlogs) - Get kubernetes audit logs dashboard data for all clusters
* [GetDashboardOperationalBar](#getdashboardoperationalbar) - Get the operation data dashboard for the given kubernetesClusterId
* [GetDashboardPermissions](#getdashboardpermissions) - Get permissions dashboard data for the given kubernetesClusterIds
* [GetDashboardPodTelemetries](#getdashboardpodtelemetries) - Get pod telemetries dashboard data for all clusters
* [GetDashboardReportDownload](#getdashboardreportdownload) - Download Secure Application security report
* [GetDashboardReportStatus](#getdashboardreportstatus) - Get Secure Application report security status
* [GetDashboardSecurityContext](#getdashboardsecuritycontext) - Get security context dashboard data for all clusters
* [GetDashboardTopSecurityRisks](#getdashboardtopsecurityrisks) - Get the top risky deployments dashboard data for the given kubernetesClusterIds
* [GetDashboardVulnerabilities](#getdashboardvulnerabilities) - Get vulnerabilities dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDConnectionTelemetries](#getdashboardkubernetesclusteridconnectiontelemetries) - Get connection telemetries dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDKubernetesAuditLogs](#getdashboardkubernetesclusteridkubernetesauditlogs) - Get kubernetes audit logs dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDPodTelemetries](#getdashboardkubernetesclusteridpodtelemetries) - Get pod telemetries dashboard data for the given kubernetesClusterId
* [GetLicensingDashboard](#getlicensingdashboard) - Get licensing dashboard data
* [PostDashboardReportGenerate](#postdashboardreportgenerate) - Generate Secure Application security report

## GetDashboardApisecRiskFindings

Get API sec risk findings widget

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardApisecRiskFindings(ctx, operations.GetDashboardApisecRiskFindingsRequest{
        APISecSource: operations.GetDashboardApisecRiskFindingsAPISecSourceExternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecRiskFindingsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetDashboardApisecRiskFindingsRequest](../../models/operations/getdashboardapisecriskfindingsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetDashboardApisecRiskFindingsResponse](../../models/operations/getdashboardapisecriskfindingsresponse.md), error**


## GetDashboardApisecRiskFindingsTrend

Get API sec risk findings trend graph widget for the last 30 days

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardApisecRiskFindingsTrend(ctx, operations.GetDashboardApisecRiskFindingsTrendRequest{
        APISecSource: operations.GetDashboardApisecRiskFindingsTrendAPISecSourceInternal,
        NumOfDays: testpango.Int64(105094),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecRiskFindingsTrendWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetDashboardApisecRiskFindingsTrendRequest](../../models/operations/getdashboardapisecriskfindingstrendrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GetDashboardApisecRiskFindingsTrendResponse](../../models/operations/getdashboardapisecriskfindingstrendresponse.md), error**


## GetDashboardApisecSpecsAndOperationsDiffs

Get API sec specs and operations diffs widget

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardApisecSpecsAndOperationsDiffs(ctx, operations.GetDashboardApisecSpecsAndOperationsDiffsRequest{
        APISecSource: operations.GetDashboardApisecSpecsAndOperationsDiffsAPISecSourceInternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SpecsAndOperationsDiffsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetDashboardApisecSpecsAndOperationsDiffsRequest](../../models/operations/getdashboardapisecspecsandoperationsdiffsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetDashboardApisecSpecsAndOperationsDiffsResponse](../../models/operations/getdashboardapisecspecsandoperationsdiffsresponse.md), error**


## GetDashboardApisecTopRiskyApis

Get API sec top risky APIs widget

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardApisecTopRiskyApis(ctx, operations.GetDashboardApisecTopRiskyApisRequest{
        APISecSource: operations.GetDashboardApisecTopRiskyApisAPISecSourceExternal,
        MaxResults: testpango.Float64(9308.77),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecTopRiskyApisWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetDashboardApisecTopRiskyApisRequest](../../models/operations/getdashboardapisectopriskyapisrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetDashboardApisecTopRiskyApisResponse](../../models/operations/getdashboardapisectopriskyapisresponse.md), error**


## GetDashboardApisecTopRiskyFindings

Get API sec top risky findings widget

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardApisecTopRiskyFindings(ctx, operations.GetDashboardApisecTopRiskyFindingsRequest{
        APISecSource: operations.GetDashboardApisecTopRiskyFindingsAPISecSourceExternal,
        MaxResults: testpango.Float64(6943.94),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecTopRiskyFindingsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetDashboardApisecTopRiskyFindingsRequest](../../models/operations/getdashboardapisectopriskyfindingsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetDashboardApisecTopRiskyFindingsResponse](../../models/operations/getdashboardapisectopriskyfindingsresponse.md), error**


## GetDashboardClusters

Get the active clusters

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardClusters(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ClustersDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardClustersResponse](../../models/operations/getdashboardclustersresponse.md), error**


## GetDashboardConnectionTelemetries

Get pod connection dashboard data for all clusters

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardConnectionTelemetries(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardConnectionTelemetriesResponse](../../models/operations/getdashboardconnectiontelemetriesresponse.md), error**


## GetDashboardKubernetesAuditLogs

Get kubernetes audit logs dashboard data for all clusters

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardKubernetesAuditLogs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardKubernetesAuditLogsResponse](../../models/operations/getdashboardkubernetesauditlogsresponse.md), error**


## GetDashboardOperationalBar

Get the operation data dashboard for the given kubernetesClusterId

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardOperationalBar(ctx, operations.GetDashboardOperationalBarRequest{
        ClustersIds: []string{
            "1c7cbdb6-eec7-4437-8ba2-5317747dc915",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OperationalBar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetDashboardOperationalBarRequest](../../models/operations/getdashboardoperationalbarrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetDashboardOperationalBarResponse](../../models/operations/getdashboardoperationalbarresponse.md), error**


## GetDashboardPermissions

Get permissions dashboard data for the given kubernetesClusterIds

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardPermissions(ctx, operations.GetDashboardPermissionsRequest{
        ClustersIds: []string{
            "ad2caf5d-d672-43dc-8f5a-e2f3a6b70087",
        },
        IncludeSystemOwners: testpango.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PermissionsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetDashboardPermissionsRequest](../../models/operations/getdashboardpermissionsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetDashboardPermissionsResponse](../../models/operations/getdashboardpermissionsresponse.md), error**


## GetDashboardPodTelemetries

Get pod telemetries dashboard data for all clusters

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardPodTelemetries(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardPodTelemetriesResponse](../../models/operations/getdashboardpodtelemetriesresponse.md), error**


## GetDashboardReportDownload

Download Secure Application security report

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardReportDownload(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDashboardReportDownload200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardReportDownloadResponse](../../models/operations/getdashboardreportdownloadresponse.md), error**


## GetDashboardReportStatus

Get Secure Application report security status

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardReportStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ReportStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDashboardReportStatusResponse](../../models/operations/getdashboardreportstatusresponse.md), error**


## GetDashboardSecurityContext

Get security context dashboard data for all clusters

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardSecurityContext(ctx, operations.GetDashboardSecurityContextRequest{
        ClustersIds: []string{
            "8756143f-5a6c-498b-9555-4080d40bcacc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SecurityContextWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetDashboardSecurityContextRequest](../../models/operations/getdashboardsecuritycontextrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetDashboardSecurityContextResponse](../../models/operations/getdashboardsecuritycontextresponse.md), error**


## GetDashboardTopSecurityRisks

Get the top risky deployments dashboard data for the given kubernetesClusterIds

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardTopSecurityRisks(ctx, operations.GetDashboardTopSecurityRisksRequest{
        ClustersIds: []string{
            "6cbd6b5f-3ec9-4093-84f9-26bad2553819",
        },
        Size: testpango.Int64(748606),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TopSecurityRisksWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetDashboardTopSecurityRisksRequest](../../models/operations/getdashboardtopsecurityrisksrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetDashboardTopSecurityRisksResponse](../../models/operations/getdashboardtopsecurityrisksresponse.md), error**


## GetDashboardVulnerabilities

Get vulnerabilities dashboard data for the given kubernetesClusterId

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardVulnerabilities(ctx, operations.GetDashboardVulnerabilitiesRequest{
        ClustersIds: []string{
            "474b0ed2-0e56-4248-bff6-39a910abdcab",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VulnerabilitiesWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetDashboardVulnerabilitiesRequest](../../models/operations/getdashboardvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetDashboardVulnerabilitiesResponse](../../models/operations/getdashboardvulnerabilitiesresponse.md), error**


## GetDashboardKubernetesClusterIDConnectionTelemetries

Get connection telemetries dashboard data for the given kubernetesClusterId

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardKubernetesClusterIDConnectionTelemetries(ctx, operations.GetDashboardKubernetesClusterIDConnectionTelemetriesRequest{
        KubernetesClusterID: "62676696-e1ec-4002-a1b3-35d89acb3ecf",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetDashboardKubernetesClusterIDConnectionTelemetriesRequest](../../models/operations/getdashboardkubernetesclusteridconnectiontelemetriesrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetDashboardKubernetesClusterIDConnectionTelemetriesResponse](../../models/operations/getdashboardkubernetesclusteridconnectiontelemetriesresponse.md), error**


## GetDashboardKubernetesClusterIDKubernetesAuditLogs

Get kubernetes audit logs dashboard data for the given kubernetesClusterId

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardKubernetesClusterIDKubernetesAuditLogs(ctx, operations.GetDashboardKubernetesClusterIDKubernetesAuditLogsRequest{
        KubernetesClusterID: "da8d0c54-9ef0-4300-8978-a61fa1cf2068",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetDashboardKubernetesClusterIDKubernetesAuditLogsRequest](../../models/operations/getdashboardkubernetesclusteridkubernetesauditlogsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetDashboardKubernetesClusterIDKubernetesAuditLogsResponse](../../models/operations/getdashboardkubernetesclusteridkubernetesauditlogsresponse.md), error**


## GetDashboardKubernetesClusterIDPodTelemetries

Get pod telemetries dashboard data for the given kubernetesClusterId

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetDashboardKubernetesClusterIDPodTelemetries(ctx, operations.GetDashboardKubernetesClusterIDPodTelemetriesRequest{
        KubernetesClusterID: "8f77c1ff-c71d-4ca1-a3f2-a3c80a97ff33",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TimeBasedWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetDashboardKubernetesClusterIDPodTelemetriesRequest](../../models/operations/getdashboardkubernetesclusteridpodtelemetriesrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.GetDashboardKubernetesClusterIDPodTelemetriesResponse](../../models/operations/getdashboardkubernetesclusteridpodtelemetriesresponse.md), error**


## GetLicensingDashboard

Get licensing dashboard data

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.GetLicensingDashboard(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LicensingDashboard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetLicensingDashboardResponse](../../models/operations/getlicensingdashboardresponse.md), error**


## PostDashboardReportGenerate

Generate Secure Application security report

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Dashboard.PostDashboardReportGenerate(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostDashboardReportGenerateResponse](../../models/operations/postdashboardreportgenerateresponse.md), error**

