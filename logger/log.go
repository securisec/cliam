package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	_, DEBUG = os.LookupEnv("DEBUG")
}

// Logger for info, error and debug
var DEBUG bool
var output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "Jan 02 15:04:05"}
var Logger = log.Output(output)

func LogSuccess(service, permission string) {
	Logger.Info().Str(service, permission).Msg("")
}

// func LogStatus(status int) {
// 	if DEBUG {
// 		Logger.Warn().Int("status", status).Msg("")
// 	}
// }

func LogDenied(status int, service, permission string) {
	if DEBUG {
		Logger.Warn().Int("status", status).Str(service, permission).Msg("")
	}
}

func LogError(err error) {
	if err != nil {
		Logger.Error().Err(err).Msg("Error")
	}
}

func LogPanic(err error) {
	if err != nil {
		Logger.Panic().Err(err).Msg("")
	}
}

func LogWarning(msg string) {
	Logger.Warn().Msg(msg)
}

func LogDebug(key string, value interface{}) {
	if DEBUG {
		Logger.Debug().Interface(key, value).Msg("")
	}
}
