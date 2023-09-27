# GetTokensRequest


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `MaxResults`                                                     | **float64*                                                       | :heavy_minus_sign:                                               | The number of entries to return (pagination)                     |
| `NoPagination`                                                   | **bool*                                                          | :heavy_minus_sign:                                               | When true, the pagination params will be ignored                 |
| `Offset`                                                         | **float64*                                                       | :heavy_minus_sign:                                               | Return entries from this offset (pagination)                     |
| `SortDir`                                                        | [*GetTokensSortDir](../../models/operations/gettokenssortdir.md) | :heavy_minus_sign:                                               | sorting direction                                                |
| `SortKey`                                                        | **string*                                                        | :heavy_minus_sign:                                               | the token sort key                                               |
| `TokenName`                                                      | **string*                                                        | :heavy_minus_sign:                                               | Defined token name                                               |