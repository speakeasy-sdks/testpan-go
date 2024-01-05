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
    res, err := s.ConnectionPolicies.GetConnectionsPolicy(ctx, operations.GetConnectionsPolicyRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionsPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetConnectionsPolicyRequest](../../pkg/models/operations/getconnectionspolicyrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetConnectionsPolicyResponse](../../pkg/models/operations/getconnectionspolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionsPolicyHistory

Get the history of the connection policies

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyHistory(ctx)
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

**[*operations.GetConnectionsPolicyHistoryResponse](../../pkg/models/operations/getconnectionspolicyhistoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionsPolicyKafkaActions

Get the a list of kafka actions

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaActions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetConnectionsPolicyKafkaActionsResponse](../../pkg/models/operations/getconnectionspolicykafkaactionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionsPolicyKafkaKubernetesClusterIDBrokers

Get the a list of kafka brokers

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaKubernetesClusterIDBrokers(ctx, operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest{
        KubernetesClusterID: "adf9ba62-2ac1-4d6f-9118-9c1b46212731",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersRequest](../../pkg/models/operations/getconnectionspolicykafkakubernetesclusteridbrokersrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |


### Response

**[*operations.GetConnectionsPolicyKafkaKubernetesClusterIDBrokersResponse](../../pkg/models/operations/getconnectionspolicykafkakubernetesclusteridbrokersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionsPolicyKafkaKubernetesClusterIDTopics

Get the a list of kafka topics

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicyKafkaKubernetesClusterIDTopics(ctx, operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest{
        KubernetesClusterID: "fa0332c3-3e86-408a-93f6-0cc1de785419",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsRequest](../../pkg/models/operations/getconnectionspolicykafkakubernetesclusteridtopicsrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetConnectionsPolicyKafkaKubernetesClusterIDTopicsResponse](../../pkg/models/operations/getconnectionspolicykafkakubernetesclusteridtopicsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConnectionsPolicySearchOptions

Get the current connection policy filter option

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
    res, err := s.ConnectionPolicies.GetConnectionsPolicySearchOptions(ctx, operations.GetConnectionsPolicySearchOptionsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyFilterSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetConnectionsPolicySearchOptionsRequest](../../pkg/models/operations/getconnectionspolicysearchoptionsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GetConnectionsPolicySearchOptionsResponse](../../pkg/models/operations/getconnectionspolicysearchoptionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessPolicyHistory

Get the history of the serverless policies

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
    res, err := s.ConnectionPolicies.GetServerlessPolicyHistory(ctx)
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

**[*operations.GetServerlessPolicyHistoryResponse](../../pkg/models/operations/getserverlesspolicyhistoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutConnectionsPolicy

Set the current connection policy

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
    res, err := s.ConnectionPolicies.PutConnectionsPolicy(ctx, shared.ConnectionsPolicy{
        DefaultRule: &shared.DefaultConnectionRule{},
        DirectPodRule: shared.DirectPodIPConnectionRule{
            Action: shared.DirectPodIPConnectionRuleActionDetect,
        },
        UserRules: []shared.ConnectionsRule{
            shared.ConnectionsRule{
                Action: shared.ConnectionRuleActionEncryptDirect,
                Destination: &shared.ConnectionRulePart{
                    ConnectionRulePartType: shared.ConnectionRulePartTypeKafkaConnectionRulePart,
                },
                Layer7Settings: &shared.Layer7SettingsPart{},
                Name: "string",
                Source: &shared.ConnectionRulePart{
                    ConnectionRulePartType: shared.ConnectionRulePartTypeAnyConnectionRulePart,
                },
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.ConnectionsPolicy](../../pkg/models/shared/connectionspolicy.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PutConnectionsPolicyResponse](../../pkg/models/operations/putconnectionspolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
