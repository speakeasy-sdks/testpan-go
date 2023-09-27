// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteDependencyElementEnvironments struct {
	ClusterEventRules []DeleteDependencyEnvironmentRuleElement `json:"clusterEventRules,omitempty"`
	ConnectionRules   []DeleteDependencyEnvironmentRuleElement `json:"connectionRules,omitempty"`
	DeploymentRules   []DeleteDependencyEnvironmentRuleElement `json:"deploymentRules,omitempty"`
	Environments      []DeleteDependencyElement                `json:"environments,omitempty"`
}

func (o *DeleteDependencyElementEnvironments) GetClusterEventRules() []DeleteDependencyEnvironmentRuleElement {
	if o == nil {
		return nil
	}
	return o.ClusterEventRules
}

func (o *DeleteDependencyElementEnvironments) GetConnectionRules() []DeleteDependencyEnvironmentRuleElement {
	if o == nil {
		return nil
	}
	return o.ConnectionRules
}

func (o *DeleteDependencyElementEnvironments) GetDeploymentRules() []DeleteDependencyEnvironmentRuleElement {
	if o == nil {
		return nil
	}
	return o.DeploymentRules
}

func (o *DeleteDependencyElementEnvironments) GetEnvironments() []DeleteDependencyElement {
	if o == nil {
		return nil
	}
	return o.Environments
}
