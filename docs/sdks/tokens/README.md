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
        TokenID: "b10248b8-2d53-4623-885c-a738cc4f3785",
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
        MaxResults: testpango.Float64(681.78),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(7366.65),
        SortDir: operations.GetTokensSortDirAsc.ToPointer(),
        SortKey: testpango.String("what"),
        TokenName: testpango.String("policy Rustic"),
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
            "17c3fed3-a3bf-4acc-82ca-e618f8f9bdf3",
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
        TokenID: "8e7108bc-33f0-4915-b49d-5a4a26ffa567",
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
            "06eb110c-ef48-45c9-b334-9ef284ebe70b",
        },
        AttributeName: testpango.String("offensive flexibility Gate"),
        AttributeType: shared.TokenAttributeTypeRequestHeader.ToPointer(),
        ExpirationDate: types.MustTimeFromString("2022-12-06T03:02:47.254Z"),
        HTTPPath: testpango.String("Gasoline Bicycle"),
        ID: testpango.String("15c6181f-23a7-44d2-945a-aec34929a503"),
        Name: "Cyclocross",
        VaultSecretPath: "intermediate array",
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
                "92f997c4-3e7b-4827-80b5-81f98e4dc9a1",
            },
            AttributeName: testpango.String("ah card Mercedes"),
            AttributeType: shared.TokenAttributeTypeRequestHeader.ToPointer(),
            ExpirationDate: types.MustTimeFromString("2023-05-19T15:11:35.927Z"),
            HTTPPath: testpango.String("Ball"),
            ID: testpango.String("6d3d22b8-df11-4333-9e59-e4bbc71d1a3f"),
            Name: "deposit Checking Rap",
            VaultSecretPath: "Chevrolet error",
        },
        TokenID: "650a257d-f5fe-417c-b321-bf84e462e505",
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

