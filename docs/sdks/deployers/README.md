# Deployers
(*Deployers*)

## Overview

APIs used to manage deployers on Secure Application

### Available Operations

* [DeleteDeployersDeployerID](#deletedeployersdeployerid) - Delete an deployer
* [GetDeployers](#getdeployers) - List all the deployers on the system
* [GetDeployersServiceAccounts](#getdeployersserviceaccounts) - List all the service account on the system
* [GetDeployersDeployerIDDeleteDependencies](#getdeployersdeployeriddeletedependencies) - get dependencies which need to be handled in order to delete specified deployer
* [PostDeployers](#postdeployers) - Create a new deployer
* [PutDeployersDeployerID](#putdeployersdeployerid) - Edit deployer definition

## DeleteDeployersDeployerID

Delete an deployer

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
    res, err := s.Deployers.DeleteDeployersDeployerID(ctx, operations.DeleteDeployersDeployerIDRequest{
        DeployerID: "3dadebf3-f849-4db8-9b45-692d5d0f8f1e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteDeployersDeployerIDRequest](../../pkg/models/operations/deletedeployersdeployeridrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteDeployersDeployerIDResponse](../../pkg/models/operations/deletedeployersdeployeridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDeployers

List all the deployers on the system

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
    res, err := s.Deployers.GetDeployers(ctx, operations.GetDeployersRequest{
        SortKey: operations.GetDeployersQueryParamSortKeyDeployer,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetDeployersRequest](../../pkg/models/operations/getdeployersrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetDeployersResponse](../../pkg/models/operations/getdeployersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDeployersServiceAccounts

List all the service account on the system

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
    res, err := s.Deployers.GetDeployersServiceAccounts(ctx, operations.GetDeployersServiceAccountsRequest{
        KubernetesClusterID: "2cec5765-5bfd-4372-88d6-c69c1df0fe41",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetDeployersServiceAccountsRequest](../../pkg/models/operations/getdeployersserviceaccountsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetDeployersServiceAccountsResponse](../../pkg/models/operations/getdeployersserviceaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDeployersDeployerIDDeleteDependencies

get dependencies which need to be handled in order to delete specified deployer

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
    res, err := s.Deployers.GetDeployersDeployerIDDeleteDependencies(ctx, operations.GetDeployersDeployerIDDeleteDependenciesRequest{
        DeployerID: "e9516b93-3d3f-40df-8a57-7f05ddcfb304",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeployerDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetDeployersDeployerIDDeleteDependenciesRequest](../../pkg/models/operations/getdeployersdeployeriddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetDeployersDeployerIDDeleteDependenciesResponse](../../pkg/models/operations/getdeployersdeployeriddeletedependenciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostDeployers

Create a new deployer

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
    res, err := s.Deployers.PostDeployers(ctx, shared.DeployerInput{
        DeployerID: "589c2d6b-f948-4feb-95d2-afb538f00cdb",
        DeployerType: shared.DeployerTypeSecureCnDeployer,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.DeployerInput](../../pkg/models/shared/deployerinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.PostDeployersResponse](../../pkg/models/operations/postdeployersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutDeployersDeployerID

Edit deployer definition

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
    res, err := s.Deployers.PutDeployersDeployerID(ctx, operations.PutDeployersDeployerIDRequest{
        Deployer: shared.DeployerInput{
            DeployerID: "bfd0fb57-ae2a-4efa-9ee4-175ba71bdf48",
            DeployerType: shared.DeployerTypeOperatorDeployer,
        },
        DeployerID: "87529aca-1222-401f-98e6-927bec6fe116",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Deployer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutDeployersDeployerIDRequest](../../pkg/models/operations/putdeployersdeployeridrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutDeployersDeployerIDResponse](../../pkg/models/operations/putdeployersdeployeridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
