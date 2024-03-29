# AgentManagement
(*AgentManagement*)

## Overview

APIs use to  interact with  agents

### Available Operations

* [GetAgents](#getagents) - List all installed agents
* [PostAgentsAgentIDUpdate](#postagentsagentidupdate) - Update the agent with the given id to the latest agent version
* [PostAgentsAgentIDUpdateState](#postagentsagentidupdatestate) - Update the status of an agent with the given id

## GetAgents

List all installed agents

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
    res, err := s.AgentManagement.GetAgents(ctx, operations.GetAgentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Agents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetAgentsRequest](../../pkg/models/operations/getagentsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetAgentsResponse](../../pkg/models/operations/getagentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAgentsAgentIDUpdate

Update the agent with the given id to the latest agent version

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
    res, err := s.AgentManagement.PostAgentsAgentIDUpdate(ctx, operations.PostAgentsAgentIDUpdateRequest{
        AgentID: "0a0835d7-8d36-4444-8529-411e73a9b7a8",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostAgentsAgentIDUpdateRequest](../../pkg/models/operations/postagentsagentidupdaterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PostAgentsAgentIDUpdateResponse](../../pkg/models/operations/postagentsagentidupdateresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.APIResponse | 402                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |

## PostAgentsAgentIDUpdateState

Update the status of an agent with the given id

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
    res, err := s.AgentManagement.PostAgentsAgentIDUpdateState(ctx, operations.PostAgentsAgentIDUpdateStateRequest{
        AgentStatusUpdate: shared.AgentStatusUpdate{},
        AgentID: "34a187e9-3552-49e2-8694-f733a8b3f850",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostAgentsAgentIDUpdateStateRequest](../../pkg/models/operations/postagentsagentidupdatestaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostAgentsAgentIDUpdateStateResponse](../../pkg/models/operations/postagentsagentidupdatestateresponse.md), error**
| Error Object          | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.APIResponse | 402                   | application/json      |
| sdkerrors.SDKError    | 4xx-5xx               | */*                   |
