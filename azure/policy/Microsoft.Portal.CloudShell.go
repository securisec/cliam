package policy

var Microsoft_Portal_CloudShell = map[string]Policy{
	"getUserSettingsWithLocation": {
		Path:   "/providers/Microsoft.Portal/locations/{{.location}}/userSettings/{{.userSettingsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "getUserSettingsWithLocation",
		Resource:    "Microsoft.Portal",
	},
	"getConsoleWithLocation": {
		Path:   "/providers/Microsoft.Portal/locations/{{.location}}/consoles/{{.consoleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "getConsoleWithLocation",
		Resource:    "Microsoft.Portal",
	},
	"keepAliveWithLocation": {
		Path:   "/providers/Microsoft.Portal/locations/{{.location}}/consoles/{{.consoleName}}/keepAlive",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "keepAliveWithLocation",
		Resource:    "Microsoft.Portal",
	},
	"GetUserSettings": {
		Path:   "/providers/Microsoft.Portal/userSettings/{{.userSettingsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "GetUserSettings",
		Resource:    "Microsoft.Portal",
	},
	"GetConsole": {
		Path:   "/providers/Microsoft.Portal/consoles/{{.consoleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "GetConsole",
		Resource:    "Microsoft.Portal",
	},
	"KeepAlive": {
		Path:   "/providers/Microsoft.Portal/consoles/{{.consoleName}}/keepAlive",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-10-01",
		},
		OperationID: "KeepAlive",
		Resource:    "Microsoft.Portal",
	},
}
