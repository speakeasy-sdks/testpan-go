# ExpansionInput

represent expansion object


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ClusterID`                                                               | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `ControllerLastActive`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | The last time that the agent sent telemetries                             |
| `Labels`                                                                  | [][shared.Label](../../../pkg/models/shared/label.md)                     | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Name`                                                                    | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `NamespaceID`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `ShouldSendMetrics`                                                       | **bool*                                                                   | :heavy_minus_sign:                                                        | N/A                                                                       |
| `WorkloadAddresses`                                                       | [][shared.WorkloadAddress](../../../pkg/models/shared/workloadaddress.md) | :heavy_check_mark:                                                        | N/A                                                                       |