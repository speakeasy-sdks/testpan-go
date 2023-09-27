# APISecTopRiskyAPI


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ID`                                                                     | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Name`                                                                   | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Type`                                                                   | [APIServiceType](../../models/shared/apiservicetype.md)                  | :heavy_check_mark:                                                       | An `enum`eration.                                                        |
| `VulnerabilitiesSummary`                                                 | [*VulnerabilitiesSummary](../../models/shared/vulnerabilitiessummary.md) | :heavy_minus_sign:                                                       | Vulnerabilities summary by severity                                      |