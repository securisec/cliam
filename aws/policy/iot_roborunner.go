package policy

// IOTRoboRunnerPolicies policy
var IOTRoboRunnerPolicies = map[string]Service{
	"ListSites": {
		Method:        "GET",
		ServiceSuffix: "listSites",
		Permission:    "ListSites",
	},

	// extra
	"GetDestination": {
		ServiceSuffix:          "/getDestination",
		Permission:             "GetDestination",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetSite": {
		ServiceSuffix:          "/getSite",
		Permission:             "GetSite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWorker": {
		ServiceSuffix:          "/getWorker",
		Permission:             "GetWorker",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWorkerFleet": {
		ServiceSuffix:          "/getWorkerFleet",
		Permission:             "GetWorkerFleet",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListDestinations": {
		ServiceSuffix:          "/listDestinations",
		Permission:             "ListDestinations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "site",
	},
	"ListWorkerFleets": {
		ServiceSuffix:          "/listWorkerFleets",
		Permission:             "ListWorkerFleets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "site",
	},
	"ListWorkers": {
		ServiceSuffix:          "/listWorkers",
		Permission:             "ListWorkers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "site",
	},
}
