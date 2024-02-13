# ServerlessPolicies
(*ServerlessPolicies*)

### Available Operations

* [GetServerlessPolicy](#getserverlesspolicy) - Get current serverless policy
* [PutServerlessPolicy](#putserverlesspolicy) - Set the current serverless policy

## GetServerlessPolicy

Get current serverless policy

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
    res, err := s.ServerlessPolicies.GetServerlessPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerlessPolicyResponse](../../pkg/models/operations/getserverlesspolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutServerlessPolicy

Set the current serverless policy

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
    res, err := s.ServerlessPolicies.PutServerlessPolicy(ctx, shared.ServerlessPolicy{
        DefaultRule: shared.ServerlessDefaultRuleDetectAll,
        UnidentifiedServerlessRule: shared.UnidentifiedServerlessRule{
            Action: shared.UnidentifiedServerlessRuleActionDetect,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.ServerlessPolicy](../../pkg/models/shared/serverlesspolicy.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PutServerlessPolicyResponse](../../pkg/models/operations/putserverlesspolicyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
