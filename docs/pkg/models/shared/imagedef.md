# ImageDef

Authorized image hash


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ImageHash`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | Valid hash for the image. * will authorize image name without validating hash |
| `ImageName`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | N/A                                                                           |
| `ImageTags`                                                                   | []*string*                                                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `TimeAdded`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |