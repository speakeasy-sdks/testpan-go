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
        EndTime: types.MustTimeFromString("2021-09-05T15:33:12.217Z"),
        EnvironmentName: []string{
            "Syrian",
        },
        Executable: []string{
            "strategy",
        },
        HighestDockerfileScanResult: []string{
            "Concrete",
        },
        Host: []string{
            "Rupee",
        },
        ImagesID: []string{
            "356669e3-0a7d-4fca-9c8c-97d1dbbd6441",
        },
        Result: []operations.GetAppTelemetriesResult{
            operations.GetAppTelemetriesResultDetect,
        },
        SortKey: operations.GetAppTelemetriesSortKeyRisk,
        StartTime: types.MustTimeFromString("2021-01-15T02:52:59.625Z"),
        Status: []string{
            "Avon",
        },
        VulnerabilityLevelsFilter: []string{
            "audition female",
        },
        WorkloadRisks: []operations.GetAppTelemetriesWorkloadRisks{
            operations.GetAppTelemetriesWorkloadRisksMedium,
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
        EndTime: types.MustTimeFromString("2023-12-02T06:47:12.551Z"),
        Result: []operations.GetConnectionTelemetriesResult{
            operations.GetConnectionTelemetriesResultAllow,
        },
        SortKey: operations.GetConnectionTelemetriesSortKeySourceAppType,
        SourceAppName: []string{
            "because",
        },
        SourceEnvironmentName: []string{
            "Cotton",
        },
        SourceExecutable: []string{
            "Hybrid",
        },
        SourceHostName: []string{
            "wireless",
        },
        SourceRisk: []operations.GetConnectionTelemetriesSourceRisk{
            operations.GetConnectionTelemetriesSourceRiskHigh,
        },
        StartTime: types.MustTimeFromString("2023-12-19T20:37:58.933Z"),
        TargetAppName: []string{
            "Moscovium",
        },
        TargetEnvironmentName: []string{
            "Regional",
        },
        TargetExecutable: []string{
            "coddle",
        },
        TargetHostName: []string{
            "Hampshire",
        },
        TargetRisk: []operations.GetConnectionTelemetriesTargetRisk{
            operations.GetConnectionTelemetriesTargetRiskMedium,
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

