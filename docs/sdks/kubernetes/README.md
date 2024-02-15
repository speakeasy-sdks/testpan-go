# Kubernetes
(*Kubernetes*)

## Overview

APIs used to manage Kubernetes clusters on Secure Application

### Available Operations

* [DeleteKubernetesClustersKubernetesClusterID](#deletekubernetesclusterskubernetesclusterid) - Delete a Kubernetes cluster
* [DeletePodDefinitionsPodID](#deletepoddefinitionspodid) - Delete a pod definition
* [GetGetControllerDataClusterID](#getgetcontrollerdataclusterid) - get controller data using clusterId
* [GetIstioSupportedVersions](#getistiosupportedversions) - Get a list of istio releases that are supported by Secure Application agent. sorted from latest to oldest
* [GetKubernetesClusters](#getkubernetesclusters) - get a list of current Kubernetes clusters
* [GetKubernetesClustersKubernetesClusterID](#getkubernetesclusterskubernetesclusterid) - get the Kubernetes cluster with the given id
* [GetKubernetesClustersKubernetesClusterIDDeleteDependencies](#getkubernetesclusterskubernetesclusteriddeletedependencies) - get dependencies which need to be handled in order to delete specified Kubernetes cluster
* [GetKubernetesClustersKubernetesClusterIDDownloadBundle](#getkubernetesclusterskubernetesclusteriddownloadbundle) - Get Secure Application installation script
* [GetKubernetesClustersKubernetesClusterIDGetHelmCommands](#getkubernetesclusterskubernetesclusteridgethelmcommands) - Get Panoptica Aug release Helm command
* [GetKubernetesClustersKubernetesClusterIDNamespaces](#getkubernetesclusterskubernetesclusteridnamespaces) - List namespaces on a specific Kubernetes cluster
* [GetKubernetesClustersKubernetesClusterIDServices](#getkubernetesclusterskubernetesclusteridservices) - List services on a specific Kubernetes cluster
* [GetLeanKubernetesClusters](#getleankubernetesclusters) - get a list of current Kubernetes clusters
* [GetNamespaces](#getnamespaces) - Get a list of current Kubernetes namespaces
* [GetPodDefinitions](#getpoddefinitions) - Get all pod definitions on the system
* [PostKubernetesClusters](#postkubernetesclusters) - Add a new Kubernetes cluster to Secure Application
* [PostPodDefinitions](#postpoddefinitions) - Create a new pod definition
* [PutKubernetesClustersKubernetesClusterID](#putkubernetesclusterskubernetesclusterid) - Update the Kubernetes cluster
* [PutKubernetesClustersKubernetesClusterIDManagedByHelm](#putkubernetesclusterskubernetesclusteridmanagedbyhelm) - Update the Kubernetes cluster which managed by HELM
* [PutPodDefinitionsPodID](#putpoddefinitionspodid) - Change pod definition

## DeleteKubernetesClustersKubernetesClusterID

Delete a Kubernetes cluster

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
    res, err := s.Kubernetes.DeleteKubernetesClustersKubernetesClusterID(ctx, operations.DeleteKubernetesClustersKubernetesClusterIDRequest{
        KubernetesClusterID: "a8fff1fd-16f4-41cf-ba2f-6b7d60168e00",
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteKubernetesClustersKubernetesClusterIDRequest](../../pkg/models/operations/deletekubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.DeleteKubernetesClustersKubernetesClusterIDResponse](../../pkg/models/operations/deletekubernetesclusterskubernetesclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeletePodDefinitionsPodID

Delete a pod definition

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
    res, err := s.Kubernetes.DeletePodDefinitionsPodID(ctx, operations.DeletePodDefinitionsPodIDRequest{
        PodID: "1d7b59d4-9a1e-473b-8e2e-c8a463501aa8",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeletePodDefinitionsPodIDRequest](../../pkg/models/operations/deletepoddefinitionspodidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeletePodDefinitionsPodIDResponse](../../pkg/models/operations/deletepoddefinitionspodidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetGetControllerDataClusterID

get controller data using clusterId

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
    res, err := s.Kubernetes.GetGetControllerDataClusterID(ctx, operations.GetGetControllerDataClusterIDRequest{
        ClusterID: "6d4f5b01-1e67-437c-a42b-e9becce491b4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ControllerDataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetGetControllerDataClusterIDRequest](../../pkg/models/operations/getgetcontrollerdataclusteridrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetGetControllerDataClusterIDResponse](../../pkg/models/operations/getgetcontrollerdataclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetIstioSupportedVersions

Get a list of istio releases that are supported by Secure Application agent. sorted from latest to oldest

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
    res, err := s.Kubernetes.GetIstioSupportedVersions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetIstioSupportedVersionsResponse](../../pkg/models/operations/getistiosupportedversionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClusters

get a list of current Kubernetes clusters

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
    res, err := s.Kubernetes.GetKubernetesClusters(ctx, operations.GetKubernetesClustersRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetKubernetesClustersRequest](../../pkg/models/operations/getkubernetesclustersrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetKubernetesClustersResponse](../../pkg/models/operations/getkubernetesclustersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterID

get the Kubernetes cluster with the given id

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterID(ctx, operations.GetKubernetesClustersKubernetesClusterIDRequest{
        KubernetesClusterID: "6ad413ce-670c-4973-bf8a-5a66b7855734",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetKubernetesClustersKubernetesClusterIDRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterIDDeleteDependencies

get dependencies which need to be handled in order to delete specified Kubernetes cluster

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDDeleteDependencies(ctx, operations.GetKubernetesClustersKubernetesClusterIDDeleteDependenciesRequest{
        KubernetesClusterID: "b06f6c06-c4ec-44f8-ba3a-a7d63215e690",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesClusterDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `request`                                                                                                                                                                        | [operations.GetKubernetesClustersKubernetesClusterIDDeleteDependenciesRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteriddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                                               | The request object to use for the request.                                                                                                                                       |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteriddeletedependenciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterIDDownloadBundle

In order to install,  extract and run "./install_bundle.sh"

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDDownloadBundle(ctx, operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest{
        KubernetesClusterID: "204aba8f-ddfd-4080-9f3c-89482d02fb76",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteriddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteriddownloadbundleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterIDGetHelmCommands

Get Panoptica Aug release Helm command

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDGetHelmCommands(ctx, operations.GetKubernetesClustersKubernetesClusterIDGetHelmCommandsRequest{
        KubernetesClusterID: "f7710c3a-4d45-43f9-898d-e19416c2b527",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HelmCommandsInstallation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.GetKubernetesClustersKubernetesClusterIDGetHelmCommandsRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridgethelmcommandsrequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDGetHelmCommandsResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridgethelmcommandsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterIDNamespaces

List namespaces on a specific Kubernetes cluster

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDNamespaces(ctx, operations.GetKubernetesClustersKubernetesClusterIDNamespacesRequest{
        KubernetesClusterID: "b18474d1-b95c-48d9-af1c-951ce15b0be6",
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

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetKubernetesClustersKubernetesClusterIDNamespacesRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridnamespacesrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDNamespacesResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridnamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKubernetesClustersKubernetesClusterIDServices

List services on a specific Kubernetes cluster

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDServices(ctx, operations.GetKubernetesClustersKubernetesClusterIDServicesRequest{
        KubernetesClusterID: "1b08f941-6d25-44a9-8a5a-94e46b69d5cc",
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetKubernetesClustersKubernetesClusterIDServicesRequest](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridservicesrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDServicesResponse](../../pkg/models/operations/getkubernetesclusterskubernetesclusteridservicesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetLeanKubernetesClusters

get a list of current Kubernetes clusters

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
    res, err := s.Kubernetes.GetLeanKubernetesClusters(ctx, operations.GetLeanKubernetesClustersRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetLeanKubernetesClustersRequest](../../pkg/models/operations/getleankubernetesclustersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetLeanKubernetesClustersResponse](../../pkg/models/operations/getleankubernetesclustersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetNamespaces

Get a list of current Kubernetes namespaces

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
    res, err := s.Kubernetes.GetNamespaces(ctx, operations.GetNamespacesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetNamespacesRequest](../../pkg/models/operations/getnamespacesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetNamespacesResponse](../../pkg/models/operations/getnamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPodDefinitions

Get all pod definitions on the system

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
    res, err := s.Kubernetes.GetPodDefinitions(ctx, operations.GetPodDefinitionsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPodDefinitionsRequest](../../pkg/models/operations/getpoddefinitionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetPodDefinitionsResponse](../../pkg/models/operations/getpoddefinitionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostKubernetesClusters

Add a new Kubernetes cluster to Secure Application

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
    res, err := s.Kubernetes.PostKubernetesClusters(ctx, shared.KubernetesCluster{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.KubernetesCluster](../../pkg/models/shared/kubernetescluster.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostKubernetesClustersResponse](../../pkg/models/operations/postkubernetesclustersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostPodDefinitions

Create a new pod definition

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
    res, err := s.Kubernetes.PostPodDefinitions(ctx, shared.PodDefinitionInput{
        ClusterID: "caf647ad-b35e-4462-baa4-a61d2b4e2410",
        Containers: []shared.Container{
            shared.Container{},
        },
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodDefinition != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.PodDefinitionInput](../../pkg/models/shared/poddefinitioninput.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostPodDefinitionsResponse](../../pkg/models/operations/postpoddefinitionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutKubernetesClustersKubernetesClusterID

Update the Kubernetes cluster

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
    res, err := s.Kubernetes.PutKubernetesClustersKubernetesClusterID(ctx, operations.PutKubernetesClustersKubernetesClusterIDRequest{
        KubernetesCluster: shared.KubernetesCluster{
            Name: "<value>",
        },
        KubernetesClusterID: "ffc49f9e-0fd9-4579-98a0-796a5f2d9d7f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutKubernetesClustersKubernetesClusterIDRequest](../../pkg/models/operations/putkubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.PutKubernetesClustersKubernetesClusterIDResponse](../../pkg/models/operations/putkubernetesclusterskubernetesclusteridresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutKubernetesClustersKubernetesClusterIDManagedByHelm

Update the Kubernetes cluster which managed by HELM

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
    res, err := s.Kubernetes.PutKubernetesClustersKubernetesClusterIDManagedByHelm(ctx, operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest{
        EditKubernetesClusterManagedByHelm: shared.EditKubernetesClusterManagedByHelm{},
        KubernetesClusterID: "98ce5ff4-7cf2-43f7-b683-e7d0be8fe575",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest](../../pkg/models/operations/putkubernetesclusterskubernetesclusteridmanagedbyhelmrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse](../../pkg/models/operations/putkubernetesclusterskubernetesclusteridmanagedbyhelmresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutPodDefinitionsPodID

Change pod definition

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
    res, err := s.Kubernetes.PutPodDefinitionsPodID(ctx, operations.PutPodDefinitionsPodIDRequest{
        PodDefinition: shared.PodDefinitionInput{
            ClusterID: "8813d860-716a-4bd1-be6a-f36d0732a688",
            Containers: []shared.Container{
                shared.Container{},
            },
            Name: "<value>",
        },
        PodID: "d5b40efc-46b3-4d79-9349-9390e36a1df3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodDefinition != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutPodDefinitionsPodIDRequest](../../pkg/models/operations/putpoddefinitionspodidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutPodDefinitionsPodIDResponse](../../pkg/models/operations/putpoddefinitionspodidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
