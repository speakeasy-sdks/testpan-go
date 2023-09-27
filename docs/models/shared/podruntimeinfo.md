# PodRuntimeInfo

runtime info of the pod (if is a pod)


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Containers`                                    | [][Container](../../models/shared/container.md) | :heavy_minus_sign:                              | runtime pod containers                          |
| `InitContainers`                                | [][Container](../../models/shared/container.md) | :heavy_minus_sign:                              | runtime pod init containers                     |
| `Labels`                                        | [][Label](../../models/shared/label.md)         | :heavy_minus_sign:                              | runtime pod labels                              |