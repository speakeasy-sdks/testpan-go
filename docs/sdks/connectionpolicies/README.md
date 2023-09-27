# ConnectionPolicies
(*ConnectionPolicies*)

## Overview

APIs used to  define and manage connection policies

### Available Operations

* [GetConnectionsPolicy](#getconnectionspolicy) - Get current connection policy
* [GetConnectionsPolicyHistory](#getconnectionspolicyhistory) - Get the history of the connection policies
* [GetConnectionsPolicyKafkaActions](#getconnectionspolicykafkaactions) - Get the a list of kafka actions
* [GetConnectionsPolicyKafkaKubernetesClusterIDBrokers](#getconnectionspolicykafkakubernetesclusteridbrokers) - Get the a list of kafka brokers
* [GetConnectionsPolicyKafkaKubernetesClusterIDTopics](#getconnectionspolicykafkakubernetesclusteridtopics) - Get the a list of kafka topics
* [GetConnectionsPolicySearchOptions](#getconnectionspolicysearchoptions) - Get the current connection policy filter option
* [GetServerlessPolicyHistory](#getserverlesspolicyhistory) - Get the history of the serverless policies
* [PutConnectionsPolicy](#putconnectionspolicy) - Set the current connection policy

## GetConnectionsPolicy

Get current connection policy

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicy(ctx, operations.GetConnectionsPolicyRequest{
        PolicyFilter: testpango.String("ipsam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionsPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetConnectionsPolicyRequest](../../models/operations/getconnectionspolicyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetConnectionsPolicyResponse](../../models/operations/getconnectionspolicyresponse.md), error**


## GetConnectionsPolicyHistory

Get the history of the connection policies

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyHistory(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionPolicyHistories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetConnectionsPolicyHistoryResponse](../../models/operations/getconnectionspolicyhistoryresponse.md), error**


## GetConnectionsPolicyKafkaActions

Get the a list of kafka actions

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaActions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetConnectionsPolicyKafkaActions200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetConnectionsPolicyKafkaActionsResponse](../../models/operations/getconnectionspolicykafkaactionsresponse.md), error**


## GetConnectionsPolicyKafkaKubernetesClusterIDBrokers

Get the a list of kafka brokers

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaKubernetesClusterIDBrokers(ctx, operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest{
        KubernetesClusterID: "a972e056-7282-427b-ad30-9470bf7a4fa8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetConnectionsPolicyKafkaKubernetesClusterIDBrokers200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest](../../models/operations/getconnectionspolicykafkakubernetesclusteridbrokersrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse](../../models/operations/getconnectionspolicykafkakubernetesclusteridbrokersresponse.md), error**


## GetConnectionsPolicyKafkaKubernetesClusterIDTopics

Get the a list of kafka topics

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaKubernetesClusterIDTopics(ctx, operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest{
        KubernetesClusterID: "7cf535a6-fae5-44eb-b60c-321f023b75d2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetConnectionsPolicyKafkaKubernetesClusterIDTopics200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest](../../models/operations/getconnectionspolicykafkakubernetesclusteridtopicsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse](../../models/operations/getconnectionspolicykafkakubernetesclusteridtopicsresponse.md), error**


## GetConnectionsPolicySearchOptions

Get the current connection policy filter option

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicySearchOptions(ctx, operations.GetConnectionsPolicySearchOptionsRequest{
        NameFilter: testpango.String("dolor"),
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetConnectionsPolicySearchOptionsRequest](../../models/operations/getconnectionspolicysearchoptionsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GetConnectionsPolicySearchOptionsResponse](../../models/operations/getconnectionspolicysearchoptionsresponse.md), error**


## GetServerlessPolicyHistory

Get the history of the serverless policies

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
    res, err := s.ConnectionPolicies.GetServerlessPolicyHistory(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionPolicyHistories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerlessPolicyHistoryResponse](../../models/operations/getserverlesspolicyhistoryresponse.md), error**


## PutConnectionsPolicy

Set the current connection policy

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
    res, err := s.ConnectionPolicies.PutConnectionsPolicy(ctx, shared.ConnectionsPolicy{
        DefaultRule: &shared.DefaultConnectionRule{
            Action: shared.ConnectionRuleActionBlock.ToPointer(),
            Type: shared.DefaultConnectionRuleTypeEnvironmentOnly.ToPointer(),
        },
        DirectPodRule: shared.DirectPodIPConnectionRule{
            Action: shared.DirectPodIPConnectionRuleActionBlock,
            IsDisabled: testpango.Bool(false),
            Name: testpango.String("Miss Gerald Orn"),
        },
        UserRules: []shared.ConnectionsRule{
            shared.ConnectionsRule{
                Action: shared.ConnectionRuleActionAllow,
                Destination: &shared.ConnectionRulePart{
                    ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeExpansionLabelsConnectionRulePart,
                },
                GroupName: testpango.String("maiores"),
                ID: testpango.String("79f0a396-d90c-4364-b7c1-5dfbace188b1"),
                IsRuleActive: testpango.Bool(false),
                Layer7Settings: &shared.Layer7SettingsPart{
                    Layer7Protocol: shared.Layer7SettingsPartLayer7ProtocolAPIServiceLayerPart.ToPointer(),
                },
                Name: "Laverne Waelchi",
                RuleOrigin: shared.ConnectionRuleOriginAutomatedPolicy.ToPointer(),
                RuleType: shared.NetworkConnectionRuleTypeDirectPodRule.ToPointer(),
                Source: &shared.ConnectionRulePart{
                    ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeEnvironmentAnyConnectionRulePart,
                },
                Status: shared.ConnectionsRuleStatusDeleted.ToPointer(),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionsPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.ConnectionsPolicy](../../models/shared/connectionspolicy.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.PutConnectionsPolicyResponse](../../models/operations/putconnectionspolicyresponse.md), error**

