# IstioInstallationParameters

istio related information


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `IsIstioAlreadyInstalled`                                                                                           | **bool*                                                                                                             | :heavy_minus_sign:                                                                                                  | indicates whether Istio is already installed on this cluster (which means Secure Application should not install it) |
| `IstioVersion`                                                                                                      | **string*                                                                                                           | :heavy_minus_sign:                                                                                                  | when istio already installed, choose the version from supported istio versions list: /istio/supportedVersions       |