package policy

// Microsoft_CognitiveServices_cognitiveservices policy
var Microsoft_CognitiveServices_cognitiveservices = map[string]Policy{
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.CognitiveServices",
	},
	"DeletedAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/locations/{{.location}}/resourceGroups/{{.resourceGroupName}}/deletedAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "DeletedAccounts_Get",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"DeletedAccounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/deletedAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "DeletedAccounts_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_ListKeys",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_RegenerateKey",
		Resource:    "Microsoft.CognitiveServices",
	},
	"ResourceSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "ResourceSkus_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_ListSkus",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_ListUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_ListUsages",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Accounts_ListModels": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/models",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Accounts_ListModels",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.CognitiveServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CheckSkuAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/locations/{{.location}}/checkSkuAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CheckSkuAvailability",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentTiers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/locations/{{.location}}/commitmentTiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentTiers_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CheckDomainAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/checkDomainAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CheckDomainAvailability",
		Resource:    "Microsoft.CognitiveServices",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.CognitiveServices",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Deployments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Deployments_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"Deployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Deployments_Get",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/commitmentPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_List",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/accounts/{{.accountName}}/commitmentPlans/{{.commitmentPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_Get",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_GetPlan": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/commitmentPlans/{{.commitmentPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_GetPlan",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_ListPlansByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/commitmentPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_ListPlansByResourceGroup",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_ListPlansBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CognitiveServices/commitmentPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_ListPlansBySubscription",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_ListAssociations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/commitmentPlans/{{.commitmentPlanName}}/accountAssociations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_ListAssociations",
		Resource:    "Microsoft.CognitiveServices",
	},
	"CommitmentPlans_GetAssociation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CognitiveServices/commitmentPlans/{{.commitmentPlanName}}/accountAssociations/{{.commitmentPlanAssociationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "CommitmentPlans_GetAssociation",
		Resource:    "Microsoft.CognitiveServices",
	},
}
