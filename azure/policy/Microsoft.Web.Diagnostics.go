package policy

// Microsoft_Web_Diagnostics policy
var Microsoft_Web_Diagnostics = map[string]Policy{
	"Diagnostics_ListHostingEnvironmentDetectorResponses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListHostingEnvironmentDetectorResponses",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetHostingEnvironmentDetectorResponse": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetHostingEnvironmentDetectorResponse",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDetectorResponses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDetectorResponses",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDetectorResponse": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDetectorResponse",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDiagnosticCategories": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDiagnosticCategories",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDiagnosticCategory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDiagnosticCategory",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteAnalyses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/analyses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteAnalyses",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteAnalysis": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/analyses/{{.analysisName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteAnalysis",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ExecuteSiteAnalysis": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/analyses/{{.analysisName}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ExecuteSiteAnalysis",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDetectors": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDetectors",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDetector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDetector",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ExecuteSiteDetector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/diagnostics/{{.diagnosticCategory}}/detectors/{{.detectorName}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ExecuteSiteDetector",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDetectorResponsesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDetectorResponsesSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDetectorResponseSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDetectorResponseSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDiagnosticCategoriesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDiagnosticCategoriesSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDiagnosticCategorySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDiagnosticCategorySlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteAnalysesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/analyses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteAnalysesSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteAnalysisSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/analyses/{{.analysisName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteAnalysisSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ExecuteSiteAnalysisSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/analyses/{{.analysisName}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ExecuteSiteAnalysisSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ListSiteDetectorsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ListSiteDetectorsSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_GetSiteDetectorSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_GetSiteDetectorSlot",
		Resource:    "Microsoft.Web",
	},
	"Diagnostics_ExecuteSiteDetectorSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/slots/{{.slot}}/diagnostics/{{.diagnosticCategory}}/detectors/{{.detectorName}}/execute",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Diagnostics_ExecuteSiteDetectorSlot",
		Resource:    "Microsoft.Web",
	},
}
