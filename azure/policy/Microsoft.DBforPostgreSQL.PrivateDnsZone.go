package policy

var Microsoft_DBforPostgreSQL_PrivateDnsZone = []Policy{
	{
		Path:   "/providers/Microsoft.DBforPostgreSQL/getPrivateDnsZoneSuffix",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "GetPrivateDnsZoneSuffix_Execute",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
}
