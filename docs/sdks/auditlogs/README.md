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
        EndTime: types.MustTimeFromString("2022-01-24T02:40:28.078Z"),
        StartTime: types.MustTimeFromString("2022-09-08T07:41:20.287Z"),
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
        EndTime: types.MustTimeFromString("2023-03-24T20:20:48.235Z"),
        KubernetesAuditAction: []string{
            "Coupe",
        },
        Result: []operations.GetAuditLogsKubernetesResult{
            operations.GetAuditLogsKubernetesResultAllow,
        },
        StartTime: types.MustTimeFromString("2023-06-24T03:55:36.271Z"),
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

