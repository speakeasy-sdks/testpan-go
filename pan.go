// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package testpango

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"github.com/speakeasy-sdks/testpan-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https:///api",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Pan
//
// https://panoptica.readme.io/reference - Product Documentation
type Pan struct {
	// APIs used for login and password management
	Users *Users
	// APIs used to define and manage  image hashes
	ImagesAndVulnerabilities *ImagesAndVulnerabilities
	// APIs used to get policy recommendations
	Advisor *Advisor
	// APIs use to  interact with  agents
	AgentManagement *AgentManagement
	// APIs to get the Secure Application API specification file
	API *API
	// APIs used to manage Api Security
	APISecurity *APISecurity
	Performance *Performance
	Bfla        *Bfla
	// APIs used to  define and manage api security policies
	APISecurityPolicies *APISecurityPolicies
	// APIs used to query for telemetries
	Telemetries *Telemetries
	// APIs used to define apps
	Apps *Apps
	// APIs used to  define and manage environment policies
	EnvironmentPolicies *EnvironmentPolicies
	// APIs used to retrieve  audit logs
	AuditLogs *AuditLogs
	// APIs used to change  credentials or return details about the  user's AWS environment
	Aws *Aws
	// APIs used to query for CD pipelines results
	Cd *Cd
	// APIs used to  define and manage CI/CD policies
	CICDPolicies *CICDPolicies
	Serverless   *Serverless
	// APIs used to  define and manage connection policies
	ConnectionPolicies *ConnectionPolicies
	// APIs to get dashboard statistics
	Dashboard *Dashboard
	// APIs used to manage deployers on Secure Application
	Deployers *Deployers
	// APIs used to define environments
	Envs *Envs
	// APIs used to manage expansions on Secure Application
	Expansions *Expansions
	Gateways   *Gateways
	// APIs used to manage Kubernetes clusters on Secure Application
	Kubernetes *Kubernetes
	// APIs to get the kubernetes cis benchmark data
	K8sCisBenchmark *K8sCisBenchmark
	// APIs used to  define and manage cluster events policies
	ClusterEventsPolicies *ClusterEventsPolicies
	Mitre                 *Mitre
	// APIs used to  query for network map
	RuntimeMap *RuntimeMap
	// APIs used to manage pod security standards profiles on Secure Application
	PspProfiles *PspProfiles
	// APIs used to  define and manage registries
	Registries *Registries
	// APIs used to manage risk assessment on Kubernetes clusters
	RiskAssessment *RiskAssessment
	// APIs used  to configure system settings
	Settings           *Settings
	ServerlessPolicies *ServerlessPolicies
	Tokens             *Tokens
	// APIs to get the Secure Application CLI
	Cli *Cli
	// APIs to delete workloads
	Truncation *Truncation
	// APIs used to  define and manage trusted signers
	TrustedSigners  *TrustedSigners
	Vulnerabilities *Vulnerabilities

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Pan)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Pan) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Pan) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Pan) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Pan) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Pan) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Pan) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Pan) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Pan {
	sdk := &Pan{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.8.1",
			GenVersion:        "2.234.3",
			UserAgent:         "speakeasy-sdk/go 0.8.1 2.234.3 1.0.0 github.com/speakeasy-sdks/testpan-go",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.ImagesAndVulnerabilities = newImagesAndVulnerabilities(sdk.sdkConfiguration)

	sdk.Advisor = newAdvisor(sdk.sdkConfiguration)

	sdk.AgentManagement = newAgentManagement(sdk.sdkConfiguration)

	sdk.API = newAPI(sdk.sdkConfiguration)

	sdk.APISecurity = newAPISecurity(sdk.sdkConfiguration)

	sdk.Performance = newPerformance(sdk.sdkConfiguration)

	sdk.Bfla = newBfla(sdk.sdkConfiguration)

	sdk.APISecurityPolicies = newAPISecurityPolicies(sdk.sdkConfiguration)

	sdk.Telemetries = newTelemetries(sdk.sdkConfiguration)

	sdk.Apps = newApps(sdk.sdkConfiguration)

	sdk.EnvironmentPolicies = newEnvironmentPolicies(sdk.sdkConfiguration)

	sdk.AuditLogs = newAuditLogs(sdk.sdkConfiguration)

	sdk.Aws = newAws(sdk.sdkConfiguration)

	sdk.Cd = newCd(sdk.sdkConfiguration)

	sdk.CICDPolicies = newCICDPolicies(sdk.sdkConfiguration)

	sdk.Serverless = newServerless(sdk.sdkConfiguration)

	sdk.ConnectionPolicies = newConnectionPolicies(sdk.sdkConfiguration)

	sdk.Dashboard = newDashboard(sdk.sdkConfiguration)

	sdk.Deployers = newDeployers(sdk.sdkConfiguration)

	sdk.Envs = newEnvs(sdk.sdkConfiguration)

	sdk.Expansions = newExpansions(sdk.sdkConfiguration)

	sdk.Gateways = newGateways(sdk.sdkConfiguration)

	sdk.Kubernetes = newKubernetes(sdk.sdkConfiguration)

	sdk.K8sCisBenchmark = newK8sCisBenchmark(sdk.sdkConfiguration)

	sdk.ClusterEventsPolicies = newClusterEventsPolicies(sdk.sdkConfiguration)

	sdk.Mitre = newMitre(sdk.sdkConfiguration)

	sdk.RuntimeMap = newRuntimeMap(sdk.sdkConfiguration)

	sdk.PspProfiles = newPspProfiles(sdk.sdkConfiguration)

	sdk.Registries = newRegistries(sdk.sdkConfiguration)

	sdk.RiskAssessment = newRiskAssessment(sdk.sdkConfiguration)

	sdk.Settings = newSettings(sdk.sdkConfiguration)

	sdk.ServerlessPolicies = newServerlessPolicies(sdk.sdkConfiguration)

	sdk.Tokens = newTokens(sdk.sdkConfiguration)

	sdk.Cli = newCli(sdk.sdkConfiguration)

	sdk.Truncation = newTruncation(sdk.sdkConfiguration)

	sdk.TrustedSigners = newTrustedSigners(sdk.sdkConfiguration)

	sdk.Vulnerabilities = newVulnerabilities(sdk.sdkConfiguration)

	return sdk
}
