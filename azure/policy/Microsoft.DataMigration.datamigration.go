package policy

// Microsoft_DataMigration_datamigration policy
var Microsoft_DataMigration_datamigration = map[string]Policy{
	"ResourceSkus_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataMigration/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "ResourceSkus_ListSkus",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_CheckStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/checkStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_CheckStatus",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_Start",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_Stop",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_ListSkus",
		Resource:    "Microsoft.DataMigration",
	},
	"Tasks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/tasks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Tasks_List",
		Resource:    "Microsoft.DataMigration",
	},
	"ServiceTasks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/serviceTasks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "ServiceTasks_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_CheckChildrenNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_CheckChildrenNameAvailability",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_ListByResourceGroup",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataMigration/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Tasks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/tasks/{{.taskName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Tasks_Get",
		Resource:    "Microsoft.DataMigration",
	},
	"ServiceTasks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/serviceTasks/{{.taskName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "ServiceTasks_Get",
		Resource:    "Microsoft.DataMigration",
	},
	"Tasks_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/tasks/{{.taskName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Tasks_Cancel",
		Resource:    "Microsoft.DataMigration",
	},
	"ServiceTasks_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/serviceTasks/{{.taskName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "ServiceTasks_Cancel",
		Resource:    "Microsoft.DataMigration",
	},
	"Tasks_Command": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/tasks/{{.taskName}}/command",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Tasks_Command",
		Resource:    "Microsoft.DataMigration",
	},
	"Projects_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Projects_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Projects_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Projects_Get",
		Resource:    "Microsoft.DataMigration",
	},
	"Services_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataMigration/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Services_CheckNameAvailability",
		Resource:    "Microsoft.DataMigration",
	},
	"Usages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataMigration/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DataMigration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Files_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/files",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Files_List",
		Resource:    "Microsoft.DataMigration",
	},
	"Files_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/files/{{.fileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Files_Get",
		Resource:    "Microsoft.DataMigration",
	},
	"Files_Read": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/files/{{.fileName}}/read",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Files_Read",
		Resource:    "Microsoft.DataMigration",
	},
	"Files_ReadWrite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.DataMigration/services/{{.serviceName}}/projects/{{.projectName}}/files/{{.fileName}}/readwrite",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-30",
		},
		OperationID: "Files_ReadWrite",
		Resource:    "Microsoft.DataMigration",
	},
}
