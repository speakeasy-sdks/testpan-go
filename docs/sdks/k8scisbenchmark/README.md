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
    res, err := s.K8sCisBenchmark.GetK8sCISBenchmark(ctx, operations.GetK8sCISBenchmarkRequest{
        ClusterIds: []string{
            "minus",
        },
        MaxResults: testpango.Float64(4196.69),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(4235.5),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.K8sCISBenchmarkClustersSummaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetK8sCISBenchmarkRequest](../../models/operations/getk8scisbenchmarkrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetK8sCISBenchmarkResponse](../../models/operations/getk8scisbenchmarkresponse.md), error**


## GetK8sCISBenchmarkSummary

Get k8s cis benchmark summary of account

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

**[*operations.GetK8sCISBenchmarkSummaryResponse](../../models/operations/getk8scisbenchmarksummaryresponse.md), error**


## GetK8sCISBenchmarkClusterID

Get k8s cis benchmark for a specific cluster

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
    res, err := s.K8sCisBenchmark.GetK8sCISBenchmarkClusterID(ctx, operations.GetK8sCISBenchmarkClusterIDRequest{
        ClusterID: "183bfe96-59eb-440e-816f-af75b0b532a4",
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
| `request`                                                                                                      | [operations.GetK8sCISBenchmarkClusterIDRequest](../../models/operations/getk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetK8sCISBenchmarkClusterIDResponse](../../models/operations/getk8scisbenchmarkclusteridresponse.md), error**


## PostK8sCISBenchmarkClusterID

initiate k8s cis benchmark scan for a specific cluster

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
    res, err := s.K8sCisBenchmark.PostK8sCISBenchmarkClusterID(ctx, operations.PostK8sCISBenchmarkClusterIDRequest{
        ClusterID: "da37cbaa-f445-42c4-842c-9b2ad32dafe8",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostK8sCISBenchmarkClusterIDRequest](../../models/operations/postk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostK8sCISBenchmarkClusterIDResponse](../../models/operations/postk8scisbenchmarkclusteridresponse.md), error**


## PutK8sCISBenchmarkClusterID

edit k8s cis benchmark for a specific cluster with test statuses

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
    res, err := s.K8sCisBenchmark.PutK8sCISBenchmarkClusterID(ctx, operations.PutK8sCISBenchmarkClusterIDRequest{
        K8sCISBenchmarkUpdateNodes: shared.K8sCISBenchmarkUpdateNodes{
            ClusterID: "1a88f444-4573-4fec-9473-53f63c820937",
            Index: "unde",
            Nodes: []shared.K8sCISBenchmarkUpdateNode{
                shared.K8sCISBenchmarkUpdateNode{
                    ID: testpango.String("aa69cd5f-bcf7-49da-98a7-822bf95894e6"),
                },
            },
            Status: shared.K8sCISBenchmarkUpdateNodeStatusPass,
        },
        ClusterID: "61adb55f-9e5d-4751-89fe-8f7502bfdc34",
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
| `request`                                                                                                      | [operations.PutK8sCISBenchmarkClusterIDRequest](../../models/operations/putk8scisbenchmarkclusteridrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutK8sCISBenchmarkClusterIDResponse](../../models/operations/putk8scisbenchmarkclusteridresponse.md), error**

