# Bfla
(*Bfla*)

### Available Operations

* [DeleteAPISecurityInternalCatalogCatalogIDBflaDetection](#deleteapisecurityinternalcatalogcatalogidbfladetection) - stop bfla detection phase
* [DeleteAPISecurityInternalCatalogCatalogIDBflaLearning](#deleteapisecurityinternalcatalogcatalogidbflalearning) - stop bfla learning phase
* [GetAPISecurityInternalCatalogCatalogIDBfla](#getapisecurityinternalcatalogcatalogidbfla) - Get bfla info for given catalogId
* [PostAPISecurityInternalCatalogCatalogIDBflaDetection](#postapisecurityinternalcatalogcatalogidbfladetection) - Start new bfla detection phase
* [PostAPISecurityInternalCatalogCatalogIDBflaLearning](#postapisecurityinternalcatalogcatalogidbflalearning) - Start new bfla learning phase
* [PostAPISecurityInternalCatalogCatalogIDBflaReset](#postapisecurityinternalcatalogcatalogidbflareset) - Reset bfla model
* [PutAPISecurityInternalCatalogCatalogIDBfla](#putapisecurityinternalcatalogcatalogidbfla) - update BFLA info for this catalogId

## DeleteAPISecurityInternalCatalogCatalogIDBflaDetection

stop bfla detection phase

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
    res, err := s.Bfla.DeleteAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        CatalogID: "82f168a3-63c8-4873-a484-380b1f6b8ca2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAPISecurityInternalCatalogCatalogIDBflaDetection204ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionRequest](../../models/operations/deleteapisecurityinternalcatalogcatalogidbfladetectionrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionResponse](../../models/operations/deleteapisecurityinternalcatalogcatalogidbfladetectionresponse.md), error**


## DeleteAPISecurityInternalCatalogCatalogIDBflaLearning

stop bfla learning phase

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
    res, err := s.Bfla.DeleteAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        CatalogID: "75a60a04-c495-4cc6-9917-1b51c1bdb1cf",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAPISecurityInternalCatalogCatalogIDBflaLearning204ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest](../../models/operations/deleteapisecurityinternalcatalogcatalogidbflalearningrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |


### Response

**[*operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningResponse](../../models/operations/deleteapisecurityinternalcatalogcatalogidbflalearningresponse.md), error**


## GetAPISecurityInternalCatalogCatalogIDBfla

Get bfla info for given catalogId

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
    res, err := s.Bfla.GetAPISecurityInternalCatalogCatalogIDBfla(ctx, operations.GetAPISecurityInternalCatalogCatalogIDBflaRequest{
        CatalogID: "4b888ebd-fc4c-4cca-99bc-7fc0b2dce108",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceBflaInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetAPISecurityInternalCatalogCatalogIDBflaRequest](../../models/operations/getapisecurityinternalcatalogcatalogidbflarequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDBflaResponse](../../models/operations/getapisecurityinternalcatalogcatalogidbflaresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDBflaDetection

Start new bfla detection phase

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
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "iusto",
        },
        CatalogID: "3e42b006-d678-4878-ba85-81a58208c54f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPISecurityInternalCatalogCatalogIDBflaDetection201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionRequest](../../models/operations/postapisecurityinternalcatalogcatalogidbfladetectionrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionResponse](../../models/operations/postapisecurityinternalcatalogcatalogidbfladetectionresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDBflaLearning

Start new bfla learning phase

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
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "saepe",
        },
        CatalogID: "fa9c95f2-eac5-4565-9307-cfee81206e28",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPISecurityInternalCatalogCatalogIDBflaLearning201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningRequest](../../models/operations/postapisecurityinternalcatalogcatalogidbflalearningrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningResponse](../../models/operations/postapisecurityinternalcatalogcatalogidbflalearningresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDBflaReset

Reset bfla model

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
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaReset(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaResetRequest{
        CatalogID: "13fa4a41-c480-4d3f-a132-af03102d514f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPISecurityInternalCatalogCatalogIDBflaReset201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.PostAPISecurityInternalCatalogCatalogIDBflaResetRequest](../../models/operations/postapisecurityinternalcatalogcatalogidbflaresetrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaResetResponse](../../models/operations/postapisecurityinternalcatalogcatalogidbflaresetresponse.md), error**


## PutAPISecurityInternalCatalogCatalogIDBfla

update BFLA info for this catalogId

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
    res, err := s.Bfla.PutAPISecurityInternalCatalogCatalogIDBfla(ctx, operations.PutAPISecurityInternalCatalogCatalogIDBflaRequest{
        APIServiceBflaInfo: shared.APIServiceBflaInfo{
            EndTime: types.MustTimeFromString("2022-03-29T11:57:17.823Z"),
            Status: shared.APIServiceBflaInfoStatusInProgressDetection,
            Tags: []shared.APIServiceBflaTagInfo{
                shared.APIServiceBflaTagInfo{
                    IsLegitimate: testpango.Bool(false),
                    Name: "Kristie Bergnaum",
                    Paths: []shared.APIServiceBflaPathInfo{
                        shared.APIServiceBflaPathInfo{
                            Clients: []shared.APIServiceBflaClientInfo{
                                shared.APIServiceBflaClientInfo{
                                    External: testpango.Bool(false),
                                    Identifier: testpango.String("f9621a6a-4f77-4a87-ae3e-4be752c65b34"),
                                    IsLegitimate: testpango.Bool(false),
                                    LastObserved: types.MustTimeFromString("2022-11-24T01:03:51.181Z"),
                                    LastStatusCode: testpango.Int64(506966),
                                    Name: "Rodney Quitzon",
                                    Namespace: testpango.String("illo"),
                                    Principles: []shared.APIServiceBflaPrincipleInfo{
                                        shared.APIServiceBflaPrincipleInfo{
                                            IP: "minus",
                                            Name: "Gustavo Mayer",
                                            PrincipleType: "accusamus",
                                        },
                                    },
                                },
                            },
                            IsLegitimate: testpango.Bool(false),
                            Method: shared.HTTPMethodGet,
                            Path: "voluptates",
                        },
                    },
                },
            },
        },
        CatalogID: "8419d8f8-4f14-44f3-a07e-dcc4aa5f3cab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutAPISecurityInternalCatalogCatalogIDBfla200ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutAPISecurityInternalCatalogCatalogIDBflaRequest](../../models/operations/putapisecurityinternalcatalogcatalogidbflarequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.PutAPISecurityInternalCatalogCatalogIDBflaResponse](../../models/operations/putapisecurityinternalcatalogcatalogidbflaresponse.md), error**

