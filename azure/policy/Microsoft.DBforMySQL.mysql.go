package policy

var Microsoft_DBforMySQL_mysql = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_Get",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_ListByResourceGroup",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforMySQL/flexibleServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_List",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/failover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_Failover",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_Restart",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_Start",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Servers_Stop",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/replicas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Replicas_ListByServer",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/backups/{{.backupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Backups_Get",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Backups_ListByServer",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/firewallRules/{{.firewallRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "FirewallRules_Get",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "FirewallRules_ListByServer",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/databases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Databases_Get",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/databases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Databases_ListByServer",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/configurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Configurations_Get",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/updateConfigurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Configurations_BatchUpdate",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMySQL/flexibleServers/{{.serverName}}/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Configurations_ListByServer",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforMySQL/locations/{{.locationName}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "LocationBasedCapabilities_List",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforMySQL/locations/{{.locationName}}/checkVirtualNetworkSubnetUsage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "CheckVirtualNetworkSubnetUsage_Execute",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DBforMySQL/locations/{{.locationName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "CheckNameAvailability_Execute",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/providers/Microsoft.DBforMySQL/getPrivateDnsZoneSuffix",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "GetPrivateDnsZoneSuffix_Execute",
		Resource:    "Microsoft.DBforMySQL",
	}, {
		Path:   "/providers/Microsoft.DBforMySQL/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DBforMySQL",
	},
}
