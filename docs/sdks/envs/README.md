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
        EnvID: "1fcb4c59-3ec1-42cd-aad0-ec7afedbd80d",
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
    res, err := s.Envs.GetEnvironments(ctx, operations.GetEnvironmentsRequest{
        DownloadAsXlsx: testpango.Bool(false),
        IncludeSystemEnvs: testpango.Bool(false),
        Name: testpango.String("Barry Funk"),
        SortDir: operations.GetEnvironmentsSortDirAsc.ToPointer(),
        SortKey: testpango.String("esse"),
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
        EnvID: "f9390c58-8809-483d-abf9-ef3ffdd9f7f0",
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
        EnvID: "79af4d35-724c-4db0-b4d2-81187d56844e",
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
                ID: testpango.String("ded85a90-65e6-428b-9fc2-032b6c879923"),
                Tags: []shared.Tag{
                    shared.Tag{
                        Key: "quidem",
                        Value: "molestiae",
                    },
                },
                Vpc: shared.VPCDescriptionInput{
                    AwsAccountID: "debitis",
                    Name: testpango.String("Rosa Hand"),
                    RegionID: "asperiores",
                    VpcID: "reprehenderit",
                },
            },
        },
        Description: testpango.String("deserunt"),
        ID: testpango.String("e12c6891-f82c-4e11-9717-2305377dcfa8"),
        IsSystemEnv: testpango.Bool(false),
        KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
            shared.KubernetesEnvironmentInput{
                ID: testpango.String("9df975e3-5668-4609-ae9c-3ddc5f111dea"),
                KubernetesCluster: "1026d541-a4d1-490f-ab21-780bccc0dbbd",
                KubernetesClusterName: testpango.String("illum"),
                NamespaceLabels: []shared.Label{
                    shared.Label{
                        Key: "soluta",
                        Value: "magnam",
                    },
                },
                Namespaces: []string{
                    "laudantium",
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
                    ID: testpango.String("4708fb4e-391e-46bc-958c-4c4e54599ea3"),
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "modi",
                            Value: "aspernatur",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "explicabo",
                        Name: testpango.String("Melissa Vandervort"),
                        RegionID: "qui",
                        VpcID: "accusantium",
                    },
                },
            },
            Description: testpango.String("consequatur"),
            ID: testpango.String("ce78a1bd-8fb7-4a0a-916c-e723d4097fa3"),
            IsSystemEnv: testpango.Bool(false),
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    ID: testpango.String("0e9af725-b291-4220-b0d8-3f5aeb7799d2"),
                    KubernetesCluster: "2e8c1f84-9382-45fd-842c-876c2c2dfb4c",
                    KubernetesClusterName: testpango.String("hic"),
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "quo",
                            Value: "illo",
                        },
                    },
                    Namespaces: []string{
                        "nobis",
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
                    ID: testpango.String("76230f84-1fb1-4bd2-bfdb-14db6be5a685"),
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "excepturi",
                            Value: "natus",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "deleniti",
                        Name: testpango.String("Adam D'Amore"),
                        RegionID: "eos",
                        VpcID: "voluptatem",
                    },
                },
            },
            Description: testpango.String("temporibus"),
            ID: testpango.String("a16fc2b2-71a2-489c-97e8-54e90439d222"),
            IsSystemEnv: testpango.Bool(false),
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    ID: testpango.String("46569462-4070-484f-bab3-7cef02225194"),
                    KubernetesCluster: "db55410a-dc66-49af-90a2-6c7cdc981f06",
                    KubernetesClusterName: testpango.String("laudantium"),
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "unde",
                            Value: "corrupti",
                        },
                    },
                    Namespaces: []string{
                        "quae",
                    },
                },
            },
            Name: "Prod",
        },
        EnvID: "d6bb33cf-aa34-48c3-9bf4-07ee4fcf0c42",
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

