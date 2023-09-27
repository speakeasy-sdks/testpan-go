# Apps
(*Apps*)

## Overview

APIs used to define apps

### Available Operations

* [GetApps](#getapps) - Returns a list of defined Apps
* [GetAppsAppID](#getappsappid) - Returns an App by its ID
* [PostApps](#postapps) - Define a New App
* [PostAppsDelete](#postappsdelete) - Delete several Apps
* [PutAppsAppID](#putappsappid) - Edit the existing App

## GetApps

Returns a list of defined Apps

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
    res, err := s.Apps.GetApps(ctx, operations.GetAppsRequest{
        DownloadAsXlsx: testpango.Bool(false),
        Name: testpango.String("Mr. Forrest Howe"),
        NoPagination: testpango.Bool(false),
        SortDir: operations.GetAppsSortDirDesc.ToPointer(),
        SortKey: operations.GetAppsSortKeyType.ToPointer(),
        Type: []string{
            "consequatur",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Apps != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetAppsRequest](../../models/operations/getappsrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetAppsResponse](../../models/operations/getappsresponse.md), error**


## GetAppsAppID

Returns an App by its ID

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
    res, err := s.Apps.GetAppsAppID(ctx, operations.GetAppsAppIDRequest{
        AppID: "01ac802e-2ec0-49ff-8f0f-816ff3477c13",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.App != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetAppsAppIDRequest](../../models/operations/getappsappidrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetAppsAppIDResponse](../../models/operations/getappsappidresponse.md), error**


## PostApps

Define a New App

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
    res, err := s.Apps.PostApps(ctx, shared.App{
        Args: []string{
            "vero",
        },
        Cwd: testpango.String("/usr/local/bin/corp"),
        Executable: "java",
        ExecutablePath: testpango.String("/usr/bin"),
        ID: testpango.String("902c1412-5b09-460a-a681-51a472af923c"),
        Labels: []shared.Label{
            shared.Label{
                Key: "quis",
                Value: "cupiditate",
            },
        },
        Name: "AccountingApp",
        ProcessName: testpango.String("accounting_app"),
        Type: "frontend",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.App != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.App](../../models/shared/app.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAppsResponse](../../models/operations/postappsresponse.md), error**


## PostAppsDelete

Delete several apps.


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
    res, err := s.Apps.PostAppsDelete(ctx, []string{
        "49f83f35-0cf8-476f-bb90-1c6ecbb4e243",
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]string](../../models//.md)                         | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAppsDeleteResponse](../../models/operations/postappsdeleteresponse.md), error**


## PutAppsAppID

Edit the existing App

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
    res, err := s.Apps.PutAppsAppID(ctx, operations.PutAppsAppIDRequest{
        App: shared.App{
            Args: []string{
                "eligendi",
            },
            Cwd: testpango.String("/usr/local/bin/corp"),
            Executable: "java",
            ExecutablePath: testpango.String("/usr/bin"),
            ID: testpango.String("f789ffaf-eda5-43e5-ae6e-0ac184c2b9c2"),
            Labels: []shared.Label{
                shared.Label{
                    Key: "magnam",
                    Value: "reprehenderit",
                },
            },
            Name: "AccountingApp",
            ProcessName: testpango.String("accounting_app"),
            Type: "frontend",
        },
        AppID: "c88373a4-0e19-442f-b2e5-5055756f5d56",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.App != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PutAppsAppIDRequest](../../models/operations/putappsappidrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PutAppsAppIDResponse](../../models/operations/putappsappidresponse.md), error**

