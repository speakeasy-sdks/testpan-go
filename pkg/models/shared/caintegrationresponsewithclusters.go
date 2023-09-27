// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CaIntegrationResponseWithClusters struct {
	Certificate     string               `json:"certificate"`
	Clusters        []CAClusterIDAndName `json:"clusters,omitempty"`
	ID              *string              `json:"id,omitempty"`
	IssuerName      string               `json:"issuerName"`
	IssuerNamespace *string              `json:"issuerNamespace,omitempty"`
	Name            string               `json:"name"`
}

func (o *CaIntegrationResponseWithClusters) GetCertificate() string {
	if o == nil {
		return ""
	}
	return o.Certificate
}

func (o *CaIntegrationResponseWithClusters) GetClusters() []CAClusterIDAndName {
	if o == nil {
		return nil
	}
	return o.Clusters
}

func (o *CaIntegrationResponseWithClusters) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CaIntegrationResponseWithClusters) GetIssuerName() string {
	if o == nil {
		return ""
	}
	return o.IssuerName
}

func (o *CaIntegrationResponseWithClusters) GetIssuerNamespace() *string {
	if o == nil {
		return nil
	}
	return o.IssuerNamespace
}

func (o *CaIntegrationResponseWithClusters) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
