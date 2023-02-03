package policy

// Microsoft_Attestation_attestation policy
var Microsoft_Attestation_attestation = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Attestation/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Attestation",
	},
	"AttestationProviders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Attestation/attestationProviders/{{.providerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "AttestationProviders_Get",
		Resource:    "Microsoft.Attestation",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Attestation/attestationProviders/{{.providerName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Attestation",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Attestation/attestationProviders/{{.providerName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Attestation",
	},
	"AttestationProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Attestation/attestationProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "AttestationProviders_List",
		Resource:    "Microsoft.Attestation",
	},
	"AttestationProviders_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Attestation/attestationProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "AttestationProviders_ListByResourceGroup",
		Resource:    "Microsoft.Attestation",
	},
	"AttestationProviders_ListDefault": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Attestation/defaultProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "AttestationProviders_ListDefault",
		Resource:    "Microsoft.Attestation",
	},
	"AttestationProviders_GetDefaultByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Attestation/locations/{{.location}}/defaultProvider",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "AttestationProviders_GetDefaultByLocation",
		Resource:    "Microsoft.Attestation",
	},
}
