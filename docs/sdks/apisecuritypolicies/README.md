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
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.DeleteAPISecurityPolicyPolicyID(ctx, operations.DeleteAPISecurityPolicyPolicyIDRequest{
        PolicyID: "04ae1a0e-dcb7-4d2b-b7a6-f7ca105f8c92",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.PostAPISecurityPolicy(ctx, shared.APISecurityPolicyInput{
        Name: "<value>",
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.PutAPISecurityPolicyPolicyID(ctx, operations.PutAPISecurityPolicyPolicyIDRequest{
        APISecurityPolicy: shared.APISecurityPolicyInput{
            Name: "<value>",
        },
        PolicyID: "735c1bef-90a3-4017-8485-58c2807a30fd",
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |
