package policy

// Search_search policy
var Search_search = map[string]Policy{
	"Search_PostArm": {
		Path:   "/providers/Microsoft.Marketplace/search",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-02",
		},
		OperationID: "Search_PostArm",
		Resource:    "Search",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Marketplace/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-02",
		},
		OperationID: "Operations_List",
		Resource:    "Search",
	},
}
