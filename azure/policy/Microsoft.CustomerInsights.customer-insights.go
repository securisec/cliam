package policy

// Microsoft_CustomerInsights_customer_insights policy
var Microsoft_CustomerInsights_customer_insights = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.CustomerInsights/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Hubs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Hubs_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Hubs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Hubs_ListByResourceGroup",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Hubs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CustomerInsights/hubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Hubs_List",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Profiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/profiles/{{.profileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Profiles_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Profiles_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/profiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Profiles_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Profiles_GetEnrichingKpis": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/profiles/{{.profileName}}/getEnrichingKpis",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Profiles_GetEnrichingKpis",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Interactions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/interactions/{{.interactionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Interactions_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Interactions_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/interactions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Interactions_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Interactions_SuggestRelationshipLinks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/interactions/{{.interactionName}}/suggestRelationshipLinks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Interactions_SuggestRelationshipLinks",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Relationships_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/relationships/{{.relationshipName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Relationships_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Relationships_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/relationships",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Relationships_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"RelationshipLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/relationshipLinks/{{.relationshipLinkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "RelationshipLinks_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"RelationshipLinks_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/relationshipLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "RelationshipLinks_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"AuthorizationPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/authorizationPolicies/{{.authorizationPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "AuthorizationPolicies_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"AuthorizationPolicies_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/authorizationPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "AuthorizationPolicies_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"AuthorizationPolicies_RegeneratePrimaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/authorizationPolicies/{{.authorizationPolicyName}}/regeneratePrimaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "AuthorizationPolicies_RegeneratePrimaryKey",
		Resource:    "Microsoft.CustomerInsights",
	},
	"AuthorizationPolicies_RegenerateSecondaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/authorizationPolicies/{{.authorizationPolicyName}}/regenerateSecondaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "AuthorizationPolicies_RegenerateSecondaryKey",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Connectors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/connectors/{{.connectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Connectors_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Connectors_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/connectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Connectors_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"ConnectorMappings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/connectors/{{.connectorName}}/mappings/{{.mappingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "ConnectorMappings_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"ConnectorMappings_ListByConnector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/connectors/{{.connectorName}}/mappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "ConnectorMappings_ListByConnector",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Kpi_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/kpi/{{.kpiName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Kpi_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Kpi_Reprocess": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/kpi/{{.kpiName}}/reprocess",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Kpi_Reprocess",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Kpi_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/kpi",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Kpi_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"WidgetTypes_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/widgetTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "WidgetTypes_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"WidgetTypes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/widgetTypes/{{.widgetTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "WidgetTypes_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Views_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/views",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Views_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Views_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/views/{{.viewName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Views_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Links_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/links/{{.linkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Links_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Links_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/links",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Links_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Roles_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/roles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Roles_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"RoleAssignments_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/roleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "RoleAssignments_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
	"RoleAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/roleAssignments/{{.assignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "RoleAssignments_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Images_GetUploadUrlForEntityType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/images/getEntityTypeImageUploadUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Images_GetUploadUrlForEntityType",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Images_GetUploadUrlForData": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/images/getDataImageUploadUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Images_GetUploadUrlForData",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Predictions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/predictions/{{.predictionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Predictions_Get",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Predictions_GetTrainingResults": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/predictions/{{.predictionName}}/getTrainingResults",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Predictions_GetTrainingResults",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Predictions_GetModelStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/predictions/{{.predictionName}}/getModelStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Predictions_GetModelStatus",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Predictions_ModelStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/predictions/{{.predictionName}}/modelStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Predictions_ModelStatus",
		Resource:    "Microsoft.CustomerInsights",
	},
	"Predictions_ListByHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CustomerInsights/hubs/{{.hubName}}/predictions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-26",
		},
		OperationID: "Predictions_ListByHub",
		Resource:    "Microsoft.CustomerInsights",
	},
}
