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
            "male",
        },
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2022-01-24T02:40:28.078Z"),
        MaxResults: testpango.Float64(5619.36),
        ObjectType: testpango.String("GB grey Folding"),
        Offset: testpango.Float64(2262.25),
        SortDir: operations.GetAuditLogsSortDirDesc.ToPointer(),
        SortKey: operations.GetAuditLogsSortKeyTime.ToPointer(),
        StartTime: types.MustTimeFromString("2022-05-15T07:17:17.108Z"),
        User: testpango.String("Lennie.Torphy0"),
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
        ClusterName: testpango.String("Coupe Southeast Southeast"),
        DownloadAsXlsx: testpango.Bool(false),
        EndTime: types.MustTimeFromString("2022-08-15T09:50:37.531Z"),
        KubernetesAuditAction: []string{
            "online",
        },
        MaxResults: testpango.Float64(3231.79),
        NamespaceName: testpango.String("Cadillac"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(5266.56),
        ResourceKind: testpango.String("Island SSL"),
        ResourceName: testpango.String("South vaguely"),
        Result: []GetAuditLogsKubernetesResult{
            operations.GetAuditLogsKubernetesResultBlock,
        },
        SortDir: operations.GetAuditLogsKubernetesSortDirDesc.ToPointer(),
        SortKey: operations.GetAuditLogsKubernetesSortKeyAction.ToPointer(),
        StartTime: types.MustTimeFromString("2022-02-10T19:51:08.510Z"),
        User: testpango.String("Rebekah66"),
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

