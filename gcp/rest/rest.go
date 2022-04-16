package rest

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"strings"
)

func joinString(s ...string) string {
	h := make([]string, 0)
	for _, i := range s {
		h = append(h, strings.ToLower(i))
	}
	return strings.Join(h, ".")
}

var emptyBody = map[string]string{}

type RestCall struct {
	ParentType       string            `json:"parent_type"`
	ParentID         string            `json:"parent_id"`
	ResourceID       string            `json:"resource_id"`
	Zone             string            `json:"zone"`
	URL              string            `json:"url"`
	PermissionMethod string            `json:"method"`
	Action           string            `json:"action"`
	ReqMethod        string            `json:"req_method"`
	ReqBody          map[string]string `json:"req_body"`
	Response         *http.Response    `json:"response"`
}

func (r RestCall) GetURL() (string, error) {
	t := template.Must(template.New("").Parse("{{.URL}}"))
	b := bytes.Buffer{}
	if err := t.Execute(&b, r); err != nil {
		return "", err
	}

	bb := bytes.Buffer{}
	if err := template.Must(template.New("").Parse(b.String())).Execute(&bb, r); err != nil {
		return "", err
	}
	url := bb.String()
	if strings.Contains(strings.Replace(url, "https://", "", -1), "//") {
		// one or more of the template variables is empty
		return "", errors.New("invalid url.")
	}
	return url, nil
}

