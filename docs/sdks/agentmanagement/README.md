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
    res, err := s.AgentManagement.GetAgents(ctx, operations.GetAgentsRequest{
        DownloadAsXlsx: testpango.Bool(false),
        EnvironmentName: []string{
            "facilis",
        },
        HostName: []string{
            "cum",
        },
        Risk: []GetAgentsRisk{
            operations.GetAgentsRiskMedium,
        },
        SortDir: operations.GetAgentsSortDirAsc.ToPointer(),
        SortKey: operations.GetAgentsSortKeyEnvironmentName.ToPointer(),
        Status: []GetAgentsStatus{
            operations.GetAgentsStatusTerminated,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Agents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetAgentsRequest](../../models/operations/getagentsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetAgentsResponse](../../models/operations/getagentsresponse.md), error**


## PostAgentsAgentIDUpdate

Update the agent with the given id to the latest agent version

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
    res, err := s.AgentManagement.PostAgentsAgentIDUpdate(ctx, operations.PostAgentsAgentIDUpdateRequest{
        AgentID: "d5e60b37-5ed4-4f6f-bee4-1f33317fe35b",
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
| `request`                                                                                              | [operations.PostAgentsAgentIDUpdateRequest](../../models/operations/postagentsagentidupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostAgentsAgentIDUpdateResponse](../../models/operations/postagentsagentidupdateresponse.md), error**


## PostAgentsAgentIDUpdateState

Update the status of an agent with the given id

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
    res, err := s.AgentManagement.PostAgentsAgentIDUpdateState(ctx, operations.PostAgentsAgentIDUpdateStateRequest{
        AgentStatusUpdate: shared.AgentStatusUpdate{
            Active: testpango.Bool(false),
        },
        AgentID: "60eb1ea4-2655-45ba-bc28-744ed53b88f3",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAgentsAgentIDUpdateStateRequest](../../models/operations/postagentsagentidupdatestaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostAgentsAgentIDUpdateStateResponse](../../models/operations/postagentsagentidupdatestateresponse.md), error**

