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
        RegistryID: "e5f28387-13f7-45c6-911a-2f9bf926c17f",
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
        SortDir: operations.GetRegistriesSortDirAsc.ToPointer(),
        SortKey: operations.GetRegistriesSortKeyURL.ToPointer(),
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
            "95b632a5-fe89-4e35-884b-fdb5be5f9972",
        },
        Credentials: &shared.RegistryCredentials{
            RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeAwsRegistryCredentials,
        },
        Type: shared.RegistryTypeAzure,
        URL: "http://svelte-curio.org",
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
            "146c9745-4e65-4c13-96b3-e4e8e106ebca",
        },
        Credentials: &shared.RegistryCredentials{
            RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeAwsRegistryCredentials,
        },
        Type: shared.RegistryTypeOther,
        URL: "http://handy-energy.net",
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
                "309c9980-a8e5-4df0-8f7a-f1c1e7dae13c",
            },
            Credentials: &shared.RegistryCredentials{
                RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeStandardRegistryCredentials,
            },
            Type: shared.RegistryTypeAws,
            URL: "http://blond-horror.org",
        },
        RegistryID: "99bd4248-5508-4257-836a-3d0e30be4155",
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
                "bb846b91-90fe-4927-838c-5cc07f3bb118",
            },
            Credentials: &shared.RegistryCredentials{
                RegistryCredentialsType: shared.RegistryCredentialsRegistryCredentialsTypeAwsRegistryCredentials,
            },
            Type: shared.RegistryTypeAws,
            URL: "https://mixed-affiliate.info",
        },
        RegistryID: "a8f4138b-1ba6-41a9-bb5f-375f1e0e9ee0",
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

