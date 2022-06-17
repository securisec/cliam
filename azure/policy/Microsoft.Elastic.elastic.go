package policy

var Microsoft_Elastic_elastic = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Elastic/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Elastic",
	},
	"Monitors_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Elastic/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "Monitors_List",
		Resource:    "Microsoft.Elastic",
	},
	"Monitors_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "Monitors_ListByResourceGroup",
		Resource:    "Microsoft.Elastic",
	},
	"Monitors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "Monitors_Get",
		Resource:    "Microsoft.Elastic",
	},
	"MonitoredResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/listMonitoredResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "MonitoredResources_List",
		Resource:    "Microsoft.Elastic",
	},
	"DeploymentInfo_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/listDeploymentInfo",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "DeploymentInfo_List",
		Resource:    "Microsoft.Elastic",
	},
	"TagRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/tagRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "TagRules_List",
		Resource:    "Microsoft.Elastic",
	},
	"TagRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/tagRules/{{.ruleSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "TagRules_Get",
		Resource:    "Microsoft.Elastic",
	},
	"VMHost_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/listVMHost",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "VMHost_List",
		Resource:    "Microsoft.Elastic",
	},
	"VMIngestion_Details": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/vmIngestionDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "VMIngestion_Details",
		Resource:    "Microsoft.Elastic",
	},
	"VMCollection_Update": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Elastic/monitors/{{.monitorName}}/vmCollectionUpdate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-01",
		},
		OperationID: "VMCollection_Update",
		Resource:    "Microsoft.Elastic",
	},
}
