# Dashboard
(*.Dashboard*)

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
        APISecSource: operations.QueryParamAPISecSourceInternal,
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
        APISecSource: operations.GetDashboardApisecRiskFindingsTrendQueryParamAPISecSourceInternal,
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
        APISecSource: operations.GetDashboardApisecSpecsAndOperationsDiffsQueryParamAPISecSourceExternal,
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
        APISecSource: operations.GetDashboardApisecTopRiskyApisQueryParamAPISecSourceExternal,
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
        APISecSource: operations.GetDashboardApisecTopRiskyFindingsQueryParamAPISecSourceInternal,
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
            "22554b3a-d14f-42dc-b8d0-c3530e8f8d65",
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
            "5a331cfa-e207-49d9-a176-e260318ece7d",
        },
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

    if res.Stream != nil {
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
            "238f2259-a31a-4eed-8f78-79faf6121c42",
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
            "e552767e-0350-4925-b7e4-39731700805c",
        },
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
            "f32607b3-fc0f-4f77-9432-186ec778c011",
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
        KubernetesClusterID: "fcc6ab38-d1d9-486d-93c4-52d73dc4b7aa",
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
        KubernetesClusterID: "6f0ce2c1-9ce6-424f-84d6-ad78557fec29",
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
        KubernetesClusterID: "6c0046f1-672f-410e-9032-dde3c7f70879",
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

