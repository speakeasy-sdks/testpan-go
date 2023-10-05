# Expansions
(*Expansions*)

## Overview

APIs used to manage expansions on Secure Application

### Available Operations

* [DeleteExpansionsExpansionID](#deleteexpansionsexpansionid) - Delete an expansion
* [GetExpansions](#getexpansions) - List all the expansions on the system
* [GetExpansionsExpansionIDInstallExpansionTarGz](#getexpansionsexpansionidinstallexpansiontargz) - Get the expansion installation
* [PostExpansions](#postexpansions) - Create a new expansion
* [PutExpansionsExpansionID](#putexpansionsexpansionid) - Edit expansion definition

## DeleteExpansionsExpansionID

Delete an expansion

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
    res, err := s.Expansions.DeleteExpansionsExpansionID(ctx, operations.DeleteExpansionsExpansionIDRequest{
        ExpansionID: "f85d88d8-8509-415c-acc6-dcdc8448694e",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteExpansionsExpansionIDRequest](../../models/operations/deleteexpansionsexpansionidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteExpansionsExpansionIDResponse](../../models/operations/deleteexpansionsexpansionidresponse.md), error**


## GetExpansions

List all the expansions on the system

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
    res, err := s.Expansions.GetExpansions(ctx, operations.GetExpansionsRequest{
        ClusterName: testpango.String("Agent interface"),
        ControllerStatus: testpango.String("Funk Flerovium Road"),
        ControllerVersion: testpango.String("Quality"),
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(2748.69),
        Name: testpango.String("male payment Kentucky"),
        NamespaceName: testpango.String("Toyota"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(4144.14),
        SortDir: operations.GetExpansionsSortDirDesc.ToPointer(),
        SortKey: operations.GetExpansionsSortKeyName,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Expansions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetExpansionsRequest](../../models/operations/getexpansionsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetExpansionsResponse](../../models/operations/getexpansionsresponse.md), error**


## GetExpansionsExpansionIDInstallExpansionTarGz

Get the expansion installation

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
    res, err := s.Expansions.GetExpansionsExpansionIDInstallExpansionTarGz(ctx, operations.GetExpansionsExpansionIDInstallExpansionTarGzRequest{
        ExpansionID: "f93ee138-484b-4ae8-9598-dc7e98557c98",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetExpansionsExpansionIDInstallExpansionTarGz200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetExpansionsExpansionIDInstallExpansionTarGzRequest](../../models/operations/getexpansionsexpansionidinstallexpansiontargzrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.GetExpansionsExpansionIDInstallExpansionTarGzResponse](../../models/operations/getexpansionsexpansionidinstallexpansiontargzresponse.md), error**


## PostExpansions

Create a new expansion

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
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
    res, err := s.Expansions.PostExpansions(ctx, shared.ExpansionInput{
        ClusterID: "ef536384-aae2-4f5a-87a4-cef022a42548",
        ControllerLastActive: types.MustTimeFromString("2023-08-24T22:57:12.780Z"),
        Labels: []shared.Label{
            shared.Label{
                Key: "<key>",
                Value: "bolívar North Frozen",
            },
        },
        Name: "postmark",
        NamespaceID: "cab88259-ba11-4df1-9d86-c530fc31b7b3",
        ShouldSendMetrics: testpango.Bool(false),
        WorkloadAddresses: []shared.WorkloadAddress{
            shared.WorkloadAddress{
                Address: "23391 Padberg Keys",
                NetworkProtocol: shared.NetworkProtocolHTTP.ToPointer(),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Expansion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.ExpansionInput](../../models/shared/expansioninput.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.PostExpansionsResponse](../../models/operations/postexpansionsresponse.md), error**


## PutExpansionsExpansionID

Edit expansion definition

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
    res, err := s.Expansions.PutExpansionsExpansionID(ctx, operations.PutExpansionsExpansionIDRequest{
        ExpansionPut: shared.ExpansionPut{
            Labels: []shared.Label{
                shared.Label{
                    Key: "<key>",
                    Value: "solemnly Money logistical",
                },
            },
            Name: "Kids Program",
            WorkloadAddresses: []shared.WorkloadAddress{
                shared.WorkloadAddress{
                    Address: "628 Will Route",
                    NetworkProtocol: shared.NetworkProtocolTCP.ToPointer(),
                },
            },
        },
        ExpansionID: "4baa7b3c-4b09-4bff-8d17-0310c446aa3a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Expansion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutExpansionsExpansionIDRequest](../../models/operations/putexpansionsexpansionidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutExpansionsExpansionIDResponse](../../models/operations/putexpansionsexpansionidresponse.md), error**

