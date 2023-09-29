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
        ID: "1532f8e0-5c4e-41fa-a5bc-5aa03e071f17",
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
            "Frozen",
        },
        ImageName: []string{
            "olive",
        },
        ImageTag: []string{
            "IB",
        },
        VulnerabilityName: testpango.String("Administrator"),
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
            "Health",
        },
        ImageName: []string{
            "male",
        },
        ImageTag: []string{
            "Ytterbium",
        },
        DownloadAsXlsx: testpango.Bool(false),
        MaxResults: testpango.Float64(9052.61),
        Offset: testpango.Float64(4508.91),
        SortDir: operations.GetImagesSortDirDesc.ToPointer(),
        SortKey: operations.GetImagesSortKeyImageName,
        VulnerabilityName: testpango.String("grow"),
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
        ImageHash: testpango.String("firewall Misty"),
        MaxResults: testpango.Float64(8276.26),
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
        ImageHash: "Dale Iowa",
        ImageName: "boutique Health Martinique",
        IsIgnored: testpango.Bool(false),
        LayerID: testpango.String("4f03eaaf-7b1e-44de-a52b-c93cd397f018"),
        MaxResults: testpango.Float64(5718.05),
        Offset: testpango.Float64(7191.45),
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
        ID: "d118d3e1-8466-490a-a3b5-54afae983ffe",
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
        ImageID: "da882664-c35f-4933-94c4-06103746e14d",
        IsIgnored: testpango.Bool(false),
        MaxResults: testpango.Float64(4943.96),
        Offset: testpango.Float64(6053.79),
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
        ImageID: "15439eb8-81bd-4ffd-8863-3eb436058a37",
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
        ImageID: "c2ee574a-71ae-4a8b-a2ff-c4930d365b60",
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
        ImageID: "32f4442d-9f10-4117-a8a5-a2e7f00c922a",
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
        ImageID: "d0043143-505e-42fc-aacc-aa87ecde910e",
        IsIgnored: testpango.Bool(false),
        LayerID: testpango.String("7b364751-2259-4df9-ae67-4e88ed17b104"),
        MaxResults: testpango.Float64(7244.87),
        Offset: testpango.Float64(6826.65),
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
        ImageHash: testpango.String("Sedan silently"),
        ImageName: testpango.String("Strontium"),
        ImageTags: []string{
            "gosh",
        },
        TimeAdded: types.MustTimeFromString("2021-03-23T13:45:56.960Z"),
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
                "afe40a3f-7c16-48da-ac19-b6309faa9f44",
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
                "7d411c9e-43ee-4529-bae8-7474b2c192fe",
            },
        },
        ActionType: operations.PostImagesImageIDDockerfileScanResultsIgnoreActionTypeAdd,
        ImageID: "fcccfe79-79d3-4058-b255-f4d4f301de39",
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
                "0f457328-8079-4ea1-b64e-d7631fc85bb9",
            },
        },
        ActionType: operations.PostImagesImageIDVulnerabilitiesIgnoreActionTypeRemove,
        ImageID: "95b06784-3712-40b3-827e-08cfaaddc5ee",
        SnoozeTime: operations.PostImagesImageIDVulnerabilitiesIgnoreSnoozeTimeMonth.ToPointer(),
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

