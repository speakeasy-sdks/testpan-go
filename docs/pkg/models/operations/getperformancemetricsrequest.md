# GetPerformanceMetricsRequest


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `EndTime`                                 | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | End date of the query                     |
| `Protocol`                                | *string*                                  | :heavy_check_mark:                        | protocol                                  |
| `SourceNamespace`                         | *string*                                  | :heavy_check_mark:                        | namespace id                              |
| `SourcePodTemplate`                       | *string*                                  | :heavy_check_mark:                        | pod template id                           |
| `StartTime`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Start date of the query                   |
| `TargetNamespace`                         | *string*                                  | :heavy_check_mark:                        | namespace id                              |
| `TargetPodTemplate`                       | *string*                                  | :heavy_check_mark:                        | pod template id                           |