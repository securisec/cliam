package policy

    var Microsoft_VirtualMachineImages_imagebuilder = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VirtualMachineImages/imageTemplates",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_List",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_ListByResourceGroup",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_Get",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/run",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_Run",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/cancel",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_Cancel",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/runOutputs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_ListRunOutputs",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/runOutputs/{{.runOutputName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "VirtualMachineImageTemplates_GetRunOutput",
    Resource:       "Microsoft.VirtualMachineImages",
},{
    Path: "/providers/Microsoft.VirtualMachineImages/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-02-14",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.VirtualMachineImages",
},
    }
    