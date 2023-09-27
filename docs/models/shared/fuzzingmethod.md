# FuzzingMethod


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Findings`                                                                 | [*FuzzingScoreFindings](../../models/shared/fuzzingscorefindings.md)       | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Method`                                                                   | [*HTTPMethod](../../models/shared/httpmethod.md)                           | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Path`                                                                     | **string*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        |
| `RequestCount`                                                             | **int64*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Severity`                                                                 | [*APISecurityRiskSeverity](../../models/shared/apisecurityriskseverity.md) | :heavy_minus_sign:                                                         | An `enum`eration.                                                          |