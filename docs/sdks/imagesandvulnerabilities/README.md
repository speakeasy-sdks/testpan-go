# ImagesAndVulnerabilities
(*ImagesAndVulnerabilities*)

## Overview

APIs used to define and manage  image hashes

### Available Operations

* [DeleteImagesID](#deleteimagesid) - Delete an image hash
* [GetAccountVulnerabilitiesXlsx](#getaccountvulnerabilitiesxlsx) - Returns a xlsx file of images alongside to their vulnerabilities.
* [GetImages](#getimages) - Returns a list of images
* [GetImagesImagesHash](#getimagesimageshash) - search for image hash in the account
* [GetImagesVulnerabilitiesByImageNameAndHash](#getimagesvulnerabilitiesbyimagenameandhash) - Returns a list of vulnerabilities detected in the image
* [GetImagesID](#getimagesid) - get an image
* [GetImagesImageIDDockerfileScanResults](#getimagesimageiddockerfilescanresults) - Returns a list of vulnerabilities detected in the  image
* [GetImagesImageIDImageLayers](#getimagesimageidimagelayers) - Returns a list of image layers
* [GetImagesImageIDPackages](#getimagesimageidpackages) - Returns a list of packages for a specific image
* [GetImagesImageIDSbomPath](#getimagesimageidsbompath) - Returns the path to the SBOM in cloud storage
* [GetImagesImageIDVulnerabilities](#getimagesimageidvulnerabilities) - Returns a list of vulnerabilities detected in the image
* [PostImages](#postimages) - Define a New image hash
* [PostImagesApprove](#postimagesapprove) - Approve an image hash
* [PostImagesImageIDDockerfileScanResultsIgnore](#postimagesimageiddockerfilescanresultsignore) - Add / remove a list of  UUIDs of dockerfileScanResults from ignored list
* [PostImagesImageIDVulnerabilitiesIgnore](#postimagesimageidvulnerabilitiesignore) - Add / remove a list of  UUIDs of vulnerabilities from ignored list

## DeleteImagesID

Delete an image hash

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
    res, err := s.ImagesAndVulnerabilities.DeleteImagesID(ctx, operations.DeleteImagesIDRequest{
        ID: "aaf9bbad-185f-4e43-9d6b-f5c838fbb8c2",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteImagesIDRequest](../../models/operations/deleteimagesidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteImagesIDResponse](../../models/operations/deleteimagesidresponse.md), error**


## GetAccountVulnerabilitiesXlsx

Returns a xlsx file of images alongside to their vulnerabilities.

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
    res, err := s.ImagesAndVulnerabilities.GetAccountVulnerabilitiesXlsx(ctx, operations.GetAccountVulnerabilitiesXlsxRequest{
        ImageHash: []string{
            "eaque",
        },
        ImageName: []string{
            "impedit",
        },
        ImageTag: []string{
            "nam",
        },
        VulnerabilityName: testpango.String("ex"),
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAccountVulnerabilitiesXlsxRequest](../../models/operations/getaccountvulnerabilitiesxlsxrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetAccountVulnerabilitiesXlsxResponse](../../models/operations/getaccountvulnerabilitiesxlsxresponse.md), error**


## GetImages

Returns a list of images

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
    res, err := s.ImagesAndVulnerabilities.GetImages(ctx, operations.GetImagesRequest{
        ImageHash: []string{
            "odio",
        },
        ImageName: []string{
            "delectus",
        },
        ImageTag: []string{
            "minus",
        },
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(2835.14),
        Offset: testpango.Float64(7160.58),
        SortDir: operations.GetImagesSortDirAsc.ToPointer(),
        SortKey: operations.GetImagesSortKeyImageName,
        VulnerabilityName: testpango.String("veniam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageDefGets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetImagesRequest](../../models/operations/getimagesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetImagesResponse](../../models/operations/getimagesresponse.md), error**


## GetImagesImagesHash

search for image hash in the account

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImagesHash(ctx, operations.GetImagesImagesHashRequest{
        ImageHash: testpango.String("repudiandae"),
        MaxResults: testpango.Float64(5723.17),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetImagesImagesHash200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetImagesImagesHashRequest](../../models/operations/getimagesimageshashrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetImagesImagesHashResponse](../../models/operations/getimagesimageshashresponse.md), error**


## GetImagesVulnerabilitiesByImageNameAndHash

Returns a list of vulnerabilities detected in the image

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
    res, err := s.ImagesAndVulnerabilities.GetImagesVulnerabilitiesByImageNameAndHash(ctx, operations.GetImagesVulnerabilitiesByImageNameAndHashRequest{
        ImageHash: "occaecati",
        ImageName: "debitis",
        IsIgnored: testpango.Bool(false),
        LayerID: testpango.String("6234c9f7-b79d-4feb-b7a5-c38d4baf91e5"),
        MaxResults: testpango.Float64(524.07),
        Offset: testpango.Float64(4064.62),
        ShowOnlyVulnerabilitiesWithFix: testpango.Bool(false),
        SortDir: operations.GetImagesVulnerabilitiesByImageNameAndHashSortDirDesc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Vulnerabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetImagesVulnerabilitiesByImageNameAndHashRequest](../../models/operations/getimagesvulnerabilitiesbyimagenameandhashrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetImagesVulnerabilitiesByImageNameAndHashResponse](../../models/operations/getimagesvulnerabilitiesbyimagenameandhashresponse.md), error**


## GetImagesID

get an image

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
    res, err := s.ImagesAndVulnerabilities.GetImagesID(ctx, operations.GetImagesIDRequest{
        ID: "f890a54b-475f-416f-96d3-85a3c4ac631b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageDefGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetImagesIDRequest](../../models/operations/getimagesidrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetImagesIDResponse](../../models/operations/getimagesidresponse.md), error**


## GetImagesImageIDDockerfileScanResults

Returns a list of vulnerabilities detected in the  image

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDDockerfileScanResults(ctx, operations.GetImagesImageIDDockerfileScanResultsRequest{
        ImageID: "99e26ced-8f9f-4db9-810f-63bbf817837b",
        IsIgnored: testpango.Bool(false),
        MaxResults: testpango.Float64(457.28),
        Offset: testpango.Float64(1127.88),
        SortDir: operations.GetImagesImageIDDockerfileScanResultsSortDirDesc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DockerfileScanResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetImagesImageIDDockerfileScanResultsRequest](../../models/operations/getimagesimageiddockerfilescanresultsrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.GetImagesImageIDDockerfileScanResultsResponse](../../models/operations/getimagesimageiddockerfilescanresultsresponse.md), error**


## GetImagesImageIDImageLayers

Returns a list of image layers

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDImageLayers(ctx, operations.GetImagesImageIDImageLayersRequest{
        ImageID: "fdd78862-4189-4eb4-8873-f5033f19dbf1",
        IsIgnored: testpango.Bool(false),
        SortDir: operations.GetImagesImageIDImageLayersSortDirAsc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageLayers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetImagesImageIDImageLayersRequest](../../models/operations/getimagesimageidimagelayersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetImagesImageIDImageLayersResponse](../../models/operations/getimagesimageidimagelayersresponse.md), error**


## GetImagesImageIDPackages

Returns a list of packages for a specific image

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDPackages(ctx, operations.GetImagesImageIDPackagesRequest{
        ImageID: "5ce4152e-ab9c-4d7e-9224-a6a0e123b784",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagePackageDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetImagesImageIDPackagesRequest](../../models/operations/getimagesimageidpackagesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetImagesImageIDPackagesResponse](../../models/operations/getimagesimageidpackagesresponse.md), error**


## GetImagesImageIDSbomPath

Returns the path to the SBOM in cloud storage

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDSbomPath(ctx, operations.GetImagesImageIDSbomPathRequest{
        ImageID: "7ec59e1f-67f3-4c4c-8e4b-6d7696ff3c57",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageSbomPathResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetImagesImageIDSbomPathRequest](../../models/operations/getimagesimageidsbompathrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetImagesImageIDSbomPathResponse](../../models/operations/getimagesimageidsbompathresponse.md), error**


## GetImagesImageIDVulnerabilities

Returns a list of vulnerabilities detected in the image

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
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDVulnerabilities(ctx, operations.GetImagesImageIDVulnerabilitiesRequest{
        ImageID: "47501357-e44f-451f-8b08-4c3197e193a2",
        IsIgnored: testpango.Bool(false),
        LayerID: testpango.String("45467f94-874c-42d5-8c49-72233e66bd8f"),
        MaxResults: testpango.Float64(8866.84),
        Offset: testpango.Float64(3578.67),
        ShowOnlyVulnerabilitiesWithFix: testpango.Bool(false),
        SortDir: operations.GetImagesImageIDVulnerabilitiesSortDirDesc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Vulnerabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetImagesImageIDVulnerabilitiesRequest](../../models/operations/getimagesimageidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetImagesImageIDVulnerabilitiesResponse](../../models/operations/getimagesimageidvulnerabilitiesresponse.md), error**


## PostImages

Define a New image hash

### Example Usage

```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/types"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.PostImages(ctx, shared.ImageDefInput{
        ImageHash: testpango.String("aut"),
        ImageName: testpango.String("voluptatem"),
        ImageTags: []string{
            "libero",
        },
        TimeAdded: types.MustTimeFromString("2022-01-11T15:07:25.289Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageDefGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.ImageDefInput](../../models/shared/imagedefinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostImagesResponse](../../models/operations/postimagesresponse.md), error**


## PostImagesApprove

Approve an image hash

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
    res, err := s.ImagesAndVulnerabilities.PostImagesApprove(ctx, operations.PostImagesApproveRequest{
        UUIDList: shared.UUIDList{
            UUIDList: []string{
                "9ef20387-3205-490c-8c10-96400313b3e5",
            },
        },
        IsImageApproved: false,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostImagesApproveRequest](../../models/operations/postimagesapproverequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostImagesApproveResponse](../../models/operations/postimagesapproveresponse.md), error**


## PostImagesImageIDDockerfileScanResultsIgnore

Add / remove a list of  UUIDs of dockerfileScanResults from ignored list

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
    res, err := s.ImagesAndVulnerabilities.PostImagesImageIDDockerfileScanResultsIgnore(ctx, operations.PostImagesImageIDDockerfileScanResultsIgnoreRequest{
        UUIDList: shared.UUIDList{
            UUIDList: []string{
                "044f65fe-72dc-4407-bd0c-c3f408efc15c",
            },
        },
        ActionType: operations.PostImagesImageIDDockerfileScanResultsIgnoreActionTypeRemove,
        ImageID: "b4d6e1ea-e0f7-45ae-9f2a-cab58b991c92",
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.PostImagesImageIDDockerfileScanResultsIgnoreRequest](../../models/operations/postimagesimageiddockerfilescanresultsignorerequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.PostImagesImageIDDockerfileScanResultsIgnoreResponse](../../models/operations/postimagesimageiddockerfilescanresultsignoreresponse.md), error**


## PostImagesImageIDVulnerabilitiesIgnore

Add / remove a list of  UUIDs of vulnerabilities from ignored list

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
    res, err := s.ImagesAndVulnerabilities.PostImagesImageIDVulnerabilitiesIgnore(ctx, operations.PostImagesImageIDVulnerabilitiesIgnoreRequest{
        UUIDList: shared.UUIDList{
            UUIDList: []string{
                "6ddb5894-61e7-4421-8be6-d9502f0ea930",
            },
        },
        ActionType: operations.PostImagesImageIDVulnerabilitiesIgnoreActionTypeRemove,
        ImageID: "69f7ac2f-72f8-4850-8904-911608207888",
        SnoozeTime: operations.PostImagesImageIDVulnerabilitiesIgnoreSnoozeTimeWeek.ToPointer(),
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PostImagesImageIDVulnerabilitiesIgnoreRequest](../../models/operations/postimagesimageidvulnerabilitiesignorerequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PostImagesImageIDVulnerabilitiesIgnoreResponse](../../models/operations/postimagesimageidvulnerabilitiesignoreresponse.md), error**

