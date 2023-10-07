# Performance
(*Performance*)

### Available Operations

* [GetAPISecurityAPICatalogIDHitCountGraph](#getapisecurityapicatalogidhitcountgraph) - Get hit count for specific spec path
* [GetPerformanceMetrics](#getperformancemetrics) - Get performance metrics for a connection between workloads

## GetAPISecurityAPICatalogIDHitCountGraph

Get hit count for specific spec path

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
    res, err := s.Performance.GetAPISecurityAPICatalogIDHitCountGraph(ctx, operations.GetAPISecurityAPICatalogIDHitCountGraphRequest{
        CatalogID: "6b535753-b47a-42be-a003-f45375c0bae8",
        SpecPath: "Tools maxime",
        SpecPathMethod: operations.GetAPISecurityAPICatalogIDHitCountGraphSpecPathMethodPost,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceSpecPathHitCountGraph != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetAPISecurityAPICatalogIDHitCountGraphRequest](../../models/operations/getapisecurityapicatalogidhitcountgraphrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetAPISecurityAPICatalogIDHitCountGraphResponse](../../models/operations/getapisecurityapicatalogidhitcountgraphresponse.md), error**


## GetPerformanceMetrics

Get performance metrics for a connection between workloads

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
    res, err := s.Performance.GetPerformanceMetrics(ctx, operations.GetPerformanceMetricsRequest{
        EndTime: types.MustTimeFromString("2023-06-20T17:24:08.589Z"),
        Protocol: "Account",
        SourceNamespace: "b8ead185-971d-4112-a105-9b89d3b7dcd1",
        SourcePodTemplate: "b894b8a5-ad8e-4111-86dd-d9453d845e92",
        StartTime: types.MustTimeFromString("2023-11-21T23:20:37.062Z"),
        TargetNamespace: "a04569a7-7dda-4601-a8cf-6d7999aa93b9",
        TargetPodTemplate: "be30176b-a06c-4558-9ceb-c600c70dc9c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PerformanceMetrics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPerformanceMetricsRequest](../../models/operations/getperformancemetricsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetPerformanceMetricsResponse](../../models/operations/getperformancemetricsresponse.md), error**

