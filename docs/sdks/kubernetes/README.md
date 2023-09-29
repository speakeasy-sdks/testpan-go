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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.DeleteKubernetesClustersKubernetesClusterIDRequest](../../models/operations/deletekubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.DeleteKubernetesClustersKubernetesClusterIDResponse](../../models/operations/deletekubernetesclusterskubernetesclusteridresponse.md), error**


## DeletePodDefinitionsPodID

Delete a pod definition

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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeletePodDefinitionsPodIDRequest](../../models/operations/deletepoddefinitionspodidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeletePodDefinitionsPodIDResponse](../../models/operations/deletepoddefinitionspodidresponse.md), error**


## GetGetControllerDataClusterID

get controller data using clusterId

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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetGetControllerDataClusterIDRequest](../../models/operations/getgetcontrollerdataclusteridrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetGetControllerDataClusterIDResponse](../../models/operations/getgetcontrollerdataclusteridresponse.md), error**


## GetIstioSupportedVersions

Get a list of istio releases that are supported by Secure Application agent. sorted from latest to oldest

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
    res, err := s.Kubernetes.GetIstioSupportedVersions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIstioSupportedVersions200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetIstioSupportedVersionsResponse](../../models/operations/getistiosupportedversionsresponse.md), error**


## GetKubernetesClusters

get a list of current Kubernetes clusters

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
    res, err := s.Kubernetes.GetKubernetesClusters(ctx, operations.GetKubernetesClustersRequest{
        ClusterName: testpango.String("handle"),
        ControllerStatus: testpango.String("Metical withdrawal unless"),
        ControllerVersion: testpango.String("Northeast Northeast"),
        DownloadAsXlsx: testpango.Bool(false),
        KubernetesVersion: testpango.String("SMTP PCI soulful"),
        MaxResults: testpango.Float64(5309.6),
        Offset: testpango.Float64(6721.13),
        OnlySpecReconstructionEnabledFilter: testpango.Bool(false),
        SortDir: operations.GetKubernetesClustersSortDirAsc.ToPointer(),
        SortKey: testpango.String("background mmm"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesClusterControllers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetKubernetesClustersRequest](../../models/operations/getkubernetesclustersrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetKubernetesClustersResponse](../../models/operations/getkubernetesclustersresponse.md), error**


## GetKubernetesClustersKubernetesClusterID

get the Kubernetes cluster with the given id

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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.GetKubernetesClustersKubernetesClusterIDRequest](../../models/operations/getkubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDResponse](../../models/operations/getkubernetesclusterskubernetesclusteridresponse.md), error**


## GetKubernetesClustersKubernetesClusterIDDeleteDependencies

get dependencies which need to be handled in order to delete specified Kubernetes cluster

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

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.GetKubernetesClustersKubernetesClusterIDDeleteDependenciesRequest](../../models/operations/getkubernetesclusterskubernetesclusteriddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDDeleteDependenciesResponse](../../models/operations/getkubernetesclusterskubernetesclusteriddeletedependenciesresponse.md), error**


## GetKubernetesClustersKubernetesClusterIDDownloadBundle

In order to install,  extract and run "./install_bundle.sh"

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDDownloadBundle(ctx, operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest{
        KubernetesClusterID: "204aba8f-ddfd-4080-9f3c-89482d02fb76",
        SendTelemetriesIntervalSec: testpango.Int64(908942),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetKubernetesClustersKubernetesClusterIDDownloadBundle200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleRequest](../../models/operations/getkubernetesclusterskubernetesclusteriddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDDownloadBundleResponse](../../models/operations/getkubernetesclusterskubernetesclusteriddownloadbundleresponse.md), error**


## GetKubernetesClustersKubernetesClusterIDGetHelmCommands

Get Panoptica Aug release Helm command

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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.GetKubernetesClustersKubernetesClusterIDGetHelmCommandsRequest](../../models/operations/getkubernetesclusterskubernetesclusteridgethelmcommandsrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDGetHelmCommandsResponse](../../models/operations/getkubernetesclusterskubernetesclusteridgethelmcommandsresponse.md), error**


## GetKubernetesClustersKubernetesClusterIDNamespaces

