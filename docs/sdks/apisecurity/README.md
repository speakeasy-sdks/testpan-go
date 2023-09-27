# APISecurity
(*APISecurity*)

## Overview

APIs used to manage Api Security

### Available Operations

* [DeleteAPISecurityAPICatalogID](#deleteapisecurityapicatalogid) - Delete an API
* [DeleteAPISecurityInternalCatalogCatalogIDBflaDetection](#deleteapisecurityinternalcatalogcatalogidbfladetection) - stop bfla detection phase
* [DeleteAPISecurityInternalCatalogCatalogIDBflaLearning](#deleteapisecurityinternalcatalogcatalogidbflalearning) - stop bfla learning phase
* [DeleteAPISecurityOpenAPISpecsCatalogID](#deleteapisecurityopenapispecscatalogid) - delete open api spec include all of it findings and scores
* [DeleteGatewaysGatewayID](#deletegatewaysgatewayid) - Delete gateway
* [GetAPISecurityExternalCatalog](#getapisecurityexternalcatalog) - Get a list of APIs and their compliance
* [GetAPISecurityExternalCatalogCount](#getapisecurityexternalcatalogcount) - Get the number of existing 3rd party APIs in the inventory
* [GetAPISecurityExternalCatalogCatalogID](#getapisecurityexternalcatalogcatalogid) - Get information about a specific API
* [GetAPISecurityInternalCatalog](#getapisecurityinternalcatalog) - Get a list of APIs and their compliance
* [GetAPISecurityInternalCatalogCount](#getapisecurityinternalcatalogcount) - Get the number of existing 3rd party APIs in the inventory
* [GetAPISecurityInternalCatalogCatalogID](#getapisecurityinternalcatalogcatalogid) - Get information about a specific API
* [GetAPISecurityInternalCatalogCatalogIDBfla](#getapisecurityinternalcatalogcatalogidbfla) - Get bfla info for given catalogId
* [GetAPISecurityInternalCatalogCatalogIDFuzzingStatus](#getapisecurityinternalcatalogcatalogidfuzzingstatus) - Get status of the last/running fuzzing test
* [GetAPISecurityInternalCatalogCatalogIDFuzzingTests](#getapisecurityinternalcatalogcatalogidfuzzingtests) - Get list of fuzzing tests
* [GetAPISecurityInternalCatalogCatalogIDTraceAnalysis](#getapisecurityinternalcatalogcatalogidtraceanalysis) - Get trace analysis details
* [GetAPISecurityOpenAPISpecsCatalogID](#getapisecurityopenapispecscatalogid) - Get provided and reconstructed open api specs for specific API
* [GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatus](#getapisecurityopenapispecscatalogiddiffdetectionstatus) - Get provided and reconstructed open api specs for specific API
* [GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatus](#getapisecurityopenapispecscatalogidgetopenapispecscorestatus) - Get open api spec score status
* [GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON](#getapisecurityopenapispecscatalogidopenapispecswaggerjson) - Get provided spec content as json
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReview](#getapisecurityopenapispecscatalogidreconstructedspecreview) - Get the suggestions of a spec reconstruction (or previously cached info)
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatus](#getapisecurityopenapispecscatalogidreconstructedspecstatus) - Get the status of a spec reconstruction
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSON](#getapisecurityopenapispecscatalogidreconstructedspecjson) - Get reconstructed open api spec as json for specific API
* [GetAPISecurityRiskFindings](#getapisecurityriskfindings) - Get a list of risk findings
* [GetAPISecurityRiskFindingsCategories](#getapisecurityriskfindingscategories) - Get a list of risk findings categories
* [GetAPISecurityRiskFindingsSources](#getapisecurityriskfindingssources) - Get a list of risk findings sources
* [GetAPISecurityRiskFindingsRiskFindingID](#getapisecurityriskfindingsriskfindingid) - Get a specific risk findings
* [GetAPISecurityCatalogIDDeleteDependencies](#getapisecuritycatalogiddeletedependencies) - get dependencies which need to be handled in order to delete specified api security service
* [GetAPISecurityCatalogIDMethods](#getapisecuritycatalogidmethods) - Get a list of an API spec methods for a specific API and its spec's tags
* [GetAPISecurityCatalogIDTags](#getapisecuritycatalogidtags) - Get a list of an API spec tags or a specific API
* [GetDashboardApisecRiskFindings](#getdashboardapisecriskfindings) - Get API sec risk findings widget
* [GetDashboardApisecRiskFindingsTrend](#getdashboardapisecriskfindingstrend) - Get API sec risk findings trend graph widget for the last 30 days
* [GetDashboardApisecSpecsAndOperationsDiffs](#getdashboardapisecspecsandoperationsdiffs) - Get API sec specs and operations diffs widget
* [GetDashboardApisecTopRiskyApis](#getdashboardapisectopriskyapis) - Get API sec top risky APIs widget
* [GetDashboardApisecTopRiskyFindings](#getdashboardapisectopriskyfindings) - Get API sec top risky findings widget
* [GetGateways](#getgateways) - Get gateways
* [GetGatewaysClusters](#getgatewaysclusters) - Get clusters info
* [GetGatewaysGatewayIDDownloadBundle](#getgatewaysgatewayiddownloadbundle) - Get a GW installation script
* [PostAPISecurityAPI](#postapisecurityapi) - Register an API for scoring
* [PostAPISecurityInternalCatalogCatalogIDBflaDetection](#postapisecurityinternalcatalogcatalogidbfladetection) - Start new bfla detection phase
* [PostAPISecurityInternalCatalogCatalogIDBflaLearning](#postapisecurityinternalcatalogcatalogidbflalearning) - Start new bfla learning phase
* [PostAPISecurityInternalCatalogCatalogIDBflaReset](#postapisecurityinternalcatalogcatalogidbflareset) - Reset bfla model
* [PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysis](#postapisecurityinternalcatalogcatalogidresettraceanalysis) - Reset trace analysis
* [PostAPISecurityInternalCatalogCatalogIDStartFuzzing](#postapisecurityinternalcatalogcatalogidstartfuzzing) - Start new fuzzing test
* [PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysis](#postapisecurityinternalcatalogcatalogidstarttraceanalysis) - Start trace analysis
* [PostAPISecurityInternalCatalogCatalogIDStopFuzzing](#postapisecurityinternalcatalogcatalogidstopfuzzing) - Stop fuzzing test
* [PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysis](#postapisecurityinternalcatalogcatalogidstoptraceanalysis) - Stop trace analysis
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbort](#postapisecurityopenapispecscatalogidreconstructedspecabort) - abort learning and reconstructing an API via API Clarity
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearn](#postapisecurityopenapispecscatalogidreconstructedspeclearn) - Start learning and reconstructing an API via API Clarity
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApprove](#postapisecurityopenapispecscatalogidreconstructedspecreviewapprove) - Approve reconstructed spec suggestions (only 1 approval per catalogId)
* [PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetection](#postapisecurityopenapispecscatalogidstartdiffsdetection) - Start spec diffs detection
* [PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetection](#postapisecurityopenapispecscatalogidstopdiffsdetection) - stop spec diffs detection
* [PostGateways](#postgateways) - Add new gateway
* [PutAPISecurityInternalCatalogCatalogIDBfla](#putapisecurityinternalcatalogcatalogidbfla) - update BFLA info for this catalogId
* [PutAPISecurityOpenAPISpecsCatalogID](#putapisecurityopenapispecscatalogid) - Add or edit a spec about a specific API for the account
* [PutGatewaysGatewayID](#putgatewaysgatewayid) - Edit gateway

## DeleteAPISecurityAPICatalogID

Delete an API

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
    res, err := s.APISecurity.DeleteAPISecurityAPICatalogID(ctx, operations.DeleteAPISecurityAPICatalogIDRequest{
        CatalogID: "a8d8f5c0-b2f2-4fb7-b194-a276b26916fe",
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
| `request`                                                                                                          | [operations.DeleteAPISecurityAPICatalogIDRequest](../../models/operations/deleteapisecurityapicatalogidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteAPISecurityAPICatalogIDResponse](../../models/operations/deleteapisecurityapicatalogidresponse.md), error**


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
    res, err := s.APISecurity.DeleteAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        CatalogID: "1f08f429-4e36-498f-847f-603e8b445e80",
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
    res, err := s.APISecurity.DeleteAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        CatalogID: "ca55efd2-0e45-47e1-858b-6a89fbe3a5aa",
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


## DeleteAPISecurityOpenAPISpecsCatalogID

delete open api spec include all of it findings and scores

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
    res, err := s.APISecurity.DeleteAPISecurityOpenAPISpecsCatalogID(ctx, operations.DeleteAPISecurityOpenAPISpecsCatalogIDRequest{
        CatalogID: "8e4824d0-ab40-4750-88e5-1862065e904f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAPISpecScoreStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.DeleteAPISecurityOpenAPISpecsCatalogIDRequest](../../models/operations/deleteapisecurityopenapispecscatalogidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.DeleteAPISecurityOpenAPISpecsCatalogIDResponse](../../models/operations/deleteapisecurityopenapispecscatalogidresponse.md), error**


## DeleteGatewaysGatewayID

Delete gateway

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
    res, err := s.APISecurity.DeleteGatewaysGatewayID(ctx, operations.DeleteGatewaysGatewayIDRequest{
        GatewayID: "3b1194b8-abf6-403a-b9f9-dfe0ab7da8a5",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteGatewaysGatewayIDRequest](../../models/operations/deletegatewaysgatewayidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.DeleteGatewaysGatewayIDResponse](../../models/operations/deletegatewaysgatewayidresponse.md), error**


## GetAPISecurityExternalCatalog

Get a list of APIs and their compliance

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
    res, err := s.APISecurity.GetAPISecurityExternalCatalog(ctx, operations.GetAPISecurityExternalCatalogRequest{
        APIPolicyProfiles: []string{
            "voluptatem",
        },
        DrillDownScore: testpango.Bool(false),
        IncludeServiceWithNoSpec: testpango.Bool(false),
        MaxResults: testpango.Float64(7908.4),
        Name: testpango.String("Ryan Littel"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(5199.52),
        SortDir: operations.GetAPISecurityExternalCatalogSortDirAsc.ToPointer(),
        SortKey: operations.GetAPISecurityExternalCatalogSortKeyRisk.ToPointer(),
        UpdatedAfter: types.MustTimeFromString("2022-09-23T11:31:21.970Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceListExternal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAPISecurityExternalCatalogRequest](../../models/operations/getapisecurityexternalcatalogrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAPISecurityExternalCatalogResponse](../../models/operations/getapisecurityexternalcatalogresponse.md), error**


## GetAPISecurityExternalCatalogCount

Get the number of existing 3rd party APIs in the inventory

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
    res, err := s.APISecurity.GetAPISecurityExternalCatalogCount(ctx, operations.GetAPISecurityExternalCatalogCountRequest{
        IncludeServiceWithNoSpec: testpango.Bool(false),
        Name: testpango.String("Grace Shields"),
        UpdatedAfter: types.MustTimeFromString("2021-03-24T07:07:12.173Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityExternalCatalogCount200ApplicationJSONInteger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetAPISecurityExternalCatalogCountRequest](../../models/operations/getapisecurityexternalcatalogcountrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetAPISecurityExternalCatalogCountResponse](../../models/operations/getapisecurityexternalcatalogcountresponse.md), error**


## GetAPISecurityExternalCatalogCatalogID

Get information about a specific API

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
    res, err := s.APISecurity.GetAPISecurityExternalCatalogCatalogID(ctx, operations.GetAPISecurityExternalCatalogCatalogIDRequest{
        APIPolicyProfiles: []string{
            "officiis",
        },
        CatalogID: "e9526f8d-986e-4881-aad4-f0e1012563f9",
        DownloadAsJSON: testpango.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceDrillDownExternal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetAPISecurityExternalCatalogCatalogIDRequest](../../models/operations/getapisecurityexternalcatalogcatalogidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetAPISecurityExternalCatalogCatalogIDResponse](../../models/operations/getapisecurityexternalcatalogcatalogidresponse.md), error**


## GetAPISecurityInternalCatalog

Get a list of APIs and their compliance

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalog(ctx, operations.GetAPISecurityInternalCatalogRequest{
        APIPolicyProfiles: []string{
            "magnam",
        },
        DrillDownScore: testpango.Bool(false),
        IncludeServiceWithNoSpec: testpango.Bool(false),
        MaxResults: testpango.Float64(9063.55),
        Name: testpango.String("Toni Torphy"),
        NamespacesFilter: testpango.String("adipisci"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(9078.76),
        SortDir: operations.GetAPISecurityInternalCatalogSortDirDesc.ToPointer(),
        SortKey: operations.GetAPISecurityInternalCatalogSortKeyName.ToPointer(),
        UpdatedAfter: types.MustTimeFromString("2022-05-04T16:36:37.699Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceListInternal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAPISecurityInternalCatalogRequest](../../models/operations/getapisecurityinternalcatalogrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAPISecurityInternalCatalogResponse](../../models/operations/getapisecurityinternalcatalogresponse.md), error**


## GetAPISecurityInternalCatalogCount

Get the number of existing 3rd party APIs in the inventory

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCount(ctx, operations.GetAPISecurityInternalCatalogCountRequest{
        IncludeServiceWithNoSpec: testpango.Bool(false),
        Name: testpango.String("Mrs. Bessie Muller"),
        NamespacesFilter: testpango.String("eveniet"),
        UpdatedAfter: types.MustTimeFromString("2022-02-14T08:24:16.303Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityInternalCatalogCount200ApplicationJSONInteger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetAPISecurityInternalCatalogCountRequest](../../models/operations/getapisecurityinternalcatalogcountrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetAPISecurityInternalCatalogCountResponse](../../models/operations/getapisecurityinternalcatalogcountresponse.md), error**


## GetAPISecurityInternalCatalogCatalogID

Get information about a specific API

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCatalogID(ctx, operations.GetAPISecurityInternalCatalogCatalogIDRequest{
        APIPolicyProfiles: []string{
            "doloremque",
        },
        CatalogID: "60807e2b-6e3a-4b88-85f0-597a60ff2a54",
        DownloadAsJSON: testpango.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceDrillDownInternal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetAPISecurityInternalCatalogCatalogIDRequest](../../models/operations/getapisecurityinternalcatalogcatalogidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDResponse](../../models/operations/getapisecurityinternalcatalogcatalogidresponse.md), error**


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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCatalogIDBfla(ctx, operations.GetAPISecurityInternalCatalogCatalogIDBflaRequest{
        CatalogID: "a31e9476-4a3e-4865-a795-6f9251a5a9da",
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


## GetAPISecurityInternalCatalogCatalogIDFuzzingStatus

Get status of the last/running fuzzing test

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCatalogIDFuzzingStatus(ctx, operations.GetAPISecurityInternalCatalogCatalogIDFuzzingStatusRequest{
        CatalogID: "660ff57b-faad-44f9-afc1-b4512c103264",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceFuzzingProgress != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.GetAPISecurityInternalCatalogCatalogIDFuzzingStatusRequest](../../models/operations/getapisecurityinternalcatalogcatalogidfuzzingstatusrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDFuzzingStatusResponse](../../models/operations/getapisecurityinternalcatalogcatalogidfuzzingstatusresponse.md), error**


## GetAPISecurityInternalCatalogCatalogIDFuzzingTests

Get list of fuzzing tests

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCatalogIDFuzzingTests(ctx, operations.GetAPISecurityInternalCatalogCatalogIDFuzzingTestsRequest{
        CatalogID: "8dc2f615-199e-4bfd-8e9f-e6c632ca3aed",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceFuzzingTests != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetAPISecurityInternalCatalogCatalogIDFuzzingTestsRequest](../../models/operations/getapisecurityinternalcatalogcatalogidfuzzingtestsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDFuzzingTestsResponse](../../models/operations/getapisecurityinternalcatalogcatalogidfuzzingtestsresponse.md), error**


## GetAPISecurityInternalCatalogCatalogIDTraceAnalysis

Get trace analysis details

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
    res, err := s.APISecurity.GetAPISecurityInternalCatalogCatalogIDTraceAnalysis(ctx, operations.GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest{
        CatalogID: "01179963-12fd-4e04-b717-78ff61d01747",
        MaxResults: testpango.Float64(4037.93),
        Offset: testpango.Float64(2352.63),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TraceAnalysisDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.GetAPISecurityInternalCatalogCatalogIDTraceAnalysisRequest](../../models/operations/getapisecurityinternalcatalogcatalogidtraceanalysisrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.GetAPISecurityInternalCatalogCatalogIDTraceAnalysisResponse](../../models/operations/getapisecurityinternalcatalogcatalogidtraceanalysisresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogID

Get provided and reconstructed open api specs for specific API

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogID(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDRequest{
        CatalogID: "60a15db6-a660-4659-a1ad-eaab5851d6c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAPISpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetAPISecurityOpenAPISpecsCatalogIDRequest](../../models/operations/getapisecurityopenapispecscatalogidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDResponse](../../models/operations/getapisecurityopenapispecscatalogidresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatus

Get provided and reconstructed open api specs for specific API

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatus(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatusRequest{
        CatalogID: "45b08b61-891b-4aa0-be1a-de008e6f8c5f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DiffDetectionStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatusRequest](../../models/operations/getapisecurityopenapispecscatalogiddiffdetectionstatusrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatusResponse](../../models/operations/getapisecurityopenapispecscatalogiddiffdetectionstatusresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatus

Get open api spec score status

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatus(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusRequest{
        CatalogID: "350d8cdb-5a34-4181-8301-0421813d5208",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAPISpecScoreStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusRequest](../../models/operations/getapisecurityopenapispecscatalogidgetopenapispecscorestatusrequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusResponse](../../models/operations/getapisecurityopenapispecscatalogidgetopenapispecscorestatusresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON

Get provided spec content as json

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONRequest{
        CatalogID: "ece7e253-b668-4451-86c6-e205e16deab3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONRequest](../../models/operations/getapisecurityopenapispecscatalogidopenapispecswaggerjsonrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSONResponse](../../models/operations/getapisecurityopenapispecscatalogidopenapispecswaggerjsonresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReview

Get the suggestions of a spec reconstruction (or previously cached info)

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReview(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewRequest{
        CatalogID: "fec9578a-6458-4427-ba84-18d162309fb0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIReconstructedSpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewRequest](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecreviewrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewResponse](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecreviewresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatus

Get the status of a spec reconstruction

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatus(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatusRequest{
        CatalogID: "929921ae-fb9f-458c-8d86-e68e4be05601",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIReconstructionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatusRequest](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecstatusrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatusResponse](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecstatusresponse.md), error**


## GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSON

Get reconstructed open api spec as json for specific API

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
    res, err := s.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSON(ctx, operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSONRequest{
        CatalogID: "3f59da75-7a59-4ecf-af66-ef1caa3383c2",
        DownloadAsJSON: testpango.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSON200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSONRequest](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecjsonrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |


### Response

**[*operations.GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSONResponse](../../models/operations/getapisecurityopenapispecscatalogidreconstructedspecjsonresponse.md), error**


## GetAPISecurityRiskFindings

Get a list of risk findings

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
    res, err := s.APISecurity.GetAPISecurityRiskFindings(ctx, operations.GetAPISecurityRiskFindingsRequest{
        APISecSource: operations.GetAPISecurityRiskFindingsAPISecSourceExternal,
        Category: testpango.String("repudiandae"),
        Detected: testpango.Bool(false),
        Element: testpango.String("nam"),
        MaxResults: testpango.Float64(2940.76),
        Name: testpango.String("Colleen Dickinson"),
        Offset: testpango.Float64(7781.72),
        Risks: []GetAPISecurityRiskFindingsRisks{
            operations.GetAPISecurityRiskFindingsRisksHigh,
        },
        SortDir: operations.GetAPISecurityRiskFindingsSortDirDesc.ToPointer(),
        SortKey: operations.GetAPISecurityRiskFindingsSortKeyName,
        Source: testpango.String("odit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskFindings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAPISecurityRiskFindingsRequest](../../models/operations/getapisecurityriskfindingsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetAPISecurityRiskFindingsResponse](../../models/operations/getapisecurityriskfindingsresponse.md), error**


## GetAPISecurityRiskFindingsCategories

Get a list of risk findings categories

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
    res, err := s.APISecurity.GetAPISecurityRiskFindingsCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityRiskFindingsCategories200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAPISecurityRiskFindingsCategoriesResponse](../../models/operations/getapisecurityriskfindingscategoriesresponse.md), error**


## GetAPISecurityRiskFindingsSources

Get a list of risk findings sources

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
    res, err := s.APISecurity.GetAPISecurityRiskFindingsSources(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPISecurityRiskFindingsSources200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAPISecurityRiskFindingsSourcesResponse](../../models/operations/getapisecurityriskfindingssourcesresponse.md), error**


## GetAPISecurityRiskFindingsRiskFindingID

Get a specific risk findings

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
    res, err := s.APISecurity.GetAPISecurityRiskFindingsRiskFindingID(ctx, operations.GetAPISecurityRiskFindingsRiskFindingIDRequest{
        RiskFindingID: "f64d1db1-f2c4-4310-a61e-96349e1cf9e0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskFinding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetAPISecurityRiskFindingsRiskFindingIDRequest](../../models/operations/getapisecurityriskfindingsriskfindingidrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetAPISecurityRiskFindingsRiskFindingIDResponse](../../models/operations/getapisecurityriskfindingsriskfindingidresponse.md), error**


## GetAPISecurityCatalogIDDeleteDependencies

get dependencies which need to be handled in order to delete specified api security service

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
    res, err := s.APISecurity.GetAPISecurityCatalogIDDeleteDependencies(ctx, operations.GetAPISecurityCatalogIDDeleteDependenciesRequest{
        CatalogID: "6e3a4370-00ae-46b6-bc9b-8f759eac55a9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetAPISecurityCatalogIDDeleteDependenciesRequest](../../models/operations/getapisecuritycatalogiddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetAPISecurityCatalogIDDeleteDependenciesResponse](../../models/operations/getapisecuritycatalogiddeletedependenciesresponse.md), error**


## GetAPISecurityCatalogIDMethods

Get a list of an API spec methods for a specific API and its spec's tags

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
    res, err := s.APISecurity.GetAPISecurityCatalogIDMethods(ctx, operations.GetAPISecurityCatalogIDMethodsRequest{
        CatalogID: "741d3113-5296-45bb-8a72-02611435e139",
        Tags: []string{
            "nulla",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceMethods != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAPISecurityCatalogIDMethodsRequest](../../models/operations/getapisecuritycatalogidmethodsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetAPISecurityCatalogIDMethodsResponse](../../models/operations/getapisecuritycatalogidmethodsresponse.md), error**


## GetAPISecurityCatalogIDTags

Get a list of an API spec tags or a specific API

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
    res, err := s.APISecurity.GetAPISecurityCatalogIDTags(ctx, operations.GetAPISecurityCatalogIDTagsRequest{
        CatalogID: "bc2259b1-abda-48c0-b0e1-084cb0672d1a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceTags != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAPISecurityCatalogIDTagsRequest](../../models/operations/getapisecuritycatalogidtagsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetAPISecurityCatalogIDTagsResponse](../../models/operations/getapisecuritycatalogidtagsresponse.md), error**


## GetDashboardApisecRiskFindings

Get API sec risk findings widget

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
    res, err := s.APISecurity.GetDashboardApisecRiskFindings(ctx, operations.GetDashboardApisecRiskFindingsRequest{
        APISecSource: operations.GetDashboardApisecRiskFindingsAPISecSourceExternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecRiskFindingsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetDashboardApisecRiskFindingsRequest](../../models/operations/getdashboardapisecriskfindingsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetDashboardApisecRiskFindingsResponse](../../models/operations/getdashboardapisecriskfindingsresponse.md), error**


## GetDashboardApisecRiskFindingsTrend

Get API sec risk findings trend graph widget for the last 30 days

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
    res, err := s.APISecurity.GetDashboardApisecRiskFindingsTrend(ctx, operations.GetDashboardApisecRiskFindingsTrendRequest{
        APISecSource: operations.GetDashboardApisecRiskFindingsTrendAPISecSourceExternal,
        NumOfDays: testpango.Int64(473143),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecRiskFindingsTrendWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetDashboardApisecRiskFindingsTrendRequest](../../models/operations/getdashboardapisecriskfindingstrendrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GetDashboardApisecRiskFindingsTrendResponse](../../models/operations/getdashboardapisecriskfindingstrendresponse.md), error**


## GetDashboardApisecSpecsAndOperationsDiffs

Get API sec specs and operations diffs widget

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
    res, err := s.APISecurity.GetDashboardApisecSpecsAndOperationsDiffs(ctx, operations.GetDashboardApisecSpecsAndOperationsDiffsRequest{
        APISecSource: operations.GetDashboardApisecSpecsAndOperationsDiffsAPISecSourceExternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SpecsAndOperationsDiffsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetDashboardApisecSpecsAndOperationsDiffsRequest](../../models/operations/getdashboardapisecspecsandoperationsdiffsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetDashboardApisecSpecsAndOperationsDiffsResponse](../../models/operations/getdashboardapisecspecsandoperationsdiffsresponse.md), error**


## GetDashboardApisecTopRiskyApis

Get API sec top risky APIs widget

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
    res, err := s.APISecurity.GetDashboardApisecTopRiskyApis(ctx, operations.GetDashboardApisecTopRiskyApisRequest{
        APISecSource: operations.GetDashboardApisecTopRiskyApisAPISecSourceExternal,
        MaxResults: testpango.Float64(8964.8),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecTopRiskyApisWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetDashboardApisecTopRiskyApisRequest](../../models/operations/getdashboardapisectopriskyapisrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetDashboardApisecTopRiskyApisResponse](../../models/operations/getdashboardapisectopriskyapisresponse.md), error**


## GetDashboardApisecTopRiskyFindings

Get API sec top risky findings widget

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
    res, err := s.APISecurity.GetDashboardApisecTopRiskyFindings(ctx, operations.GetDashboardApisecTopRiskyFindingsRequest{
        APISecSource: operations.GetDashboardApisecTopRiskyFindingsAPISecSourceExternal,
        MaxResults: testpango.Float64(5750.78),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecTopRiskyFindingsWidget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetDashboardApisecTopRiskyFindingsRequest](../../models/operations/getdashboardapisectopriskyfindingsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetDashboardApisecTopRiskyFindingsResponse](../../models/operations/getdashboardapisectopriskyfindingsresponse.md), error**


## GetGateways

Get gateways

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
    res, err := s.APISecurity.GetGateways(ctx, operations.GetGatewaysRequest{
        MaxResults: testpango.Float64(4097.26),
        Name: testpango.String("Brittany Prosacco"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(8890.6),
        SortDir: operations.GetGatewaysSortDirDesc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateways != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetGatewaysRequest](../../models/operations/getgatewaysrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetGatewaysResponse](../../models/operations/getgatewaysresponse.md), error**


## GetGatewaysClusters

Get clusters info

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
    res, err := s.APISecurity.GetGatewaysClusters(ctx, operations.GetGatewaysClustersRequest{
        GatewayType: operations.GetGatewaysClustersGatewayTypeTykInternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GatewayClusterInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetGatewaysClustersRequest](../../models/operations/getgatewaysclustersrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetGatewaysClustersResponse](../../models/operations/getgatewaysclustersresponse.md), error**


## GetGatewaysGatewayIDDownloadBundle

In order to install,  extract and run "./install_bundle.sh"

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
    res, err := s.APISecurity.GetGatewaysGatewayIDDownloadBundle(ctx, operations.GetGatewaysGatewayIDDownloadBundleRequest{
        GatewayID: "d02bae0b-e2d7-4822-99e3-ea4b5197f924",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetGatewaysGatewayIDDownloadBundle200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetGatewaysGatewayIDDownloadBundleRequest](../../models/operations/getgatewaysgatewayiddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetGatewaysGatewayIDDownloadBundleResponse](../../models/operations/getgatewaysgatewayiddownloadbundleresponse.md), error**


## PostAPISecurityAPI

Register an API for scoring

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
    res, err := s.APISecurity.PostAPISecurityAPI(ctx, shared.APISecurityAPI{
        Name: "Peggy Swift",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPISecurityAPI201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.APISecurityAPI](../../models/shared/apisecurityapi.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.PostAPISecurityAPIResponse](../../models/operations/postapisecurityapiresponse.md), error**


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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDBflaDetection(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaDetectionRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "optio",
        },
        CatalogID: "e52b895c-537c-4645-8efb-0b34896c3ca5",
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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
        BflaDurationConfiguration: shared.BflaDurationConfiguration{
            Duration: "est",
        },
        CatalogID: "cfbe2fd5-7075-4779-a917-7deac646ecb5",
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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDBflaReset(ctx, operations.PostAPISecurityInternalCatalogCatalogIDBflaResetRequest{
        CatalogID: "73409e3e-b1e5-4a2b-92eb-07f116db9954",
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


## PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysis

Reset trace analysis

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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysis(ctx, operations.PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisRequest{
        CatalogID: "5fc95fa8-8970-4e18-9dbb-30fcb33ea055",
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

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisRequest](../../models/operations/postapisecurityinternalcatalogcatalogidresettraceanalysisrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysisResponse](../../models/operations/postapisecurityinternalcatalogcatalogidresettraceanalysisresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDStartFuzzing

Start new fuzzing test

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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDStartFuzzing(ctx, operations.PostAPISecurityInternalCatalogCatalogIDStartFuzzingRequest{
        APIFuzzingTestConfiguration: shared.APIFuzzingTestConfiguration{
            Auth: shared.AuthorizationScheme{
                AuthorizationSchemeType: shared.AuthorizationSchemeAuthorizationSchemeTypeAuthorizationSchemeBearerToken,
            },
            Depth: shared.TestInputDepthEnumQuick,
        },
        CatalogID: "97cd44e2-f52d-482d-b513-bb6f48b656bc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIServiceFuzzingTest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.PostAPISecurityInternalCatalogCatalogIDStartFuzzingRequest](../../models/operations/postapisecurityinternalcatalogcatalogidstartfuzzingrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDStartFuzzingResponse](../../models/operations/postapisecurityinternalcatalogcatalogidstartfuzzingresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysis

Start trace analysis

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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysis(ctx, operations.PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysisRequest{
        TraceAnalysisConfiguration: shared.TraceAnalysisConfiguration{
            Duration: 815200,
            TimeUnit: shared.TraceAnalysisConfigurationTimeUnitDays,
        },
        CatalogID: "35ff2e4b-2753-47a8-8d9e-7319c177d525",
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

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysisRequest](../../models/operations/postapisecurityinternalcatalogcatalogidstarttraceanalysisrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysisResponse](../../models/operations/postapisecurityinternalcatalogcatalogidstarttraceanalysisresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDStopFuzzing

Stop fuzzing test

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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDStopFuzzing(ctx, operations.PostAPISecurityInternalCatalogCatalogIDStopFuzzingRequest{
        CatalogID: "f77b114e-eb52-4ff7-85fc-37814d4c98e0",
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.PostAPISecurityInternalCatalogCatalogIDStopFuzzingRequest](../../models/operations/postapisecurityinternalcatalogcatalogidstopfuzzingrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDStopFuzzingResponse](../../models/operations/postapisecurityinternalcatalogcatalogidstopfuzzingresponse.md), error**


## PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysis

Stop trace analysis

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
    res, err := s.APISecurity.PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysis(ctx, operations.PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysisRequest{
        CatalogID: "c2bb89eb-75da-4d63-ac60-0503d8bb3118",
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

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysisRequest](../../models/operations/postapisecurityinternalcatalogcatalogidstoptraceanalysisrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |


### Response

**[*operations.PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysisResponse](../../models/operations/postapisecurityinternalcatalogcatalogidstoptraceanalysisresponse.md), error**


## PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbort

abort learning and reconstructing an API via API Clarity

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
    res, err := s.APISecurity.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbort(ctx, operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbortRequest{
        CatalogID: "0f739ae9-e057-4eb8-89e2-810331f3981d",
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

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbortRequest](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspecabortrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbortResponse](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspecabortresponse.md), error**


## PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearn

Start learning and reconstructing an API via API Clarity

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
    res, err := s.APISecurity.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearn(ctx, operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearnRequest{
        APIReconstructionRequest: &shared.APIReconstructionRequest{
            ClusterID: testpango.String("4c700b60-7f3c-493c-b3b9-da3f2ceda7e2"),
            LearningDuration: testpango.String("consectetur"),
        },
        CatalogID: "f2257411-faf4-4b75-84e4-72e802857a5b",
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

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearnRequest](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspeclearnrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearnResponse](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspeclearnresponse.md), error**


## PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApprove

Approve reconstructed spec suggestions (only 1 approval per catalogId)

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
    res, err := s.APISecurity.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApprove(ctx, operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveRequest{
        APIReconstructedSpecInput: shared.APIReconstructedSpecInput{
            OASVersion: shared.OASVersionOaSv20,
            ReviewPathItems: []shared.ReviewPathItem{
                shared.ReviewPathItem{
                    APIEventsPaths: []shared.APIEventPathAndMethods{
                        shared.APIEventPathAndMethods{
                            Methods: []shared.HTTPMethod{
                                shared.HTTPMethodGet,
                            },
                            Path: testpango.String("modi"),
                        },
                    },
                    SuggestedPath: testpango.String("eum"),
                },
            },
        },
        CatalogID: "3a7d575f-1400-4e76-8ad7-334ec1b781b3",
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

| Parameter                                                                                                                                                                                    | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | The context to use for the request.                                                                                                                                                          |
| `request`                                                                                                                                                                                    | [operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveRequest](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspecreviewapproverequest.md) | :heavy_check_mark:                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                   |


### Response

**[*operations.PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApproveResponse](../../models/operations/postapisecurityopenapispecscatalogidreconstructedspecreviewapproveresponse.md), error**


## PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetection

Start spec diffs detection

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
    res, err := s.APISecurity.PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetection(ctx, operations.PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetectionRequest{
        ActionDuration: shared.ActionDuration{
            Duration: 418637,
            TimeUnit: shared.ActionDurationTimeUnitDays,
        },
        CatalogID: "08088d10-0efa-4da2-80ef-0422eb2164cf",
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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetectionRequest](../../models/operations/postapisecurityopenapispecscatalogidstartdiffsdetectionrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetectionResponse](../../models/operations/postapisecurityopenapispecscatalogidstartdiffsdetectionresponse.md), error**


## PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetection

stop spec diffs detection

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
    res, err := s.APISecurity.PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetection(ctx, operations.PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetectionRequest{
        CatalogID: "9ab8366c-723f-4fda-9e06-bee4825c1fc0",
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

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetectionRequest](../../models/operations/postapisecurityopenapispecscatalogidstopdiffsdetectionrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetectionResponse](../../models/operations/postapisecurityopenapispecscatalogidstopdiffsdetectionresponse.md), error**


## PostGateways

Add new gateway

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
    res, err := s.APISecurity.PostGateways(ctx, shared.Gateway{
        ClusterName: "officiis",
        Description: testpango.String("architecto"),
        ID: testpango.String("15c80bff-9185-444e-842d-efcce8f19777"),
        Name: "Crystal Tremblay",
        Type: shared.GatewayTypeKongInternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateway != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Gateway](../../models/shared/gateway.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostGatewaysResponse](../../models/operations/postgatewaysresponse.md), error**


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
    res, err := s.APISecurity.PutAPISecurityInternalCatalogCatalogIDBfla(ctx, operations.PutAPISecurityInternalCatalogCatalogIDBflaRequest{
        APIServiceBflaInfo: shared.APIServiceBflaInfo{
            EndTime: types.MustTimeFromString("2022-11-02T06:07:32.455Z"),
            Status: shared.APIServiceBflaInfoStatusInProgressLearning,
            Tags: []shared.APIServiceBflaTagInfo{
                shared.APIServiceBflaTagInfo{
                    IsLegitimate: testpango.Bool(false),
                    Name: "Ms. Verna Gislason",
                    Paths: []shared.APIServiceBflaPathInfo{
                        shared.APIServiceBflaPathInfo{
                            Clients: []shared.APIServiceBflaClientInfo{
                                shared.APIServiceBflaClientInfo{
                                    External: testpango.Bool(false),
                                    Identifier: testpango.String("05e3d48f-daf3-413a-9f5f-d94259c0b36f"),
                                    IsLegitimate: testpango.Bool(false),
                                    LastObserved: types.MustTimeFromString("2022-09-02T09:20:25.526Z"),
                                    LastStatusCode: testpango.Int64(893773),
                                    Name: "Kirk Goyette",
                                    Namespace: testpango.String("adipisci"),
                                    Principles: []shared.APIServiceBflaPrincipleInfo{
                                        shared.APIServiceBflaPrincipleInfo{
                                            IP: "libero",
                                            Name: "Lorraine Jacobson Sr.",
                                            PrincipleType: "hic",
                                        },
                                    },
                                },
                            },
                            IsLegitimate: testpango.Bool(false),
                            Method: shared.HTTPMethodDelete,
                            Path: "quisquam",
                        },
                    },
                },
            },
        },
        CatalogID: "37a51262-4383-45bb-805a-23a45cefc5fd",
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


## PutAPISecurityOpenAPISpecsCatalogID

Add or edit a spec about a specific API for the account

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
    res, err := s.APISecurity.PutAPISecurityOpenAPISpecsCatalogID(ctx, operations.PutAPISecurityOpenAPISpecsCatalogIDRequest{
        RequestBody: "officiis",
        CatalogID: "10a0ce21-69e5-4100-99c6-dc5e34762799",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAPISpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PutAPISecurityOpenAPISpecsCatalogIDRequest](../../models/operations/putapisecurityopenapispecscatalogidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.PutAPISecurityOpenAPISpecsCatalogIDResponse](../../models/operations/putapisecurityopenapispecscatalogidresponse.md), error**


## PutGatewaysGatewayID

Edit gateway

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
    res, err := s.APISecurity.PutGatewaysGatewayID(ctx, operations.PutGatewaysGatewayIDRequest{
        Gateway: shared.Gateway{
            ClusterName: "cum",
            Description: testpango.String("doloribus"),
            ID: testpango.String("bbe6949f-b2bb-44ec-ae6c-3d5db3adebd5"),
            Name: "Pablo Veum",
            Type: shared.GatewayTypeF5BigIP,
        },
        GatewayID: "506a8aa9-4c02-4644-8f5e-9d9a4578adc1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateway != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutGatewaysGatewayIDRequest](../../models/operations/putgatewaysgatewayidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PutGatewaysGatewayIDResponse](../../models/operations/putgatewaysgatewayidresponse.md), error**

