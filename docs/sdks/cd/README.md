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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteCdRuleIDConnectionsRuleRequest](../../pkg/models/operations/deletecdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.DeleteCdRuleIDConnectionsRuleResponse](../../pkg/models/operations/deletecdruleidconnectionsruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCdRuleIDServerlessRule

delete a cd serverless rule.

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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteCdRuleIDServerlessRuleRequest](../../pkg/models/operations/deletecdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteCdRuleIDServerlessRuleResponse](../../pkg/models/operations/deletecdruleidserverlessruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCd

Get all the CD pipelines results

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
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
    res, err := s.Cd.GetCd(ctx, operations.GetCdRequest{
        EndTime: types.MustTimeFromString("2023-11-14T04:42:16.390Z"),
        StartTime: types.MustTimeFromString("2024-12-05T23:05:32.860Z"),
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetCdRequest](../../pkg/models/operations/getcdrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetCdResponse](../../pkg/models/operations/getcdresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCdResourceID

Get A single CD pipeline results

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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCdResourceIDRequest](../../pkg/models/operations/getcdresourceidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetCdResourceIDResponse](../../pkg/models/operations/getcdresourceidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCdRuleIDConnectionsRule

get a cd connection rule.

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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCdRuleIDConnectionsRuleRequest](../../pkg/models/operations/getcdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetCdRuleIDConnectionsRuleResponse](../../pkg/models/operations/getcdruleidconnectionsruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCdRuleIDServerlessRule

get a cd serverless rule.

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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCdRuleIDServerlessRuleRequest](../../pkg/models/operations/getcdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetCdRuleIDServerlessRuleResponse](../../pkg/models/operations/getcdruleidserverlessruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCdConnectionsRule

Adds cd connection rule.

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
    res, err := s.Cd.PostCdConnectionsRule(ctx, shared.CdConnectionRule{})
    if err != nil {
        log.Fatal(err)
    }

    if res.CdConnectionRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.CdConnectionRule](../../pkg/models/shared/cdconnectionrule.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostCdConnectionsRuleResponse](../../pkg/models/operations/postcdconnectionsruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCdServerlessRule

Adds cd serverless rule.

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
    res, err := s.Cd.PostCdServerlessRule(ctx, shared.CdServerlessRule{
        Action: shared.ServerlessRuleActionDetect,
        Name: "<value>",
        Rule: shared.ServerlessRuleType{
            ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType,
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.CdServerlessRule](../../pkg/models/shared/cdserverlessrule.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostCdServerlessRuleResponse](../../pkg/models/operations/postcdserverlessruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutCdRuleIDConnectionsRule

update a cd connection rule.

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
    res, err := s.Cd.PutCdRuleIDConnectionsRule(ctx, operations.PutCdRuleIDConnectionsRuleRequest{
        CdConnectionRule: shared.CdConnectionRule{},
        RuleID: "0d3491c8-3a5a-4dc3-9212-eed6713c4812",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PutCdRuleIDConnectionsRuleRequest](../../pkg/models/operations/putcdruleidconnectionsrulerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PutCdRuleIDConnectionsRuleResponse](../../pkg/models/operations/putcdruleidconnectionsruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutCdRuleIDServerlessRule

update a cd serverless rule.

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
    res, err := s.Cd.PutCdRuleIDServerlessRule(ctx, operations.PutCdRuleIDServerlessRuleRequest{
        CdServerlessRule: shared.CdServerlessRule{
            Action: shared.ServerlessRuleActionDetect,
            Name: "<value>",
            Rule: shared.ServerlessRuleType{
                ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType,
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutCdRuleIDServerlessRuleRequest](../../pkg/models/operations/putcdruleidserverlessrulerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutCdRuleIDServerlessRuleResponse](../../pkg/models/operations/putcdruleidserverlessruleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
