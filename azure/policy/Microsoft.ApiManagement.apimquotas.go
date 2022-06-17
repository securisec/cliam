package policy

var Microsoft_ApiManagement_apimquotas = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/quotas/{{.quotaCounterKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "QuotaByCounterKeys_ListByService",
		Resource:    "Microsoft.ApiManagement",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/quotas/{{.quotaCounterKey}}/periods/{{.quotaPeriodKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "QuotaByPeriodKeys_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
