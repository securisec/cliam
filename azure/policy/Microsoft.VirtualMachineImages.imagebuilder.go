package policy

var Microsoft_VirtualMachineImages_imagebuilder = map[string]Policy{
	"VirtualMachineImageTemplates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VirtualMachineImages/imageTemplates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_List",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_ListByResourceGroup",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_Get",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_Run": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_Run",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_Cancel",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_ListRunOutputs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/runOutputs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_ListRunOutputs",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"VirtualMachineImageTemplates_GetRunOutput": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VirtualMachineImages/imageTemplates/{{.imageTemplateName}}/runOutputs/{{.runOutputName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "VirtualMachineImageTemplates_GetRunOutput",
		Resource:    "Microsoft.VirtualMachineImages",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.VirtualMachineImages/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-14",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.VirtualMachineImages",
	},
}
