# OpenAPISpecTag


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `HasDiffs`                                                                 | **bool*                                                                    | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Methods`                                                                  | [][SpecMethod](../../models/shared/specmethod.md)                          | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Name`                                                                     | **string*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Severity`                                                                 | [*APISecurityRiskSeverity](../../models/shared/apisecurityriskseverity.md) | :heavy_minus_sign:                                                         | An `enum`eration.                                                          |
| `VulnerabilitiesSummary`                                                   | [*VulnerabilitiesSummary](../../models/shared/vulnerabilitiessummary.md)   | :heavy_minus_sign:                                                         | Vulnerabilities summary by severity                                        |