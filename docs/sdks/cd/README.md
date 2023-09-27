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
        RuleID: "e1e91e45-0ad2-4abd-8426-9802d502a94b",
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
        RuleID: "b4f63c96-9e9a-43ef-a77d-fb14cd66ae39",
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
        EndTime: types.MustTimeFromString("2022-02-13T03:59:53.583Z"),
        MaxResults: testpango.Float64(9654.17),
        Offset: testpango.Float64(6925.32),
        ResourceName: testpango.String("provident"),
        SortDir: operations.GetCdSortDirDesc.ToPointer(),
        SortKey: operations.GetCdSortKeyStatus.ToPointer(),
        StartTime: types.MustTimeFromString("2021-12-07T18:13:34.827Z"),
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
        ResourceID: "f3a66997-074b-4a44-a9b6-e2141959890a",
        SortDir: operations.GetCdResourceIDSortDirDesc.ToPointer(),
        SortKey: testpango.String("mollitia"),
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
        RuleID: "563e2516-fe4c-48b7-91e5-b7fd2ed02892",
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
        RuleID: "1cddc692-601f-4b57-ab0d-5f0d30c5fbb2",
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
        Action: shared.ConnectionRuleActionBlock.ToPointer(),
        Destination: &shared.ConnectionRulePart{
            ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeIPRangeConnectionRulePart,
        },
        GroupName: testpango.String("dignissimos"),
        ID: testpango.String("053202c7-3d5f-4e9b-90c2-8909b3fe49a8"),
        Name: testpango.String("Rene Rolfson"),
        Source: &shared.ConnectionRulePart{
            ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypePodLablesConnectionRulePart,
        },
        Status: shared.CdConnectionRuleStatusDisabled.ToPointer(),
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
        Action: shared.ServerlessRuleActionAllow,
        GroupName: testpango.String("dolorem"),
        ID: testpango.String("3323f9b7-7f3a-4410-8674-ebf69280d1ba"),
        Name: "Colleen Pagac",
        Rule: shared.ServerlessRuleType{
            ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{
                DataAccessRisk: shared.ServerlessDataAccessRiskMedium.ToPointer(),
                FunctionPermissionRisk: shared.ServerlessPolicyRiskHigh.ToPointer(),
                IsUnusedFunction: testpango.Bool(false),
                PubliclyAccessibleRisk: shared.ServerlessPubliclyAccessibleRiskMedium.ToPointer(),
                Risk: shared.ServerlessFunctionRiskLevelMedium.ToPointer(),
                SecretsRisk: shared.ServerlessSecretsRiskNoKnownRisk.ToPointer(),
                Vulnerability: shared.VulnerabilitySeverityMedium.ToPointer(),
            },
            ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionNameServerlessRuleType,
        },
        Scope: []shared.ServerlessRuleScope{
            shared.ServerlessRuleScope{
                CloudAccount: "saepe",
                Regions: []string{
                    "eius",
                },
            },
        },
        Status: shared.ServerlessRuleStatusEnabled,
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
            Action: shared.ConnectionRuleActionDetect.ToPointer(),
            Destination: &shared.ConnectionRulePart{
                ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypePodNameConnectionRulePart,
            },
            GroupName: testpango.String("optio"),
            ID: testpango.String("e5e6a95d-8a0d-4446-8e2a-f7a73cf3be45"),
            Name: testpango.String("Jeannie Leannon MD"),
            Source: &shared.ConnectionRulePart{
                ConnectionRulePartType: shared.ConnectionRulePartConnectionRulePartTypeAppAnyConnectionRulePart,
            },
            Status: shared.CdConnectionRuleStatusEnabled.ToPointer(),
        },
        RuleID: "6b5a7342-9cdb-41a8-822b-b679d2322715",
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
            Action: shared.ServerlessRuleActionBlock,
            GroupName: testpango.String("hic"),
            ID: testpango.String("0cbb1e31-b8b9-40f3-843a-1108e0adcf4b"),
            Name: "Alan Bergnaum",
            Rule: shared.ServerlessRuleType{
                ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{
                    DataAccessRisk: shared.ServerlessDataAccessRiskLow.ToPointer(),
                    FunctionPermissionRisk: shared.ServerlessPolicyRiskCritical.ToPointer(),
                    IsUnusedFunction: testpango.Bool(false),
                    PubliclyAccessibleRisk: shared.ServerlessPubliclyAccessibleRiskMedium.ToPointer(),
                    Risk: shared.ServerlessFunctionRiskLevelCritical.ToPointer(),
                    SecretsRisk: shared.ServerlessSecretsRiskRiskIdentified.ToPointer(),
                    Vulnerability: shared.VulnerabilitySeverityLow.ToPointer(),
                },
                ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionAnyServerlessRuleType,
            },
            Scope: []shared.ServerlessRuleScope{
                shared.ServerlessRuleScope{
                    CloudAccount: "delectus",
                    Regions: []string{
                        "voluptate",
                    },
                },
            },
            Status: shared.ServerlessRuleStatusEnabled,
        },
        RuleID: "ef7fbc7a-bd74-4dd3-9c0f-5d2cff7c70a4",
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

