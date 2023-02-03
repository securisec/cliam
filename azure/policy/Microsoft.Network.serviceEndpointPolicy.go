package policy

// Microsoft_Network_serviceEndpointPolicy policy
var Microsoft_Network_serviceEndpointPolicy = map[string]Policy{
	"ServiceEndpointPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/serviceEndpointPolicies/{{.serviceEndpointPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceEndpointPolicies_Get",
		Resource:    "Microsoft.Network",
	},
	"ServiceEndpointPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ServiceEndpointPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceEndpointPolicies_List",
		Resource:    "Microsoft.Network",
	},
	"ServiceEndpointPolicies_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/serviceEndpointPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceEndpointPolicies_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"ServiceEndpointPolicyDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/serviceEndpointPolicies/{{.serviceEndpointPolicyName}}/serviceEndpointPolicyDefinitions/{{.serviceEndpointPolicyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceEndpointPolicyDefinitions_Get",
		Resource:    "Microsoft.Network",
	},
	"ServiceEndpointPolicyDefinitions_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/serviceEndpointPolicies/{{.serviceEndpointPolicyName}}/serviceEndpointPolicyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceEndpointPolicyDefinitions_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
