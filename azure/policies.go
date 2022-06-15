package azure

import "github.com/securisec/cliam/azure/policy"

func GetPolicyKeys() []string {
	hold := make([]string, 0, len(Policies))
	for k := range Policies {
		hold = append(hold, k)
	}
	return hold
}

var Policies = map[string][]policy.Policy{
	"Microsoft.Web.WebApps":                policy.Microsoft_Web_WebApps,
	"Microsoft.Web.StaticSites":            policy.Microsoft_Web_StaticSites,
	"Microsoft.Web.ResourceProvider":       policy.Microsoft_Web_ResourceProvider,
	"Microsoft.Web.ResourceHealthMetadata": policy.Microsoft_Web_ResourceHealthMetadata,
	"Microsoft.Storage.storage":            policy.Microsoft_Storage_storage,
}
