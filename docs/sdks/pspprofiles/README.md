# PspProfiles
(*PspProfiles*)

## Overview

APIs used to manage pod security standards profiles on Secure Application

### Available Operations

* [DeletePodSecurityPolicyProfilesProfileID](#deletepodsecuritypolicyprofilesprofileid) - Delete a pod security policy standards
* [DeleteSeccompProfilesProfileID](#deleteseccompprofilesprofileid) - Delete a seccomp profile
* [GetPodSecurityPolicyProfiles](#getpodsecuritypolicyprofiles) - Get all the pod security standards profiles on the system
* [GetSeccompProfiles](#getseccompprofiles) - Get all the seccomp profiles on the system
* [PostPodSecurityPolicyProfiles](#postpodsecuritypolicyprofiles) - Create a new pod security standards
* [PostPodSecurityPolicyProfilesBatch](#postpodsecuritypolicyprofilesbatch) - Add a number of Pod Security Standards
* [PostSeccompProfiles](#postseccompprofiles) - Add seccomp profile
* [PutPodSecurityPolicyProfilesProfileID](#putpodsecuritypolicyprofilesprofileid) - Change pod security standards profile
* [PutSeccompProfilesProfileID](#putseccompprofilesprofileid) - Change seccomp profile

## DeletePodSecurityPolicyProfilesProfileID

Delete a pod security policy standards

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
    res, err := s.PspProfiles.DeletePodSecurityPolicyProfilesProfileID(ctx, operations.DeletePodSecurityPolicyProfilesProfileIDRequest{
        ProfileID: "bb90cc1f-4444-454a-8574-313c05003108",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.DeletePodSecurityPolicyProfilesProfileIDRequest](../../pkg/models/operations/deletepodsecuritypolicyprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.DeletePodSecurityPolicyProfilesProfileIDResponse](../../pkg/models/operations/deletepodsecuritypolicyprofilesprofileidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteSeccompProfilesProfileID

Delete a seccomp profile

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
    res, err := s.PspProfiles.DeleteSeccompProfilesProfileID(ctx, operations.DeleteSeccompProfilesProfileIDRequest{
        ProfileID: "4cc0f6a2-ff01-4311-a20d-d6510a1ff69a",
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
| `request`                                                                                                                | [operations.DeleteSeccompProfilesProfileIDRequest](../../pkg/models/operations/deleteseccompprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.DeleteSeccompProfilesProfileIDResponse](../../pkg/models/operations/deleteseccompprofilesprofileidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPodSecurityPolicyProfiles

Get all the pod security standards profiles on the system

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
    res, err := s.PspProfiles.GetPodSecurityPolicyProfiles(ctx)
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

**[*operations.GetPodSecurityPolicyProfilesResponse](../../pkg/models/operations/getpodsecuritypolicyprofilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSeccompProfiles

Get all the seccomp profiles on the system

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
    res, err := s.PspProfiles.GetSeccompProfiles(ctx)
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

**[*operations.GetSeccompProfilesResponse](../../pkg/models/operations/getseccompprofilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostPodSecurityPolicyProfiles

Create a new pod security standards

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
    res, err := s.PspProfiles.PostPodSecurityPolicyProfiles(ctx, shared.PodSecurityPolicy{
        FsGroup: shared.RunAsGroupStrategyOptions{},
        Name: "<value>",
        RunAsUser: shared.RunAsUserStrategyOptions{},
        SupplementalGroups: shared.RunAsGroupStrategyOptions{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodSecurityPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.PodSecurityPolicy](../../pkg/models/shared/podsecuritypolicy.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostPodSecurityPolicyProfilesResponse](../../pkg/models/operations/postpodsecuritypolicyprofilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostPodSecurityPolicyProfilesBatch

Add a number of new Pod Security Standards Profiles.


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
    res, err := s.PspProfiles.PostPodSecurityPolicyProfilesBatch(ctx, []shared.PodSecurityPolicy{
        shared.PodSecurityPolicy{
            FsGroup: shared.RunAsGroupStrategyOptions{},
            Name: "<value>",
            RunAsUser: shared.RunAsUserStrategyOptions{},
            SupplementalGroups: shared.RunAsGroupStrategyOptions{},
        },
    })
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
| `request`                                             | [[]shared.PodSecurityPolicy](../../.md)               | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostPodSecurityPolicyProfilesBatchResponse](../../pkg/models/operations/postpodsecuritypolicyprofilesbatchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostSeccompProfiles

Add seccomp profile

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
    res, err := s.PspProfiles.PostSeccompProfiles(ctx, shared.SeccompProfileInput{})
    if err != nil {
        log.Fatal(err)
    }

    if res.SeccompProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.SeccompProfileInput](../../pkg/models/shared/seccompprofileinput.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PostSeccompProfilesResponse](../../pkg/models/operations/postseccompprofilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutPodSecurityPolicyProfilesProfileID

Change pod security standards profile

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
    res, err := s.PspProfiles.PutPodSecurityPolicyProfilesProfileID(ctx, operations.PutPodSecurityPolicyProfilesProfileIDRequest{
        PodSecurityPolicy: shared.PodSecurityPolicy{
            FsGroup: shared.RunAsGroupStrategyOptions{},
            Name: "<value>",
            RunAsUser: shared.RunAsUserStrategyOptions{},
            SupplementalGroups: shared.RunAsGroupStrategyOptions{},
        },
        ProfileID: "10ce973a-74d3-47da-b254-604edb379d9a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodSecurityPolicy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PutPodSecurityPolicyProfilesProfileIDRequest](../../pkg/models/operations/putpodsecuritypolicyprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.PutPodSecurityPolicyProfilesProfileIDResponse](../../pkg/models/operations/putpodsecuritypolicyprofilesprofileidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutSeccompProfilesProfileID

Change seccomp profile

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
    res, err := s.PspProfiles.PutSeccompProfilesProfileID(ctx, operations.PutSeccompProfilesProfileIDRequest{
        SeccompProfile: shared.SeccompProfileInput{},
        ProfileID: "7a689ccb-cc10-41fe-9293-6c76ef44408a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SeccompProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutSeccompProfilesProfileIDRequest](../../pkg/models/operations/putseccompprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutSeccompProfilesProfileIDResponse](../../pkg/models/operations/putseccompprofilesprofileidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
