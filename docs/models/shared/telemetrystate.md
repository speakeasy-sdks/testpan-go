# TelemetryState

Status of a telemetry entry


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `LastSeen`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `StartTime`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Status`                                                             | [*TelemetryStateStatus](../../models/shared/telemetrystatestatus.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `StatusReason`                                                       | **string*                                                            | :heavy_minus_sign:                                                   | will be populate only when status is unhealthy                       |