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
    res, err := s.Envs.GetEnvironments(ctx, operations.GetEnvironmentsRequest{
        DownloadAsXlsx: testpango.Bool(false),
        IncludeSystemEnvs: testpango.Bool(false),
        Name: testpango.String("Pangender"),
        SortDir: operations.GetEnvironmentsSortDirAsc.ToPointer(),
        SortKey: operations.GetEnvironmentsSortKeyName.ToPointer(),
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
                ID: testpango.String("3143be01-277d-4065-a891-6a0e622811bd"),
                Tags: []shared.Tag{
                    shared.Tag{
                        Key: "<key>",
                        Value: "blue",
                    },
                },
                Vpc: shared.VPCDescriptionInput{
                    AwsAccountID: "sequestrate Square",
                    Name: testpango.String("keyboarding Southwest"),
                    RegionID: "sievert Nepalese",
                    VpcID: "Card Electronic",
                },
            },
        },
        Description: testpango.String("Managed cohesive conglomeration"),
        ID: testpango.String("8b0df411-a017-4fec-b7c0-652306ac0dd6"),
        IsSystemEnv: testpango.Bool(false),
        KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
            shared.KubernetesEnvironmentInput{
                ID: testpango.String("2ded15fb-507e-428f-8658-431e3d4048f3"),
                KubernetesCluster: "1327f0e5-119e-46a7-bae1-a8564c416b7d",
                KubernetesClusterName: testpango.String("female Hat"),
                NamespaceLabels: []shared.Label{
                    shared.Label{
                        Key: "<key>",
                        Value: "Bedfordshire Northeast",
                    },
                },
                Namespaces: []string{
                    "Tracy",
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
                    ID: testpango.String("925b166e-fc31-42c0-aebd-81198d741c5b"),
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "<key>",
                            Value: "epithelium Architect transsexual",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "Touring",
                        Name: testpango.String("maroon Designer Rubber"),
                        RegionID: "Van",
                        VpcID: "Ouguiya Practical Sausages",
                    },
                },
            },
            Description: testpango.String("Reactive multimedia open system"),
            ID: testpango.String("a8d71159-31ea-4cc0-a751-fa912e0c2e77"),
            IsSystemEnv: testpango.Bool(false),
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    ID: testpango.String("81de4db5-179e-49e8-a850-eef1ac7cc5cf"),
                    KubernetesCluster: "7cd38ac9-5b1f-4cd3-b4df-81a2f60b1759",
                    KubernetesClusterName: testpango.String("Avon male East"),
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "Missouri magenta",
                        },
                    },
                    Namespaces: []string{
                        "preface",
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
                    ID: testpango.String("57cf89fe-5e65-409c-b415-d6003566df6f"),
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "<key>",
                            Value: "Avon",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "actually communities Soul",
                        Name: testpango.String("fatally Customer Accounts"),
                        RegionID: "Cargo Mexico",
                        VpcID: "Fiat executive",
                    },
                },
            },
            Description: testpango.String("Pre-emptive bifurcated access"),
            ID: testpango.String("4790c5f1-efb6-450d-a980-1dae0e967d9d"),
            IsSystemEnv: testpango.Bool(false),
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    ID: testpango.String("0297331b-d38c-454f-9f13-9d5f8956b63c"),
                    KubernetesCluster: "b013e34e-0fd3-4e76-a959-cf530e5cf52d",
                    KubernetesClusterName: testpango.String("Carolina Integration"),
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "including Recumbent International",
                        },
                    },
                    Namespaces: []string{
                        "Loan",
                    },
                },
            },
            Name: "Prod",
        },
        EnvID: "fd1e02d4-5afd-490e-8b66-39364458cb01",
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

