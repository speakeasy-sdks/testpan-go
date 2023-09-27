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
        TrustedSignerID: "a550a656-ed33-43bb-8ce8-aa65432a986e",
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
        SortDir: operations.GetTrustedSignersSortDirDesc.ToPointer(),
        SortKey: testpango.String("odio"),
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
        TrustedSignerID: "e14ca564-0891-4500-9701-9a48f88ece7b",
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
                Key: "hic",
                Name: "Christopher Fritsch Sr.",
            },
        },
        Name: "Karen Hammes",
        TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
            shared.TrustedSignerCloudAccountInput{
                ID: testpango.String("8908162c-6beb-468a-8f65-7b7d03a1480f"),
                Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                Validation: shared.TrustedSignerClusterValidationNone.ToPointer(),
            },
        },
        TrustedSignerClusters: []shared.TrustedSignerClusterInput{
            shared.TrustedSignerClusterInput{
                ID: testpango.String("e30f069d-8106-418d-97e1-52297510da80"),
                Status: shared.TrustedSignerClusterStatusValid.ToPointer(),
                Validation: shared.TrustedSignerClusterValidationSignature.ToPointer(),
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
                    Key: "explicabo",
                    Name: "Ramona Conroy",
                },
            },
            Name: "Frances Sauer",
            TrustedSignerCloudAccounts: []shared.TrustedSignerCloudAccountInput{
                shared.TrustedSignerCloudAccountInput{
                    ID: testpango.String("702bb97e-e102-4da2-9e35-f8e01bf33eaa"),
                    Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                    Validation: shared.TrustedSignerClusterValidationSignature.ToPointer(),
                },
            },
            TrustedSignerClusters: []shared.TrustedSignerClusterInput{
                shared.TrustedSignerClusterInput{
                    ID: testpango.String("5402ac17-04bf-41cc-9fc6-1aae5eb5f0c4"),
                    Status: shared.TrustedSignerClusterStatusWarning.ToPointer(),
                    Validation: shared.TrustedSignerClusterValidationSignature.ToPointer(),
                },
            },
        },
        TrustedSignerID: "b5744d08-a226-47aa-ae79-e3c71ad31bec",
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

