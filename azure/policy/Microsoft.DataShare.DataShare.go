package policy

// Microsoft_DataShare_DataShare policy
var Microsoft_DataShare_DataShare = map[string]Policy{
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataShare/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.DataShare",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.DataShare",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.DataShare",
	},
	"ConsumerInvitations_ListInvitations": {
		Path:   "/providers/Microsoft.DataShare/listInvitations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ConsumerInvitations_ListInvitations",
		Resource:    "Microsoft.DataShare",
	},
	"ConsumerInvitations_Get": {
		Path:   "/providers/Microsoft.DataShare/locations/{{.location}}/consumerInvitations/{{.invitationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ConsumerInvitations_Get",
		Resource:    "Microsoft.DataShare",
	},
	"ConsumerInvitations_RejectInvitation": {
		Path:   "/providers/Microsoft.DataShare/locations/{{.location}}/rejectInvitation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ConsumerInvitations_RejectInvitation",
		Resource:    "Microsoft.DataShare",
	},
	"DataSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/dataSets/{{.dataSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DataSets_Get",
		Resource:    "Microsoft.DataShare",
	},
	"DataSets_ListByShare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/dataSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DataSets_ListByShare",
		Resource:    "Microsoft.DataShare",
	},
	"DataSetMappings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/dataSetMappings/{{.dataSetMappingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DataSetMappings_Get",
		Resource:    "Microsoft.DataShare",
	},
	"DataSetMappings_ListByShareSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/dataSetMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DataSetMappings_ListByShareSubscription",
		Resource:    "Microsoft.DataShare",
	},
	"EmailRegistrations_ActivateEmail": {
		Path:   "/providers/Microsoft.DataShare/locations/{{.location}}/activateEmail",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "EmailRegistrations_ActivateEmail",
		Resource:    "Microsoft.DataShare",
	},
	"EmailRegistrations_RegisterEmail": {
		Path:   "/providers/Microsoft.DataShare/locations/{{.location}}/registerEmail",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "EmailRegistrations_RegisterEmail",
		Resource:    "Microsoft.DataShare",
	},
	"Invitations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/invitations/{{.invitationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Invitations_Get",
		Resource:    "Microsoft.DataShare",
	},
	"Invitations_ListByShare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/invitations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Invitations_ListByShare",
		Resource:    "Microsoft.DataShare",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DataShare/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DataShare",
	},
	"Shares_ListSynchronizationDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/listSynchronizationDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Shares_ListSynchronizationDetails",
		Resource:    "Microsoft.DataShare",
	},
	"Shares_ListSynchronizations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/listSynchronizations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Shares_ListSynchronizations",
		Resource:    "Microsoft.DataShare",
	},
	"ProviderShareSubscriptions_Adjust": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/providerShareSubscriptions/{{.providerShareSubscriptionId}}/adjust",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProviderShareSubscriptions_Adjust",
		Resource:    "Microsoft.DataShare",
	},
	"ProviderShareSubscriptions_Reinstate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/providerShareSubscriptions/{{.providerShareSubscriptionId}}/reinstate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProviderShareSubscriptions_Reinstate",
		Resource:    "Microsoft.DataShare",
	},
	"ProviderShareSubscriptions_Revoke": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/providerShareSubscriptions/{{.providerShareSubscriptionId}}/revoke",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProviderShareSubscriptions_Revoke",
		Resource:    "Microsoft.DataShare",
	},
	"ProviderShareSubscriptions_GetByShare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/providerShareSubscriptions/{{.providerShareSubscriptionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProviderShareSubscriptions_GetByShare",
		Resource:    "Microsoft.DataShare",
	},
	"ProviderShareSubscriptions_ListByShare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/providerShareSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProviderShareSubscriptions_ListByShare",
		Resource:    "Microsoft.DataShare",
	},
	"Shares_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Shares_Get",
		Resource:    "Microsoft.DataShare",
	},
	"Shares_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Shares_ListByAccount",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_CancelSynchronization": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/cancelSynchronization",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_CancelSynchronization",
		Resource:    "Microsoft.DataShare",
	},
	"ConsumerSourceDataSets_ListByShareSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/consumerSourceDataSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ConsumerSourceDataSets_ListByShareSubscription",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_ListSourceShareSynchronizationSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/listSourceShareSynchronizationSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_ListSourceShareSynchronizationSettings",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_ListSynchronizationDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/listSynchronizationDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_ListSynchronizationDetails",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_ListSynchronizations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/listSynchronizations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_ListSynchronizations",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_Synchronize": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/synchronize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_Synchronize",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_Get",
		Resource:    "Microsoft.DataShare",
	},
	"ShareSubscriptions_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ShareSubscriptions_ListByAccount",
		Resource:    "Microsoft.DataShare",
	},
	"SynchronizationSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/synchronizationSettings/{{.synchronizationSettingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "SynchronizationSettings_Get",
		Resource:    "Microsoft.DataShare",
	},
	"SynchronizationSettings_ListByShare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shares/{{.shareName}}/synchronizationSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "SynchronizationSettings_ListByShare",
		Resource:    "Microsoft.DataShare",
	},
	"Triggers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/triggers/{{.triggerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Triggers_Get",
		Resource:    "Microsoft.DataShare",
	},
	"Triggers_ListByShareSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataShare/accounts/{{.accountName}}/shareSubscriptions/{{.shareSubscriptionName}}/triggers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Triggers_ListByShareSubscription",
		Resource:    "Microsoft.DataShare",
	},
}
