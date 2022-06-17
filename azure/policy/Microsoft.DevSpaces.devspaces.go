package policy

    var Microsoft_DevSpaces_devspaces = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/locations/{{.location}}/checkContainerHostMapping",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "ContainerHostMappings_GetContainerHostMapping",
    Resource:       "Microsoft.DevSpaces",
},{
    Path: "/providers/Microsoft.DevSpaces/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.DevSpaces",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Controllers_Get",
    Resource:       "Microsoft.DevSpaces",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Controllers_ListByResourceGroup",
    Resource:       "Microsoft.DevSpaces",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevSpaces/controllers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Controllers_List",
    Resource:       "Microsoft.DevSpaces",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers/{{.name}}/listConnectionDetails",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Controllers_ListConnectionDetails",
    Resource:       "Microsoft.DevSpaces",
},
    }
    