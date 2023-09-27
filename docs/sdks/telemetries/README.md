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
            "impedit",
        },
        AppType: []string{
            "amet",
        },
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2022-08-01T20:59:30.133Z"),
        EnvironmentName: []string{
            "neque",
        },
        Executable: []string{
            "enim",
        },
        HideInternals: testpango.Bool(false),
        HighestDockerfileScanResult: []string{
            "eaque",
        },
        Host: []string{
            "officia",
        },
        ImagesID: []string{
            "46714378-9ce0-4e99-9594-d93a74c0252f",
        },
        IsIdentified: testpango.Bool(false),
        MaxResults: testpango.Float64(9050.46),
        NamespacesFilter: testpango.String("dolor"),
        Offset: testpango.Float64(6946.47),
        ProtectionStatus: operations.GetAppTelemetriesProtectionStatusDeploymentOnly.ToPointer(),
        Result: []GetAppTelemetriesResult{
            operations.GetAppTelemetriesResultBlock,
        },
        ShowOnlyEntriesWithAppName: testpango.Bool(false),
        ShowOnlyViolations: testpango.Bool(false),
        ShowSystemPods: testpango.Bool(false),
        SortDir: operations.GetAppTelemetriesSortDirAsc.ToPointer(),
        SortKey: operations.GetAppTelemetriesSortKeyFinishTime,
        StartTime: types.MustTimeFromString("2021-12-03T05:15:36.959Z"),
        Status: []string{
            "tempore",
        },
        VulnerabilityLevelsFilter: []string{
            "odio",
        },
        WorkloadRisks: []GetAppTelemetriesWorkloadRisks{
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
        AppTelemetryID: "8ebb6e1d-2cf5-402b-afb2-cbc4635d5e65",
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
        AppTelemetryID: "da028c3e-951a-41e3-8fda-966489d7b786",
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
        AppTelemetryID: "73e13a12-a6b9-4924-9459-4487f5c84383",
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
        AppTelemetryID: "6b86b3cd-f641-45b0-849f-9df13f4eedbe",
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
        EndTime: types.MustTimeFromString("2022-06-20T02:15:58.382Z"),
        LogicalOperator: operations.GetConnectionTelemetriesLogicalOperatorOr.ToPointer(),
        MaxResults: testpango.Float64(9709.57),
        Offset: testpango.Float64(4133.85),
        Result: []GetConnectionTelemetriesResult{
            operations.GetConnectionTelemetriesResultDetect,
        },
        ShowOnlyViolations: testpango.Bool(false),
        SortDir: operations.GetConnectionTelemetriesSortDirAsc.ToPointer(),
        SortKey: operations.GetConnectionTelemetriesSortKeyTargetEnvironmentName,
        SourceAppName: []string{
            "quia",
        },
        SourceEnvironmentName: []string{
            "ipsam",
        },
        SourceExecutable: []string{
            "rem",
        },
        SourceHostName: []string{
            "molestias",
        },
        SourceNamespacesFilter: testpango.String("eius"),
        SourceRisk: []GetConnectionTelemetriesSourceRisk{
            operations.GetConnectionTelemetriesSourceRiskLow,
        },
        StartTime: types.MustTimeFromString("2022-02-06T23:07:05.850Z"),
        TargetAppName: []string{
            "aliquid",
        },
        TargetEnvironmentName: []string{
            "amet",
        },
        TargetExecutable: []string{
            "fugiat",
        },
        TargetHostName: []string{
            "corporis",
        },
        TargetNamespacesFilter: testpango.String("impedit"),
        TargetRisk: []GetConnectionTelemetriesTargetRisk{
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
        ConnectionTelemetryID: "2795b785-148d-46d5-89e5-635b33bc0f97",
        EndTime: types.MustTimeFromString("2022-03-10T11:26:33.471Z"),
        StartTime: types.MustTimeFromString("2022-10-31T13:05:54.663Z"),
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

