# CDPipelineResult

Pipeline result for a scanned CD resource


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `DeploymentName`                                                                              | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `DeploymentSource`                                                                            | [*shared.DeploymentSource](../../../pkg/models/shared/deploymentsource.md)                    | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `DeploymentVersion`                                                                           | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `ID`                                                                                          | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `PolicyID`                                                                                    | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `PolicyName`                                                                                  | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `Result`                                                                                      | [*shared.CDResult](../../../pkg/models/shared/cdresult.md)                                    | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `SecurityFinding`                                                                             | [][shared.CDPipelineSecurityFinding](../../../pkg/models/shared/cdpipelinesecurityfinding.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `Status`                                                                                      | [*shared.CDPipelineResultStatus](../../../pkg/models/shared/cdpipelineresultstatus.md)        | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `Time`                                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                                    | :heavy_minus_sign:                                                                            | N/A                                                                                           |