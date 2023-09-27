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
        KubernetesClusterID: "50841f17-6445-4637-9f3f-b27e21f86265",
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
        PodID: "7b36fc6b-9f58-47ce-925c-67641a8312e5",
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
        ClusterID: "047b4c21-ccb4-423a-bcdc-91faabdd88e7",
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
        ClusterName: testpango.String("vitae"),
        ControllerStatus: testpango.String("delectus"),
        ControllerVersion: testpango.String("laboriosam"),
        DownloadAsXlsx: testpango.Bool(false),
        KubernetesVersion: testpango.String("minus"),
        MaxResults: testpango.Float64(2894.1),
        Offset: testpango.Float64(5587.31),
        OnlySpecReconstructionEnabledFilter: testpango.Bool(false),
        SortDir: operations.GetKubernetesClustersSortDirAsc.ToPointer(),
        SortKey: testpango.String("veniam"),
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
        KubernetesClusterID: "2d7771e7-fd07-4400-9ef8-d29de1dd7097",
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
        KubernetesClusterID: "b5da08c5-7fa6-4c78-a216-e19bafeca619",
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
        KubernetesClusterID: "1498140b-64ff-48ae-970e-f03b5f37e4aa",
        SendTelemetriesIntervalSec: testpango.Int64(504386),
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
        KubernetesClusterID: "68555966-732a-4a5d-8b66-82cb70f8cfd5",
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
        KubernetesClusterID: "fb6e91b9-a9f7-4484-ae2c-3309db0536d9",
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
        KubernetesClusterID: "5ca006f5-392c-411a-a5a8-bf92f97428ad",
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
        ClusterName: testpango.String("unde"),
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
        ClusterName: testpango.String("laborum"),
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(5904.22),
        NamespaceName: testpango.String("hic"),
        Offset: testpango.Float64(5458.54),
        ProtectionStatus: operations.GetNamespacesProtectionStatusDisabled.ToPointer(),
        SortDir: operations.GetNamespacesSortDirDesc.ToPointer(),
        SortKey: operations.GetNamespacesSortKeyRunningPods.ToPointer(),
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
        ClusterName: testpango.String("explicabo"),
        DeploymentType: []string{
            "odit",
        },
        DownloadAsXlsx: testpango.Bool(false),
        Name: testpango.String("Julie Cremin"),
        NoPagination: testpango.Bool(false),
        TemplateSource: []string{
            "ullam",
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
        AccountName: testpango.String("provident"),
        AgentFailClose: testpango.Bool(false),
        APIIntelligenceDAST: testpango.Bool(false),
        AutoLabelEnabled: testpango.Bool(false),
        AutoUpgradeControllerVersion: testpango.Bool(false),
        AutomatedPolicyRequiresDeployer: testpango.Bool(false),
        CiImageSignatureValidation: testpango.Bool(false),
        CiImageValidation: testpango.Bool(false),
        ClusterPodDefinitionSource: shared.ClusterPodDefinitionSourceCd.ToPointer(),
        ControllerDataResponse: &shared.ControllerDataResponse{
            AgentID: testpango.String("98387f7a-79cd-472c-9248-4da21729f2ac"),
            SharedKey: testpango.String("numquam"),
        },
        ControllerStatus: shared.ControllerStatusPendingInstall.ToPointer(),
        EnableConnectionsControl: testpango.Bool(false),
        HelmCommandsInstallation: &shared.HelmCommandsInstallation{
            IstioHelmCommand: testpango.String("necessitatibus"),
            PanopticaHelmCommand: testpango.String("tenetur"),
            VaultHelmCommand: testpango.String("exercitationem"),
        },
        ID: testpango.String("725f1169-ac1e-441d-8a23-c23e34f2dfa4"),
        InstallEnvoyTracingSupport: testpango.Bool(false),
        InstallTracingSupport: testpango.Bool(false),
        InstallationSource: shared.InstallationSourceHelm.ToPointer(),
        InternalRegistryParameters: &shared.InternalRegistryParameters{
            InternalRegistry: testpango.String("sunt"),
            InternalRegistryEnabled: testpango.Bool(false),
        },
        IsHoldApplicationUntilProxyStarts: testpango.Bool(false),
        IsIstioIngressEnabled: testpango.Bool(false),
        IsMultiCluster: testpango.Bool(false),
        IsPersistent: testpango.Bool(false),
        IstioIngressAnnotations: []shared.KubernetesAnnotation{
            shared.KubernetesAnnotation{
                Key: "perspiciatis",
                Value: "quam",
            },
        },
        IstioInstallationParameters: &shared.IstioInstallationParameters{
            IsIstioAlreadyInstalled: testpango.Bool(false),
            IstioVersion: testpango.String("a"),
        },
        K8sEventsEnabled: testpango.Bool(false),
        KubernetesSecurity: testpango.Bool(false),
        MinimalNumberOfControllerReplicas: testpango.Int64(434345),
        Name: "Clay Monahan",
        OrchestrationType: shared.KubernetesClusterOrchestrationTypeGke.ToPointer(),
        PreserveOriginalSourceIP: testpango.Bool(false),
        ProxyConfiguration: &shared.ProxyConfiguration{
            EnableProxy: testpango.Bool(false),
            HTTPSProxy: testpango.String("ipsam"),
        },
        RestrictRegistires: testpango.Bool(false),
        ScanConfiguration: &shared.ScanConfiguration{
            NumberOfScanners: testpango.Int64(92978),
            ScanTypes: []shared.ScanType{
                shared.ScanTypeDockerCisBenchmark,
            },
        },
        ServiceDiscoveryIsolationEnabled: testpango.Bool(false),
        SidecarsResources: &shared.SidecarsResource{
            ProxyInitLimitsCPU: testpango.String("saepe"),
            ProxyInitLimitsMemory: testpango.String("sunt"),
            ProxyInitRequestsCPU: testpango.String("in"),
            ProxyInitRequestsMemory: testpango.String("architecto"),
            ProxyLimitsCPU: testpango.String("sed"),
            ProxyLimitsMemory: testpango.String("voluptatem"),
            ProxyRequestCPU: testpango.String("perspiciatis"),
            ProxyRequestMemory: testpango.String("error"),
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
        ClusterID: "853e9f54-3d85-4443-9ee2-24460443bc15",
        Containers: []shared.Container{
            shared.Container{
                Image: &shared.Image{
                    DockerfileScanSeverity: shared.DockerfileScanSeverityInfo.ToPointer(),
                    Hash: testpango.String("vitae"),
                    Repository: testpango.String("quos"),
                    Tag: testpango.String("atque"),
                    VulnerabilitySeverityLevel: shared.VulnerabilitySeverityHigh.ToPointer(),
                },
            },
        },
        InitContainers: []shared.Container{
            shared.Container{
                Image: &shared.Image{
                    DockerfileScanSeverity: shared.DockerfileScanSeverityInfo.ToPointer(),
                    Hash: testpango.String("asperiores"),
                    Repository: testpango.String("corporis"),
                    Tag: testpango.String("vel"),
                    VulnerabilitySeverityLevel: shared.VulnerabilitySeverityCritical.ToPointer(),
                },
            },
        },
        Kind: shared.PodTemplateKindDaemonSet.ToPointer(),
        Labels: []shared.Label{
            shared.Label{
                Key: "ipsam",
                Value: "at",
            },
        },
        Name: "Jessie Larkin",
        PodDefinitionSource: shared.PodDefinitionSourceCli.ToPointer(),
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
            AccountName: testpango.String("mollitia"),
            AgentFailClose: testpango.Bool(false),
            APIIntelligenceDAST: testpango.Bool(false),
            AutoLabelEnabled: testpango.Bool(false),
            AutoUpgradeControllerVersion: testpango.Bool(false),
            AutomatedPolicyRequiresDeployer: testpango.Bool(false),
            CiImageSignatureValidation: testpango.Bool(false),
            CiImageValidation: testpango.Bool(false),
            ClusterPodDefinitionSource: shared.ClusterPodDefinitionSourceCd.ToPointer(),
            ControllerDataResponse: &shared.ControllerDataResponse{
                AgentID: testpango.String("d617c3b0-d51a-444b-b01b-ad8706d46082"),
                SharedKey: testpango.String("libero"),
            },
            ControllerStatus: shared.ControllerStatusAutoUpgradeInProgress.ToPointer(),
            EnableConnectionsControl: testpango.Bool(false),
            HelmCommandsInstallation: &shared.HelmCommandsInstallation{
                IstioHelmCommand: testpango.String("nam"),
                PanopticaHelmCommand: testpango.String("pariatur"),
                VaultHelmCommand: testpango.String("quod"),
            },
            ID: testpango.String("41ff5d4e-2ae4-4fb5-8b35-d17638f1edb7"),
            InstallEnvoyTracingSupport: testpango.Bool(false),
            InstallTracingSupport: testpango.Bool(false),
            InstallationSource: shared.InstallationSourceHelm.ToPointer(),
            InternalRegistryParameters: &shared.InternalRegistryParameters{
                InternalRegistry: testpango.String("consectetur"),
                InternalRegistryEnabled: testpango.Bool(false),
            },
            IsHoldApplicationUntilProxyStarts: testpango.Bool(false),
            IsIstioIngressEnabled: testpango.Bool(false),
            IsMultiCluster: testpango.Bool(false),
            IsPersistent: testpango.Bool(false),
            IstioIngressAnnotations: []shared.KubernetesAnnotation{
                shared.KubernetesAnnotation{
                    Key: "nemo",
                    Value: "provident",
                },
            },
            IstioInstallationParameters: &shared.IstioInstallationParameters{
                IsIstioAlreadyInstalled: testpango.Bool(false),
                IstioVersion: testpango.String("accusamus"),
            },
            K8sEventsEnabled: testpango.Bool(false),
            KubernetesSecurity: testpango.Bool(false),
            MinimalNumberOfControllerReplicas: testpango.Int64(793260),
            Name: "Warren Runolfsson",
            OrchestrationType: shared.KubernetesClusterOrchestrationTypeRancher.ToPointer(),
            PreserveOriginalSourceIP: testpango.Bool(false),
            ProxyConfiguration: &shared.ProxyConfiguration{
                EnableProxy: testpango.Bool(false),
                HTTPSProxy: testpango.String("doloremque"),
            },
            RestrictRegistires: testpango.Bool(false),
            ScanConfiguration: &shared.ScanConfiguration{
                NumberOfScanners: testpango.Int64(964052),
                ScanTypes: []shared.ScanType{
                    shared.ScanTypeDockerCisBenchmark,
                },
            },
            ServiceDiscoveryIsolationEnabled: testpango.Bool(false),
            SidecarsResources: &shared.SidecarsResource{
                ProxyInitLimitsCPU: testpango.String("impedit"),
                ProxyInitLimitsMemory: testpango.String("illum"),
                ProxyInitRequestsCPU: testpango.String("ullam"),
                ProxyInitRequestsMemory: testpango.String("praesentium"),
                ProxyLimitsCPU: testpango.String("perferendis"),
                ProxyLimitsMemory: testpango.String("soluta"),
                ProxyRequestCPU: testpango.String("animi"),
                ProxyRequestMemory: testpango.String("molestiae"),
            },
            SSHMonitorDisabled: testpango.Bool(false),
            SupportExternalTraceSource: testpango.Bool(false),
            TLSInspectionEnabled: testpango.Bool(false),
            TokenInjectionEnabled: testpango.Bool(false),
            UseExternalCA: testpango.Bool(false),
        },
        KubernetesClusterID: "3810e4fe-4447-4297-8d3b-1dd3bbce247b",
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
            Name: testpango.String("Vicki Labadie"),
        },
        KubernetesClusterID: "ff50126d-71cf-4fbd-8eb7-4b8421953b44",
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
            ClusterID: "bd3c4315-9d33-4e59-93c0-01139863aa41",
            Containers: []shared.Container{
                shared.Container{
                    Image: &shared.Image{
                        DockerfileScanSeverity: shared.DockerfileScanSeverityFatal.ToPointer(),
                        Hash: testpango.String("aliquid"),
                        Repository: testpango.String("optio"),
                        Tag: testpango.String("adipisci"),
                        VulnerabilitySeverityLevel: shared.VulnerabilitySeverityUnknown.ToPointer(),
                    },
                },
            },
            InitContainers: []shared.Container{
                shared.Container{
                    Image: &shared.Image{
                        DockerfileScanSeverity: shared.DockerfileScanSeverityFatal.ToPointer(),
                        Hash: testpango.String("porro"),
                        Repository: testpango.String("explicabo"),
                        Tag: testpango.String("reiciendis"),
                        VulnerabilitySeverityLevel: shared.VulnerabilitySeverityUnknown.ToPointer(),
                    },
                },
            },
            Kind: shared.PodTemplateKindOther.ToPointer(),
            Labels: []shared.Label{
                shared.Label{
                    Key: "porro",
                    Value: "tempore",
                },
            },
            Name: "Joan Schaefer",
            PodDefinitionSource: shared.PodDefinitionSourceHelm.ToPointer(),
        },
        PodID: "1ffbe9cb-d795-4ee6-9e07-6cc7abf616ea",
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

