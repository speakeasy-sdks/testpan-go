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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteTrustedSignersTrustedSignerIDRequest](../../models/operations/deletetrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.DeleteTrustedSignersTrustedSignerIDResponse](../../models/operations/deletetrustedsignerstrustedsigneridresponse.md), error**


## GetTrustedSigners

Get a list of defined trusted signers

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
    res, err := s.TrustedSigners.GetTrustedSigners(ctx, operations.GetTrustedSignersRequest{
        SortDir: operations.GetTrustedSignersSortDirAsc.ToPointer(),
        SortKey: operations.GetTrustedSignersSortKeyName.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrustedSigners != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTrustedSignersRequest](../../models/operations/gettrustedsignersrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTrustedSignersResponse](../../models/operations/gettrustedsignersresponse.md), error**


## GetTrustedSignersTrustedSignerID

get existing trusted signer

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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetTrustedSignersTrustedSignerIDRequest](../../models/operations/gettrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GetTrustedSignersTrustedSignerIDResponse](../../models/operations/gettrustedsignerstrustedsigneridresponse.md), error**


## PostTrustedSigners

Add new trusted signer

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
    res, err := s.TrustedSigners.PostTrustedSigners(ctx, shared.TrustedSignerInput{
        Keys: []shared.TrustedSignerKey{
            shared.TrustedSignerKey{
                Key: "<key>",
                Name: "Northeast Iowa",
            },
        },
        Name: "killer United Agender",
        TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
            shared.TrustedSignerCloudAccountInput{
                ID: testpango.String("339127fc-1c35-45d9-9677-cf34f19a2406"),
                Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                Validation: shared.TrustedSignerClusterValidationHash.ToPointer(),
            },
        },
        TrustedSignerClusters: []shared.TrustedSignerClusterInput{
            shared.TrustedSignerClusterInput{
                ID: testpango.String("22272998-72b1-4626-b1e4-e3bd500d0622"),
                Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                Validation: shared.TrustedSignerClusterValidationNone.ToPointer(),
            },
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.TrustedSignerInput](../../models/shared/trustedsignerinput.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostTrustedSignersResponse](../../models/operations/posttrustedsignersresponse.md), error**


## PutTrustedSignersTrustedSignerID

edit existing trusted signer

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
    res, err := s.TrustedSigners.PutTrustedSignersTrustedSignerID(ctx, operations.PutTrustedSignersTrustedSignerIDRequest{
        TrustedSignerInput: shared.TrustedSignerInput{
            Keys: []shared.TrustedSignerKey{
                shared.TrustedSignerKey{
                    Key: "<key>",
                    Name: "Chrysler Centers",
                },
            },
            Name: "so static never",
            TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
                shared.TrustedSignerCloudAccountInput{
                    ID: testpango.String("2390f3eb-b00b-4b0b-a50b-1429abb4df87"),
                    Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                    Validation: shared.TrustedSignerClusterValidationNone.ToPointer(),
                },
            },
            TrustedSignerClusters: []shared.TrustedSignerClusterInput{
                shared.TrustedSignerClusterInput{
                    ID: testpango.String("af5cddc0-6513-44a6-9fd6-1a9bf59409fe"),
                    Status: shared.TrustedSignerClusterStatusValid.ToPointer(),
                    Validation: shared.TrustedSignerClusterValidationSignature.ToPointer(),
                },
            },
        },
        TrustedSignerID: "c650d3f6-4267-4f45-b31f-42c8f00b005e",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutTrustedSignersTrustedSignerIDRequest](../../models/operations/puttrustedsignerstrustedsigneridrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.PutTrustedSignersTrustedSignerIDResponse](../../models/operations/puttrustedsignerstrustedsigneridresponse.md), error**

