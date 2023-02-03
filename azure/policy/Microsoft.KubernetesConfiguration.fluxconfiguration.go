package policy

// Microsoft_KubernetesConfiguration_fluxconfiguration policy
var Microsoft_KubernetesConfiguration_fluxconfiguration = map[string]Policy{
	"FluxConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{{.fluxConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "FluxConfigurations_Get",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"FluxConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "FluxConfigurations_List",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"FluxConfigOperationStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{{.fluxConfigurationName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "FluxConfigOperationStatus_Get",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
}
