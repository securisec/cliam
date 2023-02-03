package policy

// Microsoft_RedHatOpenShift_redhatopenshift policy
var Microsoft_RedHatOpenShift_redhatopenshift = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.RedHatOpenShift/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RedHatOpenShift/locations/{{.location}}/openshiftversions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftVersions_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RedHatOpenShift/openShiftClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftClusters_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"MachinePools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{{.resourceName}}/machinePools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "MachinePools_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"Secrets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{{.resourceName}}/secrets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "Secrets_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"SyncIdentityProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{{.resourceName}}/syncIdentityProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "SyncIdentityProviders_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"SyncSets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{{.resourceName}}/syncSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "SyncSets_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftClusters_ListByResourceGroup",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftClusters_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListAdminCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}/listAdminCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftClusters_ListAdminCredentials",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}/listCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "OpenShiftClusters_ListCredentials",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"MachinePools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{{.resourceName}}/machinePool/{{.childResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "MachinePools_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"Secrets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{{.resourceName}}/secret/{{.childResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "Secrets_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"SyncIdentityProviders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{{.resourceName}}/syncIdentityProvider/{{.childResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "SyncIdentityProviders_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"SyncSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{{.resourceName}}/syncSet/{{.childResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-04",
		},
		OperationID: "SyncSets_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
}
