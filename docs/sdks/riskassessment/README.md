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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.DeleteRiskAssessmentIgnoredRisksIgnoredRiskIDRequest](../../pkg/models/operations/deleteriskassessmentignoredrisksignoredriskidrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.DeleteRiskAssessmentIgnoredRisksIgnoredRiskIDResponse](../../pkg/models/operations/deleteriskassessmentignoredrisksignoredriskidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRiskAssessmentKubernetesClusterIDCancel

Cancel the runtime scan on the given cluster with the given id

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.DeleteRiskAssessmentKubernetesClusterIDCancelRequest](../../pkg/models/operations/deleteriskassessmentkubernetesclusteridcancelrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.DeleteRiskAssessmentKubernetesClusterIDCancelResponse](../../pkg/models/operations/deleteriskassessmentkubernetesclusteridcancelresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessment

Get risk assessment data for all clusters

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

**[*operations.GetRiskAssessmentResponse](../../pkg/models/operations/getriskassessmentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentIgnoredRisks

Get all the ignored risks

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

**[*operations.GetRiskAssessmentIgnoredRisksResponse](../../pkg/models/operations/getriskassessmentignoredrisksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentPermissions

Get list of clusters and their permissions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.RiskAssessment.GetRiskAssessmentPermissions(ctx, operations.GetRiskAssessmentPermissionsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetRiskAssessmentPermissionsRequest](../../pkg/models/operations/getriskassessmentpermissionsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetRiskAssessmentPermissionsResponse](../../pkg/models/operations/getriskassessmentpermissionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentPermissionsClusterID

Get all of the users permissions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetRiskAssessmentPermissionsClusterIDRequest](../../pkg/models/operations/getriskassessmentpermissionsclusteridrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDResponse](../../pkg/models/operations/getriskassessmentpermissionsclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentPermissionsClusterIDOwnerID

Get the owner permissions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRequest](../../pkg/models/operations/getriskassessmentpermissionsclusteridowneridrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDOwnerIDResponse](../../pkg/models/operations/getriskassessmentpermissionsclusteridowneridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentPermissionsClusterIDOwnerIDRoleID

Get the owner permissions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDRequest](../../pkg/models/operations/getriskassessmentpermissionsclusteridowneridroleidrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetRiskAssessmentPermissionsClusterIDOwnerIDRoleIDResponse](../../pkg/models/operations/getriskassessmentpermissionsclusteridowneridroleidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentPoll

Poll running scans

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetRiskAssessmentPollRequest](../../pkg/models/operations/getriskassessmentpollrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetRiskAssessmentPollResponse](../../pkg/models/operations/getriskassessmentpollresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentImageIDVulnerabilities

Get all images of given risk assessment pod

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetRiskAssessmentImageIDVulnerabilitiesRequest](../../pkg/models/operations/getriskassessmentimageidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetRiskAssessmentImageIDVulnerabilitiesResponse](../../pkg/models/operations/getriskassessmentimageidvulnerabilitiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRiskAssessmentKubernetesClusterIDPods

Get all risk assessments of all pods of given cluster

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.RiskAssessment.GetRiskAssessmentKubernetesClusterIDPods(ctx, operations.GetRiskAssessmentKubernetesClusterIDPodsRequest{
        KubernetesClusterID: "764514eb-01d8-4d87-972e-7065c0075222",
        SortKey: operations.GetRiskAssessmentKubernetesClusterIDPodsQueryParamSortKeyRisk,
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetRiskAssessmentKubernetesClusterIDPodsRequest](../../pkg/models/operations/getriskassessmentkubernetesclusteridpodsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetRiskAssessmentKubernetesClusterIDPodsResponse](../../pkg/models/operations/getriskassessmentkubernetesclusteridpodsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRiskAssessmentIgnoredRisks

Add ignore risk

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

**[*operations.PostRiskAssessmentIgnoredRisksResponse](../../pkg/models/operations/postriskassessmentignoredrisksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRiskAssessmentPermissionsOwnerIDApprove

add / remove permissions to /from the approved permissions list

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.RiskAssessment.PostRiskAssessmentPermissionsOwnerIDApprove(ctx, operations.PostRiskAssessmentPermissionsOwnerIDApproveRequest{
        UUIDList: shared.UUIDList{},
        ActionType: operations.PostRiskAssessmentPermissionsOwnerIDApproveQueryParamActionTypeRemove,
        OwnerID: "da4268f7-c786-4d34-be27-82993d8ba4dd",
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
| `request`                                                                                                                                          | [operations.PostRiskAssessmentPermissionsOwnerIDApproveRequest](../../pkg/models/operations/postriskassessmentpermissionsowneridapproverequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.PostRiskAssessmentPermissionsOwnerIDApproveResponse](../../pkg/models/operations/postriskassessmentpermissionsowneridapproveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRiskAssessmentKubernetesClusterIDScan

Execute a new runtime scan on the given cluster with the given configuration

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PostRiskAssessmentKubernetesClusterIDScanRequest](../../pkg/models/operations/postriskassessmentkubernetesclusteridscanrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PostRiskAssessmentKubernetesClusterIDScanResponse](../../pkg/models/operations/postriskassessmentkubernetesclusteridscanresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostRiskAssessmentKubernetesClusterIDSettings

Save the runtime scan configuration on the given cluster

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.RiskAssessment.PostRiskAssessmentKubernetesClusterIDSettings(ctx, operations.PostRiskAssessmentKubernetesClusterIDSettingsRequest{
        RiskAssessmentClusterScanConfig: shared.RiskAssessmentClusterScanConfig{
            MaxParallelism: 752536,
            MinimumSeverity: shared.VulnerabilitySeverityLow,
        },
        KubernetesClusterID: "8c1a0c98-84de-4f8c-97bd-4845665df2f7",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.PostRiskAssessmentKubernetesClusterIDSettingsRequest](../../pkg/models/operations/postriskassessmentkubernetesclusteridsettingsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PostRiskAssessmentKubernetesClusterIDSettingsResponse](../../pkg/models/operations/postriskassessmentkubernetesclusteridsettingsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutRiskAssessmentIgnoredRisksIgnoredRiskID

Edit ignore risk

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.RiskAssessment.PutRiskAssessmentIgnoredRisksIgnoredRiskID(ctx, operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest{
        CiPolicy: shared.CiPolicyInput{
            Name: "<value>",
        },
        IgnoredRiskID: "44ed1978-a3a5-42b6-af16-d35fdf8529c3",
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDRequest](../../pkg/models/operations/putriskassessmentignoredrisksignoredriskidrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.PutRiskAssessmentIgnoredRisksIgnoredRiskIDResponse](../../pkg/models/operations/putriskassessmentignoredrisksignoredriskidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
