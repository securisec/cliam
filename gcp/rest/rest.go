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
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ProjectID}}/zones/{{.Zone}}/reservations",
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
			URL:              "https://iam.googleapis.com/v1/projects/{{.ProjectID}}/roles",
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
