package policy

import "github.com/securisec/cliam/shared"

// BedrockPolicies policy
var BedrockPolicies = map[string]Service{
	"GetModelInvocationLoggingConfiguration": {
		Method:        "GET",
		ServiceSuffix: "logging/modelinvocations",
		Permission:    "GetModelInvocationLoggingConfiguration",
	},
	"ListCustomModels": {
		Method:        "GET",
		ServiceSuffix: "custom-models",
		Permission:    "ListCustomModels",
	},
	"ListEvaluationJobs": {
		Method:        "GET",
		ServiceSuffix: "evaluation-jobs",
		Permission:    "ListEvaluationJobs",
	},
	"ListFoundationModels": {
		Method:        "GET",
		ServiceSuffix: "foundation-models",
		Permission:    "ListFoundationModels",
	},
	"ListGuardrails": {
		Method:        "GET",
		ServiceSuffix: "guardrails",
		Permission:    "ListGuardrails",
	},
	"ListModelCopyJobs": {
		Method:        "GET",
		ServiceSuffix: "model-copy-jobs",
		Permission:    "ListModelCopyJobs",
	},
	"ListModelCustomizationJobs": {
		Method:        "GET",
		ServiceSuffix: "model-customization-jobs",
		Permission:    "ListModelCustomizationJobs",
	},
	"ListProvisionedModelThroughputs": {
		Method:        "GET",
		ServiceSuffix: "provisioned-model-throughputs",
		Permission:    "ListProvisionedModelThroughputs",
	},

	// extra
	"GetCustomModel": {
		ServiceSuffix:          "/custom-models/{{.model_identifier}}",
		Permission:             "GetCustomModel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "model_identifier",
	},
	"GetEvaluationJob": {
		ServiceSuffix:          "/evaluation-jobs/{{.job_identifier}}",
		Permission:             "GetEvaluationJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_identifier",
	},
	"GetFoundationModel": {
		ServiceSuffix:          "/foundation-models/{{.model_identifier}}",
		Permission:             "GetFoundationModel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "model_identifier",
	},
	"GetGuardrail": {
		ServiceSuffix:          "/guardrails/{{.guardrail_identifier}}",
		Permission:             "GetGuardrail",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "guardrail_identifier",
	},
	"GetModelCopyJob": {
		ServiceSuffix:          "/model-copy-jobs/{{.job_arn}}",
		Permission:             "GetModelCopyJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_arn",
	},
	"GetModelCustomizationJob": {
		ServiceSuffix:          "/model-customization-jobs/{{.job_identifier}}",
		Permission:             "GetModelCustomizationJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_identifier",
	},
	"GetProvisionedModelThroughput": {
		ServiceSuffix:          "/provisioned-model-throughput/{{.provisioned_model_id}}",
		Permission:             "GetProvisionedModelThroughput",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "provisioned_model_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2023-04-20",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceARN",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_a_r_n",
	},
}
