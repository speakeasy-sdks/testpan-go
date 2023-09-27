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
    res, err := s.CICDPolicies.DeleteCdPolicyPolicyID(ctx, operations.DeleteCdPolicyPolicyIDRequest{
        PolicyID: "5626d436-813f-416d-9f5f-ce6c556146c3",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteCdPolicyPolicyIDRequest](../../models/operations/deletecdpolicypolicyidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteCdPolicyPolicyIDResponse](../../models/operations/deletecdpolicypolicyidresponse.md), error**


## DeleteCiPolicyPolicyID

Delete CI policy

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
    res, err := s.CICDPolicies.DeleteCiPolicyPolicyID(ctx, operations.DeleteCiPolicyPolicyIDRequest{
        PolicyID: "e250fb00-8c42-4e14-9aac-366c8dd6b144",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteCiPolicyPolicyIDRequest](../../models/operations/deletecipolicypolicyidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteCiPolicyPolicyIDResponse](../../models/operations/deletecipolicypolicyidresponse.md), error**


## GetCdPolicy

Get the current CD policy

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
    res, err := s.CICDPolicies.GetCdPolicy(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CdPolicies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCdPolicyResponse](../../models/operations/getcdpolicyresponse.md), error**


## GetCiPolicy

Get the current CI policy

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

**[*operations.GetCiPolicyResponse](../../models/operations/getcipolicyresponse.md), error**


## PostCdPolicy

Set the current CD policy. At least one CdPolicyElement should be present

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
    res, err := s.CICDPolicies.PostCdPolicy(ctx, shared.CdPolicyInput{
        APISecurityCdPolicy: &shared.APISecurityCdPolicyElement{
            APISecurityProfile: "29074747-78a7-4bd4-a6d2-8c10ab3cdca4",
            EnforcementOption: shared.EnforcementOptionFail,
        },
        Deployers: []string{
            "51904e52-3c7e-40bc-b178-e4796f2a70c6",
        },
        Description: testpango.String("quas"),
        Name: "Eugene Leuschke",
        PermissionCDPolicy: &shared.CdPolicyElement{
            EnforcementOption: shared.EnforcementOptionIgnore,
            PermissibleVulnerabilityLevel: shared.CDPipelineFindingRiskNoRisk,
        },
        SecretCDPolicy: &shared.SecretsCdPolicyElement{
            EnforcementOption: shared.EnforcementOptionIgnore,
            PermissibleVulnerabilityLevel: shared.CDPipelineSecretsFindingRiskNoKnownRisk,
        },
        SecurityContextCDPolicy: &shared.CdPolicyElement{
            EnforcementOption: shared.EnforcementOptionFail,
            PermissibleVulnerabilityLevel: shared.CDPipelineFindingRiskMedium,
        },
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.CdPolicyInput](../../models/shared/cdpolicyinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostCdPolicyResponse](../../models/operations/postcdpolicyresponse.md), error**


## PostCiPolicy

Set the current CI policy

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
    res, err := s.CICDPolicies.PostCiPolicy(ctx, shared.CiPolicyInput{
        Description: testpango.String("fugit"),
        DockerfileScanCiPolicy: &shared.DockerfileScanCiPolicy{
            EnforcementOption: shared.EnforcementOptionIgnore,
            PermissibleDockerfileScanSeverity: shared.DockerfileScanSeverityInfo,
        },
        Name: "Rose Turner",
        VulnerabilityCiPolicy: &shared.VulnerabilityCiPolicy{
            EnforcementOption: shared.EnforcementOptionFail,
            PermissibleVulnerabilityLevel: shared.VulnerabilitySeverityMedium,
        },
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.CiPolicyInput](../../models/shared/cipolicyinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostCiPolicyResponse](../../models/operations/postcipolicyresponse.md), error**


## PutCdPolicyPolicyID

Edit CD policy. At least one CdPolicyElement should be present

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
    res, err := s.CICDPolicies.PutCdPolicyPolicyID(ctx, operations.PutCdPolicyPolicyIDRequest{
        CdPolicyInput: shared.CdPolicyInput{
            APISecurityCdPolicy: &shared.APISecurityCdPolicyElement{
                APISecurityProfile: "ee17cbe6-1e6b-47b9-9bc0-ab3c20c4f378",
                EnforcementOption: shared.EnforcementOptionIgnore,
            },
            Deployers: []string{
                "fd871f99-dd2e-4fd1-a1aa-6f1e674bdb04",
            },
            Description: testpango.String("sapiente"),
            Name: "Marion Kihn",
            PermissionCDPolicy: &shared.CdPolicyElement{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleVulnerabilityLevel: shared.CDPipelineFindingRiskMedium,
            },
            SecretCDPolicy: &shared.SecretsCdPolicyElement{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleVulnerabilityLevel: shared.CDPipelineSecretsFindingRiskRiskIdentified,
            },
            SecurityContextCDPolicy: &shared.CdPolicyElement{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleVulnerabilityLevel: shared.CDPipelineFindingRiskMedium,
            },
        },
        PolicyID: "ea19f1d1-7051-4339-9080-86a1840394c2",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PutCdPolicyPolicyIDRequest](../../models/operations/putcdpolicypolicyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PutCdPolicyPolicyIDResponse](../../models/operations/putcdpolicypolicyidresponse.md), error**


## PutCiPolicyPolicyID

Edit CI policy

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
    res, err := s.CICDPolicies.PutCiPolicyPolicyID(ctx, operations.PutCiPolicyPolicyIDRequest{
        CiPolicyInput: shared.CiPolicyInput{
            Description: testpango.String("voluptas"),
            DockerfileScanCiPolicy: &shared.DockerfileScanCiPolicy{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleDockerfileScanSeverity: shared.DockerfileScanSeverityWarn,
            },
            Name: "Elisa Mosciski",
            VulnerabilityCiPolicy: &shared.VulnerabilityCiPolicy{
                EnforcementOption: shared.EnforcementOptionFail,
                PermissibleVulnerabilityLevel: shared.VulnerabilitySeverityCritical,
            },
        },
        PolicyID: "0642dac7-af51-45cc-813a-a63aae8d6786",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PutCiPolicyPolicyIDRequest](../../models/operations/putcipolicypolicyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PutCiPolicyPolicyIDResponse](../../models/operations/putcipolicypolicyidresponse.md), error**

