# Cd
(*Cd*)

## Overview

APIs used to query for CD pipelines results

### Available Operations

* [DeleteCdRuleIDConnectionsRule](#deletecdruleidconnectionsrule) - delete a cd connection rule.
* [DeleteCdRuleIDServerlessRule](#deletecdruleidserverlessrule) - delete a cd serverless rule.
* [GetCd](#getcd) - Get all the CD pipelines results
* [GetCdResourceID](#getcdresourceid) - Get A single CD pipeline results
* [GetCdRuleIDConnectionsRule](#getcdruleidconnectionsrule) - get a cd connection rule.
* [GetCdRuleIDServerlessRule](#getcdruleidserverlessrule) - get a cd serverless rule.
* [PostCdConnectionsRule](#postcdconnectionsrule) - Adds cd connection rule.
* [PostCdServerlessRule](#postcdserverlessrule) - Adds cd serverless rule.
* [PutCdRuleIDConnectionsRule](#putcdruleidconnectionsrule) - update a cd connection rule.
* [PutCdRuleIDServerlessRule](#putcdruleidserverlessrule) - update a cd serverless rule.

## DeleteCdRuleIDConnectionsRule

delete a cd connection rule.

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
    res, err := s.Cd.DeleteCdRuleIDConnectionsRule(ctx, operations.DeleteCdRuleIDConnectionsRuleRequest{
        RuleID: "192323ea-9230-406f-b63c-3a2786be61ed",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteCdRuleIDConnectionsRuleRequest](../../models/operations/deletecdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteCdRuleIDConnectionsRuleResponse](../../models/operations/deletecdruleidconnectionsruleresponse.md), error**


## DeleteCdRuleIDServerlessRule

delete a cd serverless rule.

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
    res, err := s.Cd.DeleteCdRuleIDServerlessRule(ctx, operations.DeleteCdRuleIDServerlessRuleRequest{
        RuleID: "3a210f22-fa8a-464b-89de-71407e1ae662",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteCdRuleIDServerlessRuleRequest](../../models/operations/deletecdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.DeleteCdRuleIDServerlessRuleResponse](../../models/operations/deletecdruleidserverlessruleresponse.md), error**


## GetCd

Get all the CD pipelines results

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cd.GetCd(ctx, operations.GetCdRequest{
        EndTime: types.MustTimeFromString("2022-11-13T13:45:57.433Z"),
        StartTime: types.MustTimeFromString("2023-12-05T23:39:45.476Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CDPipelineResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [operations.GetCdRequest](../../models/operations/getcdrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.GetCdResponse](../../models/operations/getcdresponse.md), error**


## GetCdResourceID

Get A single CD pipeline results

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
    res, err := s.Cd.GetCdResourceID(ctx, operations.GetCdResourceIDRequest{
        ResourceID: "dbdc0e78-4707-4528-b885-f251b95127b5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdPipelineResourceResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetCdResourceIDRequest](../../models/operations/getcdresourceidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetCdResourceIDResponse](../../models/operations/getcdresourceidresponse.md), error**


## GetCdRuleIDConnectionsRule

get a cd connection rule.

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
    res, err := s.Cd.GetCdRuleIDConnectionsRule(ctx, operations.GetCdRuleIDConnectionsRuleRequest{
        RuleID: "d826ec74-f98b-4a09-b982-a21b9e98807d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdConnectionRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCdRuleIDConnectionsRuleRequest](../../models/operations/getcdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetCdRuleIDConnectionsRuleResponse](../../models/operations/getcdruleidconnectionsruleresponse.md), error**


## GetCdRuleIDServerlessRule

get a cd serverless rule.

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
    res, err := s.Cd.GetCdRuleIDServerlessRule(ctx, operations.GetCdRuleIDServerlessRuleRequest{
        RuleID: "fd218fe5-1dd3-4b03-aead-515501fcd459",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdServerlessRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCdRuleIDServerlessRuleRequest](../../models/operations/getcdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetCdRuleIDServerlessRuleResponse](../../models/operations/getcdruleidserverlessruleresponse.md), error**


## PostCdConnectionsRule

Adds cd connection rule.

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
    res, err := s.Cd.PostCdConnectionsRule(ctx, shared.CdConnectionRule{
        Destination: &shared.ConnectionRulePart{
            ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypePodNameConnectionRulePart,
        },
        Source: &shared.ConnectionRulePart{
            ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeAppTypeConnectionRulePart,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdConnectionRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.CdConnectionRule](../../models/shared/cdconnectionrule.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostCdConnectionsRuleResponse](../../models/operations/postcdconnectionsruleresponse.md), error**


## PostCdServerlessRule

Adds cd serverless rule.

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
    res, err := s.Cd.PostCdServerlessRule(ctx, shared.CdServerlessRule{
        Action: shared.ServerlessRuleActionDetect,
        Name: "string",
        Rule: shared.ServerlessRuleType{
            ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{},
            ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType,
        },
        Scope: []shared.ServerlessRuleScope{
            shared.ServerlessRuleScope{
                CloudAccount: "string",
                Regions: []string{
                    "string",
                },
            },
        },
        Status: shared.ServerlessRuleStatusDisabled,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdServerlessRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.CdServerlessRule](../../models/shared/cdserverlessrule.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostCdServerlessRuleResponse](../../models/operations/postcdserverlessruleresponse.md), error**


## PutCdRuleIDConnectionsRule

update a cd connection rule.

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
    res, err := s.Cd.PutCdRuleIDConnectionsRule(ctx, operations.PutCdRuleIDConnectionsRuleRequest{
        CdConnectionRule: shared.CdConnectionRule{
            Destination: &shared.ConnectionRulePart{
                ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeAppNameConnectionRulePart,
            },
            Source: &shared.ConnectionRulePart{
                ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeExpansionAnyConnectionRulePart,
            },
        },
        RuleID: "3491c83a-5adc-4392-92ee-d6713c48128c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdConnectionRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutCdRuleIDConnectionsRuleRequest](../../models/operations/putcdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutCdRuleIDConnectionsRuleResponse](../../models/operations/putcdruleidconnectionsruleresponse.md), error**


## PutCdRuleIDServerlessRule

update a cd serverless rule.

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
    res, err := s.Cd.PutCdRuleIDServerlessRule(ctx, operations.PutCdRuleIDServerlessRuleRequest{
        CdServerlessRule: shared.CdServerlessRule{
            Action: shared.ServerlessRuleActionDetect,
            Name: "string",
            Rule: shared.ServerlessRuleType{
                ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{},
                ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType,
            },
            Scope: []shared.ServerlessRuleScope{
                shared.ServerlessRuleScope{
                    CloudAccount: "string",
                    Regions: []string{
                        "string",
                    },
                },
            },
            Status: shared.ServerlessRuleStatusDisabled,
        },
        RuleID: "4bb3d5f9-13b9-454c-81f6-ceb379119c49",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CdServerlessRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCdRuleIDServerlessRuleRequest](../../models/operations/putcdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PutCdRuleIDServerlessRuleResponse](../../models/operations/putcdruleidserverlessruleresponse.md), error**

