package policy

// Microsoft_Network_applicationGateway policy
var Microsoft_Network_applicationGateway = map[string]Policy{
	"ApplicationGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_List",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAll",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_Start",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_Stop",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_BackendHealth": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/backendhealth",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_BackendHealth",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_BackendHealthOnDemand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/getBackendHealthOnDemand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_BackendHealthOnDemand",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGatewayPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGatewayPrivateLinkResources_List",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGatewayPrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/privateEndpointConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGatewayPrivateEndpointConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGatewayPrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationGateways/{{.applicationGatewayName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGatewayPrivateEndpointConnections_List",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableServerVariables": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableServerVariables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableServerVariables",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableRequestHeaders": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableRequestHeaders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableRequestHeaders",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableResponseHeaders": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableResponseHeaders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableResponseHeaders",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableWafRuleSets": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableWafRuleSets",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableSslOptions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableSslOptions",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_ListAvailableSslPredefinedPolicies": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_ListAvailableSslPredefinedPolicies",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGateways_GetSslPredefinedPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/{{.predefinedPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGateways_GetSslPredefinedPolicy",
		Resource:    "Microsoft.Network",
	},
}
