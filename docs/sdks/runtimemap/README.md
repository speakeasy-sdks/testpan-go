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
        RequestID: "d744d9fc-99bf-4fa7-bd43-977bd77decee",
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
        Apps: []string{
            "string",
        },
        EndTime: types.MustTimeFromString("2023-01-07T14:07:43.013Z"),
        Environments: []string{
            "string",
        },
        ExcludeApps: []string{
            "string",
        },
        Labels: []string{
            "string",
        },
        Namespaces: []string{
            "fe7ca0fe-d770-4727-a0bc-8a727a6cf780",
        },
        StartTime: types.MustTimeFromString("2022-08-24T13:40:09.607Z"),
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
        RequestID: "3d22620c-025b-41f7-b070-9e069dc87bfd",
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
        RequestID: "78c61cf8-31c2-407a-a949-4700ad100770",
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

