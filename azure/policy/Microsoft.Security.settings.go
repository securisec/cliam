package policy

// Microsoft_Security_settings policy
var Microsoft_Security_settings = map[string]Policy{
	"Settings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Security/settings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Settings_List",
		Resource:    "Microsoft.Security",
	},
	"Settings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Security/settings/{{.settingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Settings_Get",
		Resource:    "Microsoft.Security",
	},
}
