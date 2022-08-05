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
	t = t.Option("missingkey=error")
	b := bytes.Buffer{}

	if err := t.Execute(&b, r); err != nil {
		return "", err
	}

	bb := bytes.Buffer{}
	tt := template.Must(template.New("").Parse(b.String()))
	tt = tt.Option("missingkey=error")

	if err := tt.Execute(&bb, r); err != nil {
		return "", err
	}

	url := bb.String()
	if strings.Contains(strings.Replace(url, "https://", "", -1), "//") {
		// one or more of the template variables is empty
		return "", errors.New("invalid url. Needed parameters are missing")
	}
	return url, nil
}

var RestApiCalls = map[string]map[string]RestCall{
	AccessApproval:        {},
	AndroidManagement:     {},
	AppEngine:             {},
	Automl:                {},
	AutoMLRecommendations: {},
	Bigquery: {
		// https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets
		"datasets_get": {
			PermissionMethod: joinString(Bigquery, "datasets"),
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/datasets/{{.ResourceID}}",
		},
		"datasets_list": {
			PermissionMethod: joinString(Bigquery, "datasets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/datasets",
		},
		"jobs_get": {
			PermissionMethod: joinString(Bigquery, "jobs"),
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/jobs/{{.ResourceID}}",
		},
		"jobs_list": {
			PermissionMethod: joinString(Bigquery, "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/jobs",
		},
		"projects_get": {
			PermissionMethod: joinString(Bigquery, "projects"),
			ReqMethod:        "GET",
			Action:           "getServiceAccount",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects/{{.ParentID}}/serviceAccount",
		},
		"projects_list": {
			PermissionMethod: joinString(Bigquery, "projects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigquery.googleapis.com/bigquery/v2/projects",
		},
	},
	Bigtable: {
		"operations_list": {
			PermissionMethod: joinString(Bigtable, "operations", "projects", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/operations/projects/{{.ParentID}}/operations",
		},
		"instances_list": {
			PermissionMethod: joinString(Bigtable, "projects", "instances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/{{.ParentID}}/instances",
		},
		"locations_list": {
			PermissionMethod: joinString(Bigtable, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://bigtableadmin.googleapis.com/v2/{{.ParentID}}/locations",
		},
	},
	Billing: {},
	BinaryAuthorization: {
		"aaestors_list": {
			PermissionMethod: joinString(BinaryAuthorization, "attestors"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://binaryauthorization.googleapis.com/v1/projects/{{.ParentID}}/attestors",
		},
	},
	CloudAsset: {},
	CloudBuild: {
		"builds_list": {
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds",
		},
		"builds": {
			PermissionMethod: "cloudbuild.builds",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/builds/{{.ResourceID}}",
		},
		"githubEnterpriseConfigs_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "githubEnterpriseConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/githubEnterpriseConfigs",
		},
		"bitbucketServerConfigs_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "bitbucketServerConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{ .ParentID }}/locations/{{ .Zone }}/bitbucketServerConfigs",
		},
		"locations_builds_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "builds"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/builds",
		},
		"locations_githubEnterpriseConfigs_get": {
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "githubEnterpriseConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/githubEnterpriseConfigs",
		},
		"locations_triggers_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "triggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/triggers",
		},
		"locations_workerPools_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "locations", "workerPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/workerPools",
		},
		"triggers_list": {
			PermissionMethod: joinString(CloudBuild, "projects", "triggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudbuild.googleapis.com/v1/projects/{{.ParentID}}/triggers",
		},
	},
	CloudConfig:   {},
	CloudDebugger: {},
	CloudFunctions: {
		"functions_list": {
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/-/functions",
		},
		"functions_get": {
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/functions/{{.ResourceID}}",
		},
		"generateDownloadURL": {
			PermissionMethod: "cloudfunctions.functions",
			ReqMethod:        "POST",
			Action:           "generateDownloadUrl",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/functions/{{.ResourceID}}:generateDownloadUrl",
			ReqBody:          emptyBody,
		},
		"locations": {
			PermissionMethod: "cloudfunctions.locations",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudfunctions.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	CloudIot: {
		"registries_list": {
			PermissionMethod: joinString(CloudIot, "projects", "locations", "registries"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudiot.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/registries",
		},
	},
	CloudIotToken:     {},
	CloudJobDiscovery: {},
	CloudKMS: {
		"locations_list": {
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
		"locations_list": {
			PermissionMethod: joinString(CloudScheduler, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudscheduler.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	CloudSecurityScanner: {
		"scanConfigs_list": {
			PermissionMethod: joinString(CloudSecurityScanner, "projects", "scanConfigs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://websecurityscanner.googleapis.com/v1/projects/{{.ParentID}}/scanConfigs",
		},
	},
	CloudSQL: {
		"instances_list": {
			PermissionMethod: joinString(CloudSQL, "instances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sqladmin.googleapis.com/v1/projects/{{.ParentID}}/instances",
		},
		"operations_list": {
			PermissionMethod: joinString(CloudSQL, "projects", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sqladmin.googleapis.com/v1/projects/{{.ParentID}}/operations",
		},
	},
	CloudTasks: {
		"locations_list": {
			PermissionMethod: joinString(CloudTasks, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudtasks.googleapis.com/v2/projects/{{.ParentID}}/locations",
		},
	},
	CloudTestService: {},
	CloudToolResults: {},
	CloudTrace: {
		"traces_list": {
			PermissionMethod: joinString(CloudTrace, "projects", "traces"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudtrace.googleapis.com/v1/projects/{{.ParentID}}/traces",
		},
		"traceSinks_list": {
			PermissionMethod: joinString(CloudTrace, "projects", "traceSinks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudtrace.googleapis.com/v2beta1/projects/{{.ParentID}}/traceSinks",
		},
	},
	CloudTranslate: {},
	Composer:       {},
	Compute: {
		"zone_instances_list": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances",
		},
		"zone_instances_get": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}",
		},
		"getEffectiveFirewalls": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getEffectiveFirewalls",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getEffectiveFirewalls",
		},
		"getGuestAttributes": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getGuestAttributes",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getGuestAttributes",
		},
		"getIamPolicy": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getIamPolicy",
		},
		"screenshot": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "screenshot",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/screenshot",
		},
		"serialPort": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "serialPort",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/serialPort",
		},
		"getShieldedInstanceIdentity": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "getShieldedInstanceIdentity",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/getShieldedInstanceIdentity",
		},
		"referrers": {
			PermissionMethod: "compute.instances",
			ReqMethod:        "GET",
			Action:           "referrers",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instances/{{.ResourceID}}/referrers",
		},
		"acceleratorTypes_list": {
			PermissionMethod: joinString(Compute, "instances", "acceleratorTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/acceleratorTypes",
		},
		"addresses_list": {
			PermissionMethod: joinString(Compute, "instances", "addresses"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/addresses",
		},
		"autoscalers_list": {
			PermissionMethod: joinString(Compute, "instances", "autoscalers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/autoscalers",
		},
		"backendBuckets_list": {
			PermissionMethod: joinString(Compute, "instances", "backendBuckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/backendBuckets",
		},
		"backendServices_list": {
			PermissionMethod: joinString(Compute, "instances", "backendServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/backendServices",
		},
		"diskTypes_list": {
			PermissionMethod: joinString(Compute, "instances", "diskTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/diskTypes",
		},
		"disks_list": {
			PermissionMethod: joinString(Compute, "instances", "disks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/disks",
		},
		"externalVpnGateways_list": {
			PermissionMethod: joinString(Compute, "instances", "externalVpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/externalVpnGateways",
		},
		"firewalls_list": {
			PermissionMethod: joinString(Compute, "instances", "firewalls"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/firewalls",
		},
		"forwardingRules_list": {
			PermissionMethod: joinString(Compute, "instances", "forwardingRules"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/forwardingRules",
		},
		"globalAddresses_list": {
			PermissionMethod: joinString(Compute, "instances", "globalAddresses"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/addresses",
		},
		"globalForwardingRules_list": {
			PermissionMethod: joinString(Compute, "instances", "globalForwardingRules"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/forwardingRules",
		},
		"globalNetworkEndpointGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "globalNetworkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/networkEndpointGroups",
		},
		"globalOperations_list": {
			PermissionMethod: joinString(Compute, "instances", "globalOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/operations",
		},
		"globalPublicDelegatedPrefixes_list": {
			PermissionMethod: joinString(Compute, "instances", "globalPublicDelegatedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/publicDelegatedPrefixes",
		},
		"healthChecks_list": {
			PermissionMethod: joinString(Compute, "instances", "healthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/healthChecks",
		},
		"httpHealthChecks_list": {
			PermissionMethod: joinString(Compute, "instances", "httpHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/httpHealthChecks",
		},
		"httpsHealthChecks_list": {
			PermissionMethod: joinString(Compute, "instances", "httpsHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/httpsHealthChecks",
		},
		"images_list": {
			PermissionMethod: joinString(Compute, "instances", "images"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/images",
		},
		"instanceGroupManagers_list": {
			PermissionMethod: joinString(Compute, "instances", "instanceGroupManagers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instanceGroupManagers",
		},
		"instanceGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "instanceGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/instanceGroups",
		},
		"instanceTemplates_list": {
			PermissionMethod: joinString(Compute, "instances", "instanceTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/instanceTemplates",
		},
		"interconnectAttachment_list": {
			PermissionMethod: joinString(Compute, "instances", "interconnectAttachments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/interconnectAttachments",
		},
		"interconnectLocations_list": {
			PermissionMethod: joinString(Compute, "instances", "interconnectLocations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/interconnectLocations",
		},
		"interconnects_list": {
			PermissionMethod: joinString(Compute, "instances", "interconnects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/interconnects",
		},
		"machineImages_list": {
			PermissionMethod: joinString(Compute, "instances", "machineImages"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/machineImages",
		},
		"machineTypes_list": {
			PermissionMethod: joinString(Compute, "instances", "machineTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/machineTypes",
		},
		"networkEndpointGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "networkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/networkEndpointGroups",
		},
		"networkFirewallPolicys_list": {
			PermissionMethod: joinString(Compute, "instances", "networkFirewallPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/firewallPolicies",
		},
		"networks_list": {
			PermissionMethod: joinString(Compute, "instances", "networks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/networks",
		},
		"nodeGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "nodeGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/nodeGroups",
		},
		"nodeTemplates_list": {
			PermissionMethod: joinString(Compute, "instances", "nodeTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/nodeTemplates",
		},
		"nodeTypes_list": {
			PermissionMethod: joinString(Compute, "instances", "nodeTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/nodeTypes",
		},
		"packetMirrorings_list": {
			PermissionMethod: joinString(Compute, "instances", "packetMirrorings"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/packetMirrorings",
		},
		"publicAdvertisedPrefixes_list": {
			PermissionMethod: joinString(Compute, "instances", "publicAdvertisedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/publicAdvertisedPrefixes",
		},
		"publicDelegatedPrefixes_list": {
			PermissionMethod: joinString(Compute, "instances", "publicDelegatedPrefixes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/publicDelegatedPrefixes",
		},
		"regionAutoscalers_list": {
			PermissionMethod: joinString(Compute, "instances", "regionAutoscalers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/autoscalers",
		},
		"regionBackendServices_list": {
			PermissionMethod: joinString(Compute, "instances", "regionBackendServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/backendServices",
		},
		"regionCommitments_list": {
			PermissionMethod: joinString(Compute, "instances", "regionCommitments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/commitments",
		},
		"regionDiskTypes_list": {
			PermissionMethod: joinString(Compute, "instances", "regionDiskTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/diskTypes",
		},
		"regionDisks_list": {
			PermissionMethod: joinString(Compute, "instances", "regionDisks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/disks",
		},
		"regionHealthCheckServices_list": {
			PermissionMethod: joinString(Compute, "instances", "regionHealthCheckServices"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/healthCheckServices",
		},
		"regionHealthChecks_list": {
			PermissionMethod: joinString(Compute, "instances", "regionHealthChecks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/healthChecks",
		},
		"regionInstanceGroupManagers_list": {
			PermissionMethod: joinString(Compute, "instances", "regionInstanceGroupManagers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/instanceGroupManagers",
		},
		"regionInstanceGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "regionInstanceGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/instanceGroups",
		},
		"regionNetworkEndpointGroups_list": {
			PermissionMethod: joinString(Compute, "instances", "regionNetworkEndpointGroups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/networkEndpointGroups",
		},
		"regionNetworkFirewallPolicies_list": {
			PermissionMethod: joinString(Compute, "instances", "regionNetworkFirewallPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/firewallPolicies",
		},
		"regionNotificationEndpoints_list": {
			PermissionMethod: joinString(Compute, "instances", "regionNotificationEndpoints"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/notificationEndpoints",
		},
		"regionOperations_list": {
			PermissionMethod: joinString(Compute, "instances", "regionOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/operations",
		},
		"regionSecurityPolicies_list": {
			PermissionMethod: joinString(Compute, "instances", "regionSecurityPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/securityPolicies",
		},
		"regionSslCertificates_list": {
			PermissionMethod: joinString(Compute, "instances", "regionSslCertificates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/sslCertificates",
		},
		"regionHttpProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "regionTargetHttpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetHttpProxies",
		},
		"regionTargetHttpsProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "regionTargetHttpsProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetHttpsProxies",
		},
		"regionUrlMaps_list": {
			PermissionMethod: joinString(Compute, "instances", "regionUrlMaps"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/urlMaps",
		},
		"reservations_list": {
			PermissionMethod: joinString(Compute, "instances", "reservations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/reservations",
		},
		"resourcePolicies_list": {
			PermissionMethod: joinString(Compute, "instances", "resourcePolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/resourcePolicies",
		},
		"routers_list": {
			PermissionMethod: joinString(Compute, "instances", "routers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/routers",
		},
		"routes_list": {
			PermissionMethod: joinString(Compute, "instances", "routes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/routes",
		},
		"securityPolicies_list": {
			PermissionMethod: joinString(Compute, "instances", "securityPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/securityPolicies",
		},
		"serviceAttachments_list": {
			PermissionMethod: joinString(Compute, "instances", "serviceAttachments"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/serviceAttachments",
		},
		"snapshots_list": {
			PermissionMethod: joinString(Compute, "instances", "snapshots"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/snapshots",
		},
		"sslCertificates_list": {
			PermissionMethod: joinString(Compute, "instances", "sslCertificates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/sslCertificates",
		},
		"sslPolicies_list": {
			PermissionMethod: joinString(Compute, "instances", "sslPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/sslPolicies",
		},
		"subnetworks_list": {
			PermissionMethod: joinString(Compute, "instances", "subnetworks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/subnetworks",
		},
		"targetGrpcProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "targetGrpcProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetGrpcProxies",
		},
		"targetHttpProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "targetHttpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetHttpProxies",
		},
		"targetHttpsProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "targetHttpsProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetHttpsProxies",
		},
		"targetInstances_list": {
			PermissionMethod: joinString(Compute, "instances", "targetInstances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/targetInstances",
		},
		"targetPools_list": {
			PermissionMethod: joinString(Compute, "instances", "targetPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetPools",
		},
		"targetSslProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "targetSslProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetSslProxies",
		},
		"targetTcpProxies_list": {
			PermissionMethod: joinString(Compute, "instances", "targetTcpProxies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/targetTcpProxies",
		},
		"targetVpnGateways_list": {
			PermissionMethod: joinString(Compute, "instances", "targetVpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/targetVpnGateways",
		},
		"urlMaps_list": {
			PermissionMethod: joinString(Compute, "instances", "urlMaps"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/global/urlMaps",
		},
		"vpnGateways_list": {
			PermissionMethod: joinString(Compute, "instances", "vpnGateways"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/vpnGateways",
		},
		"vpnTunnels_list": {
			PermissionMethod: joinString(Compute, "instances", "vpnTunnels"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/regions/{{.Region}}/vpnTunnels",
		},
		"zoneOperations_list": {
			PermissionMethod: joinString(Compute, "instances", "zoneOperations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://compute.googleapis.com/compute/v1/projects/{{.ParentID}}/zones/{{.Zone}}/operations",
		},
	},
	Container: {
		"usableSubnetworks_list": {
			PermissionMethod: joinString(Container, "projects", "aggregated", "usableSubnetworks"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/aggregated/usableSubnetworks",
		},
		"clusters_list": {
			PermissionMethod: joinString(Container, "projects", "zones", "clusters"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/zones/{{.Zone}}/clusters",
		},
		"operations_list": {
			PermissionMethod: joinString(Container, "projects", "zones", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://container.googleapis.com/v1/projects/{{.ParentID}}/zones/{{.Zone}}/operations",
		},
	},
	DataCatalog: {},
	DataFlow: {
		"jobs_list": {
			PermissionMethod: joinString(DataFlow, "projects", "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataflow.googleapis.com/v1b3/projects/{{.ParentID}}/jobs",
		},
	},
	DataFusion: {
		"autoscalingPolicies_list": {
			PermissionMethod: joinString(DataFusion, "projects", "regions", "autoscalingPolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/autoscalingPolicies",
		},
		"regions_clusters_list": {
			PermissionMethod: joinString(DataFusion, "projects", "regions", "clusters"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/clusters",
		},
		"regions_jobs_list": {
			PermissionMethod: joinString(DataFusion, "projects", "regions", "jobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/jobs",
		},
		"regions_operations_list": {
			PermissionMethod: joinString(DataFusion, "projects", "regions", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/operations",
		},
		"workflowTemplates_list": {
			PermissionMethod: joinString(DataFusion, "projects", "regions", "workflowTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dataproc.googleapis.com/v1/projects/{{.ParentID}}/regions/{{.Region}}/workflowTemplates",
		},
	},
	DataLabeling: {},
	DataPrep:     {},
	DataProc: {
		"locations_list": {
			PermissionMethod: joinString(DataProc, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://datafusion.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	DataStore: {
		"operations_list": {
			PermissionMethod: joinString(DataStore, "projects", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://datastore.googleapis.com/v1/projects/{{.ParentID}}/operations",
		},
	},
	DeploymentManager: {},
	DialogFlow:        {},
	Dlp: {
		"deidentifyTemplates_list": {
			PermissionMethod: joinString(Dlp, "deidentifyTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/deidentifyTemplates",
		},
		"dlpJobs_list": {
			PermissionMethod: joinString(Dlp, "dlpJobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/dlpJobs",
		},
		"inspectTemplates_list": {
			PermissionMethod: joinString(Dlp, "inspectTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/inspectTemplates",
		},
		"jobTriggers_list": {
			PermissionMethod: joinString(Dlp, "jobTriggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/jobTriggers",
		},
		"storedInfoTypes_list": {
			PermissionMethod: joinString(Dlp, "storedInfoTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/storedInfoTypes",
		},
		"region_deidentifyTemplates_list": {
			PermissionMethod: joinString(Dlp, "deidentifyTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/locations/{{.Region}}/deidentifyTemplates",
		},
		"region_dlpJobs_list": {
			PermissionMethod: joinString(Dlp, "dlpJobs"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/locations/{{.Region}}/dlpJobs",
		},
		"region_inspectTemplates_list": {
			PermissionMethod: joinString(Dlp, "inspectTemplates"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/locations/{{.Region}}/inspectTemplates",
		},
		"region_jobTriggers_list": {
			PermissionMethod: joinString(Dlp, "jobTriggers"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/locations/{{.Region}}/jobTriggers",
		},
		"region_storedInfoTypes_list": {
			PermissionMethod: joinString(Dlp, "storedInfoTypes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dlp.googleapis.com/v2/projects/{{.ParentID}}/locations/{{.Region}}/storedInfoTypes",
		},
	},
	Dns: {
		"managedZones_list": {
			PermissionMethod: joinString(Dns, "managedZones"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/managedZones",
		},
		"policies_list": {
			PermissionMethod: joinString(Dns, "policies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/policies",
		},
		"responsePolicies_list": {
			PermissionMethod: joinString(Dns, "responsePolicies"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://dns.googleapis.com/dns/v1/projects/{{.ParentID}}/responsePolicies",
		},
	},
	Endpoints:      {},
	ErrorReporting: {},
	File: {
		"locations_list": {
			PermissionMethod: joinString(File, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://file.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
		"locations_backups_list": {
			PermissionMethod: joinString(File, "projects", "locations", "backups"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://file.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/backups",
		},
		"locations_instances_list": {
			PermissionMethod: joinString(File, "projects", "locations", "instances"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://file.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/instances",
		},
		"locations_operations_list": {
			PermissionMethod: joinString(File, "projects", "locations", "operations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://file.googleapis.com/v1/projects/{{.ParentID}}/locations/{{.Region}}/operations",
		},
	},

	Firebase: {
		// firestore
		"list_collections": {
			PermissionMethod: joinString("firestore", "documents:listCollectionIds"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firestore.clients6.google.com/v1/projects/{{.ParentID}}/databases/(default)/documents:listCollectionIds",
		},
		// https://firebase.google.com/docs/reference/hosting/rest
		"sites_list": {
			PermissionMethod: "firebase.hosting.projects.sites",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebasehosting.googleapis.com/v1beta1/projects/{{.ParentID}}/sites",
		},
		"database_list": {
			PermissionMethod: joinString(Firebase, "firestore", "projects", "databases"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firestore.googleapis.com/v1/projects/{{.ParentID}}/databases",
		},
		"buckets_list": {
			PermissionMethod: joinString(Firebase, "cloudstorage", "projects", "buckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebasestorage.googleapis.com/v1beta/projects/{{.ParentID}}/buckets",
		},
		"releases_list": {
			PermissionMethod: joinString(Firebase, "rules", "projects", "releases"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaserules.googleapis.com/v1/projects/{{.ParentID}}/releases",
		},
		"rulesets_list": {
			PermissionMethod: joinString(Firebase, "rules", "projects", "rulesets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaserules.googleapis.com/v1/projects/{{.ParentID}}/rulesets",
		},
		"listVersions_list": {
			PermissionMethod: joinString(Firebase, "config", "remoteConfig", "listVersions"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaseremoteconfig.googleapis.com/v1/projects/{{.ParentID}}/remoteConfig:listVersions",
		},
		"distribution_list": {
			PermissionMethod: joinString(Firebase, "app", "distribution"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://firebaseappdistribution.googleapis.com/v1/projects/{{.ParentID}}/testers",
		},
	},
	Genomics:   {},
	Healthcare: {},
	Iam: {
		"serviceAccounts_list": {
			PermissionMethod: "iam.serviceAccounts",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/{{.ParentType}}/{{.ParentID}}/serviceAccounts",
		},
		"oranization_roles_list": {
			PermissionMethod: joinString(Iam, "organizations", "roles"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/{{.ParentType}}/{{.ParentID}}/roles",
		},
		"workloadIdentityPools_list": {
			PermissionMethod: joinString(Iam, "projects", "locations", "workloadIdentityPools"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/projects/{{ .ParentID }}/locations/{{ .Zone }}/workloadIdentityPools",
		},
		"roles_list": {
			PermissionMethod: joinString(Iam, "projects", "roles"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iam.googleapis.com/v1/projects/{{.ParentID}}/roles",
		},
	},
	Iap: {
		"brands_list": {
			PermissionMethod: joinString(Iap, "projects", "brands"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://iap.googleapis.com/v1/projects/{{.ParentID}}/brands",
		},
	},
	Logging: {
		"logEntries_list": {
			PermissionMethod: "logging.logEntries",
			Action:           "list",
			ReqMethod:        "POST",
			URL:              "https://logging.googleapis.com/v2/entries:list",
			ReqBody:          emptyBody,
		},
		"logs_list": {
			PermissionMethod: "logging.logs",
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/logs",
			ReqMethod:        "GET",
		},
		"exclusions_list": {
			PermissionMethod: joinString(Logging, "billingAccounts", "exclusions"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/exclusions",
			ReqMethod:        "GET",
		},
		"billingAccounts_logs_list": {
			PermissionMethod: joinString(Logging, "billingAccounts", "logs"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/logs",
			ReqMethod:        "GET",
		},
		"sinks_list": {
			PermissionMethod: joinString(Logging, "billingAccounts", "sinks"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/sinks",
			ReqMethod:        "GET",
		},
		"organization_exclusions_list": {
			PermissionMethod: joinString(Logging, "exclusions"),
			Action:           "list",
			URL:              "https://logging.googleapis.com/v2/{{.ParentType}}/{{.ParentID}}/exclusions",
			ReqMethod:        "GET",
		},
		"operation_list": {
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
		"policies_list": {
			PermissionMethod: joinString(OrgPolicy, "projects", "policies"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/projects/{{.ParentID}}/policies",
			ReqMethod:        "GET",
		},
		"constraints_list": {
			PermissionMethod: joinString(OrgPolicy, "projects", "constraints"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/projects/{{.ParentID}}/constraints",
			ReqMethod:        "GET",
		},
		"organization_constraints_list": {
			PermissionMethod: joinString(OrgPolicy, "organizations", "policies", "constraints"),
			Action:           "list",
			URL:              "https://orgpolicy.googleapis.com/v2/organizations/{{.ParentID}}/constraints",
			ReqMethod:        "GET",
		},
	},
	ProximityBeacon: {},
	Pubsub: {
		"subscriptions_list": {
			PermissionMethod: "pubsub.subscriptions",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/subscriptions",
		},
		"subscriptions_get": {
			PermissionMethod: "pubsub.subscriptions",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/subscriptions/{{.ResourceID}}",
		},
		"topics_list": {
			PermissionMethod: "pubsub.topics",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/topics",
		},
		"snapsshots_list": {
			PermissionMethod: joinString(Pubsub, "projects", "snapshots"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/snapshots",
		},
		"schemas_list": {
			PermissionMethod: joinString(Pubsub, "projects", "schemas"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/schemas",
		},
		"topics_get": {
			PermissionMethod: "pubsub.topics",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://pubsub.googleapis.com/v1/projects/{{.ParentID}}/topics/{{.ResourceID}}",
		},
	},
	Redis: {
		"locations_list": {
			PermissionMethod: joinString(Redis, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://redis.googleapis.com/v1/{{.ParentID}}/locations",
		},
	},
	RemoteBuildExecution: {},
	ResourceManager: {
		"liens_list": {
			PermissionMethod: joinString(ResourceManager, "liens"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudresourcemanager.googleapis.com/v1/liens?parent=projects/{{.ParentID}}",
		},
		"projects_list": {
			PermissionMethod: joinString(ResourceManager, "projects"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://cloudresourcemanager.googleapis.com/v1/projects",
		},
	},
	Run: {
		"revisions_list": {
			PermissionMethod: joinString(Run, "namespaces", "revisions"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/revisions",
		},
		"authorizeddomains_list": {
			PermissionMethod: joinString(Run, "namespaces", "authorizeddomains"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/authorizeddomains",
		},
		"configurations_list": {
			PermissionMethod: joinString(Run, "namespaces", "configurations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/configurations",
		},
		"domainmappings_list": {
			PermissionMethod: joinString(Run, "namespaces", "domainmappings"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/domainmappings",
		},
		"routes_list": {
			PermissionMethod: joinString(Run, "namespaces", "routes"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/routes",
		},
		"services_list": {
			PermissionMethod: joinString(Run, "namespaces", "services"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://{{.Region}}-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/{{.ParentID}}/services",
		},
		"locations_list": {
			PermissionMethod: joinString(Run, "namespaces", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://run.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
	},
	RuntimeConfig: {},
	SecretManager: {
		"locations_list": {
			PermissionMethod: joinString(SecretManager, "projects", "locations"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://secretmanager.googleapis.com/v1/projects/{{.ParentID}}/locations",
		},
		"secrets_list": {
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
		"repos_getConfig": {
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "getConfig",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/config",
		},
		"repos_list": {
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos",
		},
		"repos_get": {
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "get",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos/{{.ResourceID}}",
		},
		"repos_getIamPolicy": {
			PermissionMethod: "source.repos",
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://sourcerepo.googleapis.com/v1/projects/{{.ParentID}}/repos/{{.ResourceID}}:getIamPolicy",
		},
	},
	Spanner: {
		"instanceConfigs_getIamPolicy": {
			PermissionMethod: joinString(Spanner, "projects", "instanceConfigs"),
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://spanner.googleapis.com/v1/projects/{{.ParentID}}/instanceConfigs",
		},
		"instance_getIamPolicy": {
			PermissionMethod: joinString(Spanner, "projects", "instances"),
			ReqMethod:        "GET",
			Action:           "getIamPolicy",
			URL:              "https://spanner.googleapis.com/v1/projects/{{.ParentID}}/instances",
		},
	},
	Stackdriver: {},
	Storage: {
		"buckets_list": {
			PermissionMethod: joinString(Storage, "buckets"),
			ReqMethod:        "GET",
			Action:           "list",
			URL:              "https://storage.googleapis.com/storage/v1/b?project={{.ParentID}}",
		},
		"hmacKeys_list": {
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
