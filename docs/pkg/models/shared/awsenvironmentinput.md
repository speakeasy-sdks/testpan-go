# AwsEnvironmentInput


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ID`                                                                            | **string*                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Tags`                                                                          | [][shared.Tag](../../../pkg/models/shared/tag.md)                               | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Vpc`                                                                           | [shared.VPCDescriptionInput](../../../pkg/models/shared/vpcdescriptioninput.md) | :heavy_check_mark:                                                              | Like VPC but also includes the name                                             |