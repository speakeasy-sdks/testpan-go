# User


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AccountID`                                                 | **string*                                                   | :heavy_minus_sign:                                          | The Secure Application account ID to which the user belongs |
| `Description`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Email`                                                     | **string*                                                   | :heavy_minus_sign:                                          | The email of the user.                                      |
| `FullName`                                                  | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `ID`                                                        | **string*                                                   | :heavy_minus_sign:                                          | ID of the user as created by Secure Application management. |
| `LastLogin`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | N/A                                                         |
| `Role`                                                      | [*Role](../../models/shared/role.md)                        | :heavy_minus_sign:                                          | The role of the user                                        |
| `ShouldDisplayEula`                                         | **bool*                                                     | :heavy_minus_sign:                                          | N/A                                                         |
| `ShouldDisplayProductTour`                                  | **bool*                                                     | :heavy_minus_sign:                                          | N/A                                                         |
| `Status`                                                    | [UserStatus](../../models/shared/userstatus.md)             | :heavy_check_mark:                                          | N/A                                                         |