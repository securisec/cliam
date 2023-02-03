package policy

// Microsoft_ConfidentialLedger_confidentialledger policy
var Microsoft_ConfidentialLedger_confidentialledger = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.ConfidentialLedger/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-13",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ConfidentialLedger",
	},
	"CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ConfidentialLedger/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-13",
		},
		OperationID: "CheckNameAvailability",
		Resource:    "Microsoft.ConfidentialLedger",
	},
	"Ledger_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ConfidentialLedger/ledgers/{{.ledgerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-13",
		},
		OperationID: "Ledger_Get",
		Resource:    "Microsoft.ConfidentialLedger",
	},
	"Ledger_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ConfidentialLedger/ledgers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-13",
		},
		OperationID: "Ledger_ListByResourceGroup",
		Resource:    "Microsoft.ConfidentialLedger",
	},
	"Ledger_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ConfidentialLedger/ledgers/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-13",
		},
		OperationID: "Ledger_ListBySubscription",
		Resource:    "Microsoft.ConfidentialLedger",
	},
}
