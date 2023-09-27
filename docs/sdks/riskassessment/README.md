# RiskAssessment
(*RiskAssessment*)

## Overview

APIs used to manage risk assessment on Kubernetes clusters

### Available Operations

* [DeleteRiskAssessmentIgnoredRisksIgnoredRiskID](#deleteriskassessmentignoredrisksignoredriskid) - Delete ignored risk
* [DeleteRiskAssessmentKubernetesClusterIDCancel](#deleteriskassessmentkubernetesclusteridcancel) - Cancel the runtime scan on the given cluster with the given id
* [GetRiskAssessment](#getriskassessment) - Get risk assessment data for all clusters
* [GetRiskAssessmentIgnoredRisks](#getriskassessmentignoredrisks) - Get all the ignored risks
* [GetRiskAssessmentPermissions](#getriskassessmentpermissions) - Get list of clusters and their permissions
* [GetRiskAssessmentPermissionsClusterID](#getriskassessmentpermissionsclusterid) - Get all of the users permissions
* [GetRiskAssessmentPermissionsClusterIDOwnerID](#getriskassessmentpermissionsclusteridownerid) - Get the owner permissions
* [GetRiskAssessmentPermissionsClusterIDOwnerIDRoleID](#getriskassessmentpermissionsclusteridowneridroleid) - Get the owner permissions
* [GetRiskAssessmentPoll](#getriskassessmentpoll) - Poll running scans
* [GetRiskAssessmentImageIDVulnerabilities](#getriskassessmentimageidvulnerabilities) - Get all images of given risk assessment pod
* [GetRiskAssessmentKubernetesClusterIDPods](#getriskassessmentkubernetesclusteridpods) - Get all risk assessments of all pods of given cluster
* [PostRiskAssessmentIgnoredRisks](#postriskassessmentignoredrisks) - Add ignore risk
* [PostRiskAssessmentPermissionsOwnerIDApprove](#postriskassessmentpermissionsowneridapprove) - add / remove permissions to /from the approved permissions list
* [PostRiskAssessmentKubernetesClusterIDScan](#postriskassessmentkubernetesclusteridscan) - Execute a new runtime scan on the given cluster with the given configuration
* [PostRiskAssessmentKubernetesClusterIDSettings](#postriskassessmentkubernetesclusteridsettings) - Save the runtime scan configuration on the given cluster
* [PutRiskAssessmentIgnoredRisksIgnoredRiskID](#putriskassessmentignoredrisksignoredriskid) - Edit ignore risk

## DeleteRiskAssessmentIgnoredRisksIgnoredRiskID

Delete ignored risk

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
    res, err := s.RiskAssessment.DeleteRiskAssessmentIgnoredRisksIgnoredRiskID(ctx, operations.DeleteRiskAssessmentIgnoredRisksIgnoredRiskIDRequest{
        IgnoredRiskID: "bc74b86c-ecc7-44f7-bb48-48bd6a6f0441",
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteRiskAssessmentIgnoredRisksIgnoredRiskIDRequest](../../models/operations/deleteriskassessmentignoredrisksignoredriskidrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.DeleteRiskAssessmentIgnoredRisksIgnoredRiskIDResponse](../../models/operations/deleteriskassessmentignoredrisksignoredriskidresponse.md), error**


## DeleteRiskAssessmentKubernetesClusterIDCancel

Cancel the runtime scan on the given cluster with the given id

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
    res, err := s.RiskAssessment.DeleteRiskAssessmentKubernetesClusterIDCancel(ctx, operations.DeleteRiskAssessmentKubernetesClusterIDCancelRequest{
        KubernetesClusterID: "d2c3b808-0943-473e-8604-59bebbad02f2",
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteRiskAssessmentKubernetesClusterIDCancelRequest](../../models/operations/deleteriskassessmentkubernetesclusteridcancelrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.DeleteRiskAssessmentKubernetesClusterIDCancelResponse](../../models/operations/deleteriskassessmentkubernetesclusteridcancelresponse.md), error**


## GetRiskAssessment

Get risk assessment data for all clusters

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
    res, err := s.RiskAssessment.GetRiskAssessment(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskAssessmentClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetRiskAssessmentResponse](../../models/operations/getriskassessmentresponse.md), error**


## GetRiskAssessmentIgnoredRisks

Get all the ignored risks

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
    res, err := s.RiskAssessment.GetRiskAssessmentIgnoredRisks(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.IgnoredRisks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetRiskAssessmentIgnoredRisksResponse](../../models/operations/getriskassessmentignoredrisksresponse.md), error**


## GetRiskAssessmentPermissions

Get list of clusters and their permissions

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
    res, err := s.RiskAssessment.GetRiskAssessmentPermissions(ctx, operations.GetRiskAssessmentPermissionsRequest{
        ClustersIds: []string{
            "586bcf15-2558-4daa-95be-6cd02756c354",
        },
        IncludeSystemOwners: testpango.Bool(false),
        PermissionRisk: operations.GetRiskAssessmentPermissionsPermissionRiskHigh.ToPointer(),
        SortDir: operations.GetRiskAssessmentPermissionsSortDirDesc.ToPointer(),
        SortKey: testpango.String("labore"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ClusterPermissions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetRiskAssessmentPermissionsRequest](../../models/operations/getriskassessmentpermissionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetRiskAssessmentPermissionsResponse](../../models/operations/getriskassessmentpermissionsresponse.md), error**


## GetRiskAssessmentPermissionsClusterID

Get all of the users permissions

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
    res, err := s.RiskAssessment.GetRiskAssessmentPermissionsClusterID(ctx, operations.GetRiskAssessmentPermissionsClusterIDRequest{
        ClusterID: "32b47e17-63c5-4208-823e-9802d82f0d45",
        IncludeSystemOwners: testpango.Bool(false),
        MaxResults: testpango.Float64(9204.88),
        NamespaceName: testpango.String("soluta"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(2626.14),
        Owner: testpango.String("fuga"),
        OwnerType: operations.GetRiskAssessmentPermissionsClusterIDOwnerTypeUser.ToPointer(),
        PermissionRisk: operations.GetRiskAssessmentPermissionsClusterIDPermissionRiskHigh.ToPointer(),
        SortDir: operations.GetRiskAssessmentPermissionsClusterIDSortDirAsc.ToPointer(),
        SortKey: operations.GetRiskAssessmentPermissionsClusterIDSortKeyOwner.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetRiskAssessmentPermissionsClusterIDRequest](../../models/operations/getriskassessmentpermissionsclusteridrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDResponse](../../models/operations/getriskassessmentpermissionsclusteridresponse.md), error**


## GetRiskAssessmentPermissionsClusterIDOwnerID

Get the owner permissions

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
    res, err := s.RiskAssessment.GetRiskAssessmentPermissionsClusterIDOwnerID(ctx, operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRequest{
        ClusterID: "4ee5cfc1-8edc-47f7-87e3-2e04b3d3ed0c",
        IsApproved: testpango.Bool(false),
        OwnerID: "5670ef42-bd3c-49f1-8c50-3f6c39bcd0a6",
        SortDir: operations.GetRiskAssessmentPermissionsClusterIDOwnerIDSortDirAsc.ToPointer(),
        SortKey: testpango.String("perspiciatis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PermissionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRequest](../../models/operations/getriskassessmentpermissionsclusteridowneridrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDOwnerIDResponse](../../models/operations/getriskassessmentpermissionsclusteridowneridresponse.md), error**


## GetRiskAssessmentPermissionsClusterIDOwnerIDRoleID

Get the owner permissions

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
    res, err := s.RiskAssessment.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleID(ctx, operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest{
        ClusterID: "0f957f38-5189-4ad7-af80-7aae03f33ca7",
        OwnerID: "9fb9de40-32ba-426f-9368-ba9216bcb415",
        RoleID: "835c7364-1723-4133-adc0-46bc5163bbca",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PermissionRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest](../../models/operations/getriskassessmentpermissionsclusteridowneridroleidrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse](../../models/operations/getriskassessmentpermissionsclusteridowneridroleidresponse.md), error**


## GetRiskAssessmentPoll

Poll running scans

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
    res, err := s.RiskAssessment.GetRiskAssessmentPoll(ctx, operations.GetRiskAssessmentPollRequest{
        RiskAssessmentPollKey: []string{
            "49227c42-c22c-4553-9049-5c5dbb3c57c1",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskAssessmentClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetRiskAssessmentPollRequest](../../models/operations/getriskassessmentpollrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetRiskAssessmentPollResponse](../../models/operations/getriskassessmentpollresponse.md), error**


## GetRiskAssessmentImageIDVulnerabilities

Get all images of given risk assessment pod

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
    res, err := s.RiskAssessment.GetRiskAssessmentImageIDVulnerabilities(ctx, operations.GetRiskAssessmentImageIDVulnerabilitiesRequest{
        ImageID: "e4981e8a-a257-4ddc-9912-ebde64bfcc54",
        MaxResults: testpango.Float64(4362.35),
        Offset: testpango.Float64(6168.21),
        SortDir: operations.GetRiskAssessmentImageIDVulnerabilitiesSortDirDesc.ToPointer(),
        SortKey: "quaerat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskAssessmentVulnerabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetRiskAssessmentImageIDVulnerabilitiesRequest](../../models/operations/getriskassessmentimageidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetRiskAssessmentImageIDVulnerabilitiesResponse](../../models/operations/getriskassessmentimageidvulnerabilitiesresponse.md), error**


## GetRiskAssessmentKubernetesClusterIDPods

Get all risk assessments of all pods of given cluster

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
    res, err := s.RiskAssessment.GetRiskAssessmentKubernetesClusterIDPods(ctx, operations.GetRiskAssessmentKubernetesClusterIDPodsRequest{
        DownloadAsXlsx: testpango.Bool(false),
        KubernetesClusterID: "015dfa79-6206-4bef-ab0a-3e42c1aa010e",
        MaxResults: testpango.Float64(5668.96),
        NamespacesNamesFilter: testpango.String("fuga"),
        Offset: testpango.Float64(6332.93),
        SortDir: operations.GetRiskAssessmentKubernetesClusterIDPodsSortDirDesc.ToPointer(),
        SortKey: operations.GetRiskAssessmentKubernetesClusterIDPodsSortKeyName,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RiskAssessmentPods != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.GetRiskAssessmentKubernetesClusterIDPodsRequest](../../models/operations/getriskassessmentkubernetesclusteridpodsrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.GetRiskAssessmentKubernetesClusterIDPodsResponse](../../models/operations/getriskassessmentkubernetesclusteridpodsresponse.md), error**


## PostRiskAssessmentIgnoredRisks

Add ignore risk

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
    res, err := s.RiskAssessment.PostRiskAssessmentIgnoredRisks(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.IgnoredRisk != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostRiskAssessmentIgnoredRisksResponse](../../models/operations/postriskassessmentignoredrisksresponse.md), error**


## PostRiskAssessmentPermissionsOwnerIDApprove

add / remove permissions to /from the approved permissions list

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
    res, err := s.RiskAssessment.PostRiskAssessmentPermissionsOwnerIDApprove(ctx, operations.PostRiskAssessmentPermissionsOwnerIDApproveRequest{
        UUIDList: shared.UUIDList{
            UUIDList: []string{
                "e9135586-d18f-49f9-ba4b-fad2bf7d67ca",
            },
        },
        ActionType: operations.PostRiskAssessmentPermissionsOwnerIDApproveActionTypeRemove,
        OwnerID: "4ad99b41-d612-4435-b187-0cf68b03ad42",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PostRiskAssessmentPermissionsOwnerIDApproveRequest](../../models/operations/postriskassessmentpermissionsowneridapproverequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PostRiskAssessmentPermissionsOwnerIDApproveResponse](../../models/operations/postriskassessmentpermissionsowneridapproveresponse.md), error**


## PostRiskAssessmentKubernetesClusterIDScan

Execute a new runtime scan on the given cluster with the given configuration

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
    res, err := s.RiskAssessment.PostRiskAssessmentKubernetesClusterIDScan(ctx, operations.PostRiskAssessmentKubernetesClusterIDScanRequest{
        KubernetesClusterID: "1bd43d1f-0cb0-4a00-83eb-22d9b3a70d94",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostRiskAssessmentKubernetesClusterIDScan201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PostRiskAssessmentKubernetesClusterIDScanRequest](../../models/operations/postriskassessmentkubernetesclusteridscanrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.PostRiskAssessmentKubernetesClusterIDScanResponse](../../models/operations/postriskassessmentkubernetesclusteridscanresponse.md), error**


## PostRiskAssessmentKubernetesClusterIDSettings

Save the runtime scan configuration on the given cluster

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
    res, err := s.RiskAssessment.PostRiskAssessmentKubernetesClusterIDSettings(ctx, operations.PostRiskAssessmentKubernetesClusterIDSettingsRequest{
        RiskAssessmentClusterScanConfig: shared.RiskAssessmentClusterScanConfig{
            MaxParallelism: 973256,
            MinimumSeverity: shared.VulnerabilitySeverityHigh,
            Namespaces: []string{
                "deserunt",
            },
            PeriodicJobExpression: &shared.PeriodicJobExpression{
                PeriodicJobType: shared.PeriodicJobExpressionPeriodicJobTypeByHoursPeriodicJobExpression,
            },
            RunDockerfileScan: testpango.Bool(false),
        },
        KubernetesClusterID: "41c57d1f-edc2-4050-938d-c3ce185472f9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostRiskAssessmentKubernetesClusterIDSettings201ApplicationJSONUUIDString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.PostRiskAssessmentKubernetesClusterIDSettingsRequest](../../models/operations/postriskassessmentkubernetesclusteridsettingsrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.PostRiskAssessmentKubernetesClusterIDSettingsResponse](../../models/operations/postriskassessmentkubernetesclusteridsettingsresponse.md), error**


## PutRiskAssessmentIgnoredRisksIgnoredRiskID

Edit ignore risk

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
    res, err := s.RiskAssessment.PutRiskAssessmentIgnoredRisksIgnoredRiskID(ctx, operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest{
        CiPolicyInput: shared.CiPolicyInput{
            Description: testpango.String("necessitatibus"),
            DockerfileScanCiPolicy: &shared.DockerfileScanCiPolicy{
                EnforcementOption: shared.EnforcementOptionIgnore,
                PermissibleDockerfileScanSeverity: shared.DockerfileScanSeverityWarn,
            },
            Name: "Arthur Kerluke",
            VulnerabilityCiPolicy: &shared.VulnerabilityCiPolicy{
                EnforcementOption: shared.EnforcementOptionIgnore,
                PermissibleVulnerabilityLevel: shared.VulnerabilitySeverityHigh,
            },
        },
        IgnoredRiskID: "e3444eac-8b3a-4287-9c6c-1fe606d07d2a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IgnoredRisk != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest](../../models/operations/putriskassessmentignoredrisksignoredriskidrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse](../../models/operations/putriskassessmentignoredrisksignoredriskidresponse.md), error**

