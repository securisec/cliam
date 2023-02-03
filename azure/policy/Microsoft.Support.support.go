package policy

// Microsoft_Support_support policy
var Microsoft_Support_support = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Support/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Support",
	},
	"Services_List": {
		Path:   "/providers/Microsoft.Support/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Services_List",
		Resource:    "Microsoft.Support",
	},
	"Services_Get": {
		Path:   "/providers/Microsoft.Support/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.Support",
	},
	"ProblemClassifications_List": {
		Path:   "/providers/Microsoft.Support/services/{{.serviceName}}/problemClassifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "ProblemClassifications_List",
		Resource:    "Microsoft.Support",
	},
	"ProblemClassifications_Get": {
		Path:   "/providers/Microsoft.Support/services/{{.serviceName}}/problemClassifications/{{.problemClassificationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "ProblemClassifications_Get",
		Resource:    "Microsoft.Support",
	},
	"SupportTickets_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "SupportTickets_CheckNameAvailability",
		Resource:    "Microsoft.Support",
	},
	"SupportTickets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/supportTickets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "SupportTickets_List",
		Resource:    "Microsoft.Support",
	},
	"SupportTickets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/supportTickets/{{.supportTicketName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "SupportTickets_Get",
		Resource:    "Microsoft.Support",
	},
	"Communications_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/supportTickets/{{.supportTicketName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Communications_CheckNameAvailability",
		Resource:    "Microsoft.Support",
	},
	"Communications_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/supportTickets/{{.supportTicketName}}/communications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Communications_List",
		Resource:    "Microsoft.Support",
	},
	"Communications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Support/supportTickets/{{.supportTicketName}}/communications/{{.communicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-04-01",
		},
		OperationID: "Communications_Get",
		Resource:    "Microsoft.Support",
	},
}
