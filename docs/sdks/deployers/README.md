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
    res, err := s.Deployers.DeleteDeployersDeployerID(ctx, operations.DeleteDeployersDeployerIDRequest{
        DeployerID: "3dadebf3-f849-4db8-9b45-692d5d0f8f1e",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeleteDeployersDeployerIDRequest](../../models/operations/deletedeployersdeployeridrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeleteDeployersDeployerIDResponse](../../models/operations/deletedeployersdeployeridresponse.md), error**


## GetDeployers

List all the deployers on the system

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
    res, err := s.Deployers.GetDeployers(ctx, operations.GetDeployersRequest{
        MaxResults: testpango.Float64(757.15),
        Name: testpango.String("neural Card"),
        Offset: testpango.Float64(4706.54),
        RuleCreation: testpango.Bool(false),
        SecurityCheck: testpango.Bool(false),
        SortDir: operations.GetDeployersSortDirAsc.ToPointer(),
        SortKey: operations.GetDeployersSortKeyType,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Deployers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDeployersRequest](../../models/operations/getdeployersrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetDeployersResponse](../../models/operations/getdeployersresponse.md), error**


## GetDeployersServiceAccounts

List all the service account on the system

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
    res, err := s.Deployers.GetDeployersServiceAccounts(ctx, operations.GetDeployersServiceAccountsRequest{
        KubernetesClusterID: "2cec5765-5bfd-4372-88d6-c69c1df0fe41",
        NamespaceName: testpango.String("synthesizing"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServiceAccountInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetDeployersServiceAccountsRequest](../../models/operations/getdeployersserviceaccountsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetDeployersServiceAccountsResponse](../../models/operations/getdeployersserviceaccountsresponse.md), error**


## GetDeployersDeployerIDDeleteDependencies

get dependencies which need to be handled in order to delete specified deployer

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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.GetDeployersDeployerIDDeleteDependenciesRequest](../../models/operations/getdeployersdeployeriddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.GetDeployersDeployerIDDeleteDependenciesResponse](../../models/operations/getdeployersdeployeriddeletedependenciesresponse.md), error**


## PostDeployers

Create a new deployer

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
    res, err := s.Deployers.PostDeployers(ctx, shared.DeployerInput{
        Deployer: testpango.String("pixel"),
        DeployerID: "c2d6bf94-8feb-495d-aafb-538f00cdbb74",
        DeployerType: shared.DeployerDeployerTypeSecureCnDeployer,
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.DeployerInput](../../models/shared/deployerinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostDeployersResponse](../../models/operations/postdeployersresponse.md), error**


## PutDeployersDeployerID

Edit deployer definition

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
    res, err := s.Deployers.PutDeployersDeployerID(ctx, operations.PutDeployersDeployerIDRequest{
        DeployerInput: shared.DeployerInput{
            Deployer: testpango.String("towards Vanuatu Zirconium"),
            DeployerID: "7ae2aefa-5ee4-4175-ba71-bdf48687529a",
            DeployerType: shared.DeployerDeployerTypeSecureCnDeployer,
        },
        DeployerID: "a122201f-98e6-4927-bec6-fe116f385701",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutDeployersDeployerIDRequest](../../models/operations/putdeployersdeployeridrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PutDeployersDeployerIDResponse](../../models/operations/putdeployersdeployeridresponse.md), error**

