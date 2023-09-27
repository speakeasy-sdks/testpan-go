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
        DefaultRule: shared.ServerlessDefaultRuleBlockAll,
        UnidentifiedServerlessRule: shared.UnidentifiedServerlessRule{
            Action: shared.UnidentifiedServerlessRuleActionBlock,
            Name: testpango.String("Wayne Romaguera"),
        },
        UserRules: []shared.ServerlessRule{
            shared.ServerlessRule{
                Action: shared.ServerlessRuleActionAllow,
                GroupName: testpango.String("sunt"),
                ID: testpango.String("2cb512c8-7824-40bf-948f-88f8f1bf0bc8"),
                Name: "Willie Wunsch III",
                Rule: shared.ServerlessRuleType{
                    ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{
                        DataAccessRisk: shared.ServerlessDataAccessRiskMedium.ToPointer(),
                        FunctionPermissionRisk: shared.ServerlessPolicyRiskLow.ToPointer(),
                        IsUnusedFunction: testpango.Bool(false),
                        PubliclyAccessibleRisk: shared.ServerlessPubliclyAccessibleRiskMedium.ToPointer(),
                        Risk: shared.ServerlessFunctionRiskLevelMedium.ToPointer(),
                        SecretsRisk: shared.ServerlessSecretsRiskNoKnownRisk.ToPointer(),
                        Vulnerability: shared.VulnerabilitySeverityUnknown.ToPointer(),
                    },
                    ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionArnServerlessRuleType,
                },
                RuleOrigin: shared.ServerlessRuleOriginUser.ToPointer(),
                Scope: []shared.ServerlessRuleScope{
                    shared.ServerlessRuleScope{
                        CloudAccount: "aperiam",
                        Regions: []string{
                            "laudantium",
                        },
                    },
                },
                Status: shared.ServerlessRuleStatusEnabled,
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

