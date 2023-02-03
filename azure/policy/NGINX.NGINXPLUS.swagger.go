package policy

// NGINX_NGINXPLUS_swagger policy
var NGINX_NGINXPLUS_swagger = map[string]Policy{
	"Certificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments/{{.deploymentName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Certificates_Get",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Certificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments/{{.deploymentName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Certificates_List",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Configurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments/{{.deploymentName}}/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Configurations_List",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Configurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments/{{.deploymentName}}/configurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Configurations_Get",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Deployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Deployments_Get",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Deployments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Nginx.NginxPlus/nginxDeployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Deployments_List",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Deployments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Nginx.NginxPlus/nginxDeployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Deployments_ListByResourceGroup",
		Resource:    "NGINX.NGINXPLUS",
	},
	"Operations_List": {
		Path:   "/providers/Nginx.NginxPlus/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "NGINX.NGINXPLUS",
	},
}
