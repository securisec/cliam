package policy

// Microsoft_KubernetesConfiguration_kubernetesconfiguration policy
var Microsoft_KubernetesConfiguration_kubernetesconfiguration = map[string]Policy{
	"SourceControlConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{{.sourceControlConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SourceControlConfigurations_Get",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
	"SourceControlConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.clusterRp}}/{{.clusterResourceName}}/{{.clusterName}}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SourceControlConfigurations_List",
		Resource:    "Microsoft.KubernetesConfiguration",
	},
}
