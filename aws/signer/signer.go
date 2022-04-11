package signer

import "github.com/aws/aws-sdk-go/aws/credentials"

func SetCredentials(accessKeyID, secretAccessKey, sessionToken string) *credentials.Credentials {
	return credentials.NewCredentials(&credentials.StaticProvider{
		Value: credentials.Value{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
			SessionToken:    sessionToken,
		},
	})
}
