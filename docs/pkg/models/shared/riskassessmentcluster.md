# RiskAssessmentCluster

Single cluster risk assessment


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ClusterID`                                                                                              | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ClusterName`                                                                                            | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `DockerfileScanResultsSummary`                                                                           | [*shared.DockerfileScanResultsSummary](../../../pkg/models/shared/dockerfilescanresultssummary.md)       | :heavy_minus_sign:                                                                                       | dockerfile scan results summary by severity                                                              |
| `HasFailed`                                                                                              | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ID`                                                                                                     | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `ScanConfig`                                                                                             | [*shared.RiskAssessmentClusterScanConfig](../../../pkg/models/shared/riskassessmentclusterscanconfig.md) | :heavy_minus_sign:                                                                                       | Single cluster risk assessment scan config                                                               |
| `Scanned`                                                                                                | **int64*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `Status`                                                                                                 | [*shared.RiskAssessmentScanStatus](../../../pkg/models/shared/riskassessmentscanstatus.md)               | :heavy_minus_sign:                                                                                       | Status of a risk assessment scan                                                                         |
| `Time`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `Total`                                                                                                  | **int64*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `VulnerabilitiesSummary`                                                                                 | [*shared.VulnerabilitiesSummary](../../../pkg/models/shared/vulnerabilitiessummary.md)                   | :heavy_minus_sign:                                                                                       | Vulnerabilities summary by severity                                                                      |