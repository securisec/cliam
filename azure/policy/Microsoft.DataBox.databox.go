package policy

    var Microsoft_DataBox_databox = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.DataBox/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.DataBox",
},
"Jobs_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBox/jobs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_List",
    Resource:       "Microsoft.DataBox",
},
"Mitigate": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}/mitigate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Mitigate",
    Resource:       "Microsoft.DataBox",
},
"Jobs_MarkDevicesShipped": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}/markDevicesShipped",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_MarkDevicesShipped",
    Resource:       "Microsoft.DataBox",
},
"Service_ListAvailableSkusByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/locations/{{.location}}/availableSkus",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_ListAvailableSkusByResourceGroup",
    Resource:       "Microsoft.DataBox",
},
"Service_ValidateAddress": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBox/locations/{{.location}}/validateAddress",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_ValidateAddress",
    Resource:       "Microsoft.DataBox",
},
"Service_ValidateInputsByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/locations/{{.location}}/validateInputs",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_ValidateInputsByResourceGroup",
    Resource:       "Microsoft.DataBox",
},
"Service_ValidateInputs": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBox/locations/{{.location}}/validateInputs",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_ValidateInputs",
    Resource:       "Microsoft.DataBox",
},
"Jobs_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_ListByResourceGroup",
    Resource:       "Microsoft.DataBox",
},
"Jobs_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_Get",
    Resource:       "Microsoft.DataBox",
},
"Jobs_BookShipmentPickUp": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}/bookShipmentPickUp",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_BookShipmentPickUp",
    Resource:       "Microsoft.DataBox",
},
"Jobs_Cancel": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}/cancel",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_Cancel",
    Resource:       "Microsoft.DataBox",
},
"Jobs_ListCredentials": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/jobs/{{.jobName}}/listCredentials",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Jobs_ListCredentials",
    Resource:       "Microsoft.DataBox",
},
"Service_RegionConfiguration": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBox/locations/{{.location}}/regionConfiguration",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_RegionConfiguration",
    Resource:       "Microsoft.DataBox",
},
"Service_RegionConfigurationByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBox/locations/{{.location}}/regionConfiguration",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-01",
    },
	OperationID:    "Service_RegionConfigurationByResourceGroup",
    Resource:       "Microsoft.DataBox",
},
    }
    