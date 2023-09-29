# Telemetries
(*Telemetries*)

## Overview

APIs used to query for telemetries

### Available Operations

* [GetAppTelemetries](#getapptelemetries) - Get App telemetries
* [GetAppTelemetriesAppTelemetryID](#getapptelemetriesapptelemetryid) - Get App telemetry by ID
* [GetAppTelemetriesAppTelemetryIDAPIRiskInfo](#getapptelemetriesapptelemetryidapiriskinfo) - Get API risks info of given app telemetry
* [GetAppTelemetriesAppTelemetryIDImagePackages](#getapptelemetriesapptelemetryidimagepackages) - list packages with licenses runnin on a pod
* [GetAppTelemetriesAppTelemetryIDInjectionInfo](#getapptelemetriesapptelemetryidinjectioninfo) - Get token injection info of given app telemetry
* [GetConnectionTelemetries](#getconnectiontelemetries) - Get connection telemetries
* [GetConnectionTelemetriesConnectionTelemetryID](#getconnectiontelemetriesconnectiontelemetryid) - get details for a single connection telemetry

## GetAppTelemetries

Get App telemetries

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetAppTelemetries(ctx, operations.GetAppTelemetriesRequest{
        AppName: []string{
            "Crew",
        },
        AppType: []string{
            "splendid",
        },
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2021-09-05T15:33:12.217Z"),
        EnvironmentName: []string{
            "Syrian",
        },
        Executable: []string{
            "strategy",
        },
        HideInternals: testpango.Bool(false),
        HighestDockerfileScanResult: []string{
            "Concrete",
        },
        Host: []string{
            "Rupee",
        },
        ImagesID: []string{
            "356669e3-0a7d-4fca-9c8c-97d1dbbd6441",
        },
        IsIdentified: testpango.Bool(false),
        MaxResults: testpango.Float64(3704.75),
        NamespacesFilter: testpango.String("North South"),
        Offset: testpango.Float64(9999.51),
        ProtectionStatus: operations.GetAppTelemetriesProtectionStatusFull.ToPointer(),
        Result: []GetAppTelemetriesResult{
            operations.GetAppTelemetriesResultBlock,
        },
        ShowOnlyEntriesWithAppName: testpango.Bool(false),
        ShowOnlyViolations: testpango.Bool(false),
        ShowSystemPods: testpango.Bool(false),
        SortDir: operations.GetAppTelemetriesSortDirAsc.ToPointer(),
        SortKey: operations.GetAppTelemetriesSortKeyEnvironmentName,
        StartTime: types.MustTimeFromString("2023-08-01T01:58:56.429Z"),
        Status: []string{
            "and",
        },
        VulnerabilityLevelsFilter: []string{
            "Benin Checking",
        },
        WorkloadRisks: []GetAppTelemetriesWorkloadRisks{
            operations.GetAppTelemetriesWorkloadRisksHigh,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppTelemetries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAppTelemetriesRequest](../../models/operations/getapptelemetriesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAppTelemetriesResponse](../../models/operations/getapptelemetriesresponse.md), error**


## GetAppTelemetriesAppTelemetryID

Get App telemetry by ID

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
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryID(ctx, operations.GetAppTelemetriesAppTelemetryIDRequest{
        AppTelemetryID: "6ebde0ec-647c-4eab-b779-88cc0a536506",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppTelemetry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetAppTelemetriesAppTelemetryIDRequest](../../models/operations/getapptelemetriesapptelemetryidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDResponse](../../models/operations/getapptelemetriesapptelemetryidresponse.md), error**


## GetAppTelemetriesAppTelemetryIDAPIRiskInfo

Get API risks info of given app telemetry

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
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDAPIRiskInfo(ctx, operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoRequest{
        AppTelemetryID: "6f98b895-9ca6-43b1-9526-7f6ebacd6bab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIRiskInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoRequest](../../models/operations/getapptelemetriesapptelemetryidapiriskinforequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoResponse](../../models/operations/getapptelemetriesapptelemetryidapiriskinforesponse.md), error**


## GetAppTelemetriesAppTelemetryIDImagePackages

list packages with licenses runnin on a pod

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
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDImagePackages(ctx, operations.GetAppTelemetriesAppTelemetryIDImagePackagesRequest{
        AppTelemetryID: "4668441f-326c-4b1d-9b80-38c3aa5ff203",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesWithLicenses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetAppTelemetriesAppTelemetryIDImagePackagesRequest](../../models/operations/getapptelemetriesapptelemetryidimagepackagesrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDImagePackagesResponse](../../models/operations/getapptelemetriesapptelemetryidimagepackagesresponse.md), error**


## GetAppTelemetriesAppTelemetryIDInjectionInfo

Get token injection info of given app telemetry

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
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDInjectionInfo(ctx, operations.GetAppTelemetriesAppTelemetryIDInjectionInfoRequest{
        AppTelemetryID: "e9382bba-032c-4379-a173-a7671d6d9a49",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TokenInjectionInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetAppTelemetriesAppTelemetryIDInjectionInfoRequest](../../models/operations/getapptelemetriesapptelemetryidinjectioninforequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDInjectionInfoResponse](../../models/operations/getapptelemetriesapptelemetryidinjectioninforesponse.md), error**


## GetConnectionTelemetries

Get connection telemetries

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetConnectionTelemetries(ctx, operations.GetConnectionTelemetriesRequest{
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2023-12-02T06:47:12.551Z"),
        LogicalOperator: operations.GetConnectionTelemetriesLogicalOperatorOr.ToPointer(),
        MaxResults: testpango.Float64(1093.57),
        Offset: testpango.Float64(9127.11),
        Result: []GetConnectionTelemetriesResult{
            operations.GetConnectionTelemetriesResultDetect,
        },
        ShowOnlyViolations: testpango.Bool(false),
        SortDir: operations.GetConnectionTelemetriesSortDirAsc.ToPointer(),
        SortKey: operations.GetConnectionTelemetriesSortKeyTargetAppType,
        SourceAppName: []string{
            "Hybrid",
        },
        SourceEnvironmentName: []string{
            "wireless",
        },
        SourceExecutable: []string{
            "Walks",
        },
        SourceHostName: []string{
            "circuit",
        },
        SourceNamespacesFilter: testpango.String("rich Coves"),
        SourceRisk: []GetConnectionTelemetriesSourceRisk{
            operations.GetConnectionTelemetriesSourceRiskLow,
        },
        StartTime: types.MustTimeFromString("2022-12-19T09:46:13.488Z"),
        TargetAppName: []string{
            "debunk",
        },
        TargetEnvironmentName: []string{
            "lysine",
        },
        TargetExecutable: []string{
            "athwart",
        },
        TargetHostName: []string{
            "female",
        },
        TargetNamespacesFilter: testpango.String("zero"),
        TargetRisk: []GetConnectionTelemetriesTargetRisk{
            operations.GetConnectionTelemetriesTargetRiskHigh,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionTelemetries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetConnectionTelemetriesRequest](../../models/operations/getconnectiontelemetriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetConnectionTelemetriesResponse](../../models/operations/getconnectiontelemetriesresponse.md), error**


## GetConnectionTelemetriesConnectionTelemetryID

get details for a single connection telemetry

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetConnectionTelemetriesConnectionTelemetryID(ctx, operations.GetConnectionTelemetriesConnectionTelemetryIDRequest{
        ConnectionTelemetryID: "726acf07-d01e-4a6d-9948-64b67b807f65",
        EndTime: types.MustTimeFromString("2021-06-17T09:20:57.797Z"),
        StartTime: types.MustTimeFromString("2023-04-26T13:53:39.546Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionTelemetry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetConnectionTelemetriesConnectionTelemetryIDRequest](../../models/operations/getconnectiontelemetriesconnectiontelemetryidrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.GetConnectionTelemetriesConnectionTelemetryIDResponse](../../models/operations/getconnectiontelemetriesconnectiontelemetryidresponse.md), error**

