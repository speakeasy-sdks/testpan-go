# ContainerSecurityContext


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `AllowPrivilegeEscalation` | **bool*                    | :heavy_minus_sign:         | N/A                        |
| `CapabilitiesAdd`          | []*string*                 | :heavy_minus_sign:         | N/A                        |
| `CapabilitiesDrop`         | []*string*                 | :heavy_minus_sign:         | N/A                        |
| `Name`                     | **string*                  | :heavy_minus_sign:         | N/A                        |
| `Privileged`               | **bool*                    | :heavy_minus_sign:         | N/A                        |
| `ProcMount`                | **string*                  | :heavy_minus_sign:         | N/A                        |
| `ReadOnlyRootFilesystem`   | **bool*                    | :heavy_minus_sign:         | N/A                        |
| `RunAsGroup`               | **int64*                   | :heavy_minus_sign:         | N/A                        |
| `RunAsNonRoot`             | **bool*                    | :heavy_minus_sign:         | N/A                        |
| `RunAsUser`                | **int64*                   | :heavy_minus_sign:         | N/A                        |