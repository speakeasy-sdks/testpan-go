# GetCdRequest


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `EndTime`                                                | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | End date of the query                                    |
| `MaxResults`                                             | **float64*                                               | :heavy_minus_sign:                                       | The number of entries to return (pagination)             |
| `Offset`                                                 | **float64*                                               | :heavy_minus_sign:                                       | Return entries from this offset (pagination)             |
| `ResourceName`                                           | **string*                                                | :heavy_minus_sign:                                       | Resource name                                            |
| `SortDir`                                                | [*GetCdSortDir](../../models/operations/getcdsortdir.md) | :heavy_minus_sign:                                       | sorting direction                                        |
| `SortKey`                                                | [*GetCdSortKey](../../models/operations/getcdsortkey.md) | :heavy_minus_sign:                                       | sort key                                                 |
| `StartTime`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Start date of the query                                  |