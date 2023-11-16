# OpenAPISpecTag


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `HasDiffs`                                                                               | **bool*                                                                                  | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Methods`                                                                                | [][shared.SpecMethod](../../../pkg/models/shared/specmethod.md)                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Name`                                                                                   | **string*                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `Severity`                                                                               | [*shared.APISecurityRiskSeverity](../../../pkg/models/shared/apisecurityriskseverity.md) | :heavy_minus_sign:                                                                       | An `enum`eration.                                                                        |
| `VulnerabilitiesSummary`                                                                 | [*shared.VulnerabilitiesSummary](../../../pkg/models/shared/vulnerabilitiessummary.md)   | :heavy_minus_sign:                                                                       | Vulnerabilities summary by severity                                                      |