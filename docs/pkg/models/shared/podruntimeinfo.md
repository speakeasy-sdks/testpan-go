# PodRuntimeInfo

runtime info of the pod (if is a pod)


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Containers`                                                  | [][shared.Container](../../../pkg/models/shared/container.md) | :heavy_minus_sign:                                            | runtime pod containers                                        |
| `InitContainers`                                              | [][shared.Container](../../../pkg/models/shared/container.md) | :heavy_minus_sign:                                            | runtime pod init containers                                   |
| `Labels`                                                      | [][shared.Label](../../../pkg/models/shared/label.md)         | :heavy_minus_sign:                                            | runtime pod labels                                            |