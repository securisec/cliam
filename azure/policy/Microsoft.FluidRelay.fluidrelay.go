package policy

// Microsoft_FluidRelay_fluidrelay policy
var Microsoft_FluidRelay_fluidrelay = map[string]Policy{
	"FluidRelayOperations_List": {
		Path:   "/providers/Microsoft.FluidRelay/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayOperations_List",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayServers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers/{{.fluidRelayServerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayServers_Get",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayServers_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers/{{.fluidRelayServerName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayServers_RegenerateKey",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayServers_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers/{{.fluidRelayServerName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayServers_ListKeys",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayServers_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.FluidRelay/fluidRelayServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayServers_ListBySubscription",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayServers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayServers_ListByResourceGroup",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers/{{.fluidRelayServerName}}/fluidRelayContainers/{{.fluidRelayContainerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayContainers_Get",
		Resource:    "Microsoft.FluidRelay",
	},
	"FluidRelayContainers_ListByFluidRelayServers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.FluidRelay/fluidRelayServers/{{.fluidRelayServerName}}/fluidRelayContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "FluidRelayContainers_ListByFluidRelayServers",
		Resource:    "Microsoft.FluidRelay",
	},
}
