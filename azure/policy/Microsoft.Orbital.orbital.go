package policy

// Microsoft_Orbital_orbital policy
var Microsoft_Orbital_orbital = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Orbital/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Orbital",
	},
	"Spacecrafts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Orbital/spacecrafts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Spacecrafts_ListBySubscription",
		Resource:    "Microsoft.Orbital",
	},
	"Spacecrafts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/spacecrafts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Spacecrafts_List",
		Resource:    "Microsoft.Orbital",
	},
	"Spacecrafts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/spacecrafts/{{.spacecraftName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Spacecrafts_Get",
		Resource:    "Microsoft.Orbital",
	},
	"Contacts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/spacecrafts/{{.spacecraftName}}/contacts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Contacts_List",
		Resource:    "Microsoft.Orbital",
	},
	"Contacts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/spacecrafts/{{.spacecraftName}}/contacts/{{.contactName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Contacts_Get",
		Resource:    "Microsoft.Orbital",
	},
	"Spacecrafts_ListAvailableContacts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/spacecrafts/{{.spacecraftName}}/listAvailableContacts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Spacecrafts_ListAvailableContacts",
		Resource:    "Microsoft.Orbital",
	},
	"ContactProfiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/contactProfiles/{{.contactProfileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContactProfiles_Get",
		Resource:    "Microsoft.Orbital",
	},
	"ContactProfiles_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Orbital/contactProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContactProfiles_ListBySubscription",
		Resource:    "Microsoft.Orbital",
	},
	"ContactProfiles_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Orbital/contactProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContactProfiles_List",
		Resource:    "Microsoft.Orbital",
	},
	"AvailableGroundStations_ListByCapability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Orbital/availableGroundStations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AvailableGroundStations_ListByCapability",
		Resource:    "Microsoft.Orbital",
	},
	"AvailableGroundStations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Orbital/availableGroundStations/{{.groundStationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AvailableGroundStations_Get",
		Resource:    "Microsoft.Orbital",
	},
	"OperationsResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Orbital/locations/{{.location}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "OperationsResults_Get",
		Resource:    "Microsoft.Orbital",
	},
}
