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
        SpecPath: "Dynamic",
        SpecPathMethod: operations.GetAPISecurityAPICatalogIDHitCountGraphSpecPathMethodDelete,
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
        Protocol: "technologies",
        SourceNamespace: "5fb8ead1-8597-41d1-9261-059b89d3b7dc",
        SourcePodTemplate: "d1b894b8-a5ad-48e1-9106-ddd9453d845e",
        StartTime: types.MustTimeFromString("2022-09-17T00:22:13.332Z"),
        TargetNamespace: "2fa04569-a77d-4da6-8168-cf6d7999aa93",
        TargetPodTemplate: "b9be3017-6ba0-46c5-981c-ebc600c70dc9",
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

