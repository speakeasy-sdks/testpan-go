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
        CatalogID: "a4d456ef-1031-4e68-99f0-c2001e22cd55",
        HoursInterval: testpango.Int64(798934),
        SpecPath: "quod",
        SpecPathMethod: operations.GetAPISecurityAPICatalogIDHitCountGraphSpecPathMethodGet,
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
        EndTime: types.MustTimeFromString("2022-06-10T08:13:32.523Z"),
        Protocol: "modi",
        SourceNamespace: "a184d76d-971f-4c82-8c65-b037bb8e0cc8",
        SourcePodTemplate: "85187e4d-e04a-4f28-85dd-db46aa1cfd6d",
        StartTime: types.MustTimeFromString("2022-09-01T18:19:44.846Z"),
        TargetNamespace: "8da01319-1129-4646-a45c-1d81f29042f5",
        TargetPodTemplate: "69b7aff0-ea22-416c-be07-1bc163e279a3",
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

