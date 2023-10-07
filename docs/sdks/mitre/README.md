# Mitre
(*Mitre*)

### Available Operations

* [GetMitreDashboard](#getmitredashboard) - Get data for MITRE dashboard for all clusters
* [GetMitreReportDownload](#getmitrereportdownload) - Download Mitre security report
* [GetMitreReportStatus](#getmitrereportstatus) - Get Mitre report status
* [GetMitreTechnique](#getmitretechnique) - Get data for MITRE technique of the given mitreTechniqueType
* [PostMitreReportGenerate](#postmitrereportgenerate) - Generate Mitre report
* [PostMitreTechniqueFix](#postmitretechniquefix) - Post fix for MITRE technique of the given mitreTechniqueType

## GetMitreDashboard

Get data for MITRE dashboard for all clusters

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
    res, err := s.Mitre.GetMitreDashboard(ctx, operations.GetMitreDashboardRequest{
        ClustersIds: []string{
            "a91a8587-5ae9-468d-bcc5-575f0642f99f",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MitreDashboard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetMitreDashboardRequest](../../models/operations/getmitredashboardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetMitreDashboardResponse](../../models/operations/getmitredashboardresponse.md), error**


## GetMitreReportDownload

Download Mitre security report

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
    res, err := s.Mitre.GetMitreReportDownload(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMitreReportDownload200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetMitreReportDownloadResponse](../../models/operations/getmitrereportdownloadresponse.md), error**


## GetMitreReportStatus

Get Mitre report status

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
    res, err := s.Mitre.GetMitreReportStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ReportStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetMitreReportStatusResponse](../../models/operations/getmitrereportstatusresponse.md), error**


## GetMitreTechnique

Get data for MITRE technique of the given mitreTechniqueType

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
    res, err := s.Mitre.GetMitreTechnique(ctx, operations.GetMitreTechniqueRequest{
        ClustersIds: []string{
            "15c243dc-adda-4d62-b64a-acf9b6fc450a",
        },
        MitreTechniqueType: operations.GetMitreTechniqueMitreTechniqueTypeDeployContainer,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MitreTechniqueInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetMitreTechniqueRequest](../../models/operations/getmitretechniquerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetMitreTechniqueResponse](../../models/operations/getmitretechniqueresponse.md), error**


## PostMitreReportGenerate

Generate Mitre report

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
    res, err := s.Mitre.PostMitreReportGenerate(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostMitreReportGenerateResponse](../../models/operations/postmitrereportgenerateresponse.md), error**


## PostMitreTechniqueFix

Post fix for MITRE technique of the given mitreTechniqueType

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
    res, err := s.Mitre.PostMitreTechniqueFix(ctx, operations.PostMitreTechniqueFixRequest{
        MitreTechniqueFixInfo: shared.MitreTechniqueFixInfo{
            AffectedElements: []shared.MitreTechniqueAffectedElement{
                shared.MitreTechniqueAffectedElement{},
            },
        },
        ClustersIds: []string{
            "684b4b2c-63c2-4f7d-9896-83947da0e326",
        },
        MitreTechniqueType: operations.PostMitreTechniqueFixMitreTechniqueTypeClearK8SEvents,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostMitreTechniqueFixRequest](../../models/operations/postmitretechniquefixrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostMitreTechniqueFixResponse](../../models/operations/postmitretechniquefixresponse.md), error**

