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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetAppTelemetries(ctx, operations.GetAppTelemetriesRequest{
        AppName: []string{
            "string",
        },
        AppType: []string{
            "string",
        },
        EndTime: types.MustTimeFromString("2024-07-22T12:37:40.720Z"),
        EnvironmentName: []string{
            "string",
        },
        Executable: []string{
            "string",
        },
        HighestDockerfileScanResult: []string{
            "string",
        },
        Host: []string{
            "string",
        },
        ImagesID: []string{
            "50dc36c5-6d44-446a-b356-669e30a7dfca",
        },
        Result: []operations.GetAppTelemetriesQueryParamResult{
            operations.GetAppTelemetriesQueryParamResultDetect,
        },
        SortKey: operations.GetAppTelemetriesQueryParamSortKeyFinishTime,
        StartTime: types.MustTimeFromString("2023-08-17T23:31:28.434Z"),
        Status: []string{
            "string",
        },
        VulnerabilityLevelsFilter: []string{
            "string",
        },
        WorkloadRisks: []operations.WorkloadRisks{
            operations.WorkloadRisksCritical,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAppTelemetriesRequest](../../pkg/models/operations/getapptelemetriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetAppTelemetriesResponse](../../pkg/models/operations/getapptelemetriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppTelemetriesAppTelemetryID

Get App telemetry by ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetAppTelemetriesAppTelemetryIDRequest](../../pkg/models/operations/getapptelemetriesapptelemetryidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDResponse](../../pkg/models/operations/getapptelemetriesapptelemetryidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppTelemetriesAppTelemetryIDAPIRiskInfo

Get API risks info of given app telemetry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDAPIRiskInfo(ctx, operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoRequest{
        AppTelemetryID: "6f98b895-9ca6-43b1-9526-7f6ebacd6bab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoRequest](../../pkg/models/operations/getapptelemetriesapptelemetryidapiriskinforequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDAPIRiskInfoResponse](../../pkg/models/operations/getapptelemetriesapptelemetryidapiriskinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppTelemetriesAppTelemetryIDImagePackages

list packages with licenses runnin on a pod

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDImagePackages(ctx, operations.GetAppTelemetriesAppTelemetryIDImagePackagesRequest{
        AppTelemetryID: "4668441f-326c-4b1d-9b80-38c3aa5ff203",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetAppTelemetriesAppTelemetryIDImagePackagesRequest](../../pkg/models/operations/getapptelemetriesapptelemetryidimagepackagesrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDImagePackagesResponse](../../pkg/models/operations/getapptelemetriesapptelemetryidimagepackagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppTelemetriesAppTelemetryIDInjectionInfo

Get token injection info of given app telemetry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetAppTelemetriesAppTelemetryIDInjectionInfo(ctx, operations.GetAppTelemetriesAppTelemetryIDInjectionInfoRequest{
        AppTelemetryID: "e9382bba-032c-4379-a173-a7671d6d9a49",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetAppTelemetriesAppTelemetryIDInjectionInfoRequest](../../pkg/models/operations/getapptelemetriesapptelemetryidinjectioninforequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetAppTelemetriesAppTelemetryIDInjectionInfoResponse](../../pkg/models/operations/getapptelemetriesapptelemetryidinjectioninforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionTelemetries

Get connection telemetries

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetConnectionTelemetries(ctx, operations.GetConnectionTelemetriesRequest{
        EndTime: types.MustTimeFromString("2024-12-02T06:08:07.741Z"),
        Result: []operations.GetConnectionTelemetriesQueryParamResult{
            operations.GetConnectionTelemetriesQueryParamResultAllow,
        },
        SortKey: operations.GetConnectionTelemetriesQueryParamSortKeySourceAppType,
        SourceAppName: []string{
            "string",
        },
        SourceEnvironmentName: []string{
            "string",
        },
        SourceExecutable: []string{
            "string",
        },
        SourceHostName: []string{
            "string",
        },
        SourceRisk: []operations.SourceRisk{
            operations.SourceRiskLow,
        },
        StartTime: types.MustTimeFromString("2022-04-04T13:56:07.190Z"),
        TargetAppName: []string{
            "string",
        },
        TargetEnvironmentName: []string{
            "string",
        },
        TargetExecutable: []string{
            "string",
        },
        TargetHostName: []string{
            "string",
        },
        TargetRisk: []operations.TargetRisk{
            operations.TargetRiskHigh,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetConnectionTelemetriesRequest](../../pkg/models/operations/getconnectiontelemetriesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetConnectionTelemetriesResponse](../../pkg/models/operations/getconnectiontelemetriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionTelemetriesConnectionTelemetryID

get details for a single connection telemetry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Telemetries.GetConnectionTelemetriesConnectionTelemetryID(ctx, operations.GetConnectionTelemetriesConnectionTelemetryIDRequest{
        ConnectionTelemetryID: "726acf07-d01e-4a6d-9948-64b67b807f65",
        EndTime: types.MustTimeFromString("2022-06-17T13:01:05.521Z"),
        StartTime: types.MustTimeFromString("2024-04-26T08:25:39.198Z"),
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.GetConnectionTelemetriesConnectionTelemetryIDRequest](../../pkg/models/operations/getconnectiontelemetriesconnectiontelemetryidrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.GetConnectionTelemetriesConnectionTelemetryIDResponse](../../pkg/models/operations/getconnectiontelemetriesconnectiontelemetryidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
