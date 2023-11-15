# EditUser


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Description`                                                         | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `FullName`                                                            | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `ID`                                                                  | **string*                                                             | :heavy_minus_sign:                                                    | ID of the user as created by Secure Application management.           |
| `Role`                                                                | [*shared.Role](../../../pkg/models/shared/role.md)                    | :heavy_minus_sign:                                                    | The role of the user                                                  |
| `Status`                                                              | [shared.EditUserStatus](../../../pkg/models/shared/edituserstatus.md) | :heavy_check_mark:                                                    | N/A                                                                   |