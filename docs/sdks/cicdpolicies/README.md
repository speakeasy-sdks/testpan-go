# CICDPolicies
(*CICDPolicies*)

## Overview

APIs used to  define and manage CI/CD policies

### Available Operations

* [DeleteCdPolicyPolicyID](#deletecdpolicypolicyid) - Delete CD policy
* [DeleteCiPolicyPolicyID](#deletecipolicypolicyid) - Delete CI policy
* [GetCdPolicy](#getcdpolicy) - Get the current CD policy
* [GetCiPolicy](#getcipolicy) - Get the current CI policy
* [PostCdPolicy](#postcdpolicy) - Set the current CD policy. At least one CdPolicyElement should be present
* [PostCiPolicy](#postcipolicy) - Set the current CI policy
* [PutCdPolicyPolicyID](#putcdpolicypolicyid) - Edit CD policy. At least one CdPolicyElement should be present
* [PutCiPolicyPolicyID](#putcipolicypolicyid) - Edit CI policy

## DeleteCdPolicyPolicyID

Delete CD policy

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
    res, err := s.CICDPolicies.DeleteCdPolicyPolicyID(ctx, operations.DeleteCdPolicyPolicyIDRequest{
        PolicyID: "113917a9-a33e-48f7-8502-0a3844f10696",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteCdPolicyPolicyIDRequest](../../pkg/models/operations/deletecdpolicypolicyidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.DeleteCdPolicyPolicyIDResponse](../../pkg/models/operations/deletecdpolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCiPolicyPolicyID

Delete CI policy

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
    res, err := s.CICDPolicies.DeleteCiPolicyPolicyID(ctx, operations.DeleteCiPolicyPolicyIDRequest{
        PolicyID: "04ff3f56-c960-40ce-85e7-ad7f8b628cd7",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteCiPolicyPolicyIDRequest](../../pkg/models/operations/deletecipolicypolicyidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.DeleteCiPolicyPolicyIDResponse](../../pkg/models/operations/deletecipolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCdPolicy

Get the current CD policy

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
    res, err := s.CICDPolicies.GetCdPolicy(ctx)
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

**[*operations.GetCdPolicyResponse](../../pkg/models/operations/getcdpolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCiPolicy

Get the current CI policy

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
    res, err := s.CICDPolicies.GetCiPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CiPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCiPolicyResponse](../../pkg/models/operations/getcipolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCdPolicy

Set the current CD policy. At least one CdPolicyElement should be present

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
    res, err := s.CICDPolicies.PostCdPolicy(ctx, shared.CdPolicyInput{
        Deployers: []string{
            "e20e4f6e-3e04-4f9f-8904-433d8246a999",
        },
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CdPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.CdPolicyInput](../../pkg/models/shared/cdpolicyinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.PostCdPolicyResponse](../../pkg/models/operations/postcdpolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCiPolicy

Set the current CI policy

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
    res, err := s.CICDPolicies.PostCiPolicy(ctx, shared.CiPolicyInput{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CiPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.CiPolicyInput](../../pkg/models/shared/cipolicyinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.PostCiPolicyResponse](../../pkg/models/operations/postcipolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutCdPolicyPolicyID

Edit CD policy. At least one CdPolicyElement should be present

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
    res, err := s.CICDPolicies.PutCdPolicyPolicyID(ctx, operations.PutCdPolicyPolicyIDRequest{
        CdPolicy: shared.CdPolicyInput{
            Deployers: []string{
                "75218fad-dbdc-48d5-b27f-e1d8ecd9e791",
            },
            Name: "<value>",
        },
        PolicyID: "545666e4-dfb7-44ef-ada8-1a0d950f62fe",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CdPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PutCdPolicyPolicyIDRequest](../../pkg/models/operations/putcdpolicypolicyidrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PutCdPolicyPolicyIDResponse](../../pkg/models/operations/putcdpolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutCiPolicyPolicyID

Edit CI policy

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
    res, err := s.CICDPolicies.PutCiPolicyPolicyID(ctx, operations.PutCiPolicyPolicyIDRequest{
        CiPolicy: shared.CiPolicyInput{
            Name: "<value>",
        },
        PolicyID: "a0c10c59-7151-45cd-be24-f5dcfd347fd8",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CiPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PutCiPolicyPolicyIDRequest](../../pkg/models/operations/putcipolicypolicyidrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PutCiPolicyPolicyIDResponse](../../pkg/models/operations/putcipolicypolicyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
