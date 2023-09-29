# Users
(*Users*)

## Overview

APIs used for login and password management

### Available Operations

* [DeleteUsersUserID](#deleteusersuserid) - Delete a user
* [GetOperatorCredentials](#getoperatorcredentials) - get the credentials of the Secure Application Operator service user
* [GetUsers](#getusers) - List current users
* [GetUsersUserIDAccessTokens](#getusersuseridaccesstokens) - Get the  access tokens for the user
* [GetUsersUserIDDeleteDependencies](#getusersuseriddeletedependencies) - get dependencies which need to be handled in order to delete specified user
* [PostAccountUsageStatus](#postaccountusagestatus) - an api to get the account usage status
* [PostChangePassword](#postchangepassword) - Change the password for the current user
* [PostLogin](#postlogin) - Login
* [PostLogout](#postlogout) - Log out
* [PostMe](#postme) - an api to get current logged in user info
* [PostUsers](#postusers) - Create a user
* [PostUsersAcceptEula](#postusersaccepteula) - Accept the EULA
* [PostUsersTrial](#postuserstrial) - Create a trail user
* [PutUsersUserID](#putusersuserid) - Change user details

## DeleteUsersUserID

Delete a user

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
    res, err := s.Users.DeleteUsersUserID(ctx, operations.DeleteUsersUserIDRequest{
        UserID: "2d4aef6d-76db-4c57-a2a3-fe8efd3db6e2",
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
| `request`                                                                                  | [operations.DeleteUsersUserIDRequest](../../models/operations/deleteusersuseridrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteUsersUserIDResponse](../../models/operations/deleteusersuseridresponse.md), error**


## GetOperatorCredentials

get the credentials of the Secure Application Operator service user

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
    res, err := s.Users.GetOperatorCredentials(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetOperatorCredentialsResponse](../../models/operations/getoperatorcredentialsresponse.md), error**


## GetUsers

List current users

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
    res, err := s.Users.GetUsers(ctx, operations.GetUsersRequest{
        Email: testpango.String("Martine_Welch@hotmail.com"),
        Roles: []GetUsersRoles{
            operations.GetUsersRolesAccountAdmin,
        },
        Username: testpango.String("Domenick_Schulist87"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserDisplays != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetUsersRequest](../../models/operations/getusersrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetUsersResponse](../../models/operations/getusersresponse.md), error**


## GetUsersUserIDAccessTokens

Get the access tokens for the user, to access Secure Application

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
    res, err := s.Users.GetUsersUserIDAccessTokens(ctx, operations.GetUsersUserIDAccessTokensRequest{
        UserID: "d48eca12-2332-4def-a816-08c21f5e1906",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetUsersUserIDAccessTokensRequest](../../models/operations/getusersuseridaccesstokensrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetUsersUserIDAccessTokensResponse](../../models/operations/getusersuseridaccesstokensresponse.md), error**


## GetUsersUserIDDeleteDependencies

get dependencies which need to be handled in order to delete specified user

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
    res, err := s.Users.GetUsersUserIDDeleteDependencies(ctx, operations.GetUsersUserIDDeleteDependenciesRequest{
        UserID: "8a75ffe4-62b3-4153-829f-86055d39570f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDependencyElementUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetUsersUserIDDeleteDependenciesRequest](../../models/operations/getusersuseriddeletedependenciesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GetUsersUserIDDeleteDependenciesResponse](../../models/operations/getusersuseriddeletedependenciesresponse.md), error**


## PostAccountUsageStatus

an api to get the account usage status

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
    res, err := s.Users.PostAccountUsageStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UsageStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostAccountUsageStatusResponse](../../models/operations/postaccountusagestatusresponse.md), error**


## PostChangePassword

Change the password for the current user

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
    res, err := s.Users.PostChangePassword(ctx, shared.ChangePasswordInfo{
        NewPassword: testpango.String("Clifton Tuna invoice"),
        OldPassword: testpango.String("Loan"),
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.ChangePasswordInfo](../../models/shared/changepasswordinfo.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PostChangePasswordResponse](../../models/operations/postchangepasswordresponse.md), error**


## PostLogin

Login

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
    res, err := s.Users.PostLogin(ctx, operations.PostLoginRequest{
        GoogleIDToken: testpango.String("Paradigm Touring enthusiastically"),
        Token: testpango.String("gray"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserLoginInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.PostLoginRequest](../../models/operations/postloginrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostLoginResponse](../../models/operations/postloginresponse.md), error**


## PostLogout

Log out

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
    res, err := s.Users.PostLogout(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostLogoutResponse](../../models/operations/postlogoutresponse.md), error**


## PostMe

an api to get current logged in user info

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
    res, err := s.Users.PostMe(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserLoginInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostMeResponse](../../models/operations/postmeresponse.md), error**


## PostUsers

Create a new user. Must be admin user.


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
    res, err := s.Users.PostUsers(ctx, shared.UserInput{
        Description: testpango.String("Cross-group reciprocal toolset"),
        Email: testpango.String("Erin_Wintheiser@hotmail.com"),
        FullName: "Mrs. Todd Halvorson",
        ID: testpango.String("d86a2273-6c5e-406f-80e1-8a3ed58c3e0d"),
        Role: shared.RolePortshiftAuditor.ToPointer(),
        Status: shared.UserStatusDisabled,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.UserInput](../../models/shared/userinput.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostUsersResponse](../../models/operations/postusersresponse.md), error**


## PostUsersAcceptEula

Accept the EULA

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
    res, err := s.Users.PostUsersAcceptEula(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PostUsersAcceptEulaResponse](../../models/operations/postusersaccepteularesponse.md), error**


## PostUsersTrial

Create a trail user

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
    res, err := s.Users.PostUsersTrial(ctx, operations.PostUsersTrialRequest{
        TrialUser: shared.TrialUser{
            Company: testpango.String("Rohan and Sons"),
            Email: "Tremaine.Bashirian@gmail.com",
            FirstName: "Sim",
            HowDidYouHearAboutUs: shared.TrialUserHowDidYouHearAboutUsAdvertising.ToPointer(),
            JobTitle: testpango.String("Human Program Liaison"),
            LastName: "Kovacek",
            PrivacyPolicyAndTermsAndConditionsAgreement: false,
        },
        GRecaptchaResponse: "Automated",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostUsersTrialRequest](../../models/operations/postuserstrialrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostUsersTrialResponse](../../models/operations/postuserstrialresponse.md), error**


## PutUsersUserID

Change user details

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
    res, err := s.Users.PutUsersUserID(ctx, operations.PutUsersUserIDRequest{
        EditUserInput: shared.EditUserInput{
            Description: testpango.String("Assimilated zero defect moratorium"),
            FullName: "Ms. Kathy Stanton",
            ID: testpango.String("79a1c67b-61ff-41da-8536-101e994c1527"),
            Role: shared.RoleAccountAdmin.ToPointer(),
            Status: shared.EditUserStatusDisabled,
        },
        UserID: "04672250-5e17-4edb-8b8f-fdfce4dbd0b1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PutUsersUserIDRequest](../../models/operations/putusersuseridrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PutUsersUserIDResponse](../../models/operations/putusersuseridresponse.md), error**

