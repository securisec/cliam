package policy

// S3OutpostsPolicies policy
var S3OutpostsPolicies = map[string]Service{
	"ListEndpoints": {
		Method:        "GET",
		ServiceSuffix: "S3Outposts/ListEndpoints",
		Permission:    "ListEndpoints",
	},

	// extra
	"ListSharedEndpoints": {
		ServiceSuffix:          "/S3Outposts/ListSharedEndpoints",
		Permission:             "ListSharedEndpoints",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_id",
	},
}
