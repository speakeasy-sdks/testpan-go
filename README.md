# github.com/speakeasy-sdks/testpan-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/testpan-go.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
# SDK Installation

```bash
go get github.com/speakeasy-sdks/testpan-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
)

func main() {
    s := testpango.New(
        testpango.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APISecurityPolicies.DeleteAPISecurityPolicyPolicyID(ctx, operations.DeleteAPISecurityPolicyPolicyIDRequest{
        PolicyID: "a05dfc2d-df7c-4c78-8a1b-a928fc816742",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
# Available Resources and Operations


## [APISecurityPolicies](docs/sdks/apisecuritypolicies/README.md)

* [DeleteAPISecurityPolicyPolicyID](docs/sdks/apisecuritypolicies/README.md#deleteapisecuritypolicypolicyid) - Delete api security policy
* [GetAPISecurityPolicy](docs/sdks/apisecuritypolicies/README.md#getapisecuritypolicy) - Get a list of API security policies
* [GetAPISecurityPolicyPolicyIDDeleteDependencies](docs/sdks/apisecuritypolicies/README.md#getapisecuritypolicypolicyiddeletedependencies) - get dependencies which need to be handled in order to delete specified api security service
* [PostAPISecurityPolicy](docs/sdks/apisecuritypolicies/README.md#postapisecuritypolicy) - Add new API security policy
* [PutAPISecurityPolicyPolicyID](docs/sdks/apisecuritypolicies/README.md#putapisecuritypolicypolicyid) - Edit Api security policy.

## [Cd](docs/sdks/cd/README.md)

* [DeleteCdRuleIDConnectionsRule](docs/sdks/cd/README.md#deletecdruleidconnectionsrule) - delete a cd connection rule.
* [DeleteCdRuleIDServerlessRule](docs/sdks/cd/README.md#deletecdruleidserverlessrule) - delete a cd serverless rule.
* [GetCd](docs/sdks/cd/README.md#getcd) - Get all the CD pipelines results
* [GetCdResourceID](docs/sdks/cd/README.md#getcdresourceid) - Get A single CD pipeline results
* [GetCdRuleIDConnectionsRule](docs/sdks/cd/README.md#getcdruleidconnectionsrule) - get a cd connection rule.
* [GetCdRuleIDServerlessRule](docs/sdks/cd/README.md#getcdruleidserverlessrule) - get a cd serverless rule.
* [PostCdConnectionsRule](docs/sdks/cd/README.md#postcdconnectionsrule) - Adds cd connection rule.
* [PostCdServerlessRule](docs/sdks/cd/README.md#postcdserverlessrule) - Adds cd serverless rule.
* [PutCdRuleIDConnectionsRule](docs/sdks/cd/README.md#putcdruleidconnectionsrule) - update a cd connection rule.
* [PutCdRuleIDServerlessRule](docs/sdks/cd/README.md#putcdruleidserverlessrule) - update a cd serverless rule.

## [CICDPolicies](docs/sdks/cicdpolicies/README.md)

* [DeleteCdPolicyPolicyID](docs/sdks/cicdpolicies/README.md#deletecdpolicypolicyid) - Delete CD policy
* [DeleteCiPolicyPolicyID](docs/sdks/cicdpolicies/README.md#deletecipolicypolicyid) - Delete CI policy
* [GetCdPolicy](docs/sdks/cicdpolicies/README.md#getcdpolicy) - Get the current CD policy
* [GetCiPolicy](docs/sdks/cicdpolicies/README.md#getcipolicy) - Get the current CI policy
* [PostCdPolicy](docs/sdks/cicdpolicies/README.md#postcdpolicy) - Set the current CD policy. At least one CdPolicyElement should be present
* [PostCiPolicy](docs/sdks/cicdpolicies/README.md#postcipolicy) - Set the current CI policy
* [PutCdPolicyPolicyID](docs/sdks/cicdpolicies/README.md#putcdpolicypolicyid) - Edit CD policy. At least one CdPolicyElement should be present
* [PutCiPolicyPolicyID](docs/sdks/cicdpolicies/README.md#putcipolicypolicyid) - Edit CI policy

## [Advisor](docs/sdks/advisor/README.md)

* [GetAdvisorClusterEventRules](docs/sdks/advisor/README.md#getadvisorclustereventrules) - Returns a list of suggested cluster event rules
* [GetAdvisorConnectionRules](docs/sdks/advisor/README.md#getadvisorconnectionrules) - Returns a list of suggested connection rules
* [GetAdvisorEnvironment](docs/sdks/advisor/README.md#getadvisorenvironment) - Returns a list of suggested kubernetes environments
* [GetAdvisorEnvironmentRules](docs/sdks/advisor/README.md#getadvisorenvironmentrules) - Returns a list of suggested environment rules
* [GetAdvisorPodSecurityPolicy](docs/sdks/advisor/README.md#getadvisorpodsecuritypolicy) - Returns a list of suggested kubernetes Pod Security Standards
* [GetAdvisorQueueAdvisorType](docs/sdks/advisor/README.md#getadvisorqueueadvisortype) - Get status for policy advisor background job
* [PostAdvisorRun](docs/sdks/advisor/README.md#postadvisorrun) - Runs the policy advisor

## [AgentManagement](docs/sdks/agentmanagement/README.md)

* [GetAgents](docs/sdks/agentmanagement/README.md#getagents) - List all installed agents
* [PostAgentsAgentIDUpdate](docs/sdks/agentmanagement/README.md#postagentsagentidupdate) - Update the agent with the given id to the latest agent version
* [PostAgentsAgentIDUpdateState](docs/sdks/agentmanagement/README.md#postagentsagentidupdatestate) - Update the status of an agent with the given id

## [API](docs/sdks/api/README.md)

* [GetAPI](docs/sdks/api/README.md#getapi) - Get Secure Application API as a Swagger file

## [APISecurity](docs/sdks/apisecurity/README.md)

* [DeleteAPISecurityAPICatalogID](docs/sdks/apisecurity/README.md#deleteapisecurityapicatalogid) - Delete an API
* [DeleteAPISecurityInternalCatalogCatalogIDBflaDetection](docs/sdks/apisecurity/README.md#deleteapisecurityinternalcatalogcatalogidbfladetection) - stop bfla detection phase
* [DeleteAPISecurityInternalCatalogCatalogIDBflaLearning](docs/sdks/apisecurity/README.md#deleteapisecurityinternalcatalogcatalogidbflalearning) - stop bfla learning phase
* [DeleteAPISecurityOpenAPISpecsCatalogID](docs/sdks/apisecurity/README.md#deleteapisecurityopenapispecscatalogid) - delete open api spec include all of it findings and scores
* [DeleteGatewaysGatewayID](docs/sdks/apisecurity/README.md#deletegatewaysgatewayid) - Delete gateway
* [GetAPISecurityExternalCatalog](docs/sdks/apisecurity/README.md#getapisecurityexternalcatalog) - Get a list of APIs and their compliance
* [GetAPISecurityExternalCatalogCount](docs/sdks/apisecurity/README.md#getapisecurityexternalcatalogcount) - Get the number of existing 3rd party APIs in the inventory
* [GetAPISecurityExternalCatalogCatalogID](docs/sdks/apisecurity/README.md#getapisecurityexternalcatalogcatalogid) - Get information about a specific API
* [GetAPISecurityInternalCatalog](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalog) - Get a list of APIs and their compliance
* [GetAPISecurityInternalCatalogCount](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcount) - Get the number of existing 3rd party APIs in the inventory
* [GetAPISecurityInternalCatalogCatalogID](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcatalogid) - Get information about a specific API
* [GetAPISecurityInternalCatalogCatalogIDBfla](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcatalogidbfla) - Get bfla info for given catalogId
* [GetAPISecurityInternalCatalogCatalogIDFuzzingStatus](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcatalogidfuzzingstatus) - Get status of the last/running fuzzing test
* [GetAPISecurityInternalCatalogCatalogIDFuzzingTests](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcatalogidfuzzingtests) - Get list of fuzzing tests
* [GetAPISecurityInternalCatalogCatalogIDTraceAnalysis](docs/sdks/apisecurity/README.md#getapisecurityinternalcatalogcatalogidtraceanalysis) - Get trace analysis details
* [GetAPISecurityOpenAPISpecsCatalogID](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogid) - Get provided and reconstructed open api specs for specific API
* [GetAPISecurityOpenAPISpecsCatalogIDDiffDetectionStatus](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogiddiffdetectionstatus) - Get provided and reconstructed open api specs for specific API
* [GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatus](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogidgetopenapispecscorestatus) - Get open api spec score status
* [GetAPISecurityOpenAPISpecsCatalogIDOpenAPISpecSwaggerJSON](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogidopenapispecswaggerjson) - Get provided spec content as json
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReview](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogidreconstructedspecreview) - Get the suggestions of a spec reconstruction (or previously cached info)
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecStatus](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogidreconstructedspecstatus) - Get the status of a spec reconstruction
* [GetAPISecurityOpenAPISpecsCatalogIDReconstructedSpecJSON](docs/sdks/apisecurity/README.md#getapisecurityopenapispecscatalogidreconstructedspecjson) - Get reconstructed open api spec as json for specific API
* [GetAPISecurityRiskFindings](docs/sdks/apisecurity/README.md#getapisecurityriskfindings) - Get a list of risk findings
* [GetAPISecurityRiskFindingsCategories](docs/sdks/apisecurity/README.md#getapisecurityriskfindingscategories) - Get a list of risk findings categories
* [GetAPISecurityRiskFindingsSources](docs/sdks/apisecurity/README.md#getapisecurityriskfindingssources) - Get a list of risk findings sources
* [GetAPISecurityRiskFindingsRiskFindingID](docs/sdks/apisecurity/README.md#getapisecurityriskfindingsriskfindingid) - Get a specific risk findings
* [GetAPISecurityCatalogIDDeleteDependencies](docs/sdks/apisecurity/README.md#getapisecuritycatalogiddeletedependencies) - get dependencies which need to be handled in order to delete specified api security service
* [GetAPISecurityCatalogIDMethods](docs/sdks/apisecurity/README.md#getapisecuritycatalogidmethods) - Get a list of an API spec methods for a specific API and its spec's tags
* [GetAPISecurityCatalogIDTags](docs/sdks/apisecurity/README.md#getapisecuritycatalogidtags) - Get a list of an API spec tags or a specific API
* [GetDashboardApisecRiskFindings](docs/sdks/apisecurity/README.md#getdashboardapisecriskfindings) - Get API sec risk findings widget
* [GetDashboardApisecRiskFindingsTrend](docs/sdks/apisecurity/README.md#getdashboardapisecriskfindingstrend) - Get API sec risk findings trend graph widget for the last 30 days
* [GetDashboardApisecSpecsAndOperationsDiffs](docs/sdks/apisecurity/README.md#getdashboardapisecspecsandoperationsdiffs) - Get API sec specs and operations diffs widget
* [GetDashboardApisecTopRiskyApis](docs/sdks/apisecurity/README.md#getdashboardapisectopriskyapis) - Get API sec top risky APIs widget
* [GetDashboardApisecTopRiskyFindings](docs/sdks/apisecurity/README.md#getdashboardapisectopriskyfindings) - Get API sec top risky findings widget
* [GetGateways](docs/sdks/apisecurity/README.md#getgateways) - Get gateways
* [GetGatewaysClusters](docs/sdks/apisecurity/README.md#getgatewaysclusters) - Get clusters info
* [GetGatewaysGatewayIDDownloadBundle](docs/sdks/apisecurity/README.md#getgatewaysgatewayiddownloadbundle) - Get a GW installation script
* [PostAPISecurityAPI](docs/sdks/apisecurity/README.md#postapisecurityapi) - Register an API for scoring
* [PostAPISecurityInternalCatalogCatalogIDBflaDetection](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidbfladetection) - Start new bfla detection phase
* [PostAPISecurityInternalCatalogCatalogIDBflaLearning](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidbflalearning) - Start new bfla learning phase
* [PostAPISecurityInternalCatalogCatalogIDBflaReset](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidbflareset) - Reset bfla model
* [PostAPISecurityInternalCatalogCatalogIDResetTraceAnalysis](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidresettraceanalysis) - Reset trace analysis
* [PostAPISecurityInternalCatalogCatalogIDStartFuzzing](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidstartfuzzing) - Start new fuzzing test
* [PostAPISecurityInternalCatalogCatalogIDStartTraceAnalysis](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidstarttraceanalysis) - Start trace analysis
* [PostAPISecurityInternalCatalogCatalogIDStopFuzzing](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidstopfuzzing) - Stop fuzzing test
* [PostAPISecurityInternalCatalogCatalogIDStopTraceAnalysis](docs/sdks/apisecurity/README.md#postapisecurityinternalcatalogcatalogidstoptraceanalysis) - Stop trace analysis
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecAbort](docs/sdks/apisecurity/README.md#postapisecurityopenapispecscatalogidreconstructedspecabort) - abort learning and reconstructing an API via API Clarity
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecLearn](docs/sdks/apisecurity/README.md#postapisecurityopenapispecscatalogidreconstructedspeclearn) - Start learning and reconstructing an API via API Clarity
* [PostAPISecurityOpenAPISpecsCatalogIDReconstructedSpecReviewApprove](docs/sdks/apisecurity/README.md#postapisecurityopenapispecscatalogidreconstructedspecreviewapprove) - Approve reconstructed spec suggestions (only 1 approval per catalogId)
* [PostAPISecurityOpenAPISpecsCatalogIDStartDiffsDetection](docs/sdks/apisecurity/README.md#postapisecurityopenapispecscatalogidstartdiffsdetection) - Start spec diffs detection
* [PostAPISecurityOpenAPISpecsCatalogIDStopDiffsDetection](docs/sdks/apisecurity/README.md#postapisecurityopenapispecscatalogidstopdiffsdetection) - stop spec diffs detection
* [PostGateways](docs/sdks/apisecurity/README.md#postgateways) - Add new gateway
* [PutAPISecurityInternalCatalogCatalogIDBfla](docs/sdks/apisecurity/README.md#putapisecurityinternalcatalogcatalogidbfla) - update BFLA info for this catalogId
* [PutAPISecurityOpenAPISpecsCatalogID](docs/sdks/apisecurity/README.md#putapisecurityopenapispecscatalogid) - Add or edit a spec about a specific API for the account
* [PutGatewaysGatewayID](docs/sdks/apisecurity/README.md#putgatewaysgatewayid) - Edit gateway

## [Apps](docs/sdks/apps/README.md)

* [GetApps](docs/sdks/apps/README.md#getapps) - Returns a list of defined Apps
* [GetAppsAppID](docs/sdks/apps/README.md#getappsappid) - Returns an App by its ID
* [PostApps](docs/sdks/apps/README.md#postapps) - Define a New App
* [PostAppsDelete](docs/sdks/apps/README.md#postappsdelete) - Delete several Apps
* [PutAppsAppID](docs/sdks/apps/README.md#putappsappid) - Edit the existing App

## [AuditLogs](docs/sdks/auditlogs/README.md)

* [GetAuditLogs](docs/sdks/auditlogs/README.md#getauditlogs) - Get audit logs
* [GetAuditLogsActions](docs/sdks/auditlogs/README.md#getauditlogsactions) - Get all the audit logs actions
* [GetAuditLogsKubernetes](docs/sdks/auditlogs/README.md#getauditlogskubernetes) - Get audit logs
* [GetAuditLogsKubernetesActions](docs/sdks/auditlogs/README.md#getauditlogskubernetesactions) - Get all the kubernetes audit logs actions

## [Aws](docs/sdks/aws/README.md)

* [GetAwsAccounts](docs/sdks/aws/README.md#getawsaccounts) - Get a list of AWS accounts
* [GetAwsRoles](docs/sdks/aws/README.md#getawsroles) - Lists AWS role ARNs for the account
* [GetAwsTags](docs/sdks/aws/README.md#getawstags) - Get a list of AWS tag keys
* [GetAwsAwsAccountIDRegions](docs/sdks/aws/README.md#getawsawsaccountidregions) - Get a list of regions for the  AWS account
* [GetAwsAwsAccountIDRegionIDSubnets](docs/sdks/aws/README.md#getawsawsaccountidregionidsubnets) - Get a list of AWS subnets for an AWS account and region
* [GetAwsAwsAccountIDRegionIDVpcs](docs/sdks/aws/README.md#getawsawsaccountidregionidvpcs) - Get a list of VPCs for AWS accounts.
* [PostAwsRoles](docs/sdks/aws/README.md#postawsroles) - Add AWS role to the account
* [PutAwsRolesRoleID](docs/sdks/aws/README.md#putawsrolesroleid) - Change AWS role name

## [Bfla](docs/sdks/bfla/README.md)

* [DeleteAPISecurityInternalCatalogCatalogIDBflaDetection](docs/sdks/bfla/README.md#deleteapisecurityinternalcatalogcatalogidbfladetection) - stop bfla detection phase
* [DeleteAPISecurityInternalCatalogCatalogIDBflaLearning](docs/sdks/bfla/README.md#deleteapisecurityinternalcatalogcatalogidbflalearning) - stop bfla learning phase
* [GetAPISecurityInternalCatalogCatalogIDBfla](docs/sdks/bfla/README.md#getapisecurityinternalcatalogcatalogidbfla) - Get bfla info for given catalogId
* [PostAPISecurityInternalCatalogCatalogIDBflaDetection](docs/sdks/bfla/README.md#postapisecurityinternalcatalogcatalogidbfladetection) - Start new bfla detection phase
* [PostAPISecurityInternalCatalogCatalogIDBflaLearning](docs/sdks/bfla/README.md#postapisecurityinternalcatalogcatalogidbflalearning) - Start new bfla learning phase
* [PostAPISecurityInternalCatalogCatalogIDBflaReset](docs/sdks/bfla/README.md#postapisecurityinternalcatalogcatalogidbflareset) - Reset bfla model
* [PutAPISecurityInternalCatalogCatalogIDBfla](docs/sdks/bfla/README.md#putapisecurityinternalcatalogcatalogidbfla) - update BFLA info for this catalogId

## [Cli](docs/sdks/cli/README.md)

* [GetToolsCliSecurecnDeploymentCli](docs/sdks/cli/README.md#gettoolsclisecurecndeploymentcli) - Get the Secure Application deployment cli

## [ClusterEventsPolicies](docs/sdks/clustereventspolicies/README.md)

* [GetKubernetesAPIPolicy](docs/sdks/clustereventspolicies/README.md#getkubernetesapipolicy) - Get current Kubernetes API policy
* [GetKubernetesAPIPolicyHistory](docs/sdks/clustereventspolicies/README.md#getkubernetesapipolicyhistory) - Get the history of the Kubernetes API policies
* [GetKubernetesAPIPolicyKubernetesResources](docs/sdks/clustereventspolicies/README.md#getkubernetesapipolicykubernetesresources) - Get the Kubernetes resource list
* [GetKubernetesAPIPolicyKubernetesUsers](docs/sdks/clustereventspolicies/README.md#getkubernetesapipolicykubernetesusers) - Get the Kubernetes user list
* [GetKubernetesAPIPolicyRecommendedRules](docs/sdks/clustereventspolicies/README.md#getkubernetesapipolicyrecommendedrules) - Get the recommended Kubernetes API rules
* [PutKubernetesAPIPolicy](docs/sdks/clustereventspolicies/README.md#putkubernetesapipolicy) - set the current Kubernetes API policy

## [ConnectionPolicies](docs/sdks/connectionpolicies/README.md)

* [GetConnectionsPolicy](docs/sdks/connectionpolicies/README.md#getconnectionspolicy) - Get current connection policy
* [GetConnectionsPolicyHistory](docs/sdks/connectionpolicies/README.md#getconnectionspolicyhistory) - Get the history of the connection policies
* [GetConnectionsPolicyKafkaActions](docs/sdks/connectionpolicies/README.md#getconnectionspolicykafkaactions) - Get the a list of kafka actions
* [GetConnectionsPolicyKafkaKubernetesClusterIDBrokers](docs/sdks/connectionpolicies/README.md#getconnectionspolicykafkakubernetesclusteridbrokers) - Get the a list of kafka brokers
* [GetConnectionsPolicyKafkaKubernetesClusterIDTopics](docs/sdks/connectionpolicies/README.md#getconnectionspolicykafkakubernetesclusteridtopics) - Get the a list of kafka topics
* [GetConnectionsPolicySearchOptions](docs/sdks/connectionpolicies/README.md#getconnectionspolicysearchoptions) - Get the current connection policy filter option
* [GetServerlessPolicyHistory](docs/sdks/connectionpolicies/README.md#getserverlesspolicyhistory) - Get the history of the serverless policies
* [PutConnectionsPolicy](docs/sdks/connectionpolicies/README.md#putconnectionspolicy) - Set the current connection policy

## [Dashboard](docs/sdks/dashboard/README.md)

* [GetDashboardApisecRiskFindings](docs/sdks/dashboard/README.md#getdashboardapisecriskfindings) - Get API sec risk findings widget
* [GetDashboardApisecRiskFindingsTrend](docs/sdks/dashboard/README.md#getdashboardapisecriskfindingstrend) - Get API sec risk findings trend graph widget for the last 30 days
* [GetDashboardApisecSpecsAndOperationsDiffs](docs/sdks/dashboard/README.md#getdashboardapisecspecsandoperationsdiffs) - Get API sec specs and operations diffs widget
* [GetDashboardApisecTopRiskyApis](docs/sdks/dashboard/README.md#getdashboardapisectopriskyapis) - Get API sec top risky APIs widget
* [GetDashboardApisecTopRiskyFindings](docs/sdks/dashboard/README.md#getdashboardapisectopriskyfindings) - Get API sec top risky findings widget
* [GetDashboardClusters](docs/sdks/dashboard/README.md#getdashboardclusters) - Get the active clusters
* [GetDashboardConnectionTelemetries](docs/sdks/dashboard/README.md#getdashboardconnectiontelemetries) - Get pod connection dashboard data for all clusters
* [GetDashboardKubernetesAuditLogs](docs/sdks/dashboard/README.md#getdashboardkubernetesauditlogs) - Get kubernetes audit logs dashboard data for all clusters
* [GetDashboardOperationalBar](docs/sdks/dashboard/README.md#getdashboardoperationalbar) - Get the operation data dashboard for the given kubernetesClusterId
* [GetDashboardPermissions](docs/sdks/dashboard/README.md#getdashboardpermissions) - Get permissions dashboard data for the given kubernetesClusterIds
* [GetDashboardPodTelemetries](docs/sdks/dashboard/README.md#getdashboardpodtelemetries) - Get pod telemetries dashboard data for all clusters
* [GetDashboardReportDownload](docs/sdks/dashboard/README.md#getdashboardreportdownload) - Download Secure Application security report
* [GetDashboardReportStatus](docs/sdks/dashboard/README.md#getdashboardreportstatus) - Get Secure Application report security status
* [GetDashboardSecurityContext](docs/sdks/dashboard/README.md#getdashboardsecuritycontext) - Get security context dashboard data for all clusters
* [GetDashboardTopSecurityRisks](docs/sdks/dashboard/README.md#getdashboardtopsecurityrisks) - Get the top risky deployments dashboard data for the given kubernetesClusterIds
* [GetDashboardVulnerabilities](docs/sdks/dashboard/README.md#getdashboardvulnerabilities) - Get vulnerabilities dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDConnectionTelemetries](docs/sdks/dashboard/README.md#getdashboardkubernetesclusteridconnectiontelemetries) - Get connection telemetries dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDKubernetesAuditLogs](docs/sdks/dashboard/README.md#getdashboardkubernetesclusteridkubernetesauditlogs) - Get kubernetes audit logs dashboard data for the given kubernetesClusterId
* [GetDashboardKubernetesClusterIDPodTelemetries](docs/sdks/dashboard/README.md#getdashboardkubernetesclusteridpodtelemetries) - Get pod telemetries dashboard data for the given kubernetesClusterId
* [GetLicensingDashboard](docs/sdks/dashboard/README.md#getlicensingdashboard) - Get licensing dashboard data
* [PostDashboardReportGenerate](docs/sdks/dashboard/README.md#postdashboardreportgenerate) - Generate Secure Application security report

## [Deployers](docs/sdks/deployers/README.md)

* [DeleteDeployersDeployerID](docs/sdks/deployers/README.md#deletedeployersdeployerid) - Delete an deployer
* [GetDeployers](docs/sdks/deployers/README.md#getdeployers) - List all the deployers on the system
* [GetDeployersServiceAccounts](docs/sdks/deployers/README.md#getdeployersserviceaccounts) - List all the service account on the system
* [GetDeployersDeployerIDDeleteDependencies](docs/sdks/deployers/README.md#getdeployersdeployeriddeletedependencies) - get dependencies which need to be handled in order to delete specified deployer
* [PostDeployers](docs/sdks/deployers/README.md#postdeployers) - Create a new deployer
* [PutDeployersDeployerID](docs/sdks/deployers/README.md#putdeployersdeployerid) - Edit deployer definition

## [EnvironmentPolicies](docs/sdks/environmentpolicies/README.md)

* [GetAppsPolicy](docs/sdks/environmentpolicies/README.md#getappspolicy) - Get the current Apps policy
* [GetAppsPolicyHistory](docs/sdks/environmentpolicies/README.md#getappspolicyhistory) - Get the history of Apps policies
* [GetAppsPolicySearchOptions](docs/sdks/environmentpolicies/README.md#getappspolicysearchoptions) - Get the current Apps policy filter option
* [PutAppsPolicy](docs/sdks/environmentpolicies/README.md#putappspolicy) - Set the current Apps policy

## [Envs](docs/sdks/envs/README.md)

* [DeleteEnvironmentsEnvID](docs/sdks/envs/README.md#deleteenvironmentsenvid)
* [GetEnvironments](docs/sdks/envs/README.md#getenvironments) - List all defined Secure Application environments
* [GetEnvironmentsEnvID](docs/sdks/envs/README.md#getenvironmentsenvid) - get a Secure Application environment
* [GetEnvironmentsEnvIDDeleteDependencies](docs/sdks/envs/README.md#getenvironmentsenviddeletedependencies) - get dependencies which need to be handled in order to delete specified environment
* [PostEnvironments](docs/sdks/envs/README.md#postenvironments) - Add a new environment
* [PostEnvironmentsBatch](docs/sdks/envs/README.md#postenvironmentsbatch) - Add a number of  Secure Application environments
* [PostEnvironmentsDelete](docs/sdks/envs/README.md#postenvironmentsdelete) - Delete multiple Secure Application environments
* [PutEnvironmentsEnvID](docs/sdks/envs/README.md#putenvironmentsenvid) - Edit an existing Secure Application environment

## [Expansions](docs/sdks/expansions/README.md)

* [DeleteExpansionsExpansionID](docs/sdks/expansions/README.md#deleteexpansionsexpansionid) - Delete an expansion
* [GetExpansions](docs/sdks/expansions/README.md#getexpansions) - List all the expansions on the system
* [GetExpansionsExpansionIDInstallExpansionTarGz](docs/sdks/expansions/README.md#getexpansionsexpansionidinstallexpansiontargz) - Get the expansion installation
* [PostExpansions](docs/sdks/expansions/README.md#postexpansions) - Create a new expansion
* [PutExpansionsExpansionID](docs/sdks/expansions/README.md#putexpansionsexpansionid) - Edit expansion definition

## [Gateways](docs/sdks/gateways/README.md)

* [DeleteGatewaysGatewayID](docs/sdks/gateways/README.md#deletegatewaysgatewayid) - Delete gateway
* [GetGateways](docs/sdks/gateways/README.md#getgateways) - Get gateways
* [GetGatewaysClusters](docs/sdks/gateways/README.md#getgatewaysclusters) - Get clusters info
* [GetGatewaysGatewayIDDownloadBundle](docs/sdks/gateways/README.md#getgatewaysgatewayiddownloadbundle) - Get a GW installation script
* [PostGateways](docs/sdks/gateways/README.md#postgateways) - Add new gateway
* [PutGatewaysGatewayID](docs/sdks/gateways/README.md#putgatewaysgatewayid) - Edit gateway

## [ImagesAndVulnerabilities](docs/sdks/imagesandvulnerabilities/README.md)

* [DeleteImagesID](docs/sdks/imagesandvulnerabilities/README.md#deleteimagesid) - Delete an image hash
* [GetAccountVulnerabilitiesXlsx](docs/sdks/imagesandvulnerabilities/README.md#getaccountvulnerabilitiesxlsx) - Returns a xlsx file of images alongside to their vulnerabilities.
* [GetImages](docs/sdks/imagesandvulnerabilities/README.md#getimages) - Returns a list of images
* [GetImagesImagesHash](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageshash) - search for image hash in the account
* [GetImagesVulnerabilitiesByImageNameAndHash](docs/sdks/imagesandvulnerabilities/README.md#getimagesvulnerabilitiesbyimagenameandhash) - Returns a list of vulnerabilities detected in the image
* [GetImagesID](docs/sdks/imagesandvulnerabilities/README.md#getimagesid) - get an image
* [GetImagesImageIDDockerfileScanResults](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageiddockerfilescanresults) - Returns a list of vulnerabilities detected in the  image
* [GetImagesImageIDImageLayers](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageidimagelayers) - Returns a list of image layers
* [GetImagesImageIDPackages](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageidpackages) - Returns a list of packages for a specific image
* [GetImagesImageIDSbomPath](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageidsbompath) - Returns the path to the SBOM in cloud storage
* [GetImagesImageIDVulnerabilities](docs/sdks/imagesandvulnerabilities/README.md#getimagesimageidvulnerabilities) - Returns a list of vulnerabilities detected in the image
* [PostImages](docs/sdks/imagesandvulnerabilities/README.md#postimages) - Define a New image hash
* [PostImagesApprove](docs/sdks/imagesandvulnerabilities/README.md#postimagesapprove) - Approve an image hash
* [PostImagesImageIDDockerfileScanResultsIgnore](docs/sdks/imagesandvulnerabilities/README.md#postimagesimageiddockerfilescanresultsignore) - Add / remove a list of  UUIDs of dockerfileScanResults from ignored list
* [PostImagesImageIDVulnerabilitiesIgnore](docs/sdks/imagesandvulnerabilities/README.md#postimagesimageidvulnerabilitiesignore) - Add / remove a list of  UUIDs of vulnerabilities from ignored list

## [K8sCisBenchmark](docs/sdks/k8scisbenchmark/README.md)

* [GetK8sCISBenchmark](docs/sdks/k8scisbenchmark/README.md#getk8scisbenchmark) - Get k8s cis benchmark for clusters
* [GetK8sCISBenchmarkSummary](docs/sdks/k8scisbenchmark/README.md#getk8scisbenchmarksummary) - Get k8s cis benchmark summary of account
* [GetK8sCISBenchmarkClusterID](docs/sdks/k8scisbenchmark/README.md#getk8scisbenchmarkclusterid) - Get k8s cis benchmark for a specific cluster
* [PostK8sCISBenchmarkClusterID](docs/sdks/k8scisbenchmark/README.md#postk8scisbenchmarkclusterid) - initiate k8s cis benchmark scan for a specific cluster
* [PutK8sCISBenchmarkClusterID](docs/sdks/k8scisbenchmark/README.md#putk8scisbenchmarkclusterid) - edit k8s cis benchmark for a specific cluster with test statuses

## [Kubernetes](docs/sdks/kubernetes/README.md)

* [DeleteKubernetesClustersKubernetesClusterID](docs/sdks/kubernetes/README.md#deletekubernetesclusterskubernetesclusterid) - Delete a Kubernetes cluster
* [DeletePodDefinitionsPodID](docs/sdks/kubernetes/README.md#deletepoddefinitionspodid) - Delete a pod definition
* [GetGetControllerDataClusterID](docs/sdks/kubernetes/README.md#getgetcontrollerdataclusterid) - get controller data using clusterId
* [GetIstioSupportedVersions](docs/sdks/kubernetes/README.md#getistiosupportedversions) - Get a list of istio releases that are supported by Secure Application agent. sorted from latest to oldest
* [GetKubernetesClusters](docs/sdks/kubernetes/README.md#getkubernetesclusters) - get a list of current Kubernetes clusters
* [GetKubernetesClustersKubernetesClusterID](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusterid) - get the Kubernetes cluster with the given id
* [GetKubernetesClustersKubernetesClusterIDDeleteDependencies](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusteriddeletedependencies) - get dependencies which need to be handled in order to delete specified Kubernetes cluster
* [GetKubernetesClustersKubernetesClusterIDDownloadBundle](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusteriddownloadbundle) - Get Secure Application installation script
* [GetKubernetesClustersKubernetesClusterIDGetHelmCommands](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusteridgethelmcommands) - Get Panoptica Aug release Helm command
* [GetKubernetesClustersKubernetesClusterIDNamespaces](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusteridnamespaces) - List namespaces on a specific Kubernetes cluster
* [GetKubernetesClustersKubernetesClusterIDServices](docs/sdks/kubernetes/README.md#getkubernetesclusterskubernetesclusteridservices) - List services on a specific Kubernetes cluster
* [GetLeanKubernetesClusters](docs/sdks/kubernetes/README.md#getleankubernetesclusters) - get a list of current Kubernetes clusters
* [GetNamespaces](docs/sdks/kubernetes/README.md#getnamespaces) - Get a list of current Kubernetes namespaces
* [GetPodDefinitions](docs/sdks/kubernetes/README.md#getpoddefinitions) - Get all pod definitions on the system
* [PostKubernetesClusters](docs/sdks/kubernetes/README.md#postkubernetesclusters) - Add a new Kubernetes cluster to Secure Application
* [PostPodDefinitions](docs/sdks/kubernetes/README.md#postpoddefinitions) - Create a new pod definition
* [PutKubernetesClustersKubernetesClusterID](docs/sdks/kubernetes/README.md#putkubernetesclusterskubernetesclusterid) - Update the Kubernetes cluster
* [PutKubernetesClustersKubernetesClusterIDManagedByHelm](docs/sdks/kubernetes/README.md#putkubernetesclusterskubernetesclusteridmanagedbyhelm) - Update the Kubernetes cluster which managed by HELM
* [PutPodDefinitionsPodID](docs/sdks/kubernetes/README.md#putpoddefinitionspodid) - Change pod definition

## [Mitre](docs/sdks/mitre/README.md)

* [GetMitreDashboard](docs/sdks/mitre/README.md#getmitredashboard) - Get data for MITRE dashboard for all clusters
* [GetMitreReportDownload](docs/sdks/mitre/README.md#getmitrereportdownload) - Download Mitre security report
* [GetMitreReportStatus](docs/sdks/mitre/README.md#getmitrereportstatus) - Get Mitre report status
* [GetMitreTechnique](docs/sdks/mitre/README.md#getmitretechnique) - Get data for MITRE technique of the given mitreTechniqueType
* [PostMitreReportGenerate](docs/sdks/mitre/README.md#postmitrereportgenerate) - Generate Mitre report
* [PostMitreTechniqueFix](docs/sdks/mitre/README.md#postmitretechniquefix) - Post fix for MITRE technique of the given mitreTechniqueType

## [Performance](docs/sdks/performance/README.md)

* [GetAPISecurityAPICatalogIDHitCountGraph](docs/sdks/performance/README.md#getapisecurityapicatalogidhitcountgraph) - Get hit count for specific spec path
* [GetPerformanceMetrics](docs/sdks/performance/README.md#getperformancemetrics) - Get performance metrics for a connection between workloads

## [PspProfiles](docs/sdks/pspprofiles/README.md)

* [DeletePodSecurityPolicyProfilesProfileID](docs/sdks/pspprofiles/README.md#deletepodsecuritypolicyprofilesprofileid) - Delete a pod security policy standards
* [DeleteSeccompProfilesProfileID](docs/sdks/pspprofiles/README.md#deleteseccompprofilesprofileid) - Delete a seccomp profile
* [GetPodSecurityPolicyProfiles](docs/sdks/pspprofiles/README.md#getpodsecuritypolicyprofiles) - Get all the pod security standards profiles on the system
* [GetSeccompProfiles](docs/sdks/pspprofiles/README.md#getseccompprofiles) - Get all the seccomp profiles on the system
* [PostPodSecurityPolicyProfiles](docs/sdks/pspprofiles/README.md#postpodsecuritypolicyprofiles) - Create a new pod security standards
* [PostPodSecurityPolicyProfilesBatch](docs/sdks/pspprofiles/README.md#postpodsecuritypolicyprofilesbatch) - Add a number of Pod Security Standards
* [PostSeccompProfiles](docs/sdks/pspprofiles/README.md#postseccompprofiles) - Add seccomp profile
* [PutPodSecurityPolicyProfilesProfileID](docs/sdks/pspprofiles/README.md#putpodsecuritypolicyprofilesprofileid) - Change pod security standards profile
* [PutSeccompProfilesProfileID](docs/sdks/pspprofiles/README.md#putseccompprofilesprofileid) - Change seccomp profile

## [Registries](docs/sdks/registries/README.md)

* [DeleteRegistriesRegistryID](docs/sdks/registries/README.md#deleteregistriesregistryid) - Delete a registry
* [GetRegistries](docs/sdks/registries/README.md#getregistries) - Get a list of defined registries
* [PostRegistries](docs/sdks/registries/README.md#postregistries) - Add new registry
* [PostRegistriesTestConnection](docs/sdks/registries/README.md#postregistriestestconnection) - test registry connection
* [PostRegistriesTestConnectionRegistryID](docs/sdks/registries/README.md#postregistriestestconnectionregistryid) - test registry connection
* [PutRegistriesRegistryID](docs/sdks/registries/README.md#putregistriesregistryid) - edit existing registry

## [RiskAssessment](docs/sdks/riskassessment/README.md)

* [DeleteRiskAssessmentIgnoredRisksIgnoredRiskID](docs/sdks/riskassessment/README.md#deleteriskassessmentignoredrisksignoredriskid) - Delete ignored risk
* [DeleteRiskAssessmentKubernetesClusterIDCancel](docs/sdks/riskassessment/README.md#deleteriskassessmentkubernetesclusteridcancel) - Cancel the runtime scan on the given cluster with the given id
* [GetRiskAssessment](docs/sdks/riskassessment/README.md#getriskassessment) - Get risk assessment data for all clusters
* [GetRiskAssessmentIgnoredRisks](docs/sdks/riskassessment/README.md#getriskassessmentignoredrisks) - Get all the ignored risks
* [GetRiskAssessmentPermissions](docs/sdks/riskassessment/README.md#getriskassessmentpermissions) - Get list of clusters and their permissions
* [GetRiskAssessmentPermissionsClusterID](docs/sdks/riskassessment/README.md#getriskassessmentpermissionsclusterid) - Get all of the users permissions
* [GetRiskAssessmentPermissionsClusterIDOwnerID](docs/sdks/riskassessment/README.md#getriskassessmentpermissionsclusteridownerid) - Get the owner permissions
* [GetRiskAssessmentPermissionsClusterIDOwnerIDRoleID](docs/sdks/riskassessment/README.md#getriskassessmentpermissionsclusteridowneridroleid) - Get the owner permissions
* [GetRiskAssessmentPoll](docs/sdks/riskassessment/README.md#getriskassessmentpoll) - Poll running scans
* [GetRiskAssessmentImageIDVulnerabilities](docs/sdks/riskassessment/README.md#getriskassessmentimageidvulnerabilities) - Get all images of given risk assessment pod
* [GetRiskAssessmentKubernetesClusterIDPods](docs/sdks/riskassessment/README.md#getriskassessmentkubernetesclusteridpods) - Get all risk assessments of all pods of given cluster
* [PostRiskAssessmentIgnoredRisks](docs/sdks/riskassessment/README.md#postriskassessmentignoredrisks) - Add ignore risk
* [PostRiskAssessmentPermissionsOwnerIDApprove](docs/sdks/riskassessment/README.md#postriskassessmentpermissionsowneridapprove) - add / remove permissions to /from the approved permissions list
* [PostRiskAssessmentKubernetesClusterIDScan](docs/sdks/riskassessment/README.md#postriskassessmentkubernetesclusteridscan) - Execute a new runtime scan on the given cluster with the given configuration
* [PostRiskAssessmentKubernetesClusterIDSettings](docs/sdks/riskassessment/README.md#postriskassessmentkubernetesclusteridsettings) - Save the runtime scan configuration on the given cluster
* [PutRiskAssessmentIgnoredRisksIgnoredRiskID](docs/sdks/riskassessment/README.md#putriskassessmentignoredrisksignoredriskid) - Edit ignore risk

## [RuntimeMap](docs/sdks/runtimemap/README.md)

* [DeleteNetworkMapQueueRequestID](docs/sdks/runtimemap/README.md#deletenetworkmapqueuerequestid) - Cancel the network map background job
* [GetNetworkMap](docs/sdks/runtimemap/README.md#getnetworkmap) - Get data for network map
* [GetNetworkMapQueueRequestID](docs/sdks/runtimemap/README.md#getnetworkmapqueuerequestid) - Get status for network map background job
* [GetNetworkMapResultsRequestID](docs/sdks/runtimemap/README.md#getnetworkmapresultsrequestid) - Get result for network map background job

## [Serverless](docs/sdks/serverless/README.md)

* [DeleteCloudAccountsCloudAccountID](docs/sdks/serverless/README.md#deletecloudaccountscloudaccountid) - Delete a cloud account
* [GetCloudAccounts](docs/sdks/serverless/README.md#getcloudaccounts) - List all the cloud accounts on the system
* [GetCloudAccountsAzureInstallationDetails](docs/sdks/serverless/README.md#getcloudaccountsazureinstallationdetails) - Get the Azure installation details
* [GetCloudAccountsInstallationDetails](docs/sdks/serverless/README.md#getcloudaccountsinstallationdetails) - Get the installation details
* [GetCloudAccountsRegionsAWS](docs/sdks/serverless/README.md#getcloudaccountsregionsaws) - List all the possible regions of AWS
* [GetCloudAccountsRegionsAzure](docs/sdks/serverless/README.md#getcloudaccountsregionsazure) - List all the possible regions of Azure
* [GetCloudAccountsCloudAccountIDDeleteDependencies](docs/sdks/serverless/README.md#getcloudaccountscloudaccountiddeletedependencies) - get dependencies which need to be handled in order to delete specified cloud account
* [GetCloudAccountsCloudAccountIDDownloadBundle](docs/sdks/serverless/README.md#getcloudaccountscloudaccountiddownloadbundle) - Get Secure Application installation script
* [GetServerlessFunctions](docs/sdks/serverless/README.md#getserverlessfunctions) - Get serverless functions
* [GetServerlessFunctionsArns](docs/sdks/serverless/README.md#getserverlessfunctionsarns) - Get serverless functions names
* [GetServerlessFunctionsNames](docs/sdks/serverless/README.md#getserverlessfunctionsnames) - Get serverless functions names
* [GetServerlessFunctionsFunctionID](docs/sdks/serverless/README.md#getserverlessfunctionsfunctionid) - Get Serverless Function by ID
* [GetServerlessFunctionsFunctionIDSecrets](docs/sdks/serverless/README.md#getserverlessfunctionsfunctionidsecrets) - Get Serverless Function secrets issues
* [GetServerlessFunctionsFunctionIDVulnerabilities](docs/sdks/serverless/README.md#getserverlessfunctionsfunctionidvulnerabilities) - Get Serverless Function Vulnerabilities by ID
* [GetServerlessZipFiles](docs/sdks/serverless/README.md#getserverlesszipfiles) - Get serverless zip files that was scanned by cli
* [GetServerlessZipFilesZipID](docs/sdks/serverless/README.md#getserverlesszipfileszipid) - Get specific zip file record
* [GetServerlessZipFilesZipIDPackages](docs/sdks/serverless/README.md#getserverlesszipfileszipidpackages) - Returns a list of packages for a specific serverless zip
* [GetServerlessZipFilesZipIDVulnerabilities](docs/sdks/serverless/README.md#getserverlesszipfileszipidvulnerabilities) - Returns a list of vulnerabilities detected in the serverless zip
* [PostCloudAccountsScan](docs/sdks/serverless/README.md#postcloudaccountsscan) - invoke cloud account scan
* [PutCloudAccountsCloudAccountID](docs/sdks/serverless/README.md#putcloudaccountscloudaccountid) - Edit cloud account definition

## [ServerlessPolicies](docs/sdks/serverlesspolicies/README.md)

* [GetServerlessPolicy](docs/sdks/serverlesspolicies/README.md#getserverlesspolicy) - Get current serverless policy
* [PutServerlessPolicy](docs/sdks/serverlesspolicies/README.md#putserverlesspolicy) - Set the current serverless policy

## [Settings](docs/sdks/settings/README.md)

* [DeleteSettingsIntegrationsCaID](docs/sdks/settings/README.md#deletesettingsintegrationscaid) - Delete the CA integration details
* [DeleteSettingsIntegrationsEventForwardingEventForwardingID](docs/sdks/settings/README.md#deletesettingsintegrationseventforwardingeventforwardingid) - Delete the event forwarding integration details with the given id
* [GetSettingsAgentsUpdate](docs/sdks/settings/README.md#getsettingsagentsupdate) - Get the agents update configurations
* [GetSettingsIntegrationsCa](docs/sdks/settings/README.md#getsettingsintegrationsca) - Get the CA integration details
* [GetSettingsIntegrationsEventForwarding](docs/sdks/settings/README.md#getsettingsintegrationseventforwarding) - Get the event forwarding integration details
* [PostSeccompProfilesValidateData](docs/sdks/settings/README.md#postseccompprofilesvalidatedata) - Test the seccomp profile data
* [PostSettingsAgentsUpdateUpdateNow](docs/sdks/settings/README.md#postsettingsagentsupdateupdatenow) - Update the agents of the account now
* [PostSettingsIntegrationsCa](docs/sdks/settings/README.md#postsettingsintegrationsca) - Set the CA integration details
* [PostSettingsIntegrationsEventForwarding](docs/sdks/settings/README.md#postsettingsintegrationseventforwarding) - Set the event forwarding integration details
* [PostSettingsIntegrationsOpsgenieTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationsopsgenietestintegration) - Test the connection to Opsgenie
* [PostSettingsIntegrationsSecurexTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationssecurextestintegration) - Test the SecureX integration by sending test message to the provided URL
* [PostSettingsIntegrationsSlackTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationsslacktestintegration) - Test the Slack integration by sending test message to the provided URL
* [PostSettingsIntegrationsSplunkTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationssplunktestintegration) - Test the connection to Splunk
* [PostSettingsIntegrationsSumoLogicTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationssumologictestintegration) - Test the Sumo Logic integration by sending test message to the provided URL
* [PostSettingsIntegrationsTeamsTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationsteamstestintegration) - Test the connection to Teams
* [PostSettingsIntegrationsWebexTestIntegration](docs/sdks/settings/README.md#postsettingsintegrationswebextestintegration) - Test the Webex integration by sending test message to the provided URL
* [PutSettingsAgentsUpdate](docs/sdks/settings/README.md#putsettingsagentsupdate) - get the agents update configurations.
* [PutSettingsIntegrationsCaID](docs/sdks/settings/README.md#putsettingsintegrationscaid) - Edit the CA integration details
* [PutSettingsIntegrationsEventForwardingEventForwardingID](docs/sdks/settings/README.md#putsettingsintegrationseventforwardingeventforwardingid) - Edit the event forwarding integration details

## [Telemetries](docs/sdks/telemetries/README.md)

* [GetAppTelemetries](docs/sdks/telemetries/README.md#getapptelemetries) - Get App telemetries
* [GetAppTelemetriesAppTelemetryID](docs/sdks/telemetries/README.md#getapptelemetriesapptelemetryid) - Get App telemetry by ID
* [GetAppTelemetriesAppTelemetryIDAPIRiskInfo](docs/sdks/telemetries/README.md#getapptelemetriesapptelemetryidapiriskinfo) - Get API risks info of given app telemetry
* [GetAppTelemetriesAppTelemetryIDImagePackages](docs/sdks/telemetries/README.md#getapptelemetriesapptelemetryidimagepackages) - list packages with licenses runnin on a pod
* [GetAppTelemetriesAppTelemetryIDInjectionInfo](docs/sdks/telemetries/README.md#getapptelemetriesapptelemetryidinjectioninfo) - Get token injection info of given app telemetry
* [GetConnectionTelemetries](docs/sdks/telemetries/README.md#getconnectiontelemetries) - Get connection telemetries
* [GetConnectionTelemetriesConnectionTelemetryID](docs/sdks/telemetries/README.md#getconnectiontelemetriesconnectiontelemetryid) - get details for a single connection telemetry

## [Tokens](docs/sdks/tokens/README.md)

* [DeleteTokensTokenID](docs/sdks/tokens/README.md#deletetokenstokenid) - Delete token
* [GetTokens](docs/sdks/tokens/README.md#gettokens) - Get tokens
* [GetTokensInfo](docs/sdks/tokens/README.md#gettokensinfo) - Get tokens by Ids
* [GetTokensTokenIDDeleteDependencies](docs/sdks/tokens/README.md#gettokenstokeniddeletedependencies) - get dependancies which need to be handled in order to delete specified token
* [PostTokens](docs/sdks/tokens/README.md#posttokens) - Add new token
* [PutTokensTokenID](docs/sdks/tokens/README.md#puttokenstokenid) - Edit token

## [Truncation](docs/sdks/truncation/README.md)

* [GetTruncationImages](docs/sdks/truncation/README.md#gettruncationimages) - Get workloads truncation time for account
* [GetTruncationWorkloads](docs/sdks/truncation/README.md#gettruncationworkloads) - Get workloads truncation time for account
* [PostTruncationImages](docs/sdks/truncation/README.md#posttruncationimages) - Update workloads truncation status for account
* [PostTruncationWorkloads](docs/sdks/truncation/README.md#posttruncationworkloads) - Update workloads truncation status for account

## [TrustedSigners](docs/sdks/trustedsigners/README.md)

* [DeleteTrustedSignersTrustedSignerID](docs/sdks/trustedsigners/README.md#deletetrustedsignerstrustedsignerid) - Delete a trusted signer
* [GetTrustedSigners](docs/sdks/trustedsigners/README.md#gettrustedsigners) - Get a list of defined trusted signers
* [GetTrustedSignersTrustedSignerID](docs/sdks/trustedsigners/README.md#gettrustedsignerstrustedsignerid) - get existing trusted signer
* [PostTrustedSigners](docs/sdks/trustedsigners/README.md#posttrustedsigners) - Add new trusted signer
* [PutTrustedSignersTrustedSignerID](docs/sdks/trustedsigners/README.md#puttrustedsignerstrustedsignerid) - edit existing trusted signer

## [Users](docs/sdks/users/README.md)

* [DeleteUsersUserID](docs/sdks/users/README.md#deleteusersuserid) - Delete a user
* [GetOperatorCredentials](docs/sdks/users/README.md#getoperatorcredentials) - get the credentials of the Secure Application Operator service user
* [GetUsers](docs/sdks/users/README.md#getusers) - List current users
* [GetUsersUserIDAccessTokens](docs/sdks/users/README.md#getusersuseridaccesstokens) - Get the  access tokens for the user
* [GetUsersUserIDDeleteDependencies](docs/sdks/users/README.md#getusersuseriddeletedependencies) - get dependencies which need to be handled in order to delete specified user
* [PostAccountUsageStatus](docs/sdks/users/README.md#postaccountusagestatus) - an api to get the account usage status
* [PostChangePassword](docs/sdks/users/README.md#postchangepassword) - Change the password for the current user
* [PostLogin](docs/sdks/users/README.md#postlogin) - Login
* [PostLogout](docs/sdks/users/README.md#postlogout) - Log out
* [PostMe](docs/sdks/users/README.md#postme) - an api to get current logged in user info
* [PostUsers](docs/sdks/users/README.md#postusers) - Create a user
* [PostUsersAcceptEula](docs/sdks/users/README.md#postusersaccepteula) - Accept the EULA
* [PostUsersTrial](docs/sdks/users/README.md#postuserstrial) - Create a trail user
* [PutUsersUserID](docs/sdks/users/README.md#putusersuserid) - Change user details

## [Vulnerabilities](docs/sdks/vulnerabilities/README.md)

* [GetVulnerabilities](docs/sdks/vulnerabilities/README.md#getvulnerabilities) - search for vulnerability names in the account
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->

<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:


<!-- End Pagination -->

<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
