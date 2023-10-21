# Aws
(*Aws*)

## Overview

APIs used to change  credentials or return details about the  user's AWS environment

### Available Operations

* [GetAwsAccounts](#getawsaccounts) - Get a list of AWS accounts
* [GetAwsRoles](#getawsroles) - Lists AWS role ARNs for the account
* [GetAwsTags](#getawstags) - Get a list of AWS tag keys
* [GetAwsAwsAccountIDRegions](#getawsawsaccountidregions) - Get a list of regions for the  AWS account
* [GetAwsAwsAccountIDRegionIDSubnets](#getawsawsaccountidregionidsubnets) - Get a list of AWS subnets for an AWS account and region
* [GetAwsAwsAccountIDRegionIDVpcs](#getawsawsaccountidregionidvpcs) - Get a list of VPCs for AWS accounts.
* [PostAwsRoles](#postawsroles) - Add AWS role to the account
* [PutAwsRolesRoleID](#putawsrolesroleid) - Change AWS role name

## GetAwsAccounts

Returns a list of AWS accounts for this Secure Application account.


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
    res, err := s.Aws.GetAwsAccounts(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AWSAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAwsAccountsResponse](../../models/operations/getawsaccountsresponse.md), error**


## GetAwsRoles

Lists AWS role ARNs for the account

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
    res, err := s.Aws.GetAwsRoles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AWSRoles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAwsRolesResponse](../../models/operations/getawsrolesresponse.md), error**


## GetAwsTags

Get a list of AWS tag keys

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
    res, err := s.Aws.GetAwsTags(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Tags != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAwsTagsResponse](../../models/operations/getawstagsresponse.md), error**


## GetAwsAwsAccountIDRegions

Returns a list of regions for AWS account.


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
    res, err := s.Aws.GetAwsAwsAccountIDRegions(ctx, operations.GetAwsAwsAccountIDRegionsRequest{
        AwsAccountID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAwsAwsAccountIDRegions200ApplicationJSONStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetAwsAwsAccountIDRegionsRequest](../../models/operations/getawsawsaccountidregionsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetAwsAwsAccountIDRegionsResponse](../../models/operations/getawsawsaccountidregionsresponse.md), error**


## GetAwsAwsAccountIDRegionIDSubnets

Get a list of AWS subnets for an AWS account and region

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
    res, err := s.Aws.GetAwsAwsAccountIDRegionIDSubnets(ctx, operations.GetAwsAwsAccountIDRegionIDSubnetsRequest{
        AwsAccountID: "string",
        RegionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PortshiftAwsSubnets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetAwsAwsAccountIDRegionIDSubnetsRequest](../../models/operations/getawsawsaccountidregionidsubnetsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GetAwsAwsAccountIDRegionIDSubnetsResponse](../../models/operations/getawsawsaccountidregionidsubnetsresponse.md), error**


## GetAwsAwsAccountIDRegionIDVpcs

Returns a list of VPCs for an AWS account and region. These values are used to define a Secure Application environment.


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
    res, err := s.Aws.GetAwsAwsAccountIDRegionIDVpcs(ctx, operations.GetAwsAwsAccountIDRegionIDVpcsRequest{
        AwsAccountID: "string",
        RegionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VpcResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAwsAwsAccountIDRegionIDVpcsRequest](../../models/operations/getawsawsaccountidregionidvpcsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetAwsAwsAccountIDRegionIDVpcsResponse](../../models/operations/getawsawsaccountidregionidvpcsresponse.md), error**


## PostAwsRoles

Upload a role ARN, that Secure Application will use to connect to the AWS account.

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
    res, err := s.Aws.PostAwsRoles(ctx, shared.AWSRolePost{
        Arn: "string",
        Name: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AWSRole != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.AWSRolePost](../../models/shared/awsrolepost.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.PostAwsRolesResponse](../../models/operations/postawsrolesresponse.md), error**


## PutAwsRolesRoleID

Change AWS role name

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
    res, err := s.Aws.PutAwsRolesRoleID(ctx, operations.PutAwsRolesRoleIDRequest{
        AWSRoleDetails: shared.AWSRoleDetails{
            Name: "string",
        },
        RoleID: "81eef92a-b1e0-45da-ba61-a597d143757f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AWSRole != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PutAwsRolesRoleIDRequest](../../models/operations/putawsrolesroleidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PutAwsRolesRoleIDResponse](../../models/operations/putawsrolesroleidresponse.md), error**

