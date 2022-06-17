package policy

var Microsoft_DBforPostgreSQL_postgresql = map[string]Policy{
	"Servers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_Get",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Servers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_ListByResourceGroup",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Servers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforPostgreSQL/flexibleServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_List",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Servers_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_Restart",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Servers_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_Start",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Servers_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Servers_Stop",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"FirewallRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/firewallRules/{{.firewallRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "FirewallRules_Get",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"FirewallRules_ListByServer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "FirewallRules_ListByServer",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Configurations_ListByServer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Configurations_ListByServer",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Configurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/configurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Configurations_Get",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"CheckNameAvailability_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforPostgreSQL/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "CheckNameAvailability_Execute",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"LocationBasedCapabilities_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforPostgreSQL/locations/{{.locationName}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "LocationBasedCapabilities_Execute",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"VirtualNetworkSubnetUsage_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforPostgreSQL/locations/{{.locationName}}/checkVirtualNetworkSubnetUsage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "VirtualNetworkSubnetUsage_Execute",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DBforPostgreSQL/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
}
