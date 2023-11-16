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

**[*operations.GetKubernetesAPIPolicyResponse](../../pkg/models/operations/getkubernetesapipolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetKubernetesAPIPolicyHistory

Get the history of the Kubernetes API policies

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyHistory(ctx)
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

**[*operations.GetKubernetesAPIPolicyHistoryResponse](../../pkg/models/operations/getkubernetesapipolicyhistoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetKubernetesAPIPolicyKubernetesResources

Get the Kubernetes resource list

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyKubernetesResources(ctx)
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

**[*operations.GetKubernetesAPIPolicyKubernetesResourcesResponse](../../pkg/models/operations/getkubernetesapipolicykubernetesresourcesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetKubernetesAPIPolicyKubernetesUsers

Get the Kubernetes user list

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyKubernetesUsers(ctx)
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

**[*operations.GetKubernetesAPIPolicyKubernetesUsersResponse](../../pkg/models/operations/getkubernetesapipolicykubernetesusersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetKubernetesAPIPolicyRecommendedRules

Get the recommended Kubernetes API rules

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
    res, err := s.ClusterEventsPolicies.GetKubernetesAPIPolicyRecommendedRules(ctx)
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

**[*operations.GetKubernetesAPIPolicyRecommendedRulesResponse](../../pkg/models/operations/getkubernetesapipolicyrecommendedrulesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutKubernetesAPIPolicy

set the current Kubernetes API policy

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
    res, err := s.ClusterEventsPolicies.PutKubernetesAPIPolicy(ctx, shared.KubernetesAPIPolicy{
        DefaultRule: &shared.DefaultKubernetesAPIRule{},
        UserRules: []shared.KubernetesAPIRule{
            shared.KubernetesAPIRule{
                KubernetesAPIRuleType: shared.KubernetesAPIRuleTypeKubernetesAPIRecommendedRule,
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.KubernetesAPIPolicy](../../pkg/models/shared/kubernetesapipolicy.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PutKubernetesAPIPolicyResponse](../../pkg/models/operations/putkubernetesapipolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
