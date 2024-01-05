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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteCloudAccountsCloudAccountIDRequest](../../pkg/models/operations/deletecloudaccountscloudaccountidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.DeleteCloudAccountsCloudAccountIDResponse](../../pkg/models/operations/deletecloudaccountscloudaccountidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccounts

List all the cloud accounts on the system

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
    res, err := s.Serverless.GetCloudAccounts(ctx, operations.GetCloudAccountsRequest{
        SortKey: operations.GetCloudAccountsQueryParamSortKeyLastScanned,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCloudAccountsRequest](../../pkg/models/operations/getcloudaccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetCloudAccountsResponse](../../pkg/models/operations/getcloudaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsAzureInstallationDetails

Get the Azure installation details

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

**[*operations.GetCloudAccountsAzureInstallationDetailsResponse](../../pkg/models/operations/getcloudaccountsazureinstallationdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsInstallationDetails

Get the installation details

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

**[*operations.GetCloudAccountsInstallationDetailsResponse](../../pkg/models/operations/getcloudaccountsinstallationdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsRegionsAWS

List all the possible regions of AWS

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
    res, err := s.Serverless.GetCloudAccountsRegionsAWS(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudAccountsRegionsAWSResponse](../../pkg/models/operations/getcloudaccountsregionsawsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsRegionsAzure

List all the possible regions of Azure

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
    res, err := s.Serverless.GetCloudAccountsRegionsAzure(ctx)
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

**[*operations.GetCloudAccountsRegionsAzureResponse](../../pkg/models/operations/getcloudaccountsregionsazureresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsCloudAccountIDDeleteDependencies

get dependencies which need to be handled in order to delete specified cloud account

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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetCloudAccountsCloudAccountIDDeleteDependenciesRequest](../../pkg/models/operations/getcloudaccountscloudaccountiddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetCloudAccountsCloudAccountIDDeleteDependenciesResponse](../../pkg/models/operations/getcloudaccountscloudaccountiddeletedependenciesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCloudAccountsCloudAccountIDDownloadBundle

In order to install, extract and run "./install_bundle.sh"

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
    res, err := s.Serverless.GetCloudAccountsCloudAccountIDDownloadBundle(ctx, operations.GetCloudAccountsCloudAccountIDDownloadBundleRequest{
        CloudAccountID: "78a9218a-bbdc-48dd-8c99-acaaac2d314e",
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetCloudAccountsCloudAccountIDDownloadBundleRequest](../../pkg/models/operations/getcloudaccountscloudaccountiddownloadbundlerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetCloudAccountsCloudAccountIDDownloadBundleResponse](../../pkg/models/operations/getcloudaccountscloudaccountiddownloadbundleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctions

Get serverless functions

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
    res, err := s.Serverless.GetServerlessFunctions(ctx, operations.GetServerlessFunctionsRequest{
        CloudAccountsIdsFilter: []string{
            "2d213824-3222-42f0-b89a-c2b1bfba58c4",
        },
        FuncName: []string{
            "string",
        },
        PolicyRisk: []operations.PolicyRisk{
            operations.PolicyRiskLow,
        },
        Result: []operations.QueryParamResult{
            operations.QueryParamResultAllow,
        },
        Risk: []operations.QueryParamRisk{
            operations.QueryParamRiskHigh,
        },
        SecretsRisk: []operations.SecretsRisk{
            operations.SecretsRiskRiskIdentified,
        },
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetServerlessFunctionsRequest](../../pkg/models/operations/getserverlessfunctionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetServerlessFunctionsResponse](../../pkg/models/operations/getserverlessfunctionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctionsArns

Get serverless functions names

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
    res, err := s.Serverless.GetServerlessFunctionsArns(ctx, operations.GetServerlessFunctionsArnsRequest{
        FuncArn: []string{
            "string",
        },
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetServerlessFunctionsArnsRequest](../../pkg/models/operations/getserverlessfunctionsarnsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetServerlessFunctionsArnsResponse](../../pkg/models/operations/getserverlessfunctionsarnsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctionsNames

Get serverless functions names

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
    res, err := s.Serverless.GetServerlessFunctionsNames(ctx, operations.GetServerlessFunctionsNamesRequest{
        FuncName: []string{
            "string",
        },
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetServerlessFunctionsNamesRequest](../../pkg/models/operations/getserverlessfunctionsnamesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetServerlessFunctionsNamesResponse](../../pkg/models/operations/getserverlessfunctionsnamesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctionsFunctionID

Get Serverless Function by ID

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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetServerlessFunctionsFunctionIDRequest](../../pkg/models/operations/getserverlessfunctionsfunctionidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetServerlessFunctionsFunctionIDResponse](../../pkg/models/operations/getserverlessfunctionsfunctionidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctionsFunctionIDSecrets

Get Serverless Function secrets issues

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
    res, err := s.Serverless.GetServerlessFunctionsFunctionIDSecrets(ctx, operations.GetServerlessFunctionsFunctionIDSecretsRequest{
        FunctionID: "1200d55a-1387-4753-ab90-33c09b02720e",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetServerlessFunctionsFunctionIDSecretsRequest](../../pkg/models/operations/getserverlessfunctionsfunctionidsecretsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetServerlessFunctionsFunctionIDSecretsResponse](../../pkg/models/operations/getserverlessfunctionsfunctionidsecretsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessFunctionsFunctionIDVulnerabilities

Get Serverless Function Vulnerabilities by ID

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
    res, err := s.Serverless.GetServerlessFunctionsFunctionIDVulnerabilities(ctx, operations.GetServerlessFunctionsFunctionIDVulnerabilitiesRequest{
        FunctionID: "2ad182ca-55b9-436d-935a-2584c6608fc0",
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

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.GetServerlessFunctionsFunctionIDVulnerabilitiesRequest](../../pkg/models/operations/getserverlessfunctionsfunctionidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |


### Response

**[*operations.GetServerlessFunctionsFunctionIDVulnerabilitiesResponse](../../pkg/models/operations/getserverlessfunctionsfunctionidvulnerabilitiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessZipFiles

Get serverless zip files that was scanned by cli

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
    res, err := s.Serverless.GetServerlessZipFiles(ctx, operations.GetServerlessZipFilesRequest{})
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
| `request`                                                                                              | [operations.GetServerlessZipFilesRequest](../../pkg/models/operations/getserverlesszipfilesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetServerlessZipFilesResponse](../../pkg/models/operations/getserverlesszipfilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessZipFilesZipID

Get specific zip file record

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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetServerlessZipFilesZipIDRequest](../../pkg/models/operations/getserverlesszipfileszipidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetServerlessZipFilesZipIDResponse](../../pkg/models/operations/getserverlesszipfileszipidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessZipFilesZipIDPackages

Returns a list of packages for a specific serverless zip

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
    res, err := s.Serverless.GetServerlessZipFilesZipIDPackages(ctx, operations.GetServerlessZipFilesZipIDPackagesRequest{
        ZipID: "8930b13e-93d3-4a51-9609-c730ba7092f4",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetServerlessZipFilesZipIDPackagesRequest](../../pkg/models/operations/getserverlesszipfileszipidpackagesrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetServerlessZipFilesZipIDPackagesResponse](../../pkg/models/operations/getserverlesszipfileszipidpackagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetServerlessZipFilesZipIDVulnerabilities

Returns a list of vulnerabilities detected in the serverless zip

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
    res, err := s.Serverless.GetServerlessZipFilesZipIDVulnerabilities(ctx, operations.GetServerlessZipFilesZipIDVulnerabilitiesRequest{
        ZipID: "7d640878-f682-4a69-b4d3-f1a10ecad858",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.GetServerlessZipFilesZipIDVulnerabilitiesRequest](../../pkg/models/operations/getserverlesszipfileszipidvulnerabilitiesrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.GetServerlessZipFilesZipIDVulnerabilitiesResponse](../../pkg/models/operations/getserverlesszipfileszipidvulnerabilitiesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCloudAccountsScan

invoke cloud account scan

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"context"
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.ServerlessScanConfig](../../pkg/models/shared/serverlessscanconfig.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PostCloudAccountsScanResponse](../../pkg/models/operations/postcloudaccountsscanresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutCloudAccountsCloudAccountID

Edit cloud account definition

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
    res, err := s.Serverless.PutCloudAccountsCloudAccountID(ctx, operations.PutCloudAccountsCloudAccountIDRequest{
        CloudAccount: shared.CloudAccountInput{
            PeriodicJobExpression: shared.ServerlessPeriodicJobExpression{
                PeriodicJobType: shared.ServerlessPeriodicJobExpressionPeriodicJobTypeServerlessByHoursPeriodicJobExpression,
            },
            Regions: []string{
                "string",
            },
            SecurityThreats: &shared.CloudAccountSecurityThreats{},
            VulnerabilitiesSummary: &shared.VulnerabilitiesSummary{},
        },
        CloudAccountID: "3c18ca7f-b155-4344-88dc-c7697a579ecd",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PutCloudAccountsCloudAccountIDRequest](../../pkg/models/operations/putcloudaccountscloudaccountidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.PutCloudAccountsCloudAccountIDResponse](../../pkg/models/operations/putcloudaccountscloudaccountidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
