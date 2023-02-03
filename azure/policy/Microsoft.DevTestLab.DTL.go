package policy

// Microsoft_DevTestLab_DTL policy
var Microsoft_DevTestLab_DTL = map[string]Policy{
	"ProviderOperations_List": {
		Path:   "/providers/Microsoft.DevTestLab/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ProviderOperations_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevTestLab/labs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ListBySubscription",
		Resource:    "Microsoft.DevTestLab",
	},
	"Operations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevTestLab/locations/{{.locationName}}/operations/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Operations_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"GlobalSchedules_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevTestLab/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GlobalSchedules_ListBySubscription",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ListByResourceGroup",
		Resource:    "Microsoft.DevTestLab",
	},
	"ArtifactSources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ArtifactSources_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"ArmTemplates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.artifactSourceName}}/armtemplates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ArmTemplates_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"ArmTemplates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.artifactSourceName}}/armtemplates/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ArmTemplates_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Artifacts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.artifactSourceName}}/artifacts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Artifacts_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Artifacts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.artifactSourceName}}/artifacts/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Artifacts_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Artifacts_GenerateArmTemplate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.artifactSourceName}}/artifacts/{{.name}}/generateArmTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Artifacts_GenerateArmTemplate",
		Resource:    "Microsoft.DevTestLab",
	},
	"ArtifactSources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/artifactsources/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ArtifactSources_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Costs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/costs/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Costs_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"CustomImages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/customimages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "CustomImages_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"CustomImages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/customimages/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "CustomImages_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Formulas_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/formulas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Formulas_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Formulas_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/formulas/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Formulas_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"GalleryImages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/galleryimages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GalleryImages_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"NotificationChannels_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/notificationchannels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "NotificationChannels_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"NotificationChannels_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/notificationchannels/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "NotificationChannels_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"NotificationChannels_Notify": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/notificationchannels/{{.name}}/notify",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "NotificationChannels_Notify",
		Resource:    "Microsoft.DevTestLab",
	},
	"PolicySets_EvaluatePolicies": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/policysets/{{.name}}/evaluatePolicies",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "PolicySets_EvaluatePolicies",
		Resource:    "Microsoft.DevTestLab",
	},
	"Policies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/policysets/{{.policySetName}}/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Policies_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Policies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/policysets/{{.policySetName}}/policies/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Policies_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Schedules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Schedules_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Schedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/schedules/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Schedules_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Schedules_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/schedules/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Schedules_Execute",
		Resource:    "Microsoft.DevTestLab",
	},
	"Schedules_ListApplicable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/schedules/{{.name}}/listApplicable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Schedules_ListApplicable",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceRunners_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/servicerunners/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceRunners_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Users_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Users_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Users_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Users_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Disks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/disks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Disks_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Disks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/disks/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Disks_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Disks_Attach": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/disks/{{.name}}/attach",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Disks_Attach",
		Resource:    "Microsoft.DevTestLab",
	},
	"Disks_Detach": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/disks/{{.name}}/detach",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Disks_Detach",
		Resource:    "Microsoft.DevTestLab",
	},
	"Environments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/environments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Environments_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Environments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/environments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Environments_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Secrets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/secrets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Secrets_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"Secrets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/secrets/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Secrets_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabrics_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabrics_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabrics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabrics_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabrics_ListApplicableSchedules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.name}}/listApplicableSchedules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabrics_ListApplicableSchedules",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabrics_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.name}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabrics_Start",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabrics_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.name}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabrics_Stop",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabricSchedules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.serviceFabricName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabricSchedules_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabricSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.serviceFabricName}}/schedules/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabricSchedules_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"ServiceFabricSchedules_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/users/{{.userName}}/servicefabrics/{{.serviceFabricName}}/schedules/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "ServiceFabricSchedules_Execute",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_AddDataDisk": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/addDataDisk",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_AddDataDisk",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_ApplyArtifacts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/applyArtifacts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_ApplyArtifacts",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Claim": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/claim",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Claim",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_DetachDataDisk": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/detachDataDisk",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_DetachDataDisk",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_GetRdpFileContents": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/getRdpFileContents",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_GetRdpFileContents",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_ListApplicableSchedules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/listApplicableSchedules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_ListApplicableSchedules",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Redeploy",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Resize": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/resize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Resize",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Restart",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Start",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_Stop",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_TransferDisks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/transferDisks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_TransferDisks",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachines_UnClaim": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.name}}/unClaim",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachines_UnClaim",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachineSchedules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.virtualMachineName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachineSchedules_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachineSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.virtualMachineName}}/schedules/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachineSchedules_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualMachineSchedules_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualmachines/{{.virtualMachineName}}/schedules/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualMachineSchedules_Execute",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualNetworks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualnetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualNetworks_List",
		Resource:    "Microsoft.DevTestLab",
	},
	"VirtualNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.labName}}/virtualnetworks/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "VirtualNetworks_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ClaimAnyVm": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/claimAnyVm",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ClaimAnyVm",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_CreateEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/createEnvironment",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_CreateEnvironment",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ExportResourceUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/exportResourceUsage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ExportResourceUsage",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_GenerateUploadUri": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/generateUploadUri",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_GenerateUploadUri",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ImportVirtualMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/importVirtualMachine",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ImportVirtualMachine",
		Resource:    "Microsoft.DevTestLab",
	},
	"Labs_ListVhds": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/labs/{{.name}}/listVhds",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "Labs_ListVhds",
		Resource:    "Microsoft.DevTestLab",
	},
	"GlobalSchedules_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GlobalSchedules_ListByResourceGroup",
		Resource:    "Microsoft.DevTestLab",
	},
	"GlobalSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/schedules/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GlobalSchedules_Get",
		Resource:    "Microsoft.DevTestLab",
	},
	"GlobalSchedules_Execute": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/schedules/{{.name}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GlobalSchedules_Execute",
		Resource:    "Microsoft.DevTestLab",
	},
	"GlobalSchedules_Retarget": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevTestLab/schedules/{{.name}}/retarget",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-09-15",
		},
		OperationID: "GlobalSchedules_Retarget",
		Resource:    "Microsoft.DevTestLab",
	},
}
