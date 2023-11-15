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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeleteEnvironmentsEnvIDRequest](../../pkg/models/operations/deleteenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeleteEnvironmentsEnvIDResponse](../../pkg/models/operations/deleteenvironmentsenvidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetEnvironmentsRequest](../../pkg/models/operations/getenvironmentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetEnvironmentsResponse](../../pkg/models/operations/getenvironmentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetEnvironmentsEnvIDRequest](../../pkg/models/operations/getenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetEnvironmentsEnvIDResponse](../../pkg/models/operations/getenvironmentsenvidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.GetEnvironmentsEnvIDDeleteDependenciesRequest](../../pkg/models/operations/getenvironmentsenviddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.GetEnvironmentsEnvIDDeleteDependenciesResponse](../../pkg/models/operations/getenvironmentsenviddeletedependenciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
                        Value: "string",
                    },
                },
                Vpc: shared.VPCDescriptionInput{
                    AwsAccountID: "string",
                    RegionID: "string",
                    VpcID: "string",
                },
            },
        },
        KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
            shared.KubernetesEnvironmentInput{
                KubernetesCluster: "3143be01-277d-4065-a891-6a0e622811bd",
                NamespaceLabels: []shared.Label{
                    shared.Label{
                        Key: "<key>",
                        Value: "string",
                    },
                },
                Namespaces: []string{
                    "string",
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.EnvironmentInput](../../pkg/models/shared/environmentinput.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostEnvironmentsResponse](../../pkg/models/operations/postenvironmentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
                            Value: "string",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "string",
                        RegionID: "string",
                        VpcID: "string",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "925b166e-fc31-42c0-aebd-81198d741c5b",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "string",
                        },
                    },
                    Namespaces: []string{
                        "string",
                    },
                },
            },
            Name: "Prod",
        },
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.EnvironmentInput](../../.md)                | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostEnvironmentsBatchResponse](../../pkg/models/operations/postenvironmentsbatchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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

**[*operations.PostEnvironmentsDeleteResponse](../../pkg/models/operations/postenvironmentsdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
        Environment: shared.EnvironmentInput{
            AwsEnvironments: []shared.AwsEnvironmentInput{
                shared.AwsEnvironmentInput{
                    Tags: []shared.Tag{
                        shared.Tag{
                            Key: "<key>",
                            Value: "string",
                        },
                    },
                    Vpc: shared.VPCDescriptionInput{
                        AwsAccountID: "string",
                        RegionID: "string",
                        VpcID: "string",
                    },
                },
            },
            KubernetesEnvironments: []shared.KubernetesEnvironmentInput{
                shared.KubernetesEnvironmentInput{
                    KubernetesCluster: "57cf89fe-5e65-409c-b415-d6003566df6f",
                    NamespaceLabels: []shared.Label{
                        shared.Label{
                            Key: "<key>",
                            Value: "string",
                        },
                    },
                    Namespaces: []string{
                        "string",
                    },
                },
            },
            Name: "Prod",
        },
        EnvID: "010de053-98de-43a9-acad-0419b6d55592",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutEnvironmentsEnvIDRequest](../../pkg/models/operations/putenvironmentsenvidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PutEnvironmentsEnvIDResponse](../../pkg/models/operations/putenvironmentsenvidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
