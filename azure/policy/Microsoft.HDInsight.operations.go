package policy

// Microsoft_HDInsight_operations policy
var Microsoft_HDInsight_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.HDInsight/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HDInsight",
	},
}
