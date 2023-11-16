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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteExpansionsExpansionIDRequest](../../pkg/models/operations/deleteexpansionsexpansionidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteExpansionsExpansionIDResponse](../../pkg/models/operations/deleteexpansionsexpansionidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetExpansions

List all the expansions on the system

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Expansions.GetExpansions(ctx, operations.GetExpansionsRequest{
        SortKey: operations.GetExpansionsQueryParamSortKeyName,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetExpansionsRequest](../../pkg/models/operations/getexpansionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetExpansionsResponse](../../pkg/models/operations/getexpansionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetExpansionsExpansionIDInstallExpansionTarGz

Get the expansion installation

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

    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.GetExpansionsExpansionIDInstallExpansionTarGzRequest](../../pkg/models/operations/getexpansionsexpansionidinstallexpansiontargzrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.GetExpansionsExpansionIDInstallExpansionTarGzResponse](../../pkg/models/operations/getexpansionsexpansionidinstallexpansiontargzresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostExpansions

Create a new expansion

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
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
        Labels: []shared.Label{
            shared.Label{
                Key: "<key>",
                Value: "string",
            },
        },
        Name: "string",
        NamespaceID: "ea7db024-5fda-41fa-8ab8-8259ba11df19",
        WorkloadAddresses: []shared.WorkloadAddress{
            shared.WorkloadAddress{
                Address: "547 Dean Brook",
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.ExpansionInput](../../pkg/models/shared/expansioninput.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostExpansionsResponse](../../pkg/models/operations/postexpansionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutExpansionsExpansionID

Edit expansion definition

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
                    Value: "string",
                },
            },
            Name: "string",
            WorkloadAddresses: []shared.WorkloadAddress{
                shared.WorkloadAddress{
                    Address: "218 Jenkins Gateway",
                },
            },
        },
        ExpansionID: "159739a0-b94d-4cfc-a4ba-a7b3c4b09bff",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutExpansionsExpansionIDRequest](../../pkg/models/operations/putexpansionsexpansionidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutExpansionsExpansionIDResponse](../../pkg/models/operations/putexpansionsexpansionidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
