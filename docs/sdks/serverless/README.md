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
        CloudAccountID: "df22a94d-20ec-490e-a41d-1f465e85156f",
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
        CloudAccountName: testpango.String("sapiente"),
        MaxResults: testpango.Float64(9524.11),
        NoPagination: testpango.Bool(false),
        Offset: testpango.Float64(4778.26),
        Region: testpango.String("amet"),
        SortDir: operations.GetCloudAccountsSortDirDesc.ToPointer(),
        SortKey: operations.GetCloudAccountsSortKeyName,
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
        CloudAccountID: "f54fdd5e-a954-4339-8daf-b42a8d63388e",
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
        CloudAccountID: "4d8039ea-5f9b-418a-a44f-d619039dacd3",
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
        CloudAccountName: testpango.String("voluptatum"),
        CloudAccountsIdsFilter: []string{
            "ed0dc671-dc7f-41e3-af15-920c90d1b490",
        },
        DownloadAsXlsx: testpango.Bool(false),
        FuncName: []string{
            "dicta",
        },
        MaxResults: testpango.Float64(9957.13),
        Offset: testpango.Float64(1728.07),
        PolicyRisk: []GetServerlessFunctionsPolicyRisk{
            operations.GetServerlessFunctionsPolicyRiskHigh,
        },
        Region: testpango.String("quibusdam"),
        Result: []GetServerlessFunctionsResult{
            operations.GetServerlessFunctionsResultDetect,
        },
        Risk: []GetServerlessFunctionsRisk{
            operations.GetServerlessFunctionsRiskHigh,
        },
        SecretsRisk: []GetServerlessFunctionsSecretsRisk{
            operations.GetServerlessFunctionsSecretsRiskRiskIdentified,
        },
        SortDir: operations.GetServerlessFunctionsSortDirDesc.ToPointer(),
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
        CloudAccountName: testpango.String("id"),
        FuncArn: []string{
            "neque",
        },
        Region: testpango.String("dolores"),
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
        CloudAccountName: testpango.String("vel"),
        FuncName: []string{
            "ipsum",
        },
        Region: testpango.String("occaecati"),
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
        FunctionID: "da5b7b69-02b8-481a-94f6-43664a8f0af8",
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
        FunctionID: "c691d732-d9fb-4af9-876a-2ae8dcc50c8a",
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
        FunctionID: "3512c737-8489-4307-90a0-0e966ec736d4",
        MaxResults: testpango.Float64(2068.14),
        Offset: testpango.Float64(816.89),
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
    res, err := s.Serverless.GetServerlessZipFiles(ctx, operations.GetServerlessZipFilesRequest{
        MaxResults: testpango.Float64(6064.72),
        Offset: testpango.Float64(3096.94),
        SortDir: operations.GetServerlessZipFilesSortDirAsc.ToPointer(),
        SortKey: operations.GetServerlessZipFilesSortKeyVulnerabilities.ToPointer(),
        ZipNameFilter: testpango.String("corrupti"),
        ZipSha256Filter: testpango.String("optio"),
    })
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
        ZipID: "783c9239-8ed3-4d3a-b7ca-3c5ca8649a70",
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
        ZipID: "cfd5d698-9b72-4064-9107-7d19ea83d492",
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
        MaxResults: testpango.Float64(9357.71),
        Offset: testpango.Float64(8426.89),
        SortDir: operations.GetServerlessZipFilesZipIDVulnerabilitiesSortDirAsc.ToPointer(),
        ZipID: "4b8a2c19-5454-45e9-95dc-c185ea4901c7",
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
            "c43ad2da-a784-4aba-bd23-0edf73811a11",
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
            CloudProvider: shared.CloudProviderTypeAws.ToPointer(),
            InstallVulnerabilityScanner: testpango.Bool(false),
            Name: testpango.String("Lena Cruickshank"),
            PeriodicJobExpression: shared.ServerlessPeriodicJobExpression{
                PeriodicJobType: shared.ServerlessPeriodicJobExpressionPeriodicJobTypeServerlessByHoursPeriodicJobExpression,
            },
            Regions: []string{
                "eveniet",
            },
            SecurityThreats: &shared.CloudAccountSecurityThreats{
                DataAccessRisk: shared.ServerlessDataAccessRiskMedium.ToPointer(),
                DataAccessRiskCount: testpango.Int64(313305),
                IsUnusedFunction: testpango.Bool(false),
                PolicyRisk: shared.ServerlessPolicyRiskLow.ToPointer(),
                PolicyRiskCount: testpango.Int64(331703),
                PubliclyAccessibleRisk: shared.ServerlessPubliclyAccessibleRiskNoRisk.ToPointer(),
                PubliclyAccessibleRiskCount: testpango.Int64(464242),
                SecretsRisk: shared.ServerlessSecretsRiskNoKnownRisk.ToPointer(),
                SecretsRiskCount: testpango.Int64(136036),
                UnusedFunctionCount: testpango.Int64(85002),
            },
            ValidateFunction: shared.CloudAccountValidateFunctionNone.ToPointer(),
            VulnerabilitiesSummary: &shared.VulnerabilitiesSummary{
                Critical: testpango.Int64(322054),
                High: testpango.Int64(533096),
                Low: testpango.Int64(985155),
                Medium: testpango.Int64(296712),
                Total: testpango.Int64(857355),
                Unknown: testpango.Int64(481914),
            },
        },
        CloudAccountID: "396564c2-0a07-411a-961d-24a7dbb8f532",
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

