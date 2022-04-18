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
	Region           string            `json:"region"`
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
	Bigtable: {
		{
			PermissionMethod: joinString(Bigtable, "operations", "projects", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/operations/projects/{{.ParentID}}/operations",
		},
		{
			PermissionMethod: joinString(Bigtable, "projects", "instances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/{{.ParentID}}/instances",
		},
		{
			PermissionMethod: joinString(Bigtable, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/{{.ParentID}}/locations",
		},
	},
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
	CloudIot:          {},
	CloudIotToken:     {},
	CloudJobDiscovery: {},
	CloudKMS: {
		{
			PermissionMethod: joinString(CloudKMS, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudkms.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	CloudMessaging:              {},
	CloudMigration:              {},
	CloudNotifications:          {},
	CloudPrivateCatalog:         {},
	CloudPrivateCatalogProducer: {},
	CloudProfiler:               {},
	CloudScheduler: {
		{
			PermissionMethod: joinString(CloudScheduler, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudscheduler.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	CloudSecurityScanner: {
		{
			PermissionMethod: joinString(CloudSecurityScanner, "projects", "scanConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://websecurityscanner.googleapis.com/v1/projects/{{.ParentID}}/scanConfigs",
		},
	},
	CloudSQL: {
		{
			PermissionMethod: joinString(CloudSQL, "instances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sqladmin.googleapis.com/v1/projects/{{.ParentID}}/instances",
		},
		{
			PermissionMethod: joinString(CloudSQL, "projects", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sqladmin.googleapis.com/v1/projects/{{.ParentID}}/operations",
		},
	},
	CloudTasks:       {},
	CloudTestService: {},
	CloudToolResults: {},
	CloudTrace:       {},
	CloudTranslate:   {},
	Composer:         {},
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
		{
			PermissionMethod: joinString(Compute, "instances", "acceleratorTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/acceleratorTypes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "addresses"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/addresses",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "autoscalers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/autoscalers",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "backendBuckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/backendBuckets",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "backendServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/backendServices",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "diskTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/diskTypes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "disks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/disks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "externalVpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/externalVpnGateways",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "firewalls"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/firewalls",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "forwardingRules"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/forwardingRules",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "globalAddresses"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/addresses",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "globalForwardingRules"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/forwardingRules",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "globalNetworkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/networkEndpointGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "globalOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/operations",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "globalPublicDelegatedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/publicDelegatedPrefixes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "healthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/healthChecks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "httpHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/httpHealthChecks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "httpsHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/httpsHealthChecks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "images"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/images",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "instanceGroupManagers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instanceGroupManagers",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "instanceGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instanceGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "instanceTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/instanceTemplates",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "interconnectAttachments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/interconnectAttachments",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "interconnectLocations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/interconnectLocations",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "interconnects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/interconnects",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "machineImages"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/machineImages",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "machineTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/machineTypes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "networkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/networkEndpointGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "networkFirewallPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/firewallPolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "networks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/networks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "nodeGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/nodeGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "nodeTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/nodeTemplates",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "nodeTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/nodeTypes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "packetMirrorings"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/packetMirrorings",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "publicAdvertisedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/publicAdvertisedPrefixes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "publicDelegatedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/publicDelegatedPrefixes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionAutoscalers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/autoscalers",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionBackendServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/backendServices",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionCommitments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/commitments",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionDiskTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/diskTypes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionDisks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/disks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionHealthCheckServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/healthCheckServices",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/healthChecks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionInstanceGroupManagers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/instanceGroupManagers",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionInstanceGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/instanceGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionNetworkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/networkEndpointGroups",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionNetworkFirewallPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/firewallPolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionNotificationEndpoints"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/notificationEndpoints",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/operations",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionSecurityPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/securityPolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionSslCertificates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/sslCertificates",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionTargetHttpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetHttpProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionTargetHttpsProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetHttpsProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "regionUrlMaps"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/urlMaps",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "reservations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/reservations",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "resourcePolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/resourcePolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "routers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/routers",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "routes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/routes",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "securityPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/securityPolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "serviceAttachments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/serviceAttachments",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "snapshots"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/snapshots",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "sslCertificates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/sslCertificates",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "sslPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/sslPolicies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "subnetworks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/subnetworks",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetGrpcProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetGrpcProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetHttpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetHttpProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetHttpsProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetHttpsProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetInstances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/targetInstances",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetPools",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetSslProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetSslProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetTcpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetTcpProxies",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "targetVpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetVpnGateways",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "urlMaps"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/urlMaps",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "vpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/vpnGateways",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "vpnTunnels"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/vpnTunnels",
		},
		{
			PermissionMethod: joinString(Compute, "instances", "zoneOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/operations",
		},
	},
	Container: {
		{
			PermissionMethod: joinString(Container, "projects", "aggregated", "usableSubnetworks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/aggregated/usableSubnetworks",
		},
		{
			PermissionMethod: joinString(Container, "projects", "zones", "clusters"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/zones/{{.Zone}}/clusters",
		},
		{
			PermissionMethod: joinString(Container, "projects", "zones", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/zones/{{.Zone}}/operations",
		},
	},
	DataCatalog: {},
	DataFlow: {
		{
			PermissionMethod: joinString(DataFlow, "projects", "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataflow.googleapis.com/v1b3/projects/{{.ParentID}}/jobs",
		},
	},
	DataFusion: {
		{
			PermissionMethod: joinString(DataFusion, "projects", "regions", "autoscalingPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/autoscalingPolicies",
		},
		{
			PermissionMethod: joinString(DataFusion, "projects", "regions", "clusters"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/clusters",
		},
		{
			PermissionMethod: joinString(DataFusion, "projects", "regions", "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/jobs",
		},
		{
			PermissionMethod: joinString(DataFusion, "projects", "regions", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/operations",
		},
		{
			PermissionMethod: joinString(DataFusion, "projects", "regions", "workflowTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/workflowTemplates",
		},
	},
	DataLabeling: {},
	DataPrep:     {},
	DataProc: {
		{
			PermissionMethod: joinString(DataProc, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://datafusion.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	DataStore:         {},
	DeploymentManager: {},
	DialogFlow:        {},
	Dlp:               {},
	Dns: {
		{
			PermissionMethod: joinString(Dns, "managedZones"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/managedZones",
		},
		{
			PermissionMethod: joinString(Dns, "policies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/policies",
		},
		{
			PermissionMethod: joinString(Dns, "responsePolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/responsePolicies",
		},
	},
	Endpoints:      {},
	ErrorReporting: {},
	File:           {},
	Firebase: {
		// https://firebase.google.com/docs/reference/hosting/rest
		{
			PermissionMethod: "firebase.hosting.projects.sites",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebasehosting.googleapis.com/v1beta1/projects/{{.ParentID}}/sites",
		},
		{
			PermissionMethod: joinString(Firebase, "firestore", "projects", "databases"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firestore.googleapis.com/v1/projects/{{.ParentID}}/databases",
		},
		{
			PermissionMethod: joinString(Firebase, "cloudstorage", "projects", "buckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebasestorage.googleapis.com/v1beta/projects/{{.ParentID}}/buckets",
		},
		{
			PermissionMethod: joinString(Firebase, "rules", "projects", "releases"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaserules.googleapis.com/v1/projects/{{.ParentID}}/releases",
		},
		{
			PermissionMethod: joinString(Firebase, "rules", "projects", "rulesets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaserules.googleapis.com/v1/projects/{{.ParentID}}/rulesets",
		},
		{
			PermissionMethod: joinString(Firebase, "config", "remoteConfig", "listVersions"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaseremoteconfig.googleapis.com/v1/projects/{{.ParentID}}/remoteConfig:listVersions",
		},
		{
			PermissionMethod: joinString(Firebase, "app", "distribution"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaseappdistribution.googleapis.com/v1/projects/{{.ParentID}}/testers",
		},
	},
	Genomics:   {},
	Healthcare: {},
	Iam: {
		{
			PermissionMethod: "iam.serviceAccounts",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/{{.ParentType}}/{{.ParentID}}/serviceAccounts",
		},
		{
			PermissionMethod: joinString(Iam, "organizations", "roles"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/{{.ParentType}}/{{.ParentID}}/roles",
		},
		{
			PermissionMethod: joinString(Iam, "projects", "locations", "workloadIdentityPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/projects/{{ .ParentID }}/locations/{{ .Zone }}/workloadIdentityPools",
		},
		{
			PermissionMethod: joinString(Iam, "projects", "roles"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/projects/{{.ParentID}}/roles",
		},
	},
	Iap: {
		{
			PermissionMethod: joinString(Iap, "projects", "brands"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iap.googleapis.com/v1/projects/{{.ParentID}}/brands",
		},
	},
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
		{
			PermissionMethod: joinString(Logging, "billingAccounts", "exclusions"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/exclusions",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(Logging, "billingAccounts", "logs"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/logs",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(Logging, "billingAccounts", "sinks"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/sinks",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(Logging, "exclusions"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/exclusions",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(Logging, "locations", "operations"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/operations",
			ReqMethod:        "GET",
		},
	},
	ManagedIdentities: {},
	Ml:                {},
	Monitoring:        {},
	OrgPolicy: {
		{
			PermissionMethod: joinString(OrgPolicy, "projects", "policies"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/projects/{{.ParentID}}/policies",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(OrgPolicy, "projects", "constraints"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/projects/{{.ParentID}}/constraints",
			ReqMethod:        "GET",
		},
		{
			PermissionMethod: joinString(OrgPolicy, "organizations", "policies"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/organizations/{{.ParentID}}/constraints",
			ReqMethod:        "GET",
		},
	},
	ProximityBeacon: {},
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
			PermissionMethod: joinString(Pubsub, "projects", "snapshots"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/snapshots",
		},
		{
			PermissionMethod: joinString(Pubsub, "projects", "schemas"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/schemas",
		},
		{
			PermissionMethod: "pubsub.topics",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/topics/{{.ResourceID}}",
		},
	},
	Redis: {
		{
			PermissionMethod: joinString(Redis, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://redis.googleapis.com/v1/{{.ParentID}}/locations",
		},
	},
	RemoteBuildExecution: {},
	ResourceManager: {
		{
			PermissionMethod: joinString(ResourceManager, "liens"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudresourcemanager.googleapis.com/v1/liens?parent=projects/{{.ParentID}}",
		},
		{
			PermissionMethod: joinString(ResourceManager, "projects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudresourcemanager.googleapis.com/v1/projects",
		},
	},
	Run: {
		{
			PermissionMethod: joinString(Run, "namespaces", "revisions"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/revisions",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "authorizeddomains"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/authorizeddomains",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "configurations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/configurations",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "domainmappings"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/domainmappings",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "routes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/routes",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "services"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/services",
		},
		{
			PermissionMethod: joinString(Run, "namespaces", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://run.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	RuntimeConfig: {},
	SecretManager: {
		{
			PermissionMethod: joinString(SecretManager, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://secretmanager.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
		{
			PermissionMethod: joinString(SecretManager, "projects", "secrets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://secretmanager.googleapis.com/v1/projects/{{.ParentID}}/secrets",
		},
	},
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
	Spanner: {
		{
			PermissionMethod: joinString(Spanner, "projects", "instanceConfigs"),
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://spanner.googleapis.com/v1/projects/{{.ParentID}}/instanceConfigs",
		},
		{
			PermissionMethod: joinString(Spanner, "projects", "instances"),
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://spanner.googleapis.com/v1/projects/{{.ParentID}}/instances",
		},
	},
	Stackdriver: {},
	Storage: {
		{
			PermissionMethod: joinString(Storage, "buckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://storage.googleapis.com/storage/v1/b?project={{.ParentID}}",
		},
		{
			PermissionMethod: joinString(Storage, "projects", "hmacKeys"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://storage.googleapis.com/storage/v1/projects/{{ .ParentID }}/hmacKeys",
		},
	},
	StorageTransfer:              {},
	SubscribeWithGoogleDeveloper: {},
	ThreatDetection:              {},
	Tpu:                          {},
	VpcAccess:                    {},
}
