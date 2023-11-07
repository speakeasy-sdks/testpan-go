# RiskAssessment
(*.RiskAssessment*)

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
        IgnoredRiskID: "591339a8-adc9-48b1-82f6-1d3cce316b8d",
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
        KubernetesClusterID: "0086ba38-7819-4a93-be05-2abfdf0aef92",
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

    if res.Classes != nil {
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

    if res.Classes != nil {
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
            "a89898f1-37f8-4c09-8113-2e54dc492339",
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
        ClusterID: "af306627-d7b5-4fd9-a554-cf5effac095b",
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
        ClusterID: "e6c26498-20da-481b-bc16-7b830d3456dd",
        OwnerID: "79a8a441-40ea-4cdf-8584-73da2c62b846",
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
        ClusterID: "419337aa-84cc-4a55-afba-5a213e1dbdfd",
        OwnerID: "4e00792b-c0ad-4551-9794-d6bafdbc8352",
        RoleID: "ad0304b4-bb83-44d6-bf7f-174c17202a1e",
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
            "813cbccb-b94b-4a50-afbd-d7c394a6030d",
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
        ImageID: "6cbcdb90-f642-47ed-b640-ae8227deac5c",
        SortKey: operations.GetRiskAssessmentImageIDVulnerabilitiesQueryParamSortKeySeverity,
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
        KubernetesClusterID: "764514eb-01d8-4d87-972e-7065c0075222",
        SortKey: operations.GetRiskAssessmentKubernetesClusterIDPodsQueryParamSortKeyName,
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
                "fda4268f-7c78-46d3-87e2-782993d8ba4d",
            },
        },
        ActionType: operations.PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionTypeRemove,
        OwnerID: "c3b8ad34-72d2-4f81-a59b-f81c3cc9c10e",
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
        KubernetesClusterID: "967e4e2b-c728-4357-980a-727b8f594aad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
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
            MaxParallelism: 752536,
            MinimumSeverity: shared.VulnerabilitySeverityLow,
            Namespaces: []string{
                "string",
            },
            PeriodicJobExpression: &shared.PeriodicJobExpression{
                PeriodicJobType: shared.PeriodicJobTypeByHoursPeriodicJobExpression,
            },
        },
        KubernetesClusterID: "c1a0c988-4def-48c1-bbd4-845665df2f73",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
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
        CiPolicy: shared.CiPolicyInput{
            DockerfileScanCiPolicy: &shared.DockerfileScanCiPolicy{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleDockerfileScanSeverity: shared.DockerfileScanSeverityInfo,
            },
            Name: "string",
            VulnerabilityCiPolicy: &shared.VulnerabilityCiPolicy{
                EnforcementOption: shared.EnforcementOptionIgnore,
                PermissibleVulnerabilityLevel: shared.VulnerabilitySeverityCritical,
            },
        },
        IgnoredRiskID: "1978a3a5-2b6e-4f16-935f-df8529c311ad",
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

