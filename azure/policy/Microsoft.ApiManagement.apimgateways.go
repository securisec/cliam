package policy

var Microsoft_ApiManagement_apimgateways = map[string]Policy{
	"Gateway_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Gateway_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Gateway_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Gateway_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Gateway_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Gateway_ListKeys",
		Resource:    "Microsoft.ApiManagement",
	},
	"Gateway_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Gateway_RegenerateKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"Gateway_GenerateToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/generateToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Gateway_GenerateToken",
		Resource:    "Microsoft.ApiManagement",
	},
	"GatewayHostnameConfiguration_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/hostnameConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GatewayHostnameConfiguration_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"GatewayHostnameConfiguration_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/hostnameConfigurations/{{.hcId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GatewayHostnameConfiguration_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"GatewayApi_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/apis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GatewayApi_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"GatewayCertificateAuthority_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/certificateAuthorities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GatewayCertificateAuthority_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"GatewayCertificateAuthority_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/gateways/{{.gatewayId}}/certificateAuthorities/{{.certificateId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GatewayCertificateAuthority_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
