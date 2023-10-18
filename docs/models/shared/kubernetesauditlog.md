# KubernetesAuditLog

Single kubernetes audit log


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [*KubernetesAuditLogAction](../../models/shared/kubernetesauditlogaction.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ClusterID`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ClusterName`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `EnvironmentID`                                                                  | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `EnvironmentName`                                                                | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `FirstSeen`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `LastSeen`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Namespace`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ResourceGroup`                                                                  | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ResourceKind`                                                                   | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `ResourceName`                                                                   | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Total`                                                                          | **int64*                                                                         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `User`                                                                           | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `UserGroups`                                                                     | []*string*                                                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `UserNamespace`                                                                  | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `UserType`                                                                       | [*KubernetesAuditLogUserType](../../models/shared/kubernetesauditlogusertype.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Violation`                                                                      | [*KubernetesAPIViolation](../../models/shared/kubernetesapiviolation.md)         | :heavy_minus_sign:                                                               | N/A                                                                              |