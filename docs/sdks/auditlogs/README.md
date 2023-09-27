# AuditLogs
(*AuditLogs*)

## Overview

APIs used to retrieve  audit logs

### Available Operations

* [GetAuditLogs](#getauditlogs) - Get audit logs
* [GetAuditLogsActions](#getauditlogsactions) - Get all the audit logs actions
* [GetAuditLogsKubernetes](#getauditlogskubernetes) - Get audit logs
* [GetAuditLogsKubernetesActions](#getauditlogskubernetesactions) - Get all the kubernetes audit logs actions

## GetAuditLogs

Get audit logs

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
    res, err := s.AuditLogs.GetAuditLogs(ctx, operations.GetAuditLogsRequest{
        Actions: []string{
            "pariatur",
        },
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2022-04-22T10:10:54.165Z"),
        MaxResults: testpango.Float64(8310.31),
        ObjectType: testpango.String("perferendis"),
        Offset: testpango.Float64(6605.36),
        SortDir: operations.GetAuditLogsSortDirDesc.ToPointer(),
        SortKey: operations.GetAuditLogsSortKeyTime.ToPointer(),
        StartTime: types.MustTimeFromString("2020-01-06T07:14:49.176Z"),
        User: testpango.String("repudiandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuditLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetAuditLogsRequest](../../models/operations/getauditlogsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetAuditLogsResponse](../../models/operations/getauditlogsresponse.md), error**


## GetAuditLogsActions

Get all the audit logs actions

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
    res, err := s.AuditLogs.GetAuditLogsActions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAuditLogsActions200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAuditLogsActionsResponse](../../models/operations/getauditlogsactionsresponse.md), error**


## GetAuditLogsKubernetes

Get audit logs

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
    res, err := s.AuditLogs.GetAuditLogsKubernetes(ctx, operations.GetAuditLogsKubernetesRequest{
        ClusterName: testpango.String("architecto"),
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2022-02-18T12:04:35.830Z"),
        KubernetesAuditAction: []string{
            "harum",
        },
        MaxResults: testpango.Float64(2942.66),
        NamespaceName: testpango.String("voluptatibus"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(4353.53),
        ResourceKind: testpango.String("explicabo"),
        ResourceName: testpango.String("minus"),
        Result: []GetAuditLogsKubernetesResult{
            operations.GetAuditLogsKubernetesResultBlock,
        },
        SortDir: operations.GetAuditLogsKubernetesSortDirDesc.ToPointer(),
        SortKey: operations.GetAuditLogsKubernetesSortKeyLastSeen.ToPointer(),
        StartTime: types.MustTimeFromString("2021-06-21T12:40:08.193Z"),
        User: testpango.String("error"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesAuditLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAuditLogsKubernetesRequest](../../models/operations/getauditlogskubernetesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetAuditLogsKubernetesResponse](../../models/operations/getauditlogskubernetesresponse.md), error**


## GetAuditLogsKubernetesActions

Get all the kubernetes audit logs actions

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
    res, err := s.AuditLogs.GetAuditLogsKubernetesActions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAuditLogsKubernetesActions200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAuditLogsKubernetesActionsResponse](../../models/operations/getauditlogskubernetesactionsresponse.md), error**

