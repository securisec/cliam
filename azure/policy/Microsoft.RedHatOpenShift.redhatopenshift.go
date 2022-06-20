package policy

var Microsoft_RedHatOpenShift_redhatopenshift = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.RedHatOpenShift/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RedHatOpenShift/openShiftClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "OpenShiftClusters_List",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "OpenShiftClusters_ListByResourceGroup",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "OpenShiftClusters_Get",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListAdminCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}/listAdminCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "OpenShiftClusters_ListAdminCredentials",
		Resource:    "Microsoft.RedHatOpenShift",
	},
	"OpenShiftClusters_ListCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{{.resourceName}}/listCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "OpenShiftClusters_ListCredentials",
		Resource:    "Microsoft.RedHatOpenShift",
	},
}
