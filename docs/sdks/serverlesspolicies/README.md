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

**[*operations.GetServerlessPolicyResponse](../../models/operations/getserverlesspolicyresponse.md), error**


## PutServerlessPolicy

Set the current serverless policy

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
    res, err := s.ServerlessPolicies.PutServerlessPolicy(ctx, shared.ServerlessPolicy{
        DefaultRule: shared.ServerlessDefaultRuleDetectAll,
        UnidentifiedServerlessRule: shared.UnidentifiedServerlessRule{
            Action: shared.UnidentifiedServerlessRuleActionDetect,
        },
        UserRules: []shared.ServerlessRule{
            shared.ServerlessRule{
                Action: shared.ServerlessRuleActionAllow,
                Name: "infrastructures",
                Rule: shared.ServerlessRuleType{
                    ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{},
                    ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionAnyServerlessRuleType,
                },
                Scope: []shared.ServerlessRuleScope{
                    shared.ServerlessRuleScope{
                        CloudAccount: "granular",
                        Regions: []string{
                            "Keys",
                        },
                    },
                },
                Status: shared.ServerlessRuleStatusDisabled,
            },
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.ServerlessPolicy](../../models/shared/serverlesspolicy.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PutServerlessPolicyResponse](../../models/operations/putserverlesspolicyresponse.md), error**

