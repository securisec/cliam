package policy

import "github.com/securisec/cliam/shared"

var SSMPolicies = map[string]Service{
	"DescribeActivations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeActivations",
		},
		Permission: "DescribeActivations",
	},
	"DescribeAvailablePatches": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeAvailablePatches",
		},
		Permission: "DescribeAvailablePatches",
	},
	"DescribeInventoryDeletions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInventoryDeletions",
		},
		Permission: "DescribeInventoryDeletions",
	},
	"DescribeMaintenanceWindows": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindows",
		},
		Permission: "DescribeMaintenanceWindows",
	},
	"DescribePatchBaselines": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribePatchBaselines",
		},
		Permission: "DescribePatchBaselines",
	},
	"DescribePatchGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribePatchGroups",
		},
		Permission: "DescribePatchGroups",
	},
	"GetDefaultPatchBaseline": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetDefaultPatchBaseline",
		},
		Permission: "GetDefaultPatchBaseline",
	},
	"GetInventorySchema": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetInventorySchema",
		},
		Permission: "GetInventorySchema",
	},
	"ListCommandInvocations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListCommandInvocations",
		},
		Permission: "ListCommandInvocations",
	},
	"ListCommands": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListCommands",
		},
		Permission: "ListCommands",
	},
	"ListComplianceItems": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListComplianceItems",
		},
		Permission: "ListComplianceItems",
	},
	"ListComplianceSummaries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListComplianceSummaries",
		},
		Permission: "ListComplianceSummaries",
	},
	"ListResourceComplianceSummaries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListResourceComplianceSummaries",
		},
		Permission: "ListResourceComplianceSummaries",
	},
	"ListResourceDataSync": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListResourceDataSync",
		},
		Permission: "ListResourceDataSync",
	},

	// extra
	"DescribeAssociationExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeAssociationExecutions",
		},
		Permission:             "DescribeAssociationExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AssociationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "association_id",
	},
	"DescribeAutomationStepExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeAutomationStepExecutions",
		},
		Permission:             "DescribeAutomationStepExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutomationExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "automation_execution_id",
	},
	"DescribeDocument": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeDocument",
		},
		Permission:             "DescribeDocument",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"DescribeEffectiveInstanceAssociations": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeEffectiveInstanceAssociations",
		},
		Permission:             "DescribeEffectiveInstanceAssociations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_id",
	},
	"DescribeEffectivePatchesForPatchBaseline": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeEffectivePatchesForPatchBaseline",
		},
		Permission:             "DescribeEffectivePatchesForPatchBaseline",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BaselineId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "baseline_id",
	},
	"DescribeInstanceAssociationsStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInstanceAssociationsStatus",
		},
		Permission:             "DescribeInstanceAssociationsStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_id",
	},
	"DescribeInstancePatchStates": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInstancePatchStates",
		},
		Permission:             "DescribeInstancePatchStates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceIds",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_ids",
	},
	"DescribeInstancePatchStatesForPatchGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInstancePatchStatesForPatchGroup",
		},
		Permission:             "DescribeInstancePatchStatesForPatchGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PatchGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "patch_group",
	},
	"DescribeInstancePatches": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInstancePatches",
		},
		Permission:             "DescribeInstancePatches",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_id",
	},
	"DescribeMaintenanceWindowExecutionTasks": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindowExecutionTasks",
		},
		Permission:             "DescribeMaintenanceWindowExecutionTasks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_execution_id",
	},
	"DescribeMaintenanceWindowExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindowExecutions",
		},
		Permission:             "DescribeMaintenanceWindowExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_id",
	},
	"DescribeMaintenanceWindowTargets": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindowTargets",
		},
		Permission:             "DescribeMaintenanceWindowTargets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_id",
	},
	"DescribeMaintenanceWindowTasks": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindowTasks",
		},
		Permission:             "DescribeMaintenanceWindowTasks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_id",
	},
	"DescribePatchGroupState": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribePatchGroupState",
		},
		Permission:             "DescribePatchGroupState",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PatchGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "patch_group",
	},
	"DescribeSessions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeSessions",
		},
		Permission:             "DescribeSessions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "State",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "state",
	},
	"GetAutomationExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetAutomationExecution",
		},
		Permission:             "GetAutomationExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutomationExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "automation_execution_id",
	},
	"GetCalendarState": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetCalendarState",
		},
		Permission:             "GetCalendarState",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CalendarNames",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "calendar_names",
	},
	"GetConnectionStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetConnectionStatus",
		},
		Permission:             "GetConnectionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Target",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "target",
	},
	"GetDocument": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetDocument",
		},
		Permission:             "GetDocument",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetMaintenanceWindow": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetMaintenanceWindow",
		},
		Permission:             "GetMaintenanceWindow",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_id",
	},
	"GetMaintenanceWindowExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetMaintenanceWindowExecution",
		},
		Permission:             "GetMaintenanceWindowExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WindowExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "window_execution_id",
	},
	"GetOpsItem": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetOpsItem",
		},
		Permission:             "GetOpsItem",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OpsItemId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "ops_item_id",
	},
	"GetOpsMetadata": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetOpsMetadata",
		},
		Permission:             "GetOpsMetadata",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OpsMetadataArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "ops_metadata_arn",
	},
	"GetParameter": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetParameter",
		},
		Permission:             "GetParameter",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetParameterHistory": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetParameterHistory",
		},
		Permission:             "GetParameterHistory",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetParameters": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetParameters",
		},
		Permission:             "GetParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Names",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "names",
	},
	"GetParametersByPath": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetParametersByPath",
		},
		Permission:             "GetParametersByPath",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Path",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "path",
	},
	"GetPatchBaseline": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetPatchBaseline",
		},
		Permission:             "GetPatchBaseline",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BaselineId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "baseline_id",
	},
	"GetPatchBaselineForPatchGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetPatchBaselineForPatchGroup",
		},
		Permission:             "GetPatchBaselineForPatchGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PatchGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "patch_group",
	},
	"GetServiceSetting": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetServiceSetting",
		},
		Permission:             "GetServiceSetting",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SettingId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "setting_id",
	},
	"ListAssociationVersions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListAssociationVersions",
		},
		Permission:             "ListAssociationVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AssociationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "association_id",
	},
	"ListDocumentVersions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListDocumentVersions",
		},
		Permission:             "ListDocumentVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
}
