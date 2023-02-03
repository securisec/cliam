package policy

// Microsoft_Relay_authorizationRules policy
var Microsoft_Relay_authorizationRules = map[string]Policy{
	"Namespaces_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListAuthorizationRules",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_GetAuthorizationRule",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/authorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListKeys",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_RegenerateKeys",
		Resource:    "Microsoft.Relay",
	},
	"HybridConnections_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections/{{.hybridConnectionName}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_ListAuthorizationRules",
		Resource:    "Microsoft.Relay",
	},
	"HybridConnections_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections/{{.hybridConnectionName}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_GetAuthorizationRule",
		Resource:    "Microsoft.Relay",
	},
	"HybridConnections_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections/{{.hybridConnectionName}}/authorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_ListKeys",
		Resource:    "Microsoft.Relay",
	},
	"HybridConnections_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections/{{.hybridConnectionName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_RegenerateKeys",
		Resource:    "Microsoft.Relay",
	},
	"WCFRelays_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays/{{.relayName}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_ListAuthorizationRules",
		Resource:    "Microsoft.Relay",
	},
	"WCFRelays_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays/{{.relayName}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_GetAuthorizationRule",
		Resource:    "Microsoft.Relay",
	},
	"WCFRelays_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays/{{.relayName}}/authorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_ListKeys",
		Resource:    "Microsoft.Relay",
	},
	"WCFRelays_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays/{{.relayName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_RegenerateKeys",
		Resource:    "Microsoft.Relay",
	},
}
