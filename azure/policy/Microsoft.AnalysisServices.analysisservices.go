package policy

// Microsoft_AnalysisServices_analysisservices policy
var Microsoft_AnalysisServices_analysisservices = map[string]Policy{
	"Servers_GetDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_GetDetails",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_Suspend": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}/suspend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_Suspend",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_Resume",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListByResourceGroup",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AnalysisServices/servers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_List",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListSkusForNew": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AnalysisServices/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListSkusForNew",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListSkusForExisting": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListSkusForExisting",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListGatewayStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}/listGatewayStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListGatewayStatus",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_DissociateGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AnalysisServices/servers/{{.serverName}}/dissociateGateway",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_DissociateGateway",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AnalysisServices/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_CheckNameAvailability",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListOperationResults": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AnalysisServices/locations/{{.location}}/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListOperationResults",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Servers_ListOperationStatuses": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AnalysisServices/locations/{{.location}}/operationstatuses/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Servers_ListOperationStatuses",
		Resource:    "Microsoft.AnalysisServices",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.AnalysisServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AnalysisServices",
	},
}
