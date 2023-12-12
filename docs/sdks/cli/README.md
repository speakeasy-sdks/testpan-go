# Cli
(*Cli*)

## Overview

APIs to get the Secure Application CLI

### Available Operations

* [GetToolsCliSecurecnDeploymentCli](#gettoolsclisecurecndeploymentcli) - Get the Secure Application deployment cli

## GetToolsCliSecurecnDeploymentCli

Get the Secure Application deployment cli

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
    res, err := s.Cli.GetToolsCliSecurecnDeploymentCli(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetToolsCliSecurecnDeploymentCliResponse](../../pkg/models/operations/gettoolsclisecurecndeploymentcliresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
