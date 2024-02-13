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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.DeleteAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        CatalogID: "3d2f6c1c-29ba-4502-9cb8-4666b2e993b7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionRequest](../../pkg/models/operations/deleteapisecurityinternalcatalogcatalogidbfladetectionrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |


### Response

**[*operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionResponse](../../pkg/models/operations/deleteapisecurityinternalcatalogcatalogidbfladetectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAPISecurityInternalCatalogCatalogIDBflaLearning

stop bfla learning phase

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.DeleteAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        CatalogID: "3ba770da-4891-4357-a360-0c0e060f5b0f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest](../../pkg/models/operations/deleteapisecurityinternalcatalogcatalogidbflalearningrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningResponse](../../pkg/models/operations/deleteapisecurityinternalcatalogcatalogidbflalearningresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPISecurityInternalCatalogCatalogIDBfla

Get bfla info for given catalogId

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetAPISecurityInternalCatalogCatalogIDBflaRequest](../../pkg/models/operations/getapisecurityinternalcatalogcatalogidbflarequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDBflaResponse](../../pkg/models/operations/getapisecurityinternalcatalogcatalogidbflaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPISecurityInternalCatalogCatalogIDBflaDetection

Start new bfla detection phase

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "string",
        },
        CatalogID: "35e94464-b0ff-40e6-be14-ca29162fc277",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionRequest](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbfladetectionrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionResponse](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbfladetectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPISecurityInternalCatalogCatalogIDBflaLearning

Start new bfla learning phase

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "string",
        },
        CatalogID: "08fbc82f-2c9e-4a85-a326-b52d888f26e1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningRequest](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbflalearningrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningResponse](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbflalearningresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPISecurityInternalCatalogCatalogIDBflaReset

Reset bfla model

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.PostAPISecurityInternalCatalogCatalogIDBflaReset(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaResetRequest{
        CatalogID: "2682814f-440b-4991-ae68-e4a1209728f3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.PostAPISecurityInternalCatalogCatalogIDBflaResetRequest](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbflaresetrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDBflaResetResponse](../../pkg/models/operations/postapisecurityinternalcatalogcatalogidbflaresetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutAPISecurityInternalCatalogCatalogIDBfla

update BFLA info for this catalogId

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Bfla.PutAPISecurityInternalCatalogCatalogIDBfla(ctx, operations.PutAPISecurityInternalCatalogCatalogIDBflaRequest{
        APIServiceBflaInfo: shared.APIServiceBflaInfo{
            Status: shared.APIServiceBflaInfoStatusNoSpec,
        },
        CatalogID: "aab64aef-3b3f-4caa-9d44-2012ad1459f8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.PutAPISecurityInternalCatalogCatalogIDBflaRequest](../../pkg/models/operations/putapisecurityinternalcatalogcatalogidbflarequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.PutAPISecurityInternalCatalogCatalogIDBflaResponse](../../pkg/models/operations/putapisecurityinternalcatalogcatalogidbflaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
