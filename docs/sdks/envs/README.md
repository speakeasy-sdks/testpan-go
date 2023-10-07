# Envs
(*Envs*)

## Overview

APIs used to define environments

### Available Operations

* [DeleteEnvironmentsEnvID](#deleteenvironmentsenvid)
* [GetEnvironments](#getenvironments) - List all defined Secure Application environments
* [GetEnvironmentsEnvID](#getenvironmentsenvid) - get a Secure Application environment
* [GetEnvironmentsEnvIDDeleteDependencies](#getenvironmentsenviddeletedependencies) - get dependencies which need to be handled in order to delete specified environment
* [PostEnvironments](#postenvironments) - Add a new environment
* [PostEnvironmentsBatch](#postenvironmentsbatch) - Add a number of  Secure Application environments
* [PostEnvironmentsDelete](#postenvironmentsdelete) - Delete multiple Secure Application environments
* [PutEnvironmentsEnvID](#putenvironmentsenvid) - Edit an existing Secure Application environment

## DeleteEnvironmentsEnvID

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
    res, err := s.Envs.DeleteEnvironmentsEnvID(ctx, operations.DeleteEnvironmentsEnvIDRequest{
        EnvID: "b0aabb9f-e7ee-44ba-ba12-8ad45974b0df",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteEnvironmentsEnvIDRequest](../../models/operations/deleteenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.DeleteEnvironmentsEnvIDResponse](../../models/operations/deleteenvironmentsenvidresponse.md), error**


## GetEnvironments

List all defined Secure Application environments

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
    res, err := s.Envs.GetEnvironments(ctx, operations.GetEnvironmentsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Environments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetEnvironmentsRequest](../../models/operations/getenvironmentsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetEnvironmentsResponse](../../models/operations/getenvironmentsresponse.md), error**


## GetEnvironmentsEnvID

get a Secure Application environment

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
    res, err := s.Envs.GetEnvironmentsEnvID(ctx, operations.GetEnvironmentsEnvIDRequest{
        EnvID: "559017e3-f060-4096-aa0f-bf777791d889",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Environments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetEnvironmentsEnvIDRequest](../../models/operations/getenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetEnvironmentsEnvIDResponse](../../models/operations/getenvironmentsenvidresponse.md), error**


## GetEnvironmentsEnvIDDeleteDependencies

get dependencies which need to be handled in order to delete specified environment

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
    res, err := s.Envs.GetEnvironmentsEnvIDDeleteDependencies(ctx, operations.GetEnvironmentsEnvIDDeleteDependenciesRequest{
        EnvID: "28b60297-9556-4e6f-93e6-0cc1c215a603",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDependencyElementEnvironments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetEnvironmentsEnvIDDeleteDependenciesRequest](../../models/operations/getenvironmentsenviddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetEnvironmentsEnvIDDeleteDependenciesResponse](../../models/operations/getenvironmentsenviddeletedependenciesresponse.md), error**


## PostEnvironments

Add a  Secure Application environment. An instance should be contained in a single environment. The environment is defined by a VPC and an optional tag. If a tag is supplied, only instances in the specified VPC with the given tag will belong to the new environment.


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
    res, err := s.Envs.PostEnvironments(ctx, shared.EnvironmentInput{
        AwsEnvironments: []shared.AwsEnvironmentInput{
            shared.AwsEnvironmentInput{
                Tags: []shared.Tag{
                    shared.Tag{
                        Key: "<key>",
                        Value: "South",
                    },
                },
                Vpc: shared.VPCDescriptionInput{
                    AwsAccountID: "Seaborgium",
                    RegionID: "Northeast",
                    VpcID: "invoice Honduras",
                },
            },
        },
        KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
            shared.KubernetesEnvironmentInput{
                KubernetesCluster: "5a8916a0-e622-4811-bd3c-8c7fc23dd96f",
                NamespaceLabels: []shared.Label{
                    shared.Label{
                        Key: "<key>",
                        Value: "Southwest group",
                    },
                },
                Namespaces: []string{
                    "sievert",
                },
            },
        },
        Name: "Prod",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Environment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.EnvironmentInput](../../models/shared/environmentinput.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostEnvironmentsResponse](../../models/operations/postenvironmentsresponse.md), error**


## PostEnvironmentsBatch

Add a number of new Secure Application environments. This is similar to the 'Add environment' method, but for multiple environments.


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
    res, err := s.Envs.PostEnvironmentsBatch(ctx, []shared.EnvironmentInput{
        shared.EnvironmentInput{
            AwsEnvironments: []shared.AwsEnvironmentInput{
                shared.AwsEnvironmentInput{
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "<key>",
                            Value: "Maine ampere",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "support round",
                        RegionID: "East",
                        VpcID: "Coast Facilitator bypassing",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "198d741c-5bdf-45a6-87e9-5a1cd6d26b83",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "West Customer",
                        },
                    },
                    Namespaces: []string{
                        "except",
                    },
                },
            },
            Name: "Prod",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Environments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.EnvironmentInput](../../models//.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostEnvironmentsBatchResponse](../../models/operations/postenvironmentsbatchresponse.md), error**


## PostEnvironmentsDelete

Delete multiple Secure Application environments.


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
    res, err := s.Envs.PostEnvironmentsDelete(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostEnvironmentsDeleteResponse](../../models/operations/postenvironmentsdeleteresponse.md), error**


## PutEnvironmentsEnvID

Edit an existing Secure Application environment.


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
    res, err := s.Envs.PutEnvironmentsEnvID(ctx, operations.PutEnvironmentsEnvIDRequest{
        EnvironmentInput: shared.EnvironmentInput{
            AwsEnvironments: []shared.AwsEnvironmentInput{
                shared.AwsEnvironmentInput{
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "<key>",
                            Value: "1080p into",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "tromp Upgradable",
                        RegionID: "paradigms eligendi",
                        VpcID: "Steel revolutionize Checking",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "3566df6f-010d-4e05-b98d-e3a9acad0419",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "portals Lakes West",
                        },
                    },
                    Namespaces: []string{
                        "zero",
                    },
                },
            },
            Name: "Prod",
        },
        EnvID: "efb650de-9801-4dae-8e96-7d9d0297331b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Environment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutEnvironmentsEnvIDRequest](../../models/operations/putenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PutEnvironmentsEnvIDResponse](../../models/operations/putenvironmentsenvidresponse.md), error**

