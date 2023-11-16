# FuzzingMethod


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Findings`                                                                               | [*shared.FuzzingScoreFindings](../../../pkg/models/shared/fuzzingscorefindings.md)       | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Method`                                                                                 | [*shared.HTTPMethod](../../../pkg/models/shared/httpmethod.md)                           | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Path`                                                                                   | **string*                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `RequestCount`                                                                           | **int64*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Severity`                                                                               | [*shared.APISecurityRiskSeverity](../../../pkg/models/shared/apisecurityriskseverity.md) | :heavy_minus_sign:                                                                       | An `enum`eration.                                                                        |