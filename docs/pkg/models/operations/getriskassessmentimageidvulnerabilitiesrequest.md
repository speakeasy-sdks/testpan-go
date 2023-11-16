# GetRiskAssessmentImageIDVulnerabilitiesRequest


## Fields

| Field                                                                                                                                                              | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ImageID`                                                                                                                                                          | *string*                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                 | The id of the risk assessment image                                                                                                                                |
| `MaxResults`                                                                                                                                                       | **float64*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                 | The number of entries to return (pagination)                                                                                                                       |
| `Offset`                                                                                                                                                           | **float64*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                 | Return entries from this offset (pagination)                                                                                                                       |
| `SortDir`                                                                                                                                                          | [*operations.GetRiskAssessmentImageIDVulnerabilitiesQueryParamSortDir](../../../pkg/models/operations/getriskassessmentimageidvulnerabilitiesqueryparamsortdir.md) | :heavy_minus_sign:                                                                                                                                                 | sorting direction                                                                                                                                                  |
| `SortKey`                                                                                                                                                          | [operations.GetRiskAssessmentImageIDVulnerabilitiesQueryParamSortKey](../../../pkg/models/operations/getriskassessmentimageidvulnerabilitiesqueryparamsortkey.md)  | :heavy_check_mark:                                                                                                                                                 | risk assessment image sort key.                                                                                                                                    |