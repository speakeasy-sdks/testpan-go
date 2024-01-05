# K8sCisBenchmark
(*K8sCisBenchmark*)

## Overview

APIs to get the kubernetes cis benchmark data

### Available Operations

* [GetK8sCISBenchmark](#getk8scisbenchmark) - Get k8s cis benchmark for clusters
* [GetK8sCISBenchmarkSummary](#getk8scisbenchmarksummary) - Get k8s cis benchmark summary of account
* [GetK8sCISBenchmarkClusterID](#getk8scisbenchmarkclusterid) - Get k8s cis benchmark for a specific cluster
* [PostK8sCISBenchmarkClusterID](#postk8scisbenchmarkclusterid) - initiate k8s cis benchmark scan for a specific cluster
* [PutK8sCISBenchmarkClusterID](#putk8scisbenchmarkclusterid) - edit k8s cis benchmark for a specific cluster with test statuses

## GetK8sCISBenchmark

Get k8s cis benchmark for clusters

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
    res, err := s.K8sCisBenchmark.GetK8sCISBenchmark(ctx, operations.GetK8sCISBenchmarkRequest{
        ClusterIds: []string{
            "string",
        },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetK8sCISBenchmarkRequest](../../pkg/models/operations/getk8scisbenchmarkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetK8sCISBenchmarkResponse](../../pkg/models/operations/getk8scisbenchmarkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetK8sCISBenchmarkSummary

Get k8s cis benchmark summary of account

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.K8sCisBenchmark.GetK8sCISBenchmarkSummary(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.K8sCISBenchmarkAccountSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetK8sCISBenchmarkSummaryResponse](../../pkg/models/operations/getk8scisbenchmarksummaryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetK8sCISBenchmarkClusterID

Get k8s cis benchmark for a specific cluster

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.K8sCisBenchmark.GetK8sCISBenchmarkClusterID(ctx, operations.GetK8sCISBenchmarkClusterIDRequest{
        ClusterID: "22b6e2e3-3165-4cce-b4e8-bc9db9f52c45",
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
| `request`                                                                                                          | [operations.GetK8sCISBenchmarkClusterIDRequest](../../pkg/models/operations/getk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetK8sCISBenchmarkClusterIDResponse](../../pkg/models/operations/getk8scisbenchmarkclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostK8sCISBenchmarkClusterID

initiate k8s cis benchmark scan for a specific cluster

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
    res, err := s.K8sCisBenchmark.PostK8sCISBenchmarkClusterID(ctx, operations.PostK8sCISBenchmarkClusterIDRequest{
        ClusterID: "35b87c47-9c94-446b-bd9d-9edabb3c15c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.K8sCISBenchmarkClustersSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostK8sCISBenchmarkClusterIDRequest](../../pkg/models/operations/postk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostK8sCISBenchmarkClusterIDResponse](../../pkg/models/operations/postk8scisbenchmarkclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutK8sCISBenchmarkClusterID

edit k8s cis benchmark for a specific cluster with test statuses

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.K8sCisBenchmark.PutK8sCISBenchmarkClusterID(ctx, operations.PutK8sCISBenchmarkClusterIDRequest{
        K8sCISBenchmarkUpdateNodes: shared.K8sCISBenchmarkUpdateNodes{
            ClusterID: "49b403e0-ddc8-4f07-8f6d-7b4806fe4a7b",
            Index: "string",
            Nodes: []shared.K8sCISBenchmarkUpdateNode{
                shared.K8sCISBenchmarkUpdateNode{},
            },
            Status: shared.K8sCISBenchmarkUpdateNodeStatusPass,
        },
        ClusterID: "e3c739e9-05f1-44c9-b8f5-5f4602b586f1",
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
| `request`                                                                                                          | [operations.PutK8sCISBenchmarkClusterIDRequest](../../pkg/models/operations/putk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutK8sCISBenchmarkClusterIDResponse](../../pkg/models/operations/putk8scisbenchmarkclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
