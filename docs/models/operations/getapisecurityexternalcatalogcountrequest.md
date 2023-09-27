# GetAPISecurityExternalCatalogCountRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `IncludeServiceWithNoSpec`                            | **bool*                                               | :heavy_minus_sign:                                    | When false, only services with specs wikk be returned |
| `Name`                                                | **string*                                             | :heavy_minus_sign:                                    | the Api Catalog name filter                           |
| `UpdatedAfter`                                        | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | Only Apis updated since this date                     |