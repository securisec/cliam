package policy

var Microsoft_Automanage_automanage = map[string]Policy{
	"BestPractices_Get": {
		Path:   "/providers/Microsoft.Automanage/bestPractices/{{.bestPracticeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "BestPractices_Get",
		Resource:    "Microsoft.Automanage",
	},
	"BestPractices_ListByTenant": {
		Path:   "/providers/Microsoft.Automanage/bestPractices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "BestPractices_ListByTenant",
		Resource:    "Microsoft.Automanage",
	},
	"BestPracticesVersions_Get": {
		Path:   "/providers/Microsoft.Automanage/bestPractices/{{.bestPracticeName}}/versions/{{.versionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "BestPracticesVersions_Get",
		Resource:    "Microsoft.Automanage",
	},
	"BestPracticesVersions_ListByTenant": {
		Path:   "/providers/Microsoft.Automanage/bestPractices/{{.bestPracticeName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "BestPracticesVersions_ListByTenant",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automanage/configurationProfiles/{{.configurationProfileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfiles_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfiles_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automanage/configurationProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfiles_ListByResourceGroup",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfiles_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Automanage/configurationProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfiles_ListBySubscription",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfilesVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automanage/configurationProfiles/{{.configurationProfileName}}/versions/{{.versionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfilesVersions_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfilesVersions_ListChildResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automanage/configurationProfiles/{{.configurationProfileName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfilesVersions_ListChildResources",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_ListByVirtualMachines": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.Automanage/configurationProfileAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_ListByVirtualMachines",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automanage/configurationProfileAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_List",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Automanage/configurationProfileAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_ListBySubscription",
		Resource:    "Microsoft.Automanage",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Automanage/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Automanage",
	},
	"reports_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports/{{.reportName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "reports_Get",
		Resource:    "Microsoft.Automanage",
	},
	"reports_ListByConfigurationProfileAssignments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "reports_ListByConfigurationProfileAssignments",
		Resource:    "Microsoft.Automanage",
	},
	"ServicePrincipals_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Automanage/servicePrincipals",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ServicePrincipals_ListBySubscription",
		Resource:    "Microsoft.Automanage",
	},
	"ServicePrincipals_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Automanage/servicePrincipals/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ServicePrincipals_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileHCRPAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileHCRPAssignments_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_ListByMachineName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.Automanage/configurationProfileAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_ListByMachineName",
		Resource:    "Microsoft.Automanage",
	},
	"HCRPReports_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports/{{.reportName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "HCRPReports_Get",
		Resource:    "Microsoft.Automanage",
	},
	"HCRPReports_ListByConfigurationProfileAssignments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "HCRPReports_ListByConfigurationProfileAssignments",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileHCIAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHci/clusters/{{.clusterName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileHCIAssignments_Get",
		Resource:    "Microsoft.Automanage",
	},
	"ConfigurationProfileAssignments_ListByClusterName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHci/clusters/{{.clusterName}}/providers/Microsoft.Automanage/configurationProfileAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "ConfigurationProfileAssignments_ListByClusterName",
		Resource:    "Microsoft.Automanage",
	},
	"HCIReports_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHci/clusters/{{.clusterName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports/{{.reportName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "HCIReports_Get",
		Resource:    "Microsoft.Automanage",
	},
	"HCIReports_ListByConfigurationProfileAssignments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHci/clusters/{{.clusterName}}/providers/Microsoft.Automanage/configurationProfileAssignments/{{.configurationProfileAssignmentName}}/reports",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-04",
		},
		OperationID: "HCIReports_ListByConfigurationProfileAssignments",
		Resource:    "Microsoft.Automanage",
	},
}
