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
        DeployerID: "4cddf857-a9e6-4187-ac6a-b21d29dfc94d",
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
        MaxResults: testpango.Float64(3772.69),
        Name: testpango.String("Clay Schaefer"),
        Offset: testpango.Float64(6143.46),
        RuleCreation: testpango.Bool(false),
        SecurityCheck: testpango.Bool(false),
        SortDir: operations.GetDeployersSortDirDesc.ToPointer(),
        SortKey: operations.GetDeployersSortKeyDeployer,
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
        KubernetesClusterID: "90066a6d-2d00-4035-9338-cec086fa21e9",
        NamespaceName: testpango.String("dicta"),
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
        DeployerID: "52cb3119-167b-48e3-88db-03408d6d364f",
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
        Deployer: testpango.String("tenetur"),
        DeployerID: "d455906d-1263-4d48-a935-c2c9e81f30be",
        DeployerType: shared.DeployerDeployerTypeOperatorDeployer,
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
            Deployer: testpango.String("recusandae"),
            DeployerID: "43202d72-1657-4650-a641-870d9d21f9ad",
            DeployerType: shared.DeployerDeployerTypeOperatorDeployer,
        },
        DeployerID: "30c4ecc1-1a08-4364-a906-8b8502a55e7f",
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

