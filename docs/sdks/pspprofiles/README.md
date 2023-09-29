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
        AllowPrivilegeEscalation: testpango.Bool(false),
        AllowedCapabilities: []string{
            "bypass",
        },
        AllowedHostPaths: []shared.AllowedHostPath{
            shared.AllowedHostPath{
                PathPrefix: testpango.String("turquoise"),
                ReadOnly: testpango.Bool(false),
            },
        },
        AllowedProcMountTypes: []shared.AllowedProcMountType{
            shared.AllowedProcMountTypeDefault,
        },
        AllowedUnsafeSysctls: []string{
            "Product",
        },
        DefaultAllowPrivilegeEscalation: testpango.Bool(false),
        Description: testpango.String("Down-sized empowering frame"),
        ForbiddenSysctls: []string{
            "incremental",
        },
        FsGroup: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(247342),
                    Min: testpango.Int64(407317),
                },
            },
            Rule: shared.RunAsGroupStrategyMayRunAs.ToPointer(),
        },
        HostIPC: testpango.Bool(false),
        HostNetwork: testpango.Bool(false),
        HostPID: testpango.Bool(false),
        HostPorts: []shared.HostPortRange{
            shared.HostPortRange{
                Max: testpango.Int(80641),
                Min: testpango.Int(34868),
            },
        },
        ID: testpango.String("142085d7-d484-444d-9348-fb70a8e37e62"),
        IsSecurecnDefaultProfile: testpango.Bool(false),
        Name: "Wooden repudiandae Paradigm",
        Privileged: testpango.Bool(false),
        ReadOnlyRootFileSystem: testpango.Bool(false),
        RequiredDropCapabilities: []string{
            "clearly",
        },
        RunAsGroup: &shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(307575),
                    Min: testpango.Int64(822060),
                },
            },
            Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
        },
        RunAsUser: shared.RunAsUserStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(177881),
                    Min: testpango.Int64(417586),
                },
            },
            Rule: shared.RunAsUserStrategyMustRunAs.ToPointer(),
        },
        SeccompProfile: testpango.String("e80dd252-3540-4eb3-996a-05613aad0af4"),
        SupplementalGroups: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(169335),
                    Min: testpango.Int64(820289),
                },
            },
            Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
        },
        Volumes: []shared.PSPVolumeTypes{
            shared.PSPVolumeTypesStorageos,
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
            AllowPrivilegeEscalation: testpango.Bool(false),
            AllowedCapabilities: []string{
                "Copernicium",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{
                    PathPrefix: testpango.String("Northeast Directives"),
                    ReadOnly: testpango.Bool(false),
                },
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeDefault,
            },
            AllowedUnsafeSysctls: []string{
                "mint",
            },
            DefaultAllowPrivilegeEscalation: testpango.Bool(false),
            Description: testpango.String("Self-enabling national application"),
            ForbiddenSysctls: []string{
                "THX",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(581609),
                        Min: testpango.Int64(807880),
                    },
                },
                Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
            },
            HostIPC: testpango.Bool(false),
            HostNetwork: testpango.Bool(false),
            HostPID: testpango.Bool(false),
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{
                    Max: testpango.Int(693689),
                    Min: testpango.Int(305839),
                },
            },
            ID: testpango.String("ccbec10a-9bc5-4ecc-aa49-c56c0e0d8163"),
            IsSecurecnDefaultProfile: testpango.Bool(false),
            Name: "Bicycle synthesize",
            Privileged: testpango.Bool(false),
            ReadOnlyRootFileSystem: testpango.Bool(false),
            RequiredDropCapabilities: []string{
                "Cyclocross",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(978567),
                        Min: testpango.Int64(527502),
                    },
                },
                Rule: shared.RunAsGroupStrategyMayRunAs.ToPointer(),
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(393434),
                        Min: testpango.Int64(336870),
                    },
                },
                Rule: shared.RunAsUserStrategyMustRunAs.ToPointer(),
            },
            SeccompProfile: testpango.String("33ebfbb0-cc93-4268-acf5-38bb5e43e0cc"),
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(259664),
                        Min: testpango.Int64(329554),
                    },
                },
                Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesVsphereVolume,
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
        Data: testpango.String("{steam: 44630, freedom: null, lathe: \"maiores metrics\"}"),
        Name: testpango.String("bitterly North"),
        PodSecurityPolicies: []string{
            "bandwidth",
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
            AllowPrivilegeEscalation: testpango.Bool(false),
            AllowedCapabilities: []string{
                "North",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{
                    PathPrefix: testpango.String("since DNS Central"),
                    ReadOnly: testpango.Bool(false),
                },
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeDefault,
            },
            AllowedUnsafeSysctls: []string{
                "Mercedes",
            },
            DefaultAllowPrivilegeEscalation: testpango.Bool(false),
            Description: testpango.String("Customer-focused eco-centric encryption"),
            ForbiddenSysctls: []string{
                "Checking",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(885380),
                        Min: testpango.Int64(870395),
                    },
                },
                Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
            },
            HostIPC: testpango.Bool(false),
            HostNetwork: testpango.Bool(false),
            HostPID: testpango.Bool(false),
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{
                    Max: testpango.Int(208639),
                    Min: testpango.Int(494299),
                },
            },
            ID: testpango.String("9d9afbb9-b687-4b2c-be0a-637e73494546"),
            IsSecurecnDefaultProfile: testpango.Bool(false),
            Name: "Silicon sensor",
            Privileged: testpango.Bool(false),
            ReadOnlyRootFileSystem: testpango.Bool(false),
            RequiredDropCapabilities: []string{
                "ex",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(639188),
                        Min: testpango.Int64(375698),
                    },
                },
                Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(34862),
                        Min: testpango.Int64(966964),
                    },
                },
                Rule: shared.RunAsUserStrategyMustRunAs.ToPointer(),
            },
            SeccompProfile: testpango.String("114f1e40-0857-4100-95d7-43f0167fa418"),
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(807200),
                        Min: testpango.Int64(91090),
                    },
                },
                Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesScaleIo,
            },
        },
        ProfileID: "30ce44b3-7a43-4d84-b255-0fa77d979933",
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
            Data: testpango.String("{incandescence: 66214, hint: null, licorice: \"Mobility lime\"}"),
            Name: testpango.String("BMX Cambridgeshire wherever"),
            PodSecurityPolicies: []string{
                "plum",
            },
        },
        ProfileID: "36c76ef4-4408-4a69-9bd6-c47238638eac",
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

