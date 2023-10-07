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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.DeletePodSecurityPolicyProfilesProfileIDRequest](../../models/operations/deletepodsecuritypolicyprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.DeletePodSecurityPolicyProfilesProfileIDResponse](../../models/operations/deletepodsecuritypolicyprofilesprofileidresponse.md), error**


## DeleteSeccompProfilesProfileID

Delete a seccomp profile

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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteSeccompProfilesProfileIDRequest](../../models/operations/deleteseccompprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.DeleteSeccompProfilesProfileIDResponse](../../models/operations/deleteseccompprofilesprofileidresponse.md), error**


## GetPodSecurityPolicyProfiles

Get all the pod security standards profiles on the system

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
    res, err := s.PspProfiles.GetPodSecurityPolicyProfiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PodSecurityPolicies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetPodSecurityPolicyProfilesResponse](../../models/operations/getpodsecuritypolicyprofilesresponse.md), error**


## GetSeccompProfiles

Get all the seccomp profiles on the system

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
    res, err := s.PspProfiles.GetSeccompProfiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SeccompProfiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSeccompProfilesResponse](../../models/operations/getseccompprofilesresponse.md), error**


## PostPodSecurityPolicyProfiles

Create a new pod security standards

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
    res, err := s.PspProfiles.PostPodSecurityPolicyProfiles(ctx, shared.PodSecurityPolicy{
        AllowedCapabilities: []string{
            "bypass",
        },
        AllowedHostPaths: []shared.AllowedHostPath{
            shared.AllowedHostPath{},
        },
        AllowedProcMountTypes: []shared.AllowedProcMountType{
            shared.AllowedProcMountTypeDefault,
        },
        AllowedUnsafeSysctls: []string{
            "turquoise",
        },
        ForbiddenSysctls: []string{
            "moderator",
        },
        FsGroup: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{},
            },
        },
        HostPorts: []shared.HostPortRange{
            shared.HostPortRange{},
        },
        Name: "Gloves",
        RequiredDropCapabilities: []string{
            "explicit",
        },
        RunAsGroup: &shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{},
            },
        },
        RunAsUser: shared.RunAsUserStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{},
            },
        },
        SupplementalGroups: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{},
            },
        },
        Volumes: []shared.PSPVolumeTypes{
            shared.PSPVolumeTypesHostPath,
        },
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.PodSecurityPolicy](../../models/shared/podsecuritypolicy.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.PostPodSecurityPolicyProfilesResponse](../../models/operations/postpodsecuritypolicyprofilesresponse.md), error**


## PostPodSecurityPolicyProfilesBatch

Add a number of new Pod Security Standards Profiles.


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
    res, err := s.PspProfiles.PostPodSecurityPolicyProfilesBatch(ctx, []shared.PodSecurityPolicy{
        shared.PodSecurityPolicy{
            AllowedCapabilities: []string{
                "Copernicium",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{},
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeUnmasked,
            },
            AllowedUnsafeSysctls: []string{
                "Northeast",
            },
            ForbiddenSysctls: []string{
                "Directives",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{},
            },
            Name: "mint",
            RequiredDropCapabilities: []string{
                "Gasoline",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesAzureDisk,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodSecurityPolicies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.PodSecurityPolicy](../../models//.md)       | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostPodSecurityPolicyProfilesBatchResponse](../../models/operations/postpodsecuritypolicyprofilesbatchresponse.md), error**


## PostSeccompProfiles

Add seccomp profile

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
    res, err := s.PspProfiles.PostSeccompProfiles(ctx, shared.SeccompProfileInput{
        PodSecurityPolicies: []string{
            "Cab",
        },
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.SeccompProfileInput](../../models/shared/seccompprofileinput.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostSeccompProfilesResponse](../../models/operations/postseccompprofilesresponse.md), error**


## PutPodSecurityPolicyProfilesProfileID

Change pod security standards profile

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
    res, err := s.PspProfiles.PutPodSecurityPolicyProfilesProfileID(ctx, operations.PutPodSecurityPolicyProfilesProfileIDRequest{
        PodSecurityPolicy: shared.PodSecurityPolicy{
            AllowedCapabilities: []string{
                "North",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{},
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeUnmasked,
            },
            AllowedUnsafeSysctls: []string{
                "since",
            },
            ForbiddenSysctls: []string{
                "DNS",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{},
            },
            Name: "Coupe THX",
            RequiredDropCapabilities: []string{
                "Dynamic",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{},
                },
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesFlocker,
            },
        },
        ProfileID: "04edb379-d9af-4bb9-b687-b2cfe0a637e7",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PutPodSecurityPolicyProfilesProfileIDRequest](../../models/operations/putpodsecuritypolicyprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.PutPodSecurityPolicyProfilesProfileIDResponse](../../models/operations/putpodsecuritypolicyprofilesprofileidresponse.md), error**


## PutSeccompProfilesProfileID

Change seccomp profile

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
    res, err := s.PspProfiles.PutSeccompProfilesProfileID(ctx, operations.PutSeccompProfilesProfileIDRequest{
        SeccompProfileInput: shared.SeccompProfileInput{
            PodSecurityPolicies: []string{
                "payment",
            },
        },
        ProfileID: "689ccbcc-101f-4ed2-936c-76ef44408a69",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutSeccompProfilesProfileIDRequest](../../models/operations/putseccompprofilesprofileidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PutSeccompProfilesProfileIDResponse](../../models/operations/putseccompprofilesprofileidresponse.md), error**

