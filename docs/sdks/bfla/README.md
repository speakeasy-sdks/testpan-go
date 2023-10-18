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
        CatalogID: "3d2f6c1c-29ba-4502-9cb8-4666b2e993b7",
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
        CatalogID: "3ba770da-4891-4357-a360-0c0e060f5b0f",
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
        CatalogID: "74e0c81f-ef10-4f23-b74f-8447d1953ccf",
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
            Duration: "Garden",
        },
        CatalogID: "e94464b0-ff0e-46fe-94ca-29162fc27770",
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
            Duration: "South",
        },
        CatalogID: "fbc82f2c-9ea8-45a3-a6b5-2d888f26e15a",
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
        CatalogID: "2682814f-440b-4991-ae68-e4a1209728f3",
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
            Status: shared.APIServiceBflaInfoStatusNoSpec,
            Tags: []shared.APIServiceBflaTagInfo{
                shared.APIServiceBflaTagInfo{
                    Name: "International",
                    Paths: []shared.APIServiceBflaPathInfo{
                        shared.APIServiceBflaPathInfo{
                            Clients: []shared.APIServiceBflaClientInfo{
                                shared.APIServiceBflaClientInfo{
                                    Name: "Human",
                                    Principles: []shared.APIServiceBflaPrincipleInfo{
                                        shared.APIServiceBflaPrincipleInfo{
                                            IP: "252.50.185.52",
                                            Name: "sherbet",
                                            PrincipleType: "Technician",
                                        },
                                    },
                                },
                            },
                            Method: shared.HTTPMethodTrace,
                            Path: "/usr/sbin",
                        },
                    },
                },
            },
        },
        CatalogID: "442012ad-1459-4f87-a9c2-0110831cf060",
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

