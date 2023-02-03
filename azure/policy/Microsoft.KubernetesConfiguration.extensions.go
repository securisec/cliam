package policy

// Microsoft_KubernetesConfiguration_extensions policy
var Microsoft_KubernetesConfiguration_extensions = map[string]Policy{
	"Extensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/extensions/{{.extensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Extensions_Get",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"Extensions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Extensions_List",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"OperationStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/extensions/{{.extensionName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "OperationStatus_Get",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
}
