# NetworkConnection


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `DestinationID`                                       | **string*                                             | :heavy_minus_sign:                                    | Destination App id                                    |
| `DestinationPortNumber`                               | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |
| `ID`                                                  | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `NumberOfConnections`                                 | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |
| `Protocol`                                            | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `SourceID`                                            | **string*                                             | :heavy_minus_sign:                                    | Source App id                                         |
| `StartTime`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | N/A                                                   |
| `Violation`                                           | [*shared.Violation](../../models/shared/violation.md) | :heavy_minus_sign:                                    | N/A                                                   |