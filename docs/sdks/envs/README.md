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
                        Value: "Electronic",
                    },
                },
                Vpc: shared.VPCDescriptionInput{
                    AwsAccountID: "sticky",
                    RegionID: "Seaborgium",
                    VpcID: "Bosnia",
                },
            },
        },
        KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
            shared.KubernetesEnvironmentInput{
                KubernetesCluster: "77d065a8-916a-40e6-a281-1bd3c8c7fc23",
                NamespaceLabels: []shared.Label{
                    shared.Label{
                        Key: "<key>",
                        Value: "substantial",
                    },
                },
                Namespaces: []string{
                    "Genderqueer",
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
                            Value: "Cisgender",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "ampere",
                        RegionID: "Dinar",
                        VpcID: "round",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "312c0aeb-d811-498d-b41c-5bdf5a607e95",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "Associate",
                        },
                    },
                    Namespaces: []string{
                        "Touring",
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
                            Value: "Networked",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "Tricycle",
                        RegionID: "pixel",
                        VpcID: "tromp",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "5e6509cb-415d-4600-b566-df6f010de053",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "Soul",
                        },
                    },
                    Namespaces: []string{
                        "Tesla",
                    },
                },
            },
            Name: "Prod",
        },
        EnvID: "3a9acad0-419b-46d5-9592-04790c5f1efb",
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

