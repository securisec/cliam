package policy

    var Microsoft_App_ManagedEnvironments = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/managedEnvironments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ManagedEnvironments_ListBySubscription",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ManagedEnvironments_ListByResourceGroup",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ManagedEnvironments_Get",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/certificates",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Certificates_List",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/certificates/{{.certificateName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Certificates_Get",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/checkNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Namespaces_CheckNameAvailability",
    Resource:       "Microsoft.App",
},
    }
    