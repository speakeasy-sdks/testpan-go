# Performance
(*.Performance*)

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
        SpecPath: "string",
        SpecPathMethod: operations.SpecPathMethodConnect,
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
        Protocol: "string",
        SourceNamespace: "565fb8ea-d185-4971-9112-61059b89d3b7",
        SourcePodTemplate: "dcd1b894-b8a5-4ad8-a111-06ddd9453d84",
        StartTime: types.MustTimeFromString("2022-01-22T20:26:53.423Z"),
        TargetNamespace: "e92fa045-69a7-47dd-a601-68cf6d7999aa",
        TargetPodTemplate: "93b9be30-176b-4a06-8558-1cebc600c70d",
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

