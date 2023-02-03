package policy

// Microsoft_ADHybridHealthService_ADHybridHealthService policy
var Microsoft_ADHybridHealthService_ADHybridHealthService = map[string]Policy{
	"addsServices_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"alerts_listAddsAlerts": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "alerts_listAddsAlerts",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"configuration_listAddsConfigurations": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/configuration",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "configuration_listAddsConfigurations",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"dimensions_listAddsDimensions": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/dimensions/{{.dimension}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "dimensions_listAddsDimensions",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServiceMembers_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/addsservicemembers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServiceMembers_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"adDomainServiceMembers_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/addomainservicemembers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "adDomainServiceMembers_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServicesUserPreference_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/features/{{.featureName}}/userpreference",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServicesUserPreference_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServicesUserPreference_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/features/{{.featureName}}/userpreference",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServicesUserPreference_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_getForestSummary": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/forestsummary",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_getForestSummary",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsService_getMetrics": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsService_getMetrics",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listMetricsAverage": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}/average",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listMetricsAverage",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listMetricsSum": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}/sum",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listMetricsSum",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listMetricMetadata": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metricmetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listMetricMetadata",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_getMetricMetadata": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metricmetadata/{{.metricName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_getMetricMetadata",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_getMetricMetadataForGroup": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/metricmetadata/{{.metricName}}/groups/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_getMetricMetadataForGroup",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listReplicationDetails": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/replicationdetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listReplicationDetails",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServicesReplicationStatus_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/replicationstatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServicesReplicationStatus_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listReplicationSummary": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/replicationsummary",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listReplicationSummary",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServicesServiceMembers_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/servicemembers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServicesServiceMembers_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServicesServiceMembers_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/servicemembers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServicesServiceMembers_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServiceMembers_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/servicemembers/{{.serviceMemberId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServiceMembers_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listServerAlerts": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listServerAlerts",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServiceMembers_listCredentials": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/credentials",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServiceMembers_listCredentials",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"addsServices_listPremiumServices": {
		Path:   "/providers/Microsoft.ADHybridHealthService/addsservices/premiumCheck",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "addsServices_listPremiumServices",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"operations_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "operations_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"configuration_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/configuration",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "configuration_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"configuration_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/configuration",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "configuration_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"reports_getDevOps": {
		Path:   "/providers/Microsoft.ADHybridHealthService/reports/DevOps/IsDevOps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "reports_getDevOps",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listPremium": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/premiumCheck",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listPremium",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listAlerts": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listAlerts",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_getFeatureAvailibility": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/checkServiceFeatureAvailibility/{{.featureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_getFeatureAvailibility",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listExportErrors": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/exporterrors/counts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listExportErrors",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listExportErrorsV2": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/exporterrors/listV2",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listExportErrorsV2",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listExportStatus": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/exportstatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listExportStatus",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_addAlertFeedback": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/feedbacktype/alerts/feedback",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_addAlertFeedback",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listAlertFeedback": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/feedbacktype/alerts/{{.shortName}}/alertfeedback",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listAlertFeedback",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"service_getMetrics": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "service_getMetrics",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listMetricsAverage": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}/average",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listMetricsAverage",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listMetricsSum": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metrics/{{.metricName}}/groups/{{.groupName}}/sum",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listMetricsSum",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listMetricMetadata": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metricmetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listMetricMetadata",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_getMetricMetadata": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metricmetadata/{{.metricName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_getMetricMetadata",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_getMetricMetadataForGroup": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/metricmetadata/{{.metricName}}/groups/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_getMetricMetadataForGroup",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listMonitoringConfigurations": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/monitoringconfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listMonitoringConfigurations",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listUserBadPasswordReport": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/reports/badpassword/details/user",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listUserBadPasswordReport",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_list": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_list",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_add": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_add",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_get": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_get",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listAlerts": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listAlerts",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listConnectors": {
		Path:   "/providers/Microsoft.ADHybridHealthService/service/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/connectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listConnectors",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listCredentials": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/credentials",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listCredentials",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listDataFreshness": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/datafreshness",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listDataFreshness",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listExportStatus": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/exportstatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listExportStatus",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_listGlobalConfiguration": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/globalconfiguration",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_listGlobalConfiguration",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_getMetrics": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/metrics/{{.metricName}}/groups/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_getMetrics",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_getServiceConfiguration": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/serviceconfiguration",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_getServiceConfiguration",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_getTenantWhitelisting": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/TenantWhitelisting/{{.featureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_getTenantWhitelisting",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listAllRiskyIpDownloadReport": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/reports/riskyIp/blobUris",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listAllRiskyIpDownloadReport",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"services_listCurrentRiskyIpDownloadReport": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/reports/riskyIp/generateBlobUri",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "services_listCurrentRiskyIpDownloadReport",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"serviceMembers_getConnectorMetadata": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/servicemembers/{{.serviceMemberId}}/metrics/{{.metricName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "serviceMembers_getConnectorMetadata",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"list_IPAddressAggregatesByService": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/ipAddressAggregates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "list_IPAddressAggregatesByService",
		Resource:    "Microsoft.ADHybridHealthService",
	},
	"list_IPAddressAggregateSettings": {
		Path:   "/providers/Microsoft.ADHybridHealthService/services/{{.serviceName}}/ipAddressAggregateSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2014-01-01",
		},
		OperationID: "list_IPAddressAggregateSettings",
		Resource:    "Microsoft.ADHybridHealthService",
	},
}
