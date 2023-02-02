package policy

// IOTDataPolicies policy
var IOTDataPolicies = map[string]Service{
	"ListRetainedMessages": {
		Method:        "GET",
		ServiceSuffix: "retainedMessage",
		Permission:    "ListRetainedMessages",
	},

	// extra
	"GetRetainedMessage": {
		ServiceSuffix:          "/retainedMessage/{{.topic}}",
		Permission:             "GetRetainedMessage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "topic",
	},
	"GetThingShadow": {
		ServiceSuffix:          "/things/{{.thing_name}}/shadow",
		Permission:             "GetThingShadow",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "thing_name",
	},
	"ListNamedShadowsForThing": {
		ServiceSuffix:          "/api/things/shadow/ListNamedShadowsForThing/{{.thing_name}}",
		Permission:             "ListNamedShadowsForThing",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "thing_name",
	},
}
