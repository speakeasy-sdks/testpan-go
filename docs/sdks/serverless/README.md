# Serverless
(*Serverless*)

### Available Operations

* [DeleteCloudAccountsCloudAccountID](#deletecloudaccountscloudaccountid) - Delete a cloud account
* [GetCloudAccounts](#getcloudaccounts) - List all the cloud accounts on the system
* [GetCloudAccountsAzureInstallationDetails](#getcloudaccountsazureinstallationdetails) - Get the Azure installation details
* [GetCloudAccountsInstallationDetails](#getcloudaccountsinstallationdetails) - Get the installation details
* [GetCloudAccountsRegionsAWS](#getcloudaccountsregionsaws) - List all the possible regions of AWS
* [GetCloudAccountsRegionsAzure](#getcloudaccountsregionsazure) - List all the possible regions of Azure
* [GetCloudAccountsCloudAccountIDDeleteDependencies](#getcloudaccountscloudaccountiddeletedependencies) - get dependencies which need to be handled in order to delete specified cloud account
* [GetCloudAccountsCloudAccountIDDownloadBundle](#getcloudaccountscloudaccountiddownloadbundle) - Get Secure Application installation script
* [GetServerlessFunctions](#getserverlessfunctions) - Get serverless functions
* [GetServerlessFunctionsArns](#getserverlessfunctionsarns) - Get serverless functions names
* [GetServerlessFunctionsNames](#getserverlessfunctionsnames) - Get serverless functions names
* [GetServerlessFunctionsFunctionID](#getserverlessfunctionsfunctionid) - Get Serverless Function by ID
* [GetServerlessFunctionsFunctionIDSecrets](#getserverlessfunctionsfunctionidsecrets) - Get Serverless Function secrets issues
* [GetServerlessFunctionsFunctionIDVulnerabilities](#getserverlessfunctionsfunctionidvulnerabilities) - Get Serverless Function Vulnerabilities by ID
* [GetServerlessZipFiles](#getserverlesszipfiles) - Get serverless zip files that was scanned by cli
* [GetServerlessZipFilesZipID](#getserverlesszipfileszipid) - Get specific zip file record
* [GetServerlessZipFilesZipIDPackages](#getserverlesszipfileszipidpackages) - Returns a list of packages for a specific serverless zip
* [GetServerlessZipFilesZipIDVulnerabilities](#getserverlesszipfileszipidvulnerabilities) - Returns a list of vulnerabilities detected in the serverless zip
* [PostCloudAccountsScan](#postcloudaccountsscan) - invoke cloud account scan
* [PutCloudAccountsCloudAccountID](#putcloudaccountscloudaccountid) - Edit cloud account definition

## DeleteCloudAccountsCloudAccountID

Delete a cloud account

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
    res, err := s.Serverless.DeleteCloudAccountsCloudAccountID(ctx, operations.DeleteCloudAccountsCloudAccountIDRequest{
        CloudAccountID: "a4061d9a-ae48-48cc-a173-2eca65acca85",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteCloudAccountsCloudAccountIDRequest](../../models/operations/deletecloudaccountscloudaccountidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.DeleteCloudAccountsCloudAccountIDResponse](../../models/operations/deletecloudaccountscloudaccountidresponse.md), error**


## GetCloudAccounts

List all the cloud accounts on the system

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
    res, err := s.Serverless.GetCloudAccounts(ctx, operations.GetCloudAccountsRequest{
        SortKey: operations.GetCloudAccountsSortKeyLastScanned,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CloudAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCloudAccountsRequest](../../models/operations/getcloudaccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetCloudAccountsResponse](../../models/operations/getcloudaccountsresponse.md), error**


## GetCloudAccountsAzureInstallationDetails

Get the Azure installation details

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
    res, err := s.Serverless.GetCloudAccountsAzureInstallationDetails(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessAzureInstallationDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudAccountsAzureInstallationDetailsResponse](../../models/operations/getcloudaccountsazureinstallationdetailsresponse.md), error**


## GetCloudAccountsInstallationDetails

Get the installation details

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
    res, err := s.Serverless.GetCloudAccountsInstallationDetails(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessInstallationDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudAccountsInstallationDetailsResponse](../../models/operations/getcloudaccountsinstallationdetailsresponse.md), error**


## GetCloudAccountsRegionsAWS

List all the possible regions of AWS

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
    res, err := s.Serverless.GetCloudAccountsRegionsAWS(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AwsRegions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudAccountsRegionsAWSResponse](../../models/operations/getcloudaccountsregionsawsresponse.md), error**


## GetCloudAccountsRegionsAzure

List all the possible regions of Azure

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
    res, err := s.Serverless.GetCloudAccountsRegionsAzure(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCloudAccountsRegionsAzure200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudAccountsRegionsAzureResponse](../../models/operations/getcloudaccountsregionsazureresponse.md), error**


## GetCloudAccountsCloudAccountIDDeleteDependencies

get dependencies which need to be handled in order to delete specified cloud account

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
    res, err := s.Serverless.GetCloudAccountsCloudAccountIDDeleteDependencies(ctx, operations.GetCloudAccountsCloudAccountIDDeleteDependenciesRequest{
        CloudAccountID: "64304886-9365-422a-8282-644a12d7387e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CloudAccountDeleteDependencies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.GetCloudAccountsCloudAccountIDDeleteDependenciesRequest](../../models/operations/getcloudaccountscloudaccountiddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |


### Response

**[*operations.GetCloudAccountsCloudAccountIDDeleteDependenciesResponse](../../models/operations/getcloudaccountscloudaccountiddeletedependenciesresponse.md), error**


## GetCloudAccountsCloudAccountIDDownloadBundle

In order to install, extract and run "./install_bundle.sh"

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
    res, err := s.Serverless.GetCloudAccountsCloudAccountIDDownloadBundle(ctx, operations.GetCloudAccountsCloudAccountIDDownloadBundleRequest{
        CloudAccountID: "78a9218a-bbdc-48dd-8c99-acaaac2d314e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCloudAccountsCloudAccountIDDownloadBundle200ApplicationJSONBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetCloudAccountsCloudAccountIDDownloadBundleRequest](../../models/operations/getcloudaccountscloudaccountiddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetCloudAccountsCloudAccountIDDownloadBundleResponse](../../models/operations/getcloudaccountscloudaccountiddownloadbundleresponse.md), error**


## GetServerlessFunctions

Get serverless functions

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
    res, err := s.Serverless.GetServerlessFunctions(ctx, operations.GetServerlessFunctionsRequest{
        CloudAccountsIdsFilter: []string{
            "2d213824-3222-42f0-b89a-c2b1bfba58c4",
        },
        FuncName: []string{
            "Electronics",
        },
        PolicyRisk: []operations.GetServerlessFunctionsPolicyRisk{
            operations.GetServerlessFunctionsPolicyRiskHigh,
        },
        Result: []operations.GetServerlessFunctionsResult{
            operations.GetServerlessFunctionsResultDetect,
        },
        Risk: []operations.GetServerlessFunctionsRisk{
            operations.GetServerlessFunctionsRiskCritical,
        },
        SecretsRisk: []operations.GetServerlessFunctionsSecretsRisk{
            operations.GetServerlessFunctionsSecretsRiskRiskIdentified,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessFunctions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetServerlessFunctionsRequest](../../models/operations/getserverlessfunctionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetServerlessFunctionsResponse](../../models/operations/getserverlessfunctionsresponse.md), error**


## GetServerlessFunctionsArns

Get serverless functions names

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
    res, err := s.Serverless.GetServerlessFunctionsArns(ctx, operations.GetServerlessFunctionsArnsRequest{
        FuncArn: []string{
            "Northeast",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessFunctionArns != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetServerlessFunctionsArnsRequest](../../models/operations/getserverlessfunctionsarnsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetServerlessFunctionsArnsResponse](../../models/operations/getserverlessfunctionsarnsresponse.md), error**


## GetServerlessFunctionsNames

Get serverless functions names

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
    res, err := s.Serverless.GetServerlessFunctionsNames(ctx, operations.GetServerlessFunctionsNamesRequest{
        FuncName: []string{
            "Hermaphrodite",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessFunctionNames != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetServerlessFunctionsNamesRequest](../../models/operations/getserverlessfunctionsnamesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetServerlessFunctionsNamesResponse](../../models/operations/getserverlessfunctionsnamesresponse.md), error**


## GetServerlessFunctionsFunctionID

Get Serverless Function by ID

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
    res, err := s.Serverless.GetServerlessFunctionsFunctionID(ctx, operations.GetServerlessFunctionsFunctionIDRequest{
        FunctionID: "c94885c8-e582-488b-bf74-bf3ac1e10be4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessFunction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetServerlessFunctionsFunctionIDRequest](../../models/operations/getserverlessfunctionsfunctionidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GetServerlessFunctionsFunctionIDResponse](../../models/operations/getserverlessfunctionsfunctionidresponse.md), error**


## GetServerlessFunctionsFunctionIDSecrets

Get Serverless Function secrets issues

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
    res, err := s.Serverless.GetServerlessFunctionsFunctionIDSecrets(ctx, operations.GetServerlessFunctionsFunctionIDSecretsRequest{
        FunctionID: "1200d55a-1387-4753-ab90-33c09b02720e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessFunctionSecretIssues != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetServerlessFunctionsFunctionIDSecretsRequest](../../models/operations/getserverlessfunctionsfunctionidsecretsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetServerlessFunctionsFunctionIDSecretsResponse](../../models/operations/getserverlessfunctionsfunctionidsecretsresponse.md), error**


## GetServerlessFunctionsFunctionIDVulnerabilities

Get Serverless Function Vulnerabilities by ID

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
    res, err := s.Serverless.GetServerlessFunctionsFunctionIDVulnerabilities(ctx, operations.GetServerlessFunctionsFunctionIDVulnerabilitiesRequest{
        FunctionID: "2ad182ca-55b9-436d-935a-2584c6608fc0",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.GetServerlessFunctionsFunctionIDVulnerabilitiesRequest](../../models/operations/getserverlessfunctionsfunctionidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.GetServerlessFunctionsFunctionIDVulnerabilitiesResponse](../../models/operations/getserverlessfunctionsfunctionidvulnerabilitiesresponse.md), error**


## GetServerlessZipFiles

Get serverless zip files that was scanned by cli

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
    res, err := s.Serverless.GetServerlessZipFiles(ctx, operations.GetServerlessZipFilesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessZips != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetServerlessZipFilesRequest](../../models/operations/getserverlesszipfilesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetServerlessZipFilesResponse](../../models/operations/getserverlesszipfilesresponse.md), error**


## GetServerlessZipFilesZipID

Get specific zip file record

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
    res, err := s.Serverless.GetServerlessZipFilesZipID(ctx, operations.GetServerlessZipFilesZipIDRequest{
        ZipID: "47af1679-7e54-41ca-b6c1-2fafc1426c1d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerlessZip != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetServerlessZipFilesZipIDRequest](../../models/operations/getserverlesszipfileszipidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetServerlessZipFilesZipIDResponse](../../models/operations/getserverlesszipfileszipidresponse.md), error**


## GetServerlessZipFilesZipIDPackages

Returns a list of packages for a specific serverless zip

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
    res, err := s.Serverless.GetServerlessZipFilesZipIDPackages(ctx, operations.GetServerlessZipFilesZipIDPackagesRequest{
        ZipID: "8930b13e-93d3-4a51-9609-c730ba7092f4",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetServerlessZipFilesZipIDPackagesRequest](../../models/operations/getserverlesszipfileszipidpackagesrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetServerlessZipFilesZipIDPackagesResponse](../../models/operations/getserverlesszipfileszipidpackagesresponse.md), error**


## GetServerlessZipFilesZipIDVulnerabilities

Returns a list of vulnerabilities detected in the serverless zip

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
    res, err := s.Serverless.GetServerlessZipFilesZipIDVulnerabilities(ctx, operations.GetServerlessZipFilesZipIDVulnerabilitiesRequest{
        ZipID: "7d640878-f682-4a69-b4d3-f1a10ecad858",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetServerlessZipFilesZipIDVulnerabilitiesRequest](../../models/operations/getserverlesszipfileszipidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetServerlessZipFilesZipIDVulnerabilitiesResponse](../../models/operations/getserverlesszipfileszipidvulnerabilitiesresponse.md), error**


## PostCloudAccountsScan

invoke cloud account scan

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
    res, err := s.Serverless.PostCloudAccountsScan(ctx, shared.ServerlessScanConfig{
        CloudAccounts: []string{
            "458b5913-dc88-4dbf-ab7f-5674c6953077",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.ServerlessScanConfig](../../models/shared/serverlessscanconfig.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostCloudAccountsScanResponse](../../models/operations/postcloudaccountsscanresponse.md), error**


## PutCloudAccountsCloudAccountID

Edit cloud account definition

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
    res, err := s.Serverless.PutCloudAccountsCloudAccountID(ctx, operations.PutCloudAccountsCloudAccountIDRequest{
        CloudAccountInput: shared.CloudAccountInput{
            PeriodicJobExpression: shared.ServerlessPeriodicJobExpression{
                PeriodicJobType: shared.ServerlessPeriodicJobExpressionPeriodicJobTypeServerlessByHoursPeriodicJobExpression,
            },
            Regions: []string{
                "Cheese",
            },
            SecurityThreats: &shared.CloudAccountSecurityThreats{},
            VulnerabilitiesSummary: &shared.VulnerabilitiesSummary{},
        },
        CloudAccountID: "18ca7fb1-5534-4488-9cc7-697a579ecdc3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CloudAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutCloudAccountsCloudAccountIDRequest](../../models/operations/putcloudaccountscloudaccountidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PutCloudAccountsCloudAccountIDResponse](../../models/operations/putcloudaccountscloudaccountidresponse.md), error**
