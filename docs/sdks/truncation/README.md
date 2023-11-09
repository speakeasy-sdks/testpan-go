# Truncation
(*Truncation*)

## Overview

APIs to delete workloads

### Available Operations

* [GetTruncationImages](#gettruncationimages) - Get workloads truncation time for account
* [GetTruncationWorkloads](#gettruncationworkloads) - Get workloads truncation time for account
* [PostTruncationImages](#posttruncationimages) - Update workloads truncation status for account
* [PostTruncationWorkloads](#posttruncationworkloads) - Update workloads truncation status for account

## GetTruncationImages

Get workloads truncation time for account

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
    res, err := s.Truncation.GetTruncationImages(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TruncationStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTruncationImagesResponse](../../pkg/models/operations/gettruncationimagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTruncationWorkloads

Get workloads truncation time for account

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
    res, err := s.Truncation.GetTruncationWorkloads(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TruncationStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTruncationWorkloadsResponse](../../pkg/models/operations/gettruncationworkloadsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostTruncationImages

Update workloads truncation status for account

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
    res, err := s.Truncation.PostTruncationImages(ctx, shared.TruncationStatus{
        IsTruncationEnabled: false,
        TruncateTimeInDays: 271429,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TruncationStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.TruncationStatus](../../pkg/models/shared/truncationstatus.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostTruncationImagesResponse](../../pkg/models/operations/posttruncationimagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostTruncationWorkloads

Update workloads truncation status for account

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
    res, err := s.Truncation.PostTruncationWorkloads(ctx, shared.TruncationStatus{
        IsTruncationEnabled: false,
        TruncateTimeInDays: 519889,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TruncationStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.TruncationStatus](../../pkg/models/shared/truncationstatus.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostTruncationWorkloadsResponse](../../pkg/models/operations/posttruncationworkloadsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
