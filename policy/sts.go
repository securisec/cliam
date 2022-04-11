package policy

var STS = []Service{
	{
		ServiceSuffix: "",
		Permission:    "GetCallerIdentity",
		FormData: map[string]string{
			"Action":  "GetCallerIdentity",
			"Version": "2011-06-15",
		},
		Method: "POST",
	},
	{
		ServiceSuffix: "?Action=GetSessionToken&Version=2011-06-15",
		Permission:    "GetSessionToken",
	},
}
