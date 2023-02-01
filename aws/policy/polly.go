package policy

// PollyPolicies policies
var PollyPolicies = map[string]Service{
	"DescribeVoices": {
		Method:        "GET",
		ServiceSuffix: "v1/voices",
		Permission:    "DescribeVoices",
	},
	"ListLexicons": {
		Method:        "GET",
		ServiceSuffix: "v1/lexicons",
		Permission:    "ListLexicons",
	},
	"ListSpeechSynthesisTasks": {
		Method:        "GET",
		ServiceSuffix: "v1/synthesisTasks",
		Permission:    "ListSpeechSynthesisTasks",
	},

	// extra
	"GetLexicon": {
		ServiceSuffix:          "/v1/lexicons/{{.name}}",
		Permission:             "GetLexicon",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"GetSpeechSynthesisTask": {
		ServiceSuffix:          "/v1/synthesisTasks/{{.task_id}}",
		Permission:             "GetSpeechSynthesisTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "task_id",
	},
}
