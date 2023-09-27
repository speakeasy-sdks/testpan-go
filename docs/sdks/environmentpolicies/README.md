# EnvironmentPolicies
(*EnvironmentPolicies*)

## Overview

APIs used to  define and manage environment policies

### Available Operations

* [GetAppsPolicy](#getappspolicy) - Get the current Apps policy
* [GetAppsPolicyHistory](#getappspolicyhistory) - Get the history of Apps policies
* [GetAppsPolicySearchOptions](#getappspolicysearchoptions) - Get the current Apps policy filter option
* [PutAppsPolicy](#putappspolicy) - Set the current Apps policy

## GetAppsPolicy

Get the current Apps policy

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
    res, err := s.EnvironmentPolicies.GetAppsPolicy(ctx, operations.GetAppsPolicyRequest{
        PolicyFilter: testpango.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAppsPolicyRequest](../../models/operations/getappspolicyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetAppsPolicyResponse](../../models/operations/getappspolicyresponse.md), error**


## GetAppsPolicyHistory

Get the history of Apps policies

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
    res, err := s.EnvironmentPolicies.GetAppsPolicyHistory(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppPolicyHistories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAppsPolicyHistoryResponse](../../models/operations/getappspolicyhistoryresponse.md), error**


## GetAppsPolicySearchOptions

Get the current Apps policy filter option

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
    res, err := s.EnvironmentPolicies.GetAppsPolicySearchOptions(ctx, operations.GetAppsPolicySearchOptionsRequest{
        NameFilter: testpango.String("neque"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyFilterSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAppsPolicySearchOptionsRequest](../../models/operations/getappspolicysearchoptionsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetAppsPolicySearchOptionsResponse](../../models/operations/getappspolicysearchoptionsresponse.md), error**


## PutAppsPolicy

Set the current Apps policy

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
    res, err := s.EnvironmentPolicies.PutAppsPolicy(ctx, shared.AppPolicy{
        DefaultRule: shared.DefaultRuleBlockAll.ToPointer(),
        UnidentifiedPodsRule: &shared.UnidentifiedPodsRule{
            Action: shared.UnidentifiedPodsRuleActionBlock,
            Name: testpango.String("Ricky Hilpert"),
        },
        UserRules: []shared.AppRule{
            shared.AppRule{
                App: &shared.WorkloadRuleType{
                    WorkloadRuleType: shared.WorkloadRuleTypeWorkloadRuleTypeAppTypeWorkloadRuleType,
                },
                GroupName: testpango.String("voluptatem"),
                ID: testpango.String("a319f4ba-df94-47c9-a867-bc4242666581"),
                Name: "Rosemarie Steuber",
                RuleOrigin: shared.AppRuleOriginAutomatedPolicy.ToPointer(),
                RuleTypeProperties: shared.AppRuleType{
                    RuleType: shared.AppRuleTypeRuleTypeViolationRuleType,
                },
                Scope: &shared.WorkloadRuleScopeType{
                    WorkloadRuleScopeType: shared.WorkloadRuleScopeTypeEnumEnvironmentNameRuleType,
                },
                Status: shared.AppRuleStatusDisabled,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.AppPolicy](../../models/shared/apppolicy.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PutAppsPolicyResponse](../../models/operations/putappspolicyresponse.md), error**

