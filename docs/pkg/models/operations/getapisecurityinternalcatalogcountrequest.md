# GetAPISecurityInternalCatalogCountRequest


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `IncludeServiceWithNoSpec`                                                                 | **bool*                                                                                    | :heavy_minus_sign:                                                                         | When false, only services with specs wikk be returned                                      |
| `Name`                                                                                     | **string*                                                                                  | :heavy_minus_sign:                                                                         | the Api Catalog name filter                                                                |
| `NamespacesFilter`                                                                         | **string*                                                                                  | :heavy_minus_sign:                                                                         | namespace filter. a base 64 representation of a list of NamespacesFilter definition object |
| `UpdatedAfter`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | Only Apis updated since this date                                                          |