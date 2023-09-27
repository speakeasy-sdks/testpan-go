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
        ProfileID: "b084da99-257d-404f-8084-7a742d84496c",
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
        ProfileID: "bdeecf6b-99bc-4635-a2eb-fdf55c294c06",
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
            "consequatur",
        },
        AllowedHostPaths: []shared.AllowedHostPath{
            shared.AllowedHostPath{
                PathPrefix: testpango.String("nobis"),
                ReadOnly: testpango.Bool(false),
            },
        },
        AllowedProcMountTypes: []shared.AllowedProcMountType{
            shared.AllowedProcMountTypeDefault,
        },
        AllowedUnsafeSysctls: []string{
            "ea",
        },
        DefaultAllowPrivilegeEscalation: testpango.Bool(false),
        Description: testpango.String("laborum"),
        ForbiddenSysctls: []string{
            "et",
        },
        FsGroup: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(144856),
                    Min: testpango.Int64(550579),
                },
            },
            Rule: shared.RunAsGroupStrategyMayRunAs.ToPointer(),
        },
        HostIPC: testpango.Bool(false),
        HostNetwork: testpango.Bool(false),
        HostPID: testpango.Bool(false),
        HostPorts: []shared.HostPortRange{
            shared.HostPortRange{
                Max: testpango.Int(454329),
                Min: testpango.Int(421018),
            },
        },
        ID: testpango.String("4eef6d0c-6d6e-4d9c-b3dd-634571509a8e"),
        IsSecurecnDefaultProfile: testpango.Bool(false),
        Name: "Julio Bauch",
        Privileged: testpango.Bool(false),
        ReadOnlyRootFileSystem: testpango.Bool(false),
        RequiredDropCapabilities: []string{
            "minus",
        },
        RunAsGroup: &shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(330908),
                    Min: testpango.Int64(664679),
                },
            },
            Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
        },
        RunAsUser: shared.RunAsUserStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(978229),
                    Min: testpango.Int64(598497),
                },
            },
            Rule: shared.RunAsUserStrategyRunAsAny.ToPointer(),
        },
        SeccompProfile: testpango.String("242c7b66-a1f3-40c7-bdf5-b6719890f42a"),
        SupplementalGroups: shared.RunAsGroupStrategyOptions{
            Ranges: []shared.IDRange{
                shared.IDRange{
                    Max: testpango.Int64(272396),
                    Min: testpango.Int64(749863),
                },
            },
            Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
        },
        Volumes: []shared.PSPVolumeTypes{
            shared.PSPVolumeTypesDownwardAPI,
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
                "adipisci",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{
                    PathPrefix: testpango.String("atque"),
                    ReadOnly: testpango.Bool(false),
                },
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeUnmasked,
            },
            AllowedUnsafeSysctls: []string{
                "rem",
            },
            DefaultAllowPrivilegeEscalation: testpango.Bool(false),
            Description: testpango.String("exercitationem"),
            ForbiddenSysctls: []string{
                "tempore",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(180839),
                        Min: testpango.Int64(389585),
                    },
                },
                Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
            },
            HostIPC: testpango.Bool(false),
            HostNetwork: testpango.Bool(false),
            HostPID: testpango.Bool(false),
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{
                    Max: testpango.Int(326894),
                    Min: testpango.Int(595595),
                },
            },
            ID: testpango.String("1d745e3c-2059-4c9c-bf56-7e0e252765b1"),
            IsSecurecnDefaultProfile: testpango.Bool(false),
            Name: "Reginald Cruickshank",
            Privileged: testpango.Bool(false),
            ReadOnlyRootFileSystem: testpango.Bool(false),
            RequiredDropCapabilities: []string{
                "possimus",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(670710),
                        Min: testpango.Int64(761835),
                    },
                },
                Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(100926),
                        Min: testpango.Int64(968792),
                    },
                },
                Rule: shared.RunAsUserStrategyMustRunAs.ToPointer(),
            },
            SeccompProfile: testpango.String("1216ce22-39e8-4f25-8d0d-19d959f439e3"),
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(573816),
                        Min: testpango.Int64(177632),
                    },
                },
                Rule: shared.RunAsGroupStrategyMayRunAs.ToPointer(),
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesFlocker,
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
        Data: testpango.String("{"foo":"boZ>zM]_.a","bar":30034,"bike":9270,"a":39110,"b":"An)tEFVT}l","name":">?XG1+Mk1U","prop":47319}"),
        Name: testpango.String("Allison Kemmer"),
        PodSecurityPolicies: []string{
            "tempora",
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
                "aspernatur",
            },
            AllowedHostPaths: []shared.AllowedHostPath{
                shared.AllowedHostPath{
                    PathPrefix: testpango.String("ad"),
                    ReadOnly: testpango.Bool(false),
                },
            },
            AllowedProcMountTypes: []shared.AllowedProcMountType{
                shared.AllowedProcMountTypeDefault,
            },
            AllowedUnsafeSysctls: []string{
                "alias",
            },
            DefaultAllowPrivilegeEscalation: testpango.Bool(false),
            Description: testpango.String("adipisci"),
            ForbiddenSysctls: []string{
                "atque",
            },
            FsGroup: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(734296),
                        Min: testpango.Int64(989913),
                    },
                },
                Rule: shared.RunAsGroupStrategyRunAsAny.ToPointer(),
            },
            HostIPC: testpango.Bool(false),
            HostNetwork: testpango.Bool(false),
            HostPID: testpango.Bool(false),
            HostPorts: []shared.HostPortRange{
                shared.HostPortRange{
                    Max: testpango.Int(328217),
                    Min: testpango.Int(584483),
                },
            },
            ID: testpango.String("71e98190-5573-489c-adba-c7fda39594d6"),
            IsSecurecnDefaultProfile: testpango.Bool(false),
            Name: "Patty Schinner",
            Privileged: testpango.Bool(false),
            ReadOnlyRootFileSystem: testpango.Bool(false),
            RequiredDropCapabilities: []string{
                "officiis",
            },
            RunAsGroup: &shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(304571),
                        Min: testpango.Int64(559392),
                    },
                },
                Rule: shared.RunAsGroupStrategyMustRunAs.ToPointer(),
            },
            RunAsUser: shared.RunAsUserStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(422215),
                        Min: testpango.Int64(209920),
                    },
                },
                Rule: shared.RunAsUserStrategyMustRunAs.ToPointer(),
            },
            SeccompProfile: testpango.String("b9954b6f-a220-4636-9828-553cb10006be"),
            SupplementalGroups: shared.RunAsGroupStrategyOptions{
                Ranges: []shared.IDRange{
                    shared.IDRange{
                        Max: testpango.Int64(968591),
                        Min: testpango.Int64(277569),
                    },
                },
                Rule: shared.RunAsGroupStrategyMayRunAs.ToPointer(),
            },
            Volumes: []shared.PSPVolumeTypes{
                shared.PSPVolumeTypesCinder,
            },
        },
        ProfileID: "1ec2053b-7493-466a-88ee-0f2bf19588d4",
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
            Data: testpango.String("{"foo":87372,"bar":19647,"bike":"4nuf_/ZMbx","a":93359,"b":"!af8!o|TGO","name":"qC2<\">hd53","prop":86405}"),
            Name: testpango.String("Ramona Crona"),
            PodSecurityPolicies: []string{
                "doloribus",
            },
        },
        ProfileID: "127fb0e0-bf1f-4821-b978-d0acca77aeb7",
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

