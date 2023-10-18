# Settings
(*Settings*)

## Overview

APIs used  to configure system settings

### Available Operations

* [DeleteSettingsIntegrationsCaID](#deletesettingsintegrationscaid) - Delete the CA integration details
* [DeleteSettingsIntegrationsEventForwardingEventForwardingID](#deletesettingsintegrationseventforwardingeventforwardingid) - Delete the event forwarding integration details with the given id
* [GetSettingsAgentsUpdate](#getsettingsagentsupdate) - Get the agents update configurations
* [GetSettingsIntegrationsCa](#getsettingsintegrationsca) - Get the CA integration details
* [GetSettingsIntegrationsEventForwarding](#getsettingsintegrationseventforwarding) - Get the event forwarding integration details
* [PostSeccompProfilesValidateData](#postseccompprofilesvalidatedata) - Test the seccomp profile data
* [PostSettingsAgentsUpdateUpdateNow](#postsettingsagentsupdateupdatenow) - Update the agents of the account now
* [PostSettingsIntegrationsCa](#postsettingsintegrationsca) - Set the CA integration details
* [PostSettingsIntegrationsEventForwarding](#postsettingsintegrationseventforwarding) - Set the event forwarding integration details
* [PostSettingsIntegrationsOpsgenieTestIntegration](#postsettingsintegrationsopsgenietestintegration) - Test the connection to Opsgenie
* [PostSettingsIntegrationsSecurexTestIntegration](#postsettingsintegrationssecurextestintegration) - Test the SecureX integration by sending test message to the provided URL
* [PostSettingsIntegrationsSlackTestIntegration](#postsettingsintegrationsslacktestintegration) - Test the Slack integration by sending test message to the provided URL
* [PostSettingsIntegrationsSplunkTestIntegration](#postsettingsintegrationssplunktestintegration) - Test the connection to Splunk
* [PostSettingsIntegrationsSumoLogicTestIntegration](#postsettingsintegrationssumologictestintegration) - Test the Sumo Logic integration by sending test message to the provided URL
* [PostSettingsIntegrationsTeamsTestIntegration](#postsettingsintegrationsteamstestintegration) - Test the connection to Teams
* [PostSettingsIntegrationsWebexTestIntegration](#postsettingsintegrationswebextestintegration) - Test the Webex integration by sending test message to the provided URL
* [PutSettingsAgentsUpdate](#putsettingsagentsupdate) - get the agents update configurations.
* [PutSettingsIntegrationsCaID](#putsettingsintegrationscaid) - Edit the CA integration details
* [PutSettingsIntegrationsEventForwardingEventForwardingID](#putsettingsintegrationseventforwardingeventforwardingid) - Edit the event forwarding integration details

## DeleteSettingsIntegrationsCaID

Delete the CA integration details

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
    res, err := s.Settings.DeleteSettingsIntegrationsCaID(ctx, operations.DeleteSettingsIntegrationsCaIDRequest{
        ID: "2b23989d-8ad5-450f-8a76-f44a753018c5",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteSettingsIntegrationsCaIDRequest](../../models/operations/deletesettingsintegrationscaidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteSettingsIntegrationsCaIDResponse](../../models/operations/deletesettingsintegrationscaidresponse.md), error**


## DeleteSettingsIntegrationsEventForwardingEventForwardingID

Delete the event forwarding integration details with the given id

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
    res, err := s.Settings.DeleteSettingsIntegrationsEventForwardingEventForwardingID(ctx, operations.DeleteSettingsIntegrationsEventForwardingEventForwardingIDRequest{
        EventForwardingID: "a54e5100-941b-4501-ab74-aa7fd2e0ccc3",
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

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.DeleteSettingsIntegrationsEventForwardingEventForwardingIDRequest](../../models/operations/deletesettingsintegrationseventforwardingeventforwardingidrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.DeleteSettingsIntegrationsEventForwardingEventForwardingIDResponse](../../models/operations/deletesettingsintegrationseventforwardingeventforwardingidresponse.md), error**


## GetSettingsAgentsUpdate

Get the agents update configurations

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
    res, err := s.Settings.GetSettingsAgentsUpdate(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AgentsUpdateSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSettingsAgentsUpdateResponse](../../models/operations/getsettingsagentsupdateresponse.md), error**


## GetSettingsIntegrationsCa

Get the CA integration details

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
    res, err := s.Settings.GetSettingsIntegrationsCa(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CaIntegrationResponseWithClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSettingsIntegrationsCaResponse](../../models/operations/getsettingsintegrationscaresponse.md), error**


## GetSettingsIntegrationsEventForwarding

Get the event forwarding integration details

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
    res, err := s.Settings.GetSettingsIntegrationsEventForwarding(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventsForwardingDetailsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSettingsIntegrationsEventForwardingResponse](../../models/operations/getsettingsintegrationseventforwardingresponse.md), error**


## PostSeccompProfilesValidateData

Test the seccomp profile data

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
    res, err := s.Settings.PostSeccompProfilesValidateData(ctx, shared.SeccompProfileData{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.SeccompProfileData](../../models/shared/seccompprofiledata.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostSeccompProfilesValidateDataResponse](../../models/operations/postseccompprofilesvalidatedataresponse.md), error**


## PostSettingsAgentsUpdateUpdateNow

Update the agents of the account now

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
    res, err := s.Settings.PostSettingsAgentsUpdateUpdateNow(ctx)
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


### Response

**[*operations.PostSettingsAgentsUpdateUpdateNowResponse](../../models/operations/postsettingsagentsupdateupdatenowresponse.md), error**


## PostSettingsIntegrationsCa

Set the CA integration details

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
    res, err := s.Settings.PostSettingsIntegrationsCa(ctx, shared.CaIntegrationRequestInput{
        Certificate: "Utah",
        IssuerName: "Internal",
        Name: "Loan",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CaIntegrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.CaIntegrationRequestInput](../../models/shared/caintegrationrequestinput.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostSettingsIntegrationsCaResponse](../../models/operations/postsettingsintegrationscaresponse.md), error**


## PostSettingsIntegrationsEventForwarding

Set the event forwarding integration details

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
    res, err := s.Settings.PostSettingsIntegrationsEventForwarding(ctx, shared.EventsForwardingDetailsInput{
        EventsForwardingDetailsType: shared.EventsForwardingDetailsTypeEnumTeamsVulnerabilityEventsForwardingDetails,
        EventsToForward: []shared.EventsToForward{
            shared.EventsToForwardAttackPath,
        },
        Name: "Gallium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventsForwardingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.EventsForwardingDetailsInput](../../models/shared/eventsforwardingdetailsinput.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostSettingsIntegrationsEventForwardingResponse](../../models/operations/postsettingsintegrationseventforwardingresponse.md), error**


## PostSettingsIntegrationsOpsgenieTestIntegration

Test the connection to Opsgenie

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
    res, err := s.Settings.PostSettingsIntegrationsOpsgenieTestIntegration(ctx, shared.TestOpsgenieConnectionRequest{
        Token: "invoice",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TestOpsgenieConnectionRequest](../../models/shared/testopsgenieconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsOpsgenieTestIntegrationResponse](../../models/operations/postsettingsintegrationsopsgenietestintegrationresponse.md), error**


## PostSettingsIntegrationsSecurexTestIntegration

Test the SecureX integration by sending test message to the provided URL

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
    res, err := s.Settings.PostSettingsIntegrationsSecurexTestIntegration(ctx, shared.TestSecureXIntegrationRequest{
        URL: "https://angelic-mortgage.name",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TestSecureXIntegrationRequest](../../models/shared/testsecurexintegrationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsSecurexTestIntegrationResponse](../../models/operations/postsettingsintegrationssecurextestintegrationresponse.md), error**


## PostSettingsIntegrationsSlackTestIntegration

Test the Slack integration by sending test message to the provided URL

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
    res, err := s.Settings.PostSettingsIntegrationsSlackTestIntegration(ctx, shared.TestSlackIntegrationRequest{
        URL: "https://every-vibration.name",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TestSlackIntegrationRequest](../../models/shared/testslackintegrationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostSettingsIntegrationsSlackTestIntegrationResponse](../../models/operations/postsettingsintegrationsslacktestintegrationresponse.md), error**


## PostSettingsIntegrationsSplunkTestIntegration

Test the connection to Splunk

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
    res, err := s.Settings.PostSettingsIntegrationsSplunkTestIntegration(ctx, shared.TestSplunkConnectionRequest{
        Token: "Realigned",
        URL: "https://forked-investigation.org",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TestSplunkConnectionRequest](../../models/shared/testsplunkconnectionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostSettingsIntegrationsSplunkTestIntegrationResponse](../../models/operations/postsettingsintegrationssplunktestintegrationresponse.md), error**


## PostSettingsIntegrationsSumoLogicTestIntegration

Test the Sumo Logic integration by sending test message to the provided URL

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
    res, err := s.Settings.PostSettingsIntegrationsSumoLogicTestIntegration(ctx, shared.TestSumoLogicIntegrationRequest{
        URL: "http://sudden-concern.name",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.TestSumoLogicIntegrationRequest](../../models/shared/testsumologicintegrationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostSettingsIntegrationsSumoLogicTestIntegrationResponse](../../models/operations/postsettingsintegrationssumologictestintegrationresponse.md), error**


## PostSettingsIntegrationsTeamsTestIntegration

Test the connection to Teams

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
    res, err := s.Settings.PostSettingsIntegrationsTeamsTestIntegration(ctx, shared.TestTeamsIntegrationRequest{
        URL: "http://faraway-rayon.info",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TestTeamsIntegrationRequest](../../models/shared/testteamsintegrationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostSettingsIntegrationsTeamsTestIntegrationResponse](../../models/operations/postsettingsintegrationsteamstestintegrationresponse.md), error**


## PostSettingsIntegrationsWebexTestIntegration

Test the Webex integration by sending test message to the provided URL

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
    res, err := s.Settings.PostSettingsIntegrationsWebexTestIntegration(ctx, shared.TestWebexIntegrationRequest{
        URL: "http://charming-gadget.info",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TestWebexIntegrationRequest](../../models/shared/testwebexintegrationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostSettingsIntegrationsWebexTestIntegrationResponse](../../models/operations/postsettingsintegrationswebextestintegrationresponse.md), error**


## PutSettingsAgentsUpdate

get the agents update configurations.

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
    res, err := s.Settings.PutSettingsAgentsUpdate(ctx, shared.AgentsUpdateSettingsInput{})
    if err != nil {
        log.Fatal(err)
    }

    if res.AgentsUpdateSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.AgentsUpdateSettingsInput](../../models/shared/agentsupdatesettingsinput.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PutSettingsAgentsUpdateResponse](../../models/operations/putsettingsagentsupdateresponse.md), error**


## PutSettingsIntegrationsCaID

Edit the CA integration details

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
    res, err := s.Settings.PutSettingsIntegrationsCaID(ctx, operations.PutSettingsIntegrationsCaIDRequest{
        CaIntegrationRequestInput: shared.CaIntegrationRequestInput{
            Certificate: "Ergonomic",
            IssuerName: "Advanced",
            Name: "Bicycle",
        },
        ID: "da77d5f8-0b65-48d2-bdca-eae5506a19b8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CaIntegrationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutSettingsIntegrationsCaIDRequest](../../models/operations/putsettingsintegrationscaidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutSettingsIntegrationsCaIDResponse](../../models/operations/putsettingsintegrationscaidresponse.md), error**


## PutSettingsIntegrationsEventForwardingEventForwardingID

Edit the event forwarding integration details

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
    res, err := s.Settings.PutSettingsIntegrationsEventForwardingEventForwardingID(ctx, operations.PutSettingsIntegrationsEventForwardingEventForwardingIDRequest{
        SplunkEventsForwardingDetailsInput: shared.SplunkEventsForwardingDetailsInput{
            EventsForwardingDetailsType: shared.EventsForwardingDetailsTypeEnumWebexEventsForwardingDetails,
            EventsToForward: []shared.EventsToForward{
                shared.EventsToForwardNotification,
            },
            Name: "Communications",
            Token: "Guinea",
        },
        EventForwardingID: "3efb0fd0-479c-42dc-9b40-8bce49d238c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventsForwardingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.PutSettingsIntegrationsEventForwardingEventForwardingIDRequest](../../models/operations/putsettingsintegrationseventforwardingeventforwardingidrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.PutSettingsIntegrationsEventForwardingEventForwardingIDResponse](../../models/operations/putsettingsintegrationseventforwardingeventforwardingidresponse.md), error**

