# Gateways
(*Gateways*)

### Available Operations

* [DeleteGatewaysGatewayID](#deletegatewaysgatewayid) - Delete gateway
* [GetGateways](#getgateways) - Get gateways
* [GetGatewaysClusters](#getgatewaysclusters) - Get clusters info
* [GetGatewaysGatewayIDDownloadBundle](#getgatewaysgatewayiddownloadbundle) - Get a GW installation script
* [PostGateways](#postgateways) - Add new gateway
* [PutGatewaysGatewayID](#putgatewaysgatewayid) - Edit gateway

## DeleteGatewaysGatewayID

Delete gateway

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
    res, err := s.Gateways.DeleteGatewaysGatewayID(ctx, operations.DeleteGatewaysGatewayIDRequest{
        GatewayID: "0059ee9e-2eb4-40ca-97cc-9e2e4879b4b7",
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
| `request`                                                                                              | [operations.DeleteGatewaysGatewayIDRequest](../../models/operations/deletegatewaysgatewayidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.DeleteGatewaysGatewayIDResponse](../../models/operations/deletegatewaysgatewayidresponse.md), error**


## GetGateways

Get gateways

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
    res, err := s.Gateways.GetGateways(ctx, operations.GetGatewaysRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateways != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetGatewaysRequest](../../models/operations/getgatewaysrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetGatewaysResponse](../../models/operations/getgatewaysresponse.md), error**


## GetGatewaysClusters

Get clusters info

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
    res, err := s.Gateways.GetGatewaysClusters(ctx, operations.GetGatewaysClustersRequest{
        GatewayType: operations.GetGatewaysClustersGatewayTypeTykInternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GatewayClusterInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetGatewaysClustersRequest](../../models/operations/getgatewaysclustersrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetGatewaysClustersResponse](../../models/operations/getgatewaysclustersresponse.md), error**


## GetGatewaysGatewayIDDownloadBundle

In order to install,  extract and run "./install_bundle.sh"

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
    res, err := s.Gateways.GetGatewaysGatewayIDDownloadBundle(ctx, operations.GetGatewaysGatewayIDDownloadBundleRequest{
        GatewayID: "d7df551a-98f0-4f5b-9bdc-d69676cf90f0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetGatewaysGatewayIDDownloadBundle200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetGatewaysGatewayIDDownloadBundleRequest](../../models/operations/getgatewaysgatewayiddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetGatewaysGatewayIDDownloadBundleResponse](../../models/operations/getgatewaysgatewayiddownloadbundleresponse.md), error**


## PostGateways

Add new gateway

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
    res, err := s.Gateways.PostGateways(ctx, shared.Gateway{
        ClusterName: "Tools",
        Name: "Indiana",
        Type: shared.GatewayTypeTykInternal,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateway != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Gateway](../../models/shared/gateway.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostGatewaysResponse](../../models/operations/postgatewaysresponse.md), error**


## PutGatewaysGatewayID

Edit gateway

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
    res, err := s.Gateways.PutGatewaysGatewayID(ctx, operations.PutGatewaysGatewayIDRequest{
        Gateway: shared.Gateway{
            ClusterName: "magenta",
            Name: "Southwest",
            Type: shared.GatewayTypeTykInternal,
        },
        GatewayID: "c87e1369-f02c-4e6c-b577-9c9c8e935111",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Gateway != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutGatewaysGatewayIDRequest](../../models/operations/putgatewaysgatewayidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PutGatewaysGatewayIDResponse](../../models/operations/putgatewaysgatewayidresponse.md), error**

