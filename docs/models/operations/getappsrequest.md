# GetAppsRequest


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `DownloadAsXlsx`                                                            | **bool*                                                                     | :heavy_minus_sign:                                                          | When true, the API will return an xlsx file, and pagination will be ignored |
| `Name`                                                                      | **string*                                                                   | :heavy_minus_sign:                                                          | Filter Apps by name                                                         |
| `NoPagination`                                                              | **bool*                                                                     | :heavy_minus_sign:                                                          | When true, the pagination params will be ignored                            |
| `SortDir`                                                                   | [*GetAppsSortDir](../../models/operations/getappssortdir.md)                | :heavy_minus_sign:                                                          | sorting direction                                                           |
| `SortKey`                                                                   | [*GetAppsSortKey](../../models/operations/getappssortkey.md)                | :heavy_minus_sign:                                                          | App sort key                                                                |
| `Type`                                                                      | []*string*                                                                  | :heavy_minus_sign:                                                          | Filter Apps by type                                                         |