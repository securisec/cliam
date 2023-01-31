package signer

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/cliam/logger"
)

// SetCredentials sets the credentials for the signer. If profile is passed,
// all other credentials are ignored.
func SetCredentials(accessKeyID, secretAccessKey, sessionToken, profile string) *credentials.Credentials {
	var creds *credentials.Credentials
	// when using service creds which starts with ASIA, a session token is required
	if strings.HasPrefix(accessKeyID, "ASIA") && sessionToken == "" {
		logger.LoggerStdErr.Fatal().Msg("Session tokens missing for ASIA credentials")
	}

	if profile != "" {
		creds = credentials.NewSharedCredentials("", profile)
	} else {
		creds = credentials.NewStaticCredentials(accessKeyID, secretAccessKey, sessionToken)
	}

	gc, err := creds.Get()
	if err != nil {
		logger.Logger.Fatal().Err(err).Send()
	}

	if logger.DEBUG {
		logger.LoggerStdErr.Debug().
			Str("profile", profile).
			Str("key-id", gc.AccessKeyID).
			Str("secret", gc.SecretAccessKey).
			Str("session-token", gc.SessionToken).Send()
	}

	return creds
	// return credentials.NewCredentials(&credentials.StaticProvider{
	// 	Value: credentials.Value{
	// 		AccessKeyID:     accessKeyID,
	// 		SecretAccessKey: secretAccessKey,
	// 		SessionToken:    sessionToken,
	// 	},
	// })
}
