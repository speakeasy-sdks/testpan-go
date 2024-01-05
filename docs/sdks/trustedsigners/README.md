# TrustedSigners
(*TrustedSigners*)

## Overview

APIs used to  define and manage trusted signers

### Available Operations

* [DeleteTrustedSignersTrustedSignerID](#deletetrustedsignerstrustedsignerid) - Delete a trusted signer
* [GetTrustedSigners](#gettrustedsigners) - Get a list of defined trusted signers
* [GetTrustedSignersTrustedSignerID](#gettrustedsignerstrustedsignerid) - get existing trusted signer
* [PostTrustedSigners](#posttrustedsigners) - Add new trusted signer
* [PutTrustedSignersTrustedSignerID](#puttrustedsignerstrustedsignerid) - edit existing trusted signer

## DeleteTrustedSignersTrustedSignerID

Delete a trusted signer

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
    res, err := s.TrustedSigners.DeleteTrustedSignersTrustedSignerID(ctx, operations.DeleteTrustedSignersTrustedSignerIDRequest{
        TrustedSignerID: "72a0c357-f09f-4b54-84ce-1c9dc40be52b",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.DeleteTrustedSignersTrustedSignerIDRequest](../../pkg/models/operations/deletetrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.DeleteTrustedSignersTrustedSignerIDResponse](../../pkg/models/operations/deletetrustedsignerstrustedsigneridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTrustedSigners

Get a list of defined trusted signers

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
    res, err := s.TrustedSigners.GetTrustedSigners(ctx, operations.GetTrustedSignersRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetTrustedSignersRequest](../../pkg/models/operations/gettrustedsignersrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetTrustedSignersResponse](../../pkg/models/operations/gettrustedsignersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTrustedSignersTrustedSignerID

get existing trusted signer

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
    res, err := s.TrustedSigners.GetTrustedSignersTrustedSignerID(ctx, operations.GetTrustedSignersTrustedSignerIDRequest{
        TrustedSignerID: "3d8f6f9a-2b35-44bb-82c6-f820a67d0e04",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrustedSigner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetTrustedSignersTrustedSignerIDRequest](../../pkg/models/operations/gettrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetTrustedSignersTrustedSignerIDResponse](../../pkg/models/operations/gettrustedsignerstrustedsigneridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostTrustedSigners

Add new trusted signer

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
    res, err := s.TrustedSigners.PostTrustedSigners(ctx, shared.TrustedSignerInput{
        Keys: []shared.TrustedSignerKey{
            shared.TrustedSignerKey{
                Key: "<key>",
                Name: "string",
            },
        },
        Name: "string",
        TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
            shared.TrustedSignerCloudAccountInput{},
        },
        TrustedSignerClusters: []shared.TrustedSignerClusterInput{
            shared.TrustedSignerClusterInput{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrustedSigner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.TrustedSignerInput](../../pkg/models/shared/trustedsignerinput.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostTrustedSignersResponse](../../pkg/models/operations/posttrustedsignersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutTrustedSignersTrustedSignerID

edit existing trusted signer

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
    res, err := s.TrustedSigners.PutTrustedSignersTrustedSignerID(ctx, operations.PutTrustedSignersTrustedSignerIDRequest{
        TrustedSigner: shared.TrustedSignerInput{
            Keys: []shared.TrustedSignerKey{
                shared.TrustedSignerKey{
                    Key: "<key>",
                    Name: "string",
                },
            },
            Name: "string",
            TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
                shared.TrustedSignerCloudAccountInput{},
            },
            TrustedSignerClusters: []shared.TrustedSignerClusterInput{
                shared.TrustedSignerClusterInput{},
            },
        },
        TrustedSignerID: "8d323c1d-de95-475d-a823-90f3ebb00bb0",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PutTrustedSignersTrustedSignerIDRequest](../../pkg/models/operations/puttrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PutTrustedSignersTrustedSignerIDResponse](../../pkg/models/operations/puttrustedsignerstrustedsigneridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
