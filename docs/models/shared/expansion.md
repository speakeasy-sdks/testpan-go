# Expansion

represent expansion object used in put method


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AccountName`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ClusterID`                                                 | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `ClusterName`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ControllerEnabled`                                         | **bool*                                                     | :heavy_minus_sign:                                          | N/A                                                         |
| `ControllerID`                                              | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ControllerIsUpdateEnabled`                                 | **bool*                                                     | :heavy_minus_sign:                                          | N/A                                                         |
| `ControllerLastActive`                                      | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | The last time that the agent sent telemetries               |
| `ControllerStatus`                                          | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ControllerVersion`                                         | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ID`                                                        | **string*                                                   | :heavy_minus_sign:                                          | unique Id                                                   |
| `InstanceID`                                                | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Labels`                                                    | [][Label](../../models/shared/label.md)                     | :heavy_minus_sign:                                          | N/A                                                         |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `NamespaceID`                                               | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `ShouldSendMetrics`                                         | **bool*                                                     | :heavy_minus_sign:                                          | N/A                                                         |
| `WorkloadAddresses`                                         | [][WorkloadAddress](../../models/shared/workloadaddress.md) | :heavy_check_mark:                                          | N/A                                                         |