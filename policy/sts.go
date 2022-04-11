package policy

var STS = []Service{
	{
		ServiceSuffix: "",
		Permission:    "GetCallerIdentity",
		FormData: map[string]string{
			"Action":  "GetCallerIdentity",
			"Version": "2011-06-15",
		},
	},
	{
		ServiceSuffix: "?Action=GetSessionToken&Version=2011-06-15",
		Permission:    "GetSessionToken",
	},
}
