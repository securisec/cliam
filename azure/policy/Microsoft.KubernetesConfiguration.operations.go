package policy

// Microsoft_KubernetesConfiguration_operations policy
var Microsoft_KubernetesConfiguration_operations = map[string]Policy{
	"OperationStatus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "OperationStatus_List",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.KubernetesConfiguration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
}
