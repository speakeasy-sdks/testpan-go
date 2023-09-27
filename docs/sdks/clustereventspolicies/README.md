# ClusterEventsPolicies
(*ClusterEventsPolicies*)

## Overview

APIs used to  define and manage cluster events policies

### Available Operations

* [GetKubernetesAPIPolicy](#getkubernetesapipolicy) - Get current Kubernetes API policy
* [GetKubernetesAPIPolicyHistory](#getkubernetesapipolicyhistory) - Get the history of the Kubernetes API policies
* [GetKubernetesAPIPolicyKubernetesResources](#getkubernetesapipolicykubernetesresources) - Get the Kubernetes resource list
* [GetKubernetesAPIPolicyKubernetesUsers](#getkubernetesapipolicykubernetesusers) - Get the Kubernetes user list
* [GetKubernetesAPIPolicyRecommendedRules](#getkubernetesapipolicyrecommendedrules) - Get the recommended Kubernetes API rules
* [PutKubernetesAPIPolicy](#putkubernetesapipolicy) - set the current Kubernetes API policy

## GetKubernetesAPIPolicy

Get current Kubernetes API policy

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesAPIPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKubernetesAPIPolicyResponse](../../models/operations/getkubernetesapipolicyresponse.md), error**


## GetKubernetesAPIPolicyHistory

Get the history of the Kubernetes API policies

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyHistory(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesAPIPolicyHistories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKubernetesAPIPolicyHistoryResponse](../../models/operations/getkubernetesapipolicyhistoryresponse.md), error**


## GetKubernetesAPIPolicyKubernetesResources

Get the Kubernetes resource list

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyKubernetesResources(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesResources != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKubernetesAPIPolicyKubernetesResourcesResponse](../../models/operations/getkubernetesapipolicykubernetesresourcesresponse.md), error**


## GetKubernetesAPIPolicyKubernetesUsers

Get the Kubernetes user list

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyKubernetesUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesUsersByTypes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKubernetesAPIPolicyKubernetesUsersResponse](../../models/operations/getkubernetesapipolicykubernetesusersresponse.md), error**


## GetKubernetesAPIPolicyRecommendedRules

Get the recommended Kubernetes API rules

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyRecommendedRules(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RecommendedKubernetesAPIRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKubernetesAPIPolicyRecommendedRulesResponse](../../models/operations/getkubernetesapipolicyrecommendedrulesresponse.md), error**


## PutKubernetesAPIPolicy

set the current Kubernetes API policy

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
    res, err := s.ClusterEventsPolicies.PutKubernetesAPIPolicy(ctx, shared.KubernetesAPIPolicy{
        DefaultRule: &shared.DefaultKubernetesAPIRule{
            Action: shared.KubernetesAPIRuleActionIgnore.ToPointer(),
        },
        UserRules: []shared.KubernetesAPIRule{
            shared.KubernetesAPIRule{
                KubernetesAPIRuleType: shared.KubernetesAPIRuleKubernetesAPIRuleTypeKubernetesAPIRecommendedRule,
                RuleOrigin: shared.KubernetesAPIRuleOriginUser.ToPointer(),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesAPIPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.KubernetesAPIPolicy](../../models/shared/kubernetesapipolicy.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PutKubernetesAPIPolicyResponse](../../models/operations/putkubernetesapipolicyresponse.md), error**

