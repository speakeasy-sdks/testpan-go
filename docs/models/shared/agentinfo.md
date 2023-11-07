# AgentInfo


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Enabled`                                                           | **bool*                                                             | :heavy_minus_sign:                                                  | N/A                                                                 |
| `LastActive`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | The last time that the agent sent telemetries                       |
| `Status`                                                            | [*shared.ControllerStatus](../../models/shared/controllerstatus.md) | :heavy_minus_sign:                                                  | The current controller state.                                       |
| `Version`                                                           | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |