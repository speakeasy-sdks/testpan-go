# RiskAssessmentClusterScanConfig

Single cluster risk assessment scan config


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `MaxParallelism`                                                                     | *int64*                                                                              | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `MinimumSeverity`                                                                    | [shared.VulnerabilitySeverity](../../../pkg/models/shared/vulnerabilityseverity.md)  | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `Namespaces`                                                                         | []*string*                                                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `PeriodicJobExpression`                                                              | [*shared.PeriodicJobExpression](../../../pkg/models/shared/periodicjobexpression.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `RunDockerfileScan`                                                                  | **bool*                                                                              | :heavy_minus_sign:                                                                   | N/A                                                                                  |