# CdConnectionRule

A rule that states what apps are allowed to communicate with each other.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Action`                                                                    | [*shared.ConnectionRuleAction](../../models/shared/connectionruleaction.md) | :heavy_minus_sign:                                                          | ENCRYPT is not allowed in default rule                                      |
| `Destination`                                                               | [*shared.ConnectionRulePart](../../models/shared/connectionrulepart.md)     | :heavy_minus_sign:                                                          | N/A                                                                         |
| `GroupName`                                                                 | **string*                                                                   | :heavy_minus_sign:                                                          | N/A                                                                         |
| `ID`                                                                        | **string*                                                                   | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Name`                                                                      | **string*                                                                   | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Source`                                                                    | [*shared.ConnectionRulePart](../../models/shared/connectionrulepart.md)     | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Status`                                                                    | [*shared.Status](../../models/shared/status.md)                             | :heavy_minus_sign:                                                          | N/A                                                                         |