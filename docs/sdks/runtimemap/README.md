# RuntimeMap
(*RuntimeMap*)

## Overview

APIs used to  query for network map

### Available Operations

* [DeleteNetworkMapQueueRequestID](#deletenetworkmapqueuerequestid) - Cancel the network map background job
* [GetNetworkMap](#getnetworkmap) - Get data for network map
* [GetNetworkMapQueueRequestID](#getnetworkmapqueuerequestid) - Get status for network map background job
* [GetNetworkMapResultsRequestID](#getnetworkmapresultsrequestid) - Get result for network map background job

## DeleteNetworkMapQueueRequestID

Cancel the network map background job

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
    res, err := s.RuntimeMap.DeleteNetworkMapQueueRequestID(ctx, operations.DeleteNetworkMapQueueRequestIDRequest{
        RequestID: "9c87ae50-c166-461a-9d91-36a7e8d53213",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteNetworkMapQueueRequestIDRequest](../../models/operations/deletenetworkmapqueuerequestidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteNetworkMapQueueRequestIDResponse](../../models/operations/deletenetworkmapqueuerequestidresponse.md), error**


## GetNetworkMap

Get data for network map

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
    res, err := s.RuntimeMap.GetNetworkMap(ctx, operations.GetNetworkMapRequest{
        APIRisk: operations.GetNetworkMapAPIRiskUnknown.ToPointer(),
        Apps: []string{
            "velit",
        },
        EndTime: types.MustTimeFromString("2021-10-04T11:34:13.847Z"),
        Environments: []string{
            "voluptas",
        },
        ExcludeApps: []string{
            "quos",
        },
        GroupAppsOnTheSameEnvironment: testpango.Bool(false),
        IgnoreExternalConnection: testpango.Bool(false),
        IsBackgroundJob: testpango.Bool(false),
        Labels: []string{
            "esse",
        },
        Namespaces: []string{
            "52db764c-59f0-4a56-8ebc-ada29ca79181",
        },
        ShowOnlyAppsWithConnections: testpango.Bool(false),
        ShowOnlyAppsWithViolations: testpango.Bool(false),
        ShowOnlyConnectionsBetweenEnvironments: testpango.Bool(false),
        ShowOnlyConnectionsWithViolations: testpango.Bool(false),
        StartTime: types.MustTimeFromString("2021-04-13T03:05:21.776Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NetworkMap != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetNetworkMapRequest](../../models/operations/getnetworkmaprequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetNetworkMapResponse](../../models/operations/getnetworkmapresponse.md), error**


## GetNetworkMapQueueRequestID

Get status for network map background job

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
    res, err := s.RuntimeMap.GetNetworkMapQueueRequestID(ctx, operations.GetNetworkMapQueueRequestIDRequest{
        RequestID: "5671663c-530b-4566-9163-a3638512ab25",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BackgroundJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetNetworkMapQueueRequestIDRequest](../../models/operations/getnetworkmapqueuerequestidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetNetworkMapQueueRequestIDResponse](../../models/operations/getnetworkmapqueuerequestidresponse.md), error**


## GetNetworkMapResultsRequestID

Get result for network map background job

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
    res, err := s.RuntimeMap.GetNetworkMapResultsRequestID(ctx, operations.GetNetworkMapResultsRequestIDRequest{
        RequestID: "21b9f2e0-7246-47b8-a40b-c05fab0d650e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NetworkMaps != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetNetworkMapResultsRequestIDRequest](../../models/operations/getnetworkmapresultsrequestidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetNetworkMapResultsRequestIDResponse](../../models/operations/getnetworkmapresultsrequestidresponse.md), error**

