package signer

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/cliam/logger"
)

// SetCredentials sets the credentials for the signer. If profile is passed,
// all other credentials are ignored.
func SetCredentials(accessKeyID, secretAccessKey, sessionToken, profile string) *credentials.Credentials {
	if logger.DEBUG {
		logger.LoggerStdErr.Debug().
			Str("profile", profile).
			Str("key-id", accessKeyID).
			Str("secret", secretAccessKey).
			Str("session-token", sessionToken).Send()
	}
	if strings.HasPrefix(accessKeyID, "ASIA") && sessionToken == "" {
		logger.LoggerStdErr.Fatal().Msg("Session tokens missing for ASIA credentials")
	}
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
