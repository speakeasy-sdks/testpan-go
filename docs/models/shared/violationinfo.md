# ViolationInfo

If the the App is running on an environment on which it is not allowed to run, this object contains the rule it violated.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Action`                                                             | [*AppRuleAction](../../models/shared/appruleaction.md)               | :heavy_minus_sign:                                                   | App rule action                                                      |
| `Comment`                                                            | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DefaultRule`                                                        | [*DefaultRule](../../models/shared/defaultrule.md)                   | :heavy_minus_sign:                                                   | N/A                                                                  |
| `LastViolation`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `MutateRule`                                                         | [*UserRule](../../models/shared/userrule.md)                         | :heavy_minus_sign:                                                   | N/A                                                                  |
| `PspViolationReasons`                                                | []*string*                                                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UnidentifiedPodReasons`                                             | []*string*                                                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UnidentifiedPodsRule`                                               | [*UnidentifiedPodsRule](../../models/shared/unidentifiedpodsrule.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UserRule`                                                           | [*UserRule](../../models/shared/userrule.md)                         | :heavy_minus_sign:                                                   | N/A                                                                  |
| `ViolationReasons`                                                   | [][ViolationReason](../../models/shared/violationreason.md)          | :heavy_minus_sign:                                                   | N/A                                                                  |