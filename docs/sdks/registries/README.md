# Registries
(*Registries*)

## Overview

APIs used to  define and manage registries

### Available Operations

* [DeleteRegistriesRegistryID](#deleteregistriesregistryid) - Delete a registry
* [GetRegistries](#getregistries) - Get a list of defined registries
* [PostRegistries](#postregistries) - Add new registry
* [PostRegistriesTestConnection](#postregistriestestconnection) - test registry connection
* [PostRegistriesTestConnectionRegistryID](#postregistriestestconnectionregistryid) - test registry connection
* [PutRegistriesRegistryID](#putregistriesregistryid) - edit existing registry

## DeleteRegistriesRegistryID

Delete a registry

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
    res, err := s.Registries.DeleteRegistriesRegistryID(ctx, operations.DeleteRegistriesRegistryIDRequest{
        RegistryID: "b7021a52-046b-464e-99fb-0e67e094fdfe",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteRegistriesRegistryIDRequest](../../models/operations/deleteregistriesregistryidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.DeleteRegistriesRegistryIDResponse](../../models/operations/deleteregistriesregistryidresponse.md), error**


## GetRegistries

Get a list of defined registries

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
    res, err := s.Registries.GetRegistries(ctx, operations.GetRegistriesRequest{
        SortDir: operations.GetRegistriesSortDirDesc.ToPointer(),
        SortKey: testpango.String("exercitationem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Registries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetRegistriesRequest](../../models/operations/getregistriesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetRegistriesResponse](../../models/operations/getregistriesresponse.md), error**


## PostRegistries

Add new registry

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
    res, err := s.Registries.PostRegistries(ctx, shared.RegistryInput{
        ClusterIds: []string{
            "540ef53a-34a1-4b8f-a997-31adc05d85ae",
        },
        Credentials: &shared.RegistryCredentials{
            RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeAwsRegistryCredentials,
        },
        Type: shared.RegistryTypeOther,
        URL: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Registry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.RegistryInput](../../models/shared/registryinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostRegistriesResponse](../../models/operations/postregistriesresponse.md), error**


## PostRegistriesTestConnection

test registry connection

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
    res, err := s.Registries.PostRegistriesTestConnection(ctx, shared.RegistryInput{
        ClusterIds: []string{
            "b70fb387-4290-4d33-a561-eca16ef89451",
        },
        Credentials: &shared.RegistryCredentials{
            RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeJfrogRegistryCredentials,
        },
        Type: shared.RegistryTypeOther,
        URL: "ducimus",
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.RegistryInput](../../models/shared/registryinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostRegistriesTestConnectionResponse](../../models/operations/postregistriestestconnectionresponse.md), error**


## PostRegistriesTestConnectionRegistryID

test registry connection

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
    res, err := s.Registries.PostRegistriesTestConnectionRegistryID(ctx, operations.PostRegistriesTestConnectionRegistryIDRequest{
        RegistryInput: shared.RegistryInput{
            ClusterIds: []string{
                "6eeeb518-c4da-41fa-9355-12f06d4e5b72",
            },
            Credentials: &shared.RegistryCredentials{
                RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeJfrogRegistryCredentials,
            },
            Type: shared.RegistryTypeAws,
            URL: "doloribus",
        },
        RegistryID: "548568a0-424e-400a-9d6e-b9434645d030",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PostRegistriesTestConnectionRegistryIDRequest](../../models/operations/postregistriestestconnectionregistryidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PostRegistriesTestConnectionRegistryIDResponse](../../models/operations/postregistriestestconnectionregistryidresponse.md), error**


## PutRegistriesRegistryID

edit existing registry

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
    res, err := s.Registries.PutRegistriesRegistryID(ctx, operations.PutRegistriesRegistryIDRequest{
        RegistryInput: shared.RegistryInput{
            ClusterIds: []string{
                "84fbba5c-ceff-45cb-81fe-51e528a45ac8",
            },
            Credentials: &shared.RegistryCredentials{
                RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeAwsRegistryCredentials,
            },
            Type: shared.RegistryTypeJfrog,
            URL: "praesentium",
        },
        RegistryID: "5f8bc2ca-ba8d-4a41-a7dd-597ff4711aa1",
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
| `request`                                                                                              | [operations.PutRegistriesRegistryIDRequest](../../models/operations/putregistriesregistryidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PutRegistriesRegistryIDResponse](../../models/operations/putregistriesregistryidresponse.md), error**

