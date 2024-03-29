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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Apps.GetApps(ctx, operations.GetAppsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetAppsRequest](../../pkg/models/operations/getappsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetAppsResponse](../../pkg/models/operations/getappsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAppsAppID

Returns an App by its ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetAppsAppIDRequest](../../pkg/models/operations/getappsappidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetAppsAppIDResponse](../../pkg/models/operations/getappsappidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostApps

Define a New App

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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
| `request`                                             | [shared.App](../../pkg/models/shared/app.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAppsResponse](../../pkg/models/operations/postappsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAppsDelete

Delete several apps.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Apps.PostAppsDelete(ctx, []string{
        "928e88cb-3f2c-4e95-af8f-afbfe2029fce",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]string](../../.md)                                 | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAppsDeleteResponse](../../pkg/models/operations/postappsdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutAppsAppID

Edit the existing App

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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
            Name: "AccountingApp",
            ProcessName: testpango.String("accounting_app"),
            Type: "frontend",
        },
        AppID: "7da0b1a3-d78c-4a7d-8e68-bca8a0d7f811",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PutAppsAppIDRequest](../../pkg/models/operations/putappsappidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PutAppsAppIDResponse](../../pkg/models/operations/putappsappidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
