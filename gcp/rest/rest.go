package rest

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"strings"
)

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
	Accessapproval:        {},
	Androidmanagement:     {},
	Appengine:             {},
	Automl:                {},
	Automlrecommendations: {},
	Bigquery:              {},
	Bigtable:              {},
	Billing:               {},
	Binaryauthorization:   {},
	Cloudasset:            {},
	Cloudbuild: {
		{
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds",
		},
		{
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds/{{.ResourceID}}",
		},
	},
	Cloudconfig:   {},
	Clouddebugger: {},
	Cloudfunctions: {
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
	Cloudiot:                    {},
	Cloudiottoken:               {},
	Cloudjobdiscovery:           {},
	Cloudkms:                    {},
	Cloudmessaging:              {},
	Cloudmigration:              {},
	Cloudnotifications:          {},
	Cloudprivatecatalog:         {},
	Cloudprivatecatalogproducer: {},
	Cloudprofiler:               {},
	Cloudscheduler:              {},
	Cloudsecurityscanner:        {},
	Cloudsql:                    {},
	Cloudtasks:                  {},
	Cloudtestservice:            {},
	Cloudtoolresults:            {},
	Cloudtrace:                  {},
	Cloudtranslate:              {},
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
	Container:              {},
	Datacatalog:            {},
	Dataflow:               {},
	Datafusion:             {},
	Datalabeling:           {},
	Dataprep:               {},
	Dataproc:               {},
	Datastore:              {},
	Deploymentmanager:      {},
	Dialogflow:             {},
	Dlp:                    {},
	Dns:                    {},
	Endpoints:              {},
	Errorreporting:         {},
	File:                   {},
	Firebase:               {},
	Firebaseabt:            {},
	Firebaseanalytics:      {},
	Firebaseauth:           {},
	Firebasecrash:          {},
	Firebasecrashlytics:    {},
	Firebasedatabase:       {},
	Firebasedynamiclinks:   {},
	Firebaseextensions:     {},
	Firebasehosting:        {},
	Firebaseinappmessaging: {},
	Firebaseml:             {},
	Firebasenotifications:  {},
	Firebaseperformance:    {},
	Firebasepredictions:    {},
	Firebaserules:          {},
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
	Managedidentities: {},
	Ml:                {},
	Monitoring:        {},
	Orgpolicy:         {},
	Proximitybeacon:   {},
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
	Redis:                     {},
	Remotebuildexecution:      {},
	Resourcemanager:           {},
	Run:                       {},
	Runtimeconfig:             {},
	Servicebroker:             {},
	Serviceconsumermanagement: {},
	Servicemanagement:         {},
	Servicenetworking:         {},
	Serviceusage:              {},
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
	Storagetransfer:              {},
	Subscribewithgoogledeveloper: {},
	Threatdetection:              {},
	Tpu:                          {},
	Vpcaccess:                    {},
}