var RestApiCalls = map[string][]RestCall{
	AccessApproval:        {},
	AndroidManagement:     {},
	AppEngine:             {},
	Automl:                {},
	AutoMLRecommendations: {},
	Bigquery: {
		// https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets
		{
			PermissionMethod: joinString(Bigquery, "datasets"),
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/datasets/{{.ResourceID}}",
		},
		{
			PermissionMethod: joinString(Bigquery, "datasets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/datasets",
		},
		{
			PermissionMethod: joinString(Bigquery, "jobs"),
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/jobs/{{.ResourceID}}",
		},
		{
			PermissionMethod: joinString(Bigquery, "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/jobs",
		},
		{
			PermissionMethod: joinString(Bigquery, "projects"),
			ReqMethod:        "GET",
			Action:           "getServiceAccount",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/serviceAccount",
		},
		{
			PermissionMethod: joinString(Bigquery, "projects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects",
		},
	},
	Bigtable:            {},
	Billing:             {},
	BinaryAuthorization: {},
	CloudAsset:          {},
	CloudBuild: {
		{
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds",
		},
		{
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds/{{.ResourceID}}",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "githubEnterpriseConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/githubEnterpriseConfigs",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "bitbucketServerConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{ .ParentID }}/locations/{{ .Zone }}/bitbucketServerConfigs",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "builds"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/builds",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "githubEnterpriseConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/githubEnterpriseConfigs",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "triggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/triggers",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "workerPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/workerPools",
		},
		{
			PermissionMethod: joinString(CloudBuild, "projects", "triggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/triggers",
		},
	},
	CloudConfig:   {},
	CloudDebugger: {},
	CloudFunctions: {
		{
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/-/functions",
		},
		{
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/functions/{{.ResourceID}}",
		},
		{
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "POST",
			Action:           "generateDownloadUrl",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Zone}}/functions/{{.ResourceID}}:generateDownloadUrl",
			ReqBody:          emptyBody,
		},
		{
			PermissionMethod: "cloudfunctions.locations",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	CloudIot:                    {},
	CloudIotToken:               {},
	CloudJobDiscovery:           {},
	CloudKMS:                    {},
	CloudMessaging:              {},
	CloudMigration:              {},
	CloudNotifications:          {},
	CloudPrivateCatalog:         {},
	CloudPrivateCatalogProducer: {},
	CloudProfiler:               {},
	CloudScheduler:              {},
	CloudSecurityScanner:        {},
	CloudSQL:                    {},
	CloudTasks:                  {},
	CloudTestService:            {},
	CloudToolResults:            {},
	CloudTrace:                  {},
	CloudTranslate:              {},
	Composer:                    {},
	Compute: {
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getEffectiveFirewalls",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getEffectiveFirewalls",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getGuestAttributes",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getGuestAttributes",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getIamPolicy",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "screenshot",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/screenshot",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "serialPort",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/serialPort",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getShieldedInstanceIdentity",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getShieldedInstanceIdentity",
		},
		{
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "referrers",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/referrers",
		},
	},
	Container:            {},
	DataCatalog:          {},
	DataFlow:             {},
	DataFusion:           {},
	DataLabeling:         {},
	DataPrep:             {},
	DataProc:             {},
	DataStore:            {},
	DeploymentManager:    {},
	DialogFlow:           {},
	Dlp:                  {},
	Dns:                  {},
	Endpoints:            {},
	ErrorReporting:       {},
	File:                 {},
	Firebase:             {},
	FirebaseAbt:          {},
	FirebaseAnalytics:    {},
	FirebaseAuth:         {},
	FirebaseCrash:        {},
	FirebaseCrashlytics:  {},
	FirebaseDatabase:     {},
	FirebaseDynamicLinks: {},
	FirebaseExtensions:   {},
	FirebaseHosting: {
		// https://firebase.google.com/docs/reference/hosting/rest
		{
			PermissionMethod: "firebasehosting.projects.sites",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebasehosting.googleapis.com/v1beta1/projects/{{.ParentID}}/sites",
		},
	},
	FirebaseInAppMessaging: {},
	FirebaseML:             {},
	FirebaseNotifications:  {},
	FirebasePerformance:    {},
	FirebasePredictions:    {},
	FirebaseRules:          {},
	Genomics:               {},
	Healthcare:             {},
	Iam: {
		{
			PermissionMethod: "iam.serviceAccounts",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/{{.ParentType}}/{{.ParentID}}/serviceAccounts",
		},
	},
	Iap: {},
	Logging: {
		{
			PermissionMethod: "logging.logEntries",
			Action:           "list",
			ReqMethod:        "POST",
			URL:              "https://logging.googleapis.com/v2/entries:list",
			ReqBody:          emptyBody,
		},
		{
			PermissionMethod: "logging.logs",
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/logs",
			ReqMethod:        "GET",
		},
	},
	ManagedIdentities: {},
	Ml:                {},
	Monitoring:        {},
	OrgPolicy:         {},
	ProximityBeacon:   {},
	Pubsub: {
		{
			PermissionMethod: "pubsub.subscriptions",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/subscriptions",
		},
		{
			PermissionMethod: "pubsub.subscriptions",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/subscriptions/{{.ResourceID}}",
		},
		{
			PermissionMethod: "pubsub.topics",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/topics",
		},
		{
			PermissionMethod: "pubsub.topics",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/topics/{{.ResourceID}}",
		},
	},
	Redis:                {},
	RemoteBuildExecution: {},
	ResourceManager:      {},
	Run: {
		{
			PermissionMethod: joinString(Run, "namespaces", "revisions"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/revisions",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "authorizeddomains"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/authorizeddomains",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "configurations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/configurations",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "domainmappings"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/domainmappings",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "routes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/routes",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "services"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Zone}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/services",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://run.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	RuntimeConfig:             {},
	ServiceBroker:             {},
	ServiceConsumerManagement: {},
	ServiceManagement:         {},
	ServiceNetworking:         {},
	ServiceUsage:              {},
	Source: {
		{
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "getConfig",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/config",
		},
		{
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos",
		},
		{
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos/{{.ResourceID}}",
		},
		{
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos/{{.ResourceID}}:getIamPolicy",
		},
	},
	Spanner:                      {},
	Stackdriver:                  {},
	Storage:                      {},
	StorageTransfer:              {},
	SubscribeWithGoogleDeveloper: {},
	ThreatDetection:              {},
	Tpu:                          {},
	VpcAccess:                    {},
}
