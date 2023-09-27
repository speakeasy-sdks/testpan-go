# RiskAssessmentClusterScanConfig

Single cluster risk assessment scan config


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `MaxParallelism`                                                       | *int64*                                                                | :heavy_check_mark:                                                     | N/A                                                                    |
| `MinimumSeverity`                                                      | [VulnerabilitySeverity](../../models/shared/vulnerabilityseverity.md)  | :heavy_check_mark:                                                     | N/A                                                                    |
| `Namespaces`                                                           | []*string*                                                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `PeriodicJobExpression`                                                | [*PeriodicJobExpression](../../models/shared/periodicjobexpression.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `RunDockerfileScan`                                                    | **bool*                                                                | :heavy_minus_sign:                                                     | N/A                                                                    |