package policy

// Microsoft_HardwareSecurityModules_dedicatedhsm policy
var Microsoft_HardwareSecurityModules_dedicatedhsm = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.HardwareSecurityModules/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-30",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HardwareSecurityModules",
	},
	"DedicatedHsm_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-30",
		},
		OperationID: "DedicatedHsm_Get",
		Resource:    "Microsoft.HardwareSecurityModules",
	},
	"DedicatedHsm_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-30",
		},
		OperationID: "DedicatedHsm_ListByResourceGroup",
		Resource:    "Microsoft.HardwareSecurityModules",
	},
	"DedicatedHsm_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-30",
		},
		OperationID: "DedicatedHsm_ListBySubscription",
		Resource:    "Microsoft.HardwareSecurityModules",
	},
	"DedicatedHsm_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{{.name}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-30",
		},
		OperationID: "DedicatedHsm_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.HardwareSecurityModules",
	},
}
