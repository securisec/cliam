package policy

    var Microsoft_StorageCache_storagecache = []Policy{
        {
    Path: "/providers/Microsoft.StorageCache/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageCache/skus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Skus_List",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageCache/usageModels",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "UsageModels_List",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageCache/locations/{{.location}}/ascOperations/{{.operationId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "AscOperations_Get",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageCache/locations/{{.location}}/usages",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "AscUsages_List",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageCache/caches",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_List",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_ListByResourceGroup",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_Get",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}/dnsRefresh",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTargets_DnsRefresh",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/debugInfo",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_DebugInfo",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/flush",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_Flush",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_Start",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_Stop",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTargets_ListByCache",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTargets_Get",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}/flush",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTarget_Flush",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}/suspend",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTarget_Suspend",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}/resume",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTarget_Resume",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/storageTargets/{{.storageTargetName}}/invalidate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "StorageTarget_Invalidate",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/startPrimingJob",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_StartPrimingJob",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/stopPrimingJob",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_StopPrimingJob",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/pausePrimingJob",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_PausePrimingJob",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/resumePrimingJob",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_ResumePrimingJob",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/upgrade",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_UpgradeFirmware",
    Resource:       "Microsoft.StorageCache",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StorageCache/caches/{{.cacheName}}/spaceAllocation",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Caches_SpaceAllocation",
    Resource:       "Microsoft.StorageCache",
},
    }
    