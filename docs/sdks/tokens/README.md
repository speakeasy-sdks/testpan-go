# Tokens
(*Tokens*)

### Available Operations

* [DeleteTokensTokenID](#deletetokenstokenid) - Delete token
* [GetTokens](#gettokens) - Get tokens
* [GetTokensInfo](#gettokensinfo) - Get tokens by Ids
* [GetTokensTokenIDDeleteDependencies](#gettokenstokeniddeletedependencies) - get dependancies which need to be handled in order to delete specified token
* [PostTokens](#posttokens) - Add new token
* [PutTokensTokenID](#puttokenstokenid) - Edit token

## DeleteTokensTokenID

Delete token

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
    res, err := s.Tokens.DeleteTokensTokenID(ctx, operations.DeleteTokensTokenIDRequest{
        TokenID: "fc9f4844-225e-475b-b960-65c0efa6f93b",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteTokensTokenIDRequest](../../models/operations/deletetokenstokenidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteTokensTokenIDResponse](../../models/operations/deletetokenstokenidresponse.md), error**


## GetTokens

Get tokens

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
    res, err := s.Tokens.GetTokens(ctx, operations.GetTokensRequest{
        MaxResults: testpango.Float64(5639.37),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(35.7),
        SortDir: operations.GetTokensSortDirDesc.ToPointer(),
        SortKey: testpango.String("beatae"),
        TokenName: testpango.String("distinctio"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Tokens != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetTokensRequest](../../models/operations/gettokensrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetTokensResponse](../../models/operations/gettokensresponse.md), error**


## GetTokensInfo

Get tokens by Ids

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
    res, err := s.Tokens.GetTokensInfo(ctx, operations.GetTokensInfoRequest{
        TokensIds: []string{
            "8c95be12-54b7-439f-8fe7-7210d1f6558c",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APITokenInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetTokensInfoRequest](../../models/operations/gettokensinforequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetTokensInfoResponse](../../models/operations/gettokensinforesponse.md), error**


## GetTokensTokenIDDeleteDependencies

get dependancies which need to be handled in order to delete specified token

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
    res, err := s.Tokens.GetTokensTokenIDDeleteDependencies(ctx, operations.GetTokensTokenIDDeleteDependenciesRequest{
        TokenID: "99c722d2-bc0f-4940-87d9-caae042dd7ca",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TokenDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetTokensTokenIDDeleteDependenciesRequest](../../models/operations/gettokenstokeniddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetTokensTokenIDDeleteDependenciesResponse](../../models/operations/gettokenstokeniddeletedependenciesresponse.md), error**


## PostTokens

Add new token

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
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
    res, err := s.Tokens.PostTokens(ctx, shared.Token{
        Apis: []string{
            "ac9b4caa-1cfe-49e1-9df9-03907f378319",
        },
        AttributeName: testpango.String("quos"),
        AttributeType: shared.TokenAttributeTypeRequestHeader.ToPointer(),
        ExpirationDate: types.MustTimeFromString("2022-03-19T04:59:23.224Z"),
        HTTPPath: testpango.String("odit"),
        ID: testpango.String("e54a8546-6597-4c50-a33c-1471d51aaa6d"),
        Name: "Terrell Hettinger",
        VaultSecretPath: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Token != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Token](../../models/shared/token.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostTokensResponse](../../models/operations/posttokensresponse.md), error**


## PutTokensTokenID

Edit token

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
    res, err := s.Tokens.PutTokensTokenID(ctx, operations.PutTokensTokenIDRequest{
        Token: shared.Token{
            Apis: []string{
                "6487c5fc-2b86-42a0-8bef-69e100157630",
            },
            AttributeName: testpango.String("libero"),
            AttributeType: shared.TokenAttributeTypeQueryParam.ToPointer(),
            ExpirationDate: types.MustTimeFromString("2022-01-23T22:28:56.652Z"),
            HTTPPath: testpango.String("similique"),
            ID: testpango.String("fded84a3-5a41-4238-a1a7-35ac26ae33be"),
            Name: "Miss Terrence Kulas",
            VaultSecretPath: "repellat",
        },
        TokenID: "46bca110-6fe9-465b-b11d-08cf88ec9f7b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Token != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PutTokensTokenIDRequest](../../models/operations/puttokenstokenidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PutTokensTokenIDResponse](../../models/operations/puttokenstokenidresponse.md), error**

