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
        Type: []string{
            "Bermuda",
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
        AppID: "b5ec2f78-8d75-415b-825b-1520f3bb2d0d",
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
            "-cp",
            "-jar",
            "./*",
        },
        Cwd: testpango.String("/usr/local/bin/corp"),
        Executable: "java",
        ExecutablePath: testpango.String("/usr/bin"),
        Labels: []shared.Label{
            shared.Label{
                Key: "<key>",
                Value: "Pennsylvania",
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
        "928e88cb-3f2c-4e95-af8f-afbfe2029fce",
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
                "-cp",
                "-jar",
                "./*",
            },
            Cwd: testpango.String("/usr/local/bin/corp"),
            Executable: "java",
            ExecutablePath: testpango.String("/usr/bin"),
            Labels: []shared.Label{
                shared.Label{
                    Key: "<key>",
                    Value: "invoice",
                },
            },
            Name: "AccountingApp",
            ProcessName: testpango.String("accounting_app"),
            Type: "frontend",
        },
        AppID: "a0b1a3d7-8ca7-4d0e-a8bc-a8a0d7f81190",
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
