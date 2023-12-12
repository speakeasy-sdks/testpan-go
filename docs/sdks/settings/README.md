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
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteSettingsIntegrationsCaIDRequest](../../pkg/models/operations/deletesettingsintegrationscaidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeleteSettingsIntegrationsCaIDResponse](../../pkg/models/operations/deletesettingsintegrationscaidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteSettingsIntegrationsEventForwardingEventForwardingID

Delete the event forwarding integration details with the given id

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.DeleteSettingsIntegrationsEventForwardingEventForwardingIDRequest](../../pkg/models/operations/deletesettingsintegrationseventforwardingeventforwardingidrequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |


### Response

**[*operations.DeleteSettingsIntegrationsEventForwardingEventForwardingIDResponse](../../pkg/models/operations/deletesettingsintegrationseventforwardingeventforwardingidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSettingsAgentsUpdate

Get the agents update configurations

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

**[*operations.GetSettingsAgentsUpdateResponse](../../pkg/models/operations/getsettingsagentsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSettingsIntegrationsCa

Get the CA integration details

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
    res, err := s.Settings.GetSettingsIntegrationsCa(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSettingsIntegrationsCaResponse](../../pkg/models/operations/getsettingsintegrationscaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSettingsIntegrationsEventForwarding

Get the event forwarding integration details

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

**[*operations.GetSettingsIntegrationsEventForwardingResponse](../../pkg/models/operations/getsettingsintegrationseventforwardingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSeccompProfilesValidateData

Test the seccomp profile data

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.SeccompProfileData](../../pkg/models/shared/seccompprofiledata.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostSeccompProfilesValidateDataResponse](../../pkg/models/operations/postseccompprofilesvalidatedataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsAgentsUpdateUpdateNow

Update the agents of the account now

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

**[*operations.PostSettingsAgentsUpdateUpdateNowResponse](../../pkg/models/operations/postsettingsagentsupdateupdatenowresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsCa

Set the CA integration details

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
    res, err := s.Settings.PostSettingsIntegrationsCa(ctx, shared.CaIntegrationRequest{
        Certificate: "string",
        IssuerName: "string",
        Name: "string",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CaIntegrationRequest](../../pkg/models/shared/caintegrationrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PostSettingsIntegrationsCaResponse](../../pkg/models/operations/postsettingsintegrationscaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsEventForwarding

Set the event forwarding integration details

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
    res, err := s.Settings.PostSettingsIntegrationsEventForwarding(ctx, shared.EventsForwardingDetailsInput{
        EventsForwardingDetailsType: shared.EventsForwardingDetailsTypeEnumTeamsVulnerabilityEventsForwardingDetails,
        EventsToForward: []shared.EventsToForward{
            shared.EventsToForwardAttackPath,
        },
        Name: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.EventsForwardingDetailsInput](../../pkg/models/shared/eventsforwardingdetailsinput.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostSettingsIntegrationsEventForwardingResponse](../../pkg/models/operations/postsettingsintegrationseventforwardingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsOpsgenieTestIntegration

Test the connection to Opsgenie

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.PostSettingsIntegrationsOpsgenieTestIntegration(ctx, shared.TestOpsgenieConnectionRequest{
        Token: "string",
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
| `request`                                                                                        | [shared.TestOpsgenieConnectionRequest](../../pkg/models/shared/testopsgenieconnectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostSettingsIntegrationsOpsgenieTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationsopsgenietestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsSecurexTestIntegration

Test the SecureX integration by sending test message to the provided URL

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.TestSecureXIntegrationRequest](../../pkg/models/shared/testsecurexintegrationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostSettingsIntegrationsSecurexTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationssecurextestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsSlackTestIntegration

Test the Slack integration by sending test message to the provided URL

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TestSlackIntegrationRequest](../../pkg/models/shared/testslackintegrationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsSlackTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationsslacktestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsSplunkTestIntegration

Test the connection to Splunk

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.PostSettingsIntegrationsSplunkTestIntegration(ctx, shared.TestSplunkConnectionRequest{
        Token: "string",
        URL: "http://same-shopper.biz",
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
| `request`                                                                                    | [shared.TestSplunkConnectionRequest](../../pkg/models/shared/testsplunkconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsSplunkTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationssplunktestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsSumoLogicTestIntegration

Test the Sumo Logic integration by sending test message to the provided URL

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.TestSumoLogicIntegrationRequest](../../pkg/models/shared/testsumologicintegrationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PostSettingsIntegrationsSumoLogicTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationssumologictestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsTeamsTestIntegration

Test the connection to Teams

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TestTeamsIntegrationRequest](../../pkg/models/shared/testteamsintegrationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsTeamsTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationsteamstestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettingsIntegrationsWebexTestIntegration

Test the Webex integration by sending test message to the provided URL

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: testpango.String("<YOUR_PASSWORD_HERE>"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.TestWebexIntegrationRequest](../../pkg/models/shared/testwebexintegrationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostSettingsIntegrationsWebexTestIntegrationResponse](../../pkg/models/operations/postsettingsintegrationswebextestintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutSettingsAgentsUpdate

get the agents update configurations.

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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.AgentsUpdateSettingsInput](../../pkg/models/shared/agentsupdatesettingsinput.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PutSettingsAgentsUpdateResponse](../../pkg/models/operations/putsettingsagentsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutSettingsIntegrationsCaID

Edit the CA integration details

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
    res, err := s.Settings.PutSettingsIntegrationsCaID(ctx, operations.PutSettingsIntegrationsCaIDRequest{
        CaIntegrationRequest: shared.CaIntegrationRequest{
            Certificate: "string",
            IssuerName: "string",
            Name: "string",
        },
        ID: "3150c8bd-a77d-45f8-8b65-8d2fdcaeae55",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutSettingsIntegrationsCaIDRequest](../../pkg/models/operations/putsettingsintegrationscaidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutSettingsIntegrationsCaIDResponse](../../pkg/models/operations/putsettingsintegrationscaidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutSettingsIntegrationsEventForwardingEventForwardingID

Edit the event forwarding integration details

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
    res, err := s.Settings.PutSettingsIntegrationsEventForwardingEventForwardingID(ctx, operations.PutSettingsIntegrationsEventForwardingEventForwardingIDRequest{
        SplunkEventsForwardingDetails: shared.SplunkEventsForwardingDetails{
            EventsForwardingDetailsType: shared.EventsForwardingDetailsTypeEnumWebexEventsForwardingDetails,
            EventsToForward: []shared.EventsToForward{
                shared.EventsToForwardNotification,
            },
            Name: "string",
            Token: "string",
        },
        EventForwardingID: "a96473ef-b0fd-4047-9c2d-c1b408bce49d",
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

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.PutSettingsIntegrationsEventForwardingEventForwardingIDRequest](../../pkg/models/operations/putsettingsintegrationseventforwardingeventforwardingidrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |


### Response

**[*operations.PutSettingsIntegrationsEventForwardingEventForwardingIDResponse](../../pkg/models/operations/putsettingsintegrationseventforwardingeventforwardingidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
