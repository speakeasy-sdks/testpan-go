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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteAPISecurityPolicyPolicyIDRequest](../../models/operations/deleteapisecuritypolicypolicyidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.DeleteAPISecurityPolicyPolicyIDResponse](../../models/operations/deleteapisecuritypolicypolicyidresponse.md), error**


## GetAPISecurityPolicy

Get a list of API security policies

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
    res, err := s.APISecurityPolicies.GetAPISecurityPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.APISecurityPolicies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAPISecurityPolicyResponse](../../models/operations/getapisecuritypolicyresponse.md), error**


## GetAPISecurityPolicyPolicyIDDeleteDependencies

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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetAPISecurityPolicyPolicyIDDeleteDependenciesRequest](../../models/operations/getapisecuritypolicypolicyiddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetAPISecurityPolicyPolicyIDDeleteDependenciesResponse](../../models/operations/getapisecuritypolicypolicyiddeletedependenciesresponse.md), error**


## PostAPISecurityPolicy

Add new API security policy

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
    res, err := s.APISecurityPolicies.PostAPISecurityPolicy(ctx, shared.APISecurityPolicyInput{
        CategoryConditions: &shared.APISecurityPolicyCategoryConditions{
            Conditions: []shared.APISecurityPolicyCategoryCondition{
                shared.APISecurityPolicyCategoryCondition{
                    Category: "solid",
                    HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityNoRisk,
                },
            },
        },
        GlobalCondition: &shared.APISecurityPolicyGlobalCondition{
            HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityMedium,
        },
        Name: "to",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.APISecurityPolicyInput](../../models/shared/apisecuritypolicyinput.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PostAPISecurityPolicyResponse](../../models/operations/postapisecuritypolicyresponse.md), error**


## PutAPISecurityPolicyPolicyID

Edit Api security policy.

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
    res, err := s.APISecurityPolicies.PutAPISecurityPolicyPolicyID(ctx, operations.PutAPISecurityPolicyPolicyIDRequest{
        APISecurityPolicyInput: shared.APISecurityPolicyInput{
            CategoryConditions: &shared.APISecurityPolicyCategoryConditions{
                Conditions: []shared.APISecurityPolicyCategoryCondition{
                    shared.APISecurityPolicyCategoryCondition{
                        Category: "deposit",
                        HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityLow,
                    },
                },
            },
            GlobalCondition: &shared.APISecurityPolicyGlobalCondition{
                HighestAcceptedSeverity: shared.APISecurityPolicyRiskSeverityHigh,
            },
            Name: "Buckinghamshire",
        },
        PolicyID: "ef90a301-7448-4558-8280-7a30fd0b43b1",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutAPISecurityPolicyPolicyIDRequest](../../models/operations/putapisecuritypolicypolicyidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PutAPISecurityPolicyPolicyIDResponse](../../models/operations/putapisecuritypolicypolicyidresponse.md), error**
