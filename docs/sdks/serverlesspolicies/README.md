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
            Name: testpango.String("infrastructures solutions"),
        },
        UserRules: []shared.ServerlessRule{
            shared.ServerlessRule{
                Action: shared.ServerlessRuleActionDetect,
                GroupName: testpango.String("Dodge"),
                ID: testpango.String("6eaf47ac-061b-4704-854c-19159eac10ed"),
                Name: "Convertible Shilling",
                Rule: shared.ServerlessRuleType{
                    ServerlessFunctionValidation: &shared.ServerlessFunctionValidation{
                        DataAccessRisk: shared.ServerlessDataAccessRiskLow.ToPointer(),
                        FunctionPermissionRisk: shared.ServerlessPolicyRiskHigh.ToPointer(),
                        IsUnusedFunction: testpango.Bool(false),
                        PubliclyAccessibleRisk: shared.ServerlessPubliclyAccessibleRiskLow.ToPointer(),
                        Risk: shared.ServerlessFunctionRiskLevelLow.ToPointer(),
                        SecretsRisk: shared.ServerlessSecretsRiskNoKnownRisk.ToPointer(),
                        Vulnerability: shared.VulnerabilitySeverityLow.ToPointer(),
                    },
                    ServerlessRuleType: shared.ServerlessRuleTypeServerlessRuleTypeFunctionNameServerlessRuleType,
                },
                RuleOrigin: shared.ServerlessRuleOriginSystem.ToPointer(),
                Scope: []shared.ServerlessRuleScope{
                    shared.ServerlessRuleScope{
                        CloudAccount: "Latin West",
                        Regions: []string{
                            "Configuration",
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

