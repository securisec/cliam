package signer

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// SetCredentials sets the credentials for the signer. If profile is passed,
// all other credentials are ignored.
func SetCredentials(accessKeyID, secretAccessKey, sessionToken, profile string) *credentials.Credentials {
	if profile != "" {
		return credentials.NewSharedCredentials("", profile)
	}
	return credentials.NewStaticCredentials(accessKeyID, secretAccessKey, sessionToken)
	// return credentials.NewCredentials(&credentials.StaticProvider{
	// 	Value: credentials.Value{
	// 		AccessKeyID:     accessKeyID,
	// 		SecretAccessKey: secretAccessKey,
	// 		SessionToken:    sessionToken,
	// 	},
	// })
}
