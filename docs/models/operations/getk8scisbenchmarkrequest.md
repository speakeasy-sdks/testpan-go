# GetK8sCISBenchmarkRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `ClusterIds`                                     | []*string*                                       | :heavy_minus_sign:                               | cluster ids to filter                            |
| `MaxResults`                                     | **float64*                                       | :heavy_minus_sign:                               | The number of entries to return (pagination)     |
| `NoPagination`                                   | **bool*                                          | :heavy_minus_sign:                               | When true, the pagination params will be ignored |
| `Offset`                                         | **float64*                                       | :heavy_minus_sign:                               | Return entries from this offset (pagination)     |