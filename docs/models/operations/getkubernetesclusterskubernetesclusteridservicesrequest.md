# GetKubernetesClustersKubernetesClusterIDServicesRequest


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `KubernetesClusterID`                                                                  | *string*                                                                               | :heavy_check_mark:                                                                     | Secure Application Kubernetes cluster ID                                               |
| `ShowIstioOnly`                                                                        | **bool*                                                                                | :heavy_minus_sign:                                                                     | if true, return only services deployed on namespace with label istio-injection=enabled |