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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteImagesIDRequest](../../pkg/models/operations/deleteimagesidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteImagesIDResponse](../../pkg/models/operations/deleteimagesidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAccountVulnerabilitiesXlsx

Returns a xlsx file of images alongside to their vulnerabilities.

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetAccountVulnerabilitiesXlsx(ctx, operations.GetAccountVulnerabilitiesXlsxRequest{
        ImageHash: []string{
            "string",
        },
        ImageName: []string{
            "string",
        },
        ImageTag: []string{
            "string",
        },
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetAccountVulnerabilitiesXlsxRequest](../../pkg/models/operations/getaccountvulnerabilitiesxlsxrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetAccountVulnerabilitiesXlsxResponse](../../pkg/models/operations/getaccountvulnerabilitiesxlsxresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImages

Returns a list of images

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImages(ctx, operations.GetImagesRequest{
        ImageHash: []string{
            "string",
        },
        ImageName: []string{
            "string",
        },
        ImageTag: []string{
            "string",
        },
        SortKey: operations.GetImagesQueryParamSortKeyImageName,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetImagesRequest](../../pkg/models/operations/getimagesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetImagesResponse](../../pkg/models/operations/getimagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImagesHash

search for image hash in the account

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImagesImagesHash(ctx, operations.GetImagesImagesHashRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetImagesImagesHashRequest](../../pkg/models/operations/getimagesimageshashrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetImagesImagesHashResponse](../../pkg/models/operations/getimagesimageshashresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesVulnerabilitiesByImageNameAndHash

Returns a list of vulnerabilities detected in the image

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImagesVulnerabilitiesByImageNameAndHash(ctx, operations.GetImagesVulnerabilitiesByImageNameAndHashRequest{
        ImageHash: "string",
        ImageName: "string",
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetImagesVulnerabilitiesByImageNameAndHashRequest](../../pkg/models/operations/getimagesvulnerabilitiesbyimagenameandhashrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetImagesVulnerabilitiesByImageNameAndHashResponse](../../pkg/models/operations/getimagesvulnerabilitiesbyimagenameandhashresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesID

get an image

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetImagesIDRequest](../../pkg/models/operations/getimagesidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetImagesIDResponse](../../pkg/models/operations/getimagesidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImageIDDockerfileScanResults

Returns a list of vulnerabilities detected in the  image

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDDockerfileScanResults(ctx, operations.GetImagesImageIDDockerfileScanResultsRequest{
        ImageID: "da882664-c35f-4933-94c4-06103746e14d",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetImagesImageIDDockerfileScanResultsRequest](../../pkg/models/operations/getimagesimageiddockerfilescanresultsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetImagesImageIDDockerfileScanResultsResponse](../../pkg/models/operations/getimagesimageiddockerfilescanresultsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImageIDImageLayers

Returns a list of image layers

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDImageLayers(ctx, operations.GetImagesImageIDImageLayersRequest{
        ImageID: "15439eb8-81bd-4ffd-8863-3eb436058a37",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetImagesImageIDImageLayersRequest](../../pkg/models/operations/getimagesimageidimagelayersrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetImagesImageIDImageLayersResponse](../../pkg/models/operations/getimagesimageidimagelayersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImageIDPackages

Returns a list of packages for a specific image

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

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetImagesImageIDPackagesRequest](../../pkg/models/operations/getimagesimageidpackagesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetImagesImageIDPackagesResponse](../../pkg/models/operations/getimagesimageidpackagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImageIDSbomPath

Returns the path to the SBOM in cloud storage

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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetImagesImageIDSbomPathRequest](../../pkg/models/operations/getimagesimageidsbompathrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetImagesImageIDSbomPathResponse](../../pkg/models/operations/getimagesimageidsbompathresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetImagesImageIDVulnerabilities

Returns a list of vulnerabilities detected in the image

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.GetImagesImageIDVulnerabilities(ctx, operations.GetImagesImageIDVulnerabilitiesRequest{
        ImageID: "d0043143-505e-42fc-aacc-aa87ecde910e",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetImagesImageIDVulnerabilitiesRequest](../../pkg/models/operations/getimagesimageidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GetImagesImageIDVulnerabilitiesResponse](../../pkg/models/operations/getimagesimageidvulnerabilitiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostImages

Define a New image hash

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
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ImagesAndVulnerabilities.PostImages(ctx, shared.ImageDef{
        ImageTags: []string{
            "string",
        },
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

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.ImageDef](../../pkg/models/shared/imagedef.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.PostImagesResponse](../../pkg/models/operations/postimagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostImagesApprove

Approve an image hash

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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostImagesApproveRequest](../../pkg/models/operations/postimagesapproverequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostImagesApproveResponse](../../pkg/models/operations/postimagesapproveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostImagesImageIDDockerfileScanResultsIgnore

Add / remove a list of  UUIDs of dockerfileScanResults from ignored list

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
        ActionType: operations.ActionTypeAdd,
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PostImagesImageIDDockerfileScanResultsIgnoreRequest](../../pkg/models/operations/postimagesimageiddockerfilescanresultsignorerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.PostImagesImageIDDockerfileScanResultsIgnoreResponse](../../pkg/models/operations/postimagesimageiddockerfilescanresultsignoreresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostImagesImageIDVulnerabilitiesIgnore

Add / remove a list of  UUIDs of vulnerabilities from ignored list

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
        ActionType: operations.QueryParamActionTypeRemove,
        ImageID: "95b06784-3712-40b3-827e-08cfaaddc5ee",
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
| `request`                                                                                                                                | [operations.PostImagesImageIDVulnerabilitiesIgnoreRequest](../../pkg/models/operations/postimagesimageidvulnerabilitiesignorerequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PostImagesImageIDVulnerabilitiesIgnoreResponse](../../pkg/models/operations/postimagesimageidvulnerabilitiesignoreresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
