package policy

// Microsoft_CostManagement_scheduledActions policy
var Microsoft_CostManagement_scheduledActions = map[string]Policy{
	"ScheduledActions_List": {
		Path:   "/providers/Microsoft.CostManagement/scheduledActions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_List",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_ListByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/scheduledActions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_ListByScope",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_Get": {
		Path:   "/providers/Microsoft.CostManagement/scheduledActions/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_Get",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_GetByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/scheduledActions/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_GetByScope",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_Run": {
		Path:   "/providers/Microsoft.CostManagement/scheduledActions/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_Run",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_RunByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/scheduledActions/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_RunByScope",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_CheckNameAvailability": {
		Path:   "/providers/Microsoft.CostManagement/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_CheckNameAvailability",
		Resource:    "Microsoft.CostManagement",
	},
	"ScheduledActions_CheckNameAvailabilityByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ScheduledActions_CheckNameAvailabilityByScope",
		Resource:    "Microsoft.CostManagement",
	},
}
