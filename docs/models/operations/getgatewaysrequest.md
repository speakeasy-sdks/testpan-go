# GetGatewaysRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `MaxResults`                                                         | **float64*                                                           | :heavy_minus_sign:                                                   | The number of entries to return (pagination)                         |
| `Name`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Filter gateways by name                                              |
| `NoPagination`                                                       | **bool*                                                              | :heavy_minus_sign:                                                   | When true, the pagination params will be ignored                     |
| `Offset`                                                             | **float64*                                                           | :heavy_minus_sign:                                                   | Return entries from this offset (pagination)                         |
| `SortDir`                                                            | [*GetGatewaysSortDir](../../models/operations/getgatewayssortdir.md) | :heavy_minus_sign:                                                   | sorting direction                                                    |