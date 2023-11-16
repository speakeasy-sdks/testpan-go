# APISecurityPolicies
(*APISecurityPolicies*)

## Overview

APIs used to  define and manage api security policies

### Available Operations

* [DeleteAPISecurityPolicyPolicyID](#deleteapisecuritypolicypolicyid) - Delete api security policy
* [GetAPISecurityPolicy](#getapisecuritypolicy) - Get a list of API security policies
* [GetAPISecurityPolicyPolicyIDDeleteDependencies](#getapisecuritypolicypolicyiddeletedependencies) - get dependencies which need to be handled in order to delete specified api security service
* [PostAPISecurityPolicy](#postapisecuritypolicy) - Add new API security policy
* [PutAPISecurityPolicyPolicyID](#putapisecuritypolicypolicyid) - Edit Api security policy.

## DeleteAPISecurityPolicyPolicyID

Delete api security policy

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.DeleteAPISecurityPolicyPolicyID(ctx, operations.DeleteAPISecurityPolicyPolicyIDRequest{
        PolicyID: "04ae1a0e-dcb7-4d2b-b7a6-f7ca105f8c92",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAPISecurityPolicyPolicyIDRequest](../../pkg/models/operations/deleteapisecuritypolicypolicyidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.DeleteAPISecurityPolicyPolicyIDResponse](../../pkg/models/operations/deleteapisecuritypolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAPISecurityPolicy

Get a list of API security policies

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.GetAPISecurityPolicy(ctx)
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

**[*operations.GetAPISecurityPolicyResponse](../../pkg/models/operations/getapisecuritypolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAPISecurityPolicyPolicyIDDeleteDependencies

get dependencies which need to be handled in order to delete specified api security service

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.GetAPISecurityPolicyPolicyIDDeleteDependencies(ctx, operations.GetAPISecurityPolicyPolicyIDDeleteDependenciesRequest{
        PolicyID: "d3b9d79d-ed4f-4420-a583-071cdb02c3fd",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecurityPolicyDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.GetAPISecurityPolicyPolicyIDDeleteDependenciesRequest](../../pkg/models/operations/getapisecuritypolicypolicyiddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |


### Response

**[*operations.GetAPISecurityPolicyPolicyIDDeleteDependenciesResponse](../../pkg/models/operations/getapisecuritypolicypolicyiddeletedependenciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostAPISecurityPolicy

Add new API security policy

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.PostAPISecurityPolicy(ctx, shared.APISecurityPolicyInput{
        CategoryConditions: &shared.APISecurityPolicyCategoryConditions{
            Conditions: []shared.APISecurityPolicyCategoryCondition{
                shared.APISecurityPolicyCategoryCondition{
                    Category: "string",
                    HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityMedium,
                },
            },
        },
        GlobalCondition: &shared.APISecurityPolicyGlobalCondition{
            HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityCritical,
        },
        Name: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecurityPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.APISecurityPolicyInput](../../pkg/models/shared/apisecuritypolicyinput.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PostAPISecurityPolicyResponse](../../pkg/models/operations/postapisecuritypolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutAPISecurityPolicyPolicyID

Edit Api security policy.

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.PutAPISecurityPolicyPolicyID(ctx, operations.PutAPISecurityPolicyPolicyIDRequest{
        APISecurityPolicy: shared.APISecurityPolicyInput{
            CategoryConditions: &shared.APISecurityPolicyCategoryConditions{
                Conditions: []shared.APISecurityPolicyCategoryCondition{
                    shared.APISecurityPolicyCategoryCondition{
                        Category: "string",
                        HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityMedium,
                    },
                },
            },
            GlobalCondition: &shared.APISecurityPolicyGlobalCondition{
                HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityLow,
            },
            Name: "string",
        },
        PolicyID: "5c1bef90-a301-4744-8558-c2807a30fd0b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecurityPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutAPISecurityPolicyPolicyIDRequest](../../pkg/models/operations/putapisecuritypolicypolicyidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PutAPISecurityPolicyPolicyIDResponse](../../pkg/models/operations/putapisecuritypolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
