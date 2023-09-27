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
        GatewayID: "4257b992-c8db-4da6-a61e-fa2198258fd0",
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
    res, err := s.Gateways.GetGateways(ctx, operations.GetGatewaysRequest{
        MaxResults: testpango.Float64(6302.86),
        Name: testpango.String("Ramiro Reilly"),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(4554),
        SortDir: operations.GetGatewaysSortDirDesc.ToPointer(),
    })
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
        GatewayType: operations.GetGatewaysClustersGatewayTypeKongInternal,
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
        GatewayID: "d3ef0496-40d6-4a18-b1c8-7adf596fdf1a",
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
        ClusterName: "possimus",
        Description: testpango.String("praesentium"),
        ID: testpango.String("37ae80c1-c19c-495b-a998-678fa3f69699"),
        Name: "Kristine Wilderman",
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
            ClusterName: "quod",
            Description: testpango.String("repudiandae"),
            ID: testpango.String("03614448-c797-47a0-af2f-536028efeef9"),
            Name: "Alicia Bernier",
            Type: shared.GatewayTypeF5BigIP,
        },
        GatewayID: "d7e253f4-c157-4dea-a717-0f445accf667",
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

