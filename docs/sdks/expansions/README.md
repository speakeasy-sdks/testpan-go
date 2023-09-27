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
        ExpansionID: "b78f1562-6398-4a0d-8766-324ccb06c8ca",
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
        ClusterName: testpango.String("inventore"),
        ControllerStatus: testpango.String("consequuntur"),
        ControllerVersion: testpango.String("repellendus"),
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(271.97),
        Name: testpango.String("Jill Collins"),
        NamespaceName: testpango.String("odio"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(27.58),
        SortDir: operations.GetExpansionsSortDirDesc.ToPointer(),
        SortKey: "deleniti",
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
        ExpansionID: "d5722dd8-95b8-4bcf-a4db-959693352f74",
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
        ClusterID: "533994d7-8de3-4b6e-9389-f5abb7f66255",
        ControllerLastActive: types.MustTimeFromString("2022-05-02T06:02:30.053Z"),
        Labels: []shared.Label{
            shared.Label{
                Key: "qui",
                Value: "praesentium",
            },
        },
        Name: "Christy Deckow",
        NamespaceID: "483afd23-15bb-4a65-8164-e06f5bf6ae59",
        ShouldSendMetrics: testpango.Bool(false),
        WorkloadAddresses: []shared.WorkloadAddress{
            shared.WorkloadAddress{
                Address: "77578 Wunsch Extension",
                NetworkProtocol: shared.NetworkProtocolTCP.ToPointer(),
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
                    Key: "vitae",
                    Value: "fugit",
                },
            },
            Name: "Marc Doyle",
            WorkloadAddresses: []shared.WorkloadAddress{
                shared.WorkloadAddress{
                    Address: "39865 Alvina Lock",
                    NetworkProtocol: shared.NetworkProtocolTCP.ToPointer(),
                },
            },
        },
        ExpansionID: "4a68a9a3-5d08-46b6-b66f-ef020e9f443b",
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

