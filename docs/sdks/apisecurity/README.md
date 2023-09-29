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
        CatalogID: "842f34dd-ffa1-41b2-85e1-21c9635c9fcb",
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
    res, err := s.APISecurity.DeleteAPISecurityInternalCatalogCatalogIDBflaLearning(ctx, operations.DeleteAPISecurityInternalCatalogCatalogIDBflaLearningRequest{
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
        CatalogID: "95e2243c-a46e-4f56-870a-883817cd0574",
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
        GatewayID: "0059ee9e-2eb4-40ca-97cc-9e2e4879b4b7",
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
            "Bronze",
        },
        DrillDownScore: testpango.Bool(false),
        IncludeServiceWithNoSpec: testpango.Bool(false),
        MaxResults: testpango.Float64(4178.64),
        Name: testpango.String("Colton Tin programming"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(9915.74),
        SortDir: operations.GetAPISecurityExternalCatalogSortDirAsc.ToPointer(),
        SortKey: operations.GetAPISecurityExternalCatalogSortKeyName.ToPointer(),
        UpdatedAfter: types.MustTimeFromString("2023-06-28T23:23:21.281Z"),
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
        Name: testpango.String("revolutionary alienated Chair"),
        UpdatedAfter: types.MustTimeFromString("2021-01-16T05:22:36.516Z"),
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
            "Sports",
        },
        CatalogID: "833a138a-8e36-4df0-9dd8-6e879c685823",
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
            "systemic",
        },
        DrillDownScore: testpango.Bool(false),
        IncludeServiceWithNoSpec: testpango.Bool(false),
        MaxResults: testpango.Float64(5658.49),
        Name: testpango.String("tranquilize pascal"),
        NamespacesFilter: testpango.String("Cis capacitor"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(1528.82),
        SortDir: operations.GetAPISecurityInternalCatalogSortDirAsc.ToPointer(),
        SortKey: operations.GetAPISecurityInternalCatalogSortKeyName.ToPointer(),
        UpdatedAfter: types.MustTimeFromString("2021-11-28T12:36:29.977Z"),
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
        Name: testpango.String("Bugatti"),
        NamespacesFilter: testpango.String("3rd Hassium"),
        UpdatedAfter: types.MustTimeFromString("2023-12-03T10:33:57.779Z"),
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
            "Future",
        },
        CatalogID: "b30b167e-21dd-44fc-9f71-8b9d3086136a",
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
        CatalogID: "b8ae405f-27d4-4986-b826-818df6837b3d",
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
        CatalogID: "03b031d5-bcbe-4041-865f-3cf281949555",
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
        CatalogID: "14f8fff8-1410-462d-9de3-1aa0e5094377",
        MaxResults: testpango.Float64(970.24),
        Offset: testpango.Float64(4969.86),
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
        CatalogID: "9766bdb5-2002-47a6-9ab1-0c9e7ba174e7",
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
        CatalogID: "89ba7d71-938d-4e0d-ae77-2227f0233a60",
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
        CatalogID: "e844ac3e-b37a-4e44-810c-fb03dafad57b",
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
        CatalogID: "c1e78882-5e52-4c34-a91e-a88ab3af7faf",
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
        CatalogID: "96eed0ac-9db4-4a1c-babc-e6affa5427c5",
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
        CatalogID: "0bfa0230-126d-4965-9d51-cb2d6c612293",
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
        CatalogID: "4022b3ba-9ecc-4767-9469-71ea890137cd",
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
        Category: testpango.String("becquerel"),
        Detected: testpango.Bool(false),
        Element: testpango.String("green"),
        MaxResults: testpango.Float64(7763.82),
        Name: testpango.String("SMTP Southwest"),
        Offset: testpango.Float64(5947.99),
        Risks: []GetAPISecurityRiskFindingsRisks{
            operations.GetAPISecurityRiskFindingsRisksHigh,
        },
        SortDir: operations.GetAPISecurityRiskFindingsSortDirAsc.ToPointer(),
        SortKey: operations.GetAPISecurityRiskFindingsSortKeyRisk,
        Source: testpango.String("alliance"),
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
        RiskFindingID: "a6f70519-fb6c-4b3a-af18-7a0de05d5783",
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
        CatalogID: "a8274b04-711c-4a08-ae55-c1ca905077f8",
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
        CatalogID: "bc418193-54db-45b9-831c-a3b7fcd92288",
        Tags: []string{
            "purple",
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
        CatalogID: "3f2d279c-52dc-45ca-8a35-0c4c47a5deff",
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
        APISecSource: operations.GetDashboardApisecRiskFindingsAPISecSourceInternal,
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
        APISecSource: operations.GetDashboardApisecRiskFindingsTrendAPISecSourceInternal,
        NumOfDays: testpango.Int64(738205),
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
        MaxResults: testpango.Float64(9119.67),
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
        APISecSource: operations.GetDashboardApisecTopRiskyFindingsAPISecSourceInternal,
        MaxResults: testpango.Float64(5232.51),
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
        MaxResults: testpango.Float64(5622.15),
        Name: testpango.String("Bicycle programming"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(2121.34),
        SortDir: operations.GetGatewaysSortDirAsc.ToPointer(),
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
        GatewayID: "d7df551a-98f0-4f5b-9bdc-d69676cf90f0",
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
        Name: "array Metal",
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
            Duration: "FTM",
        },
        CatalogID: "464b0ff0-e6fe-414c-a291-62fc27770a3f",
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
            Duration: "parse",
        },
        CatalogID: "bc82f2c9-ea85-4a32-ab52-d888f26e15ab",
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
        CatalogID: "3e5cd1f7-bf20-45d2-aba9-73c37afbfaa8",
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
                AuthorizationSchemeType: shared.AuthorizationSchemeAuthorizationSchemeTypeAuthorizationSchemeAPIToken,
            },
            Depth: shared.TestInputDepthEnumDeep,
        },
        CatalogID: "312eede1-be03-42b8-bb19-27700a035c99",
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
            Duration: 916267,
            TimeUnit: shared.TraceAnalysisConfigurationTimeUnitDays,
        },
        CatalogID: "815f76cb-d054-45bf-9982-26857956700d",
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
        CatalogID: "c198f457-4849-4dd7-8320-e57f81ab2f63",
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
        CatalogID: "a026be33-c18b-419f-9481-083f4a844dc4",
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
        CatalogID: "c8574da1-a228-4204-81bd-8ea8316c46d5",
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
            ClusterID: testpango.String("4fd7941d-9d44-44e5-b609-cdd9605335d7"),
            LearningDuration: testpango.String("male Administrator"),
        },
        CatalogID: "1ecf6d42-4fd0-4f29-96f3-b4a797f266b2",
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
            OASVersion: shared.OASVersionOaSv30,
            ReviewPathItems: []shared.ReviewPathItem{
                shared.ReviewPathItem{
                    APIEventsPaths: []shared.APIEventPathAndMethods{
                        shared.APIEventPathAndMethods{
                            Methods: []shared.HTTPMethod{
                                shared.HTTPMethodPut,
                            },
                            Path: testpango.String("/home/user/dir"),
                        },
                    },
                    SuggestedPath: testpango.String("DNS"),
                },
            },
        },
        CatalogID: "b3e10c33-f3cc-4f1f-a2c5-b724ecc18aab",
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
            Duration: 672614,
            TimeUnit: shared.ActionDurationTimeUnitHours,
        },
        CatalogID: "84658317-7e1f-4e27-b84c-670056ef2d95",
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
        CatalogID: "733d2638-f8e8-462b-aa51-67f9d1bf9b92",
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
        ClusterName: "budgetary",
        Description: testpango.String("Organized non-volatile migration"),
        ID: testpango.String("569c810a-5247-4536-a053-f148b119db42"),
        Name: "Savings distinctio blue",
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
            EndTime: types.MustTimeFromString("2021-04-10T03:52:46.500Z"),
            Status: shared.APIServiceBflaInfoStatusInProgressLearning,
            Tags: []shared.APIServiceBflaTagInfo{
                shared.APIServiceBflaTagInfo{
                    IsLegitimate: testpango.Bool(false),
                    Name: "Rhodium Rubber yearly",
                    Paths: []shared.APIServiceBflaPathInfo{
                        shared.APIServiceBflaPathInfo{
                            Clients: []shared.APIServiceBflaClientInfo{
                                shared.APIServiceBflaClientInfo{
                                    External: testpango.Bool(false),
                                    Identifier: testpango.String("3b3fcaad-d442-4012-ad14-59f8769c2011"),
                                    IsLegitimate: testpango.Bool(false),
                                    LastObserved: types.MustTimeFromString("2021-02-25T06:09:36.065Z"),
                                    LastStatusCode: testpango.Int64(508539),
                                    Name: "Carolina",
                                    Namespace: testpango.String("Account"),
                                    Principles: []shared.APIServiceBflaPrincipleInfo{
                                        shared.APIServiceBflaPrincipleInfo{
                                            IP: "56.91.79.204",
                                            Name: "Hyundai punctually",
                                            PrincipleType: "Granite Rustic",
                                        },
                                    },
                                },
                            },
                            IsLegitimate: testpango.Bool(false),
                            Method: shared.HTTPMethodTrace,
                            Path: "/usr/local/bin",
                        },
                    },
                },
            },
        },
        CatalogID: "ac789245-d57c-4c22-8dbd-c1e1df3ce330",
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
        RequestBody: "rerum",
        CatalogID: "6c1e7d20-9d8c-4f8d-a690-7b8b2d3e446c",
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
            ClusterName: "Executive Minivan sky",
            Description: testpango.String("Upgradable bandwidth-monitored contingency"),
            ID: testpango.String("69f02ce6-c757-479c-9c8e-935111463ea3"),
            Name: "Directives pink Central",
            Type: shared.GatewayTypeF5BigIP,
        },
        GatewayID: "95ead58c-f0fb-4b52-ac2b-9a404d8ad2eb",
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