List namespaces on a specific Kubernetes cluster

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDNamespaces(ctx, operations.GetKubernetesClustersKubernetesClusterIDNamespacesRequest{
        IncludeScannable: testpango.Bool(false),
        KubernetesClusterID: "b18474d1-b95c-48d9-af1c-951ce15b0be6",
        SortDir: operations.GetKubernetesClustersKubernetesClusterIDNamespacesSortDirDesc.ToPointer(),
        SortKey: operations.GetKubernetesClustersKubernetesClusterIDNamespacesSortKeyName.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesNamespaceResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetKubernetesClustersKubernetesClusterIDNamespacesRequest](../../models/operations/getkubernetesclusterskubernetesclusteridnamespacesrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDNamespacesResponse](../../models/operations/getkubernetesclusterskubernetesclusteridnamespacesresponse.md), error**


## GetKubernetesClustersKubernetesClusterIDServices

List services on a specific Kubernetes cluster

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
    res, err := s.Kubernetes.GetKubernetesClustersKubernetesClusterIDServices(ctx, operations.GetKubernetesClustersKubernetesClusterIDServicesRequest{
        KubernetesClusterID: "1b08f941-6d25-44a9-8a5a-94e46b69d5cc",
        ShowIstioOnly: testpango.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KubernetesServices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.GetKubernetesClustersKubernetesClusterIDServicesRequest](../../models/operations/getkubernetesclusterskubernetesclusteridservicesrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |


### Response

**[*operations.GetKubernetesClustersKubernetesClusterIDServicesResponse](../../models/operations/getkubernetesclusterskubernetesclusteridservicesresponse.md), error**


## GetLeanKubernetesClusters

get a list of current Kubernetes clusters

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
    res, err := s.Kubernetes.GetLeanKubernetesClusters(ctx, operations.GetLeanKubernetesClustersRequest{
        ClusterName: testpango.String("sensor Granite"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LeanKubernetesClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetLeanKubernetesClustersRequest](../../models/operations/getleankubernetesclustersrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetLeanKubernetesClustersResponse](../../models/operations/getleankubernetesclustersresponse.md), error**


## GetNamespaces

Get a list of current Kubernetes namespaces

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
    res, err := s.Kubernetes.GetNamespaces(ctx, operations.GetNamespacesRequest{
        ClusterName: testpango.String("Thorium"),
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(6640.03),
        NamespaceName: testpango.String("mindshare"),
        Offset: testpango.Float64(617.97),
        ProtectionStatus: operations.GetNamespacesProtectionStatusDisabled.ToPointer(),
        SortDir: operations.GetNamespacesSortDirDesc.ToPointer(),
        SortKey: operations.GetNamespacesSortKeyProtectionStatus.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Namespaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetNamespacesRequest](../../models/operations/getnamespacesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetNamespacesResponse](../../models/operations/getnamespacesresponse.md), error**


## GetPodDefinitions

Get all pod definitions on the system

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
    res, err := s.Kubernetes.GetPodDefinitions(ctx, operations.GetPodDefinitionsRequest{
        ClusterName: testpango.String("Illinois Savings"),
        DeploymentType: []string{
            "disclaimer",
        },
        DownloadAsXlsx: testpango.Bool(false),
        Name: testpango.String("Multigender M2F"),
        NoPagination: testpango.Bool(false),
        TemplateSource: []string{
            "Boulder",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PodDefinitions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetPodDefinitionsRequest](../../models/operations/getpoddefinitionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetPodDefinitionsResponse](../../models/operations/getpoddefinitionsresponse.md), error**


## PostKubernetesClusters

Add a new Kubernetes cluster to Secure Application

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
    res, err := s.Kubernetes.PostKubernetesClusters(ctx, shared.KubernetesCluster{
        AccountName: testpango.String("withdrawal"),
        AgentFailClose: testpango.Bool(false),
        APIIntelligenceDAST: testpango.Bool(false),
        AutoLabelEnabled: testpango.Bool(false),
        AutoUpgradeControllerVersion: testpango.Bool(false),
        AutomatedPolicyRequiresDeployer: testpango.Bool(false),
        CiImageSignatureValidation: testpango.Bool(false),
        CiImageValidation: testpango.Bool(false),
        ClusterPodDefinitionSource: shared.ClusterPodDefinitionSourceKubernetes.ToPointer(),
        ControllerDataResponse: &shared.ControllerDataResponse{
            AgentID: testpango.String("81eb935a-67c1-4676-82ab-92109ab1c541"),
            SharedKey: testpango.String("Movies"),
        },
        ControllerStatus: shared.ControllerStatusUnknown.ToPointer(),
        EnableConnectionsControl: testpango.Bool(false),
        HelmCommandsInstallation: &shared.HelmCommandsInstallation{
            IstioHelmCommand: testpango.String("wherever Bedfordshire"),
            PanopticaHelmCommand: testpango.String("Personal generation Small"),
            VaultHelmCommand: testpango.String("Generic than"),
        },
        ID: testpango.String("6e76a9b0-598f-49f4-b42b-b8cf86555cd8"),
        InstallEnvoyTracingSupport: testpango.Bool(false),
        InstallTracingSupport: testpango.Bool(false),
        InstallationSource: shared.InstallationSourceHelm.ToPointer(),
        InternalRegistryParameters: &shared.InternalRegistryParameters{
            InternalRegistry: testpango.String("Pines Trial"),
            InternalRegistryEnabled: testpango.Bool(false),
        },
        IsHoldApplicationUntilProxyStarts: testpango.Bool(false),
        IsIstioIngressEnabled: testpango.Bool(false),
        IsMultiCluster: testpango.Bool(false),
        IsPersistent: testpango.Bool(false),
        IstioIngressAnnotations: []shared.KubernetesAnnotation{
            shared.KubernetesAnnotation{
                Key: "<key>",
                Value: "upright maximize verbally",
            },
        },
        IstioInstallationParameters: &shared.IstioInstallationParameters{
            IsIstioAlreadyInstalled: testpango.Bool(false),
            IstioVersion: testpango.String("ladybug"),
        },
        K8sEventsEnabled: testpango.Bool(false),
        KubernetesSecurity: testpango.Bool(false),
        MinimalNumberOfControllerReplicas: testpango.Int64(225610),
        Name: "female Branding Architect",
        OrchestrationType: shared.KubernetesClusterOrchestrationTypeRancher.ToPointer(),
        PreserveOriginalSourceIP: testpango.Bool(false),
        ProxyConfiguration: &shared.ProxyConfiguration{
            EnableProxy: testpango.Bool(false),
            HTTPSProxy: testpango.String("Customer befriend"),
        },
        RestrictRegistires: testpango.Bool(false),
        ScanConfiguration: &shared.ScanConfiguration{
            NumberOfScanners: testpango.Int64(127892),
            ScanTypes: []shared.ScanType{
                shared.ScanTypeVulnerabilities,
            },
        },
        ServiceDiscoveryIsolationEnabled: testpango.Bool(false),
        SidecarsResources: &shared.SidecarsResource{
            ProxyInitLimitsCPU: testpango.String("meter deposit South"),
            ProxyInitLimitsMemory: testpango.String("compress Lake"),
            ProxyInitRequestsCPU: testpango.String("as Tactics mobile"),
            ProxyInitRequestsMemory: testpango.String("blue Manager Program"),
            ProxyLimitsCPU: testpango.String("achievement Optimization"),
            ProxyLimitsMemory: testpango.String("Zirconium"),
            ProxyRequestCPU: testpango.String("fuchsia Texas Ergonomic"),
            ProxyRequestMemory: testpango.String("Longview"),
        },
        SSHMonitorDisabled: testpango.Bool(false),
        SupportExternalTraceSource: testpango.Bool(false),
        TLSInspectionEnabled: testpango.Bool(false),
        TokenInjectionEnabled: testpango.Bool(false),
        UseExternalCA: testpango.Bool(false),
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.KubernetesCluster](../../models/shared/kubernetescluster.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.PostKubernetesClustersResponse](../../models/operations/postkubernetesclustersresponse.md), error**


## PostPodDefinitions

Create a new pod definition

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
    res, err := s.Kubernetes.PostPodDefinitions(ctx, shared.PodDefinitionInput{
        ClusterID: "caf647ad-b35e-4462-baa4-a61d2b4e2410",
        Containers: []shared.Container{
            shared.Container{
                Image: &shared.Image{
                    DockerfileScanSeverity: shared.DockerfileScanSeverityFatal.ToPointer(),
                    Hash: testpango.String("Folk"),
                    Repository: testpango.String("Smart Trigender"),
                    Tag: testpango.String("Analyst"),
                    VulnerabilitySeverityLevel: shared.VulnerabilitySeverityHigh.ToPointer(),
                },
            },
        },
        InitContainers: []shared.Container{
            shared.Container{
                Image: &shared.Image{
                    DockerfileScanSeverity: shared.DockerfileScanSeverityInfo.ToPointer(),
                    Hash: testpango.String("yellow Solutions geez"),
                    Repository: testpango.String("resound mobile applications"),
                    Tag: testpango.String("meh"),
                    VulnerabilitySeverityLevel: shared.VulnerabilitySeverityLow.ToPointer(),
                },
            },
        },
        Kind: shared.PodTemplateKindReplicaSet.ToPointer(),
        Labels: []shared.Label{
            shared.Label{
                Key: "<key>",
                Value: "unless",
            },
        },
        Name: "payment Michigan",
        PodDefinitionSource: shared.PodDefinitionSourceKubernetes.ToPointer(),
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.PodDefinitionInput](../../models/shared/poddefinitioninput.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostPodDefinitionsResponse](../../models/operations/postpoddefinitionsresponse.md), error**


## PutKubernetesClustersKubernetesClusterID

Update the Kubernetes cluster

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
    res, err := s.Kubernetes.PutKubernetesClustersKubernetesClusterID(ctx, operations.PutKubernetesClustersKubernetesClusterIDRequest{
        KubernetesCluster: shared.KubernetesCluster{
            AccountName: testpango.String("savior normalization absentmindedly"),
            AgentFailClose: testpango.Bool(false),
            APIIntelligenceDAST: testpango.Bool(false),
            AutoLabelEnabled: testpango.Bool(false),
            AutoUpgradeControllerVersion: testpango.Bool(false),
            AutomatedPolicyRequiresDeployer: testpango.Bool(false),
            CiImageSignatureValidation: testpango.Bool(false),
            CiImageValidation: testpango.Bool(false),
            ClusterPodDefinitionSource: shared.ClusterPodDefinitionSourceCd.ToPointer(),
            ControllerDataResponse: &shared.ControllerDataResponse{
                AgentID: testpango.String("d957958a-0796-4a5f-ad9d-7f9c73b14840"),
                SharedKey: testpango.String("Fantastic Berkshire"),
            },
            ControllerStatus: shared.ControllerStatusPendingInstall.ToPointer(),
            EnableConnectionsControl: testpango.Bool(false),
            HelmCommandsInstallation: &shared.HelmCommandsInstallation{
                IstioHelmCommand: testpango.String("orange Central invoice"),
                PanopticaHelmCommand: testpango.String("pascal to squeegee"),
                VaultHelmCommand: testpango.String("parsing blissfully pine"),
            },
            ID: testpango.String("f781f7b5-14b5-45e8-9b88-a92334e6deac"),
            InstallEnvoyTracingSupport: testpango.Bool(false),
            InstallTracingSupport: testpango.Bool(false),
            InstallationSource: shared.InstallationSourceHelm.ToPointer(),
            InternalRegistryParameters: &shared.InternalRegistryParameters{
                InternalRegistry: testpango.String("North Highlands"),
                InternalRegistryEnabled: testpango.Bool(false),
            },
            IsHoldApplicationUntilProxyStarts: testpango.Bool(false),
            IsIstioIngressEnabled: testpango.Bool(false),
            IsMultiCluster: testpango.Bool(false),
            IsPersistent: testpango.Bool(false),
            IstioIngressAnnotations: []shared.KubernetesAnnotation{
                shared.KubernetesAnnotation{
                    Key: "<key>",
                    Value: "ADP Seamless",
                },
            },
            IstioInstallationParameters: &shared.IstioInstallationParameters{
                IsIstioAlreadyInstalled: testpango.Bool(false),
                IstioVersion: testpango.String("Southeast"),
            },
            K8sEventsEnabled: testpango.Bool(false),
            KubernetesSecurity: testpango.Bool(false),
            MinimalNumberOfControllerReplicas: testpango.Int64(890421),
            Name: "delectus Rap Awesome",
            OrchestrationType: shared.KubernetesClusterOrchestrationTypeIks.ToPointer(),
            PreserveOriginalSourceIP: testpango.Bool(false),
            ProxyConfiguration: &shared.ProxyConfiguration{
                EnableProxy: testpango.Bool(false),
                HTTPSProxy: testpango.String("Directives spar Handcrafted"),
            },
            RestrictRegistires: testpango.Bool(false),
            ScanConfiguration: &shared.ScanConfiguration{
                NumberOfScanners: testpango.Int64(663885),
                ScanTypes: []shared.ScanType{
                    shared.ScanTypeVulnerabilities,
                },
            },
            ServiceDiscoveryIsolationEnabled: testpango.Bool(false),
            SidecarsResources: &shared.SidecarsResource{
                ProxyInitLimitsCPU: testpango.String("tan comic tan"),
                ProxyInitLimitsMemory: testpango.String("Rubber"),
                ProxyInitRequestsCPU: testpango.String("Sleek drain sensor"),
                ProxyInitRequestsMemory: testpango.String("indexing Hackensack"),
                ProxyLimitsCPU: testpango.String("superstructure"),
                ProxyLimitsMemory: testpango.String("generating"),
                ProxyRequestCPU: testpango.String("micronutrient"),
                ProxyRequestMemory: testpango.String("pro Cambridgeshire"),
            },
            SSHMonitorDisabled: testpango.Bool(false),
            SupportExternalTraceSource: testpango.Bool(false),
            TLSInspectionEnabled: testpango.Bool(false),
            TokenInjectionEnabled: testpango.Bool(false),
            UseExternalCA: testpango.Bool(false),
        },
        KubernetesClusterID: "96a74b13-6974-4b67-a067-bd753e149f7b",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PutKubernetesClustersKubernetesClusterIDRequest](../../models/operations/putkubernetesclusterskubernetesclusteridrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PutKubernetesClustersKubernetesClusterIDResponse](../../models/operations/putkubernetesclusterskubernetesclusteridresponse.md), error**


## PutKubernetesClustersKubernetesClusterIDManagedByHelm

Update the Kubernetes cluster which managed by HELM

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
    res, err := s.Kubernetes.PutKubernetesClustersKubernetesClusterIDManagedByHelm(ctx, operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest{
        EditKubernetesClusterManagedByHelm: shared.EditKubernetesClusterManagedByHelm{
            Name: testpango.String("transmitter innocently"),
        },
        KubernetesClusterID: "ff47cf23-f7b6-483e-bd0b-e8fe575edb2f",
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

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmRequest](../../models/operations/putkubernetesclusterskubernetesclusteridmanagedbyhelmrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |


### Response

**[*operations.PutKubernetesClustersKubernetesClusterIDManagedByHelmResponse](../../models/operations/putkubernetesclusterskubernetesclusteridmanagedbyhelmresponse.md), error**


## PutPodDefinitionsPodID

Change pod definition

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
    res, err := s.Kubernetes.PutPodDefinitionsPodID(ctx, operations.PutPodDefinitionsPodIDRequest{
        PodDefinitionInput: shared.PodDefinitionInput{
            ClusterID: "8813d860-716a-4bd1-be6a-f36d0732a688",
            Containers: []shared.Container{
                shared.Container{
                    Image: &shared.Image{
                        DockerfileScanSeverity: shared.DockerfileScanSeverityFatal.ToPointer(),
                        Hash: testpango.String("female West"),
                        Repository: testpango.String("Cyclocross female jubilant"),
                        Tag: testpango.String("Rubber Tamiami"),
                        VulnerabilitySeverityLevel: shared.VulnerabilitySeverityLow.ToPointer(),
                    },
                },
            },
            InitContainers: []shared.Container{
                shared.Container{
                    Image: &shared.Image{
                        DockerfileScanSeverity: shared.DockerfileScanSeverityWarn.ToPointer(),
                        Hash: testpango.String("Southeast custody"),
                        Repository: testpango.String("Music Unbranded Visionary"),
                        Tag: testpango.String("Squares towards"),
                        VulnerabilitySeverityLevel: shared.VulnerabilitySeverityLow.ToPointer(),
                    },
                },
            },
            Kind: shared.PodTemplateKindStatefulSet.ToPointer(),
            Labels: []shared.Label{
                shared.Label{
                    Key: "<key>",
                    Value: "actualize Dollar",
                },
            },
            Name: "opt virtual copying",
            PodDefinitionSource: shared.PodDefinitionSourceHelm.ToPointer(),
        },
        PodID: "d7ad7bdf-9c15-4917-b40e-79415cfed383",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutPodDefinitionsPodIDRequest](../../models/operations/putpoddefinitionspodidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PutPodDefinitionsPodIDResponse](../../models/operations/putpoddefinitionspodidresponse.md), error**

