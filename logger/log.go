package logger

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	_, DEBUG = os.LookupEnv("DEBUG")
	_, VERBOSE = os.LookupEnv("VERBOSE")
}

// Logger for info, error and debug
var DEBUG bool
var VERBOSE bool
var stdOut = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "Jan 02 15:04:05"}
var stdErr = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "Jan 02 15:04:05"}
var Logger = log.Output(stdOut)
var LoggerStdErr = log.Output(stdErr)

func LogSuccess(service, permission string) {
	Logger.Info().Str(service, toSnakeCase(permission)).Msg("Success")
}

func LogMaybe(service, permission string) {
	if DEBUG {
		Logger.Debug().Str(service, toSnakeCase(permission)).Msg("Maybe")
	}
}

func LogDenied(status int, service, permission string) {
	if DEBUG {
		LoggerStdErr.Error().Int("status", status).Str(service, permission).Msg("Error")
	}
}

func LogError(err error) {
	if err != nil {
		if DEBUG {
			LoggerStdErr.Error().Err(err).Msg("Error")
		} else {
			Logger.Error().Err(err).Msg("Error")
		}
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
		LoggerStdErr.Debug().Interface(key, value).Msg("")
	}
}

func LogDebugVerbose(key string, value interface{}) {
	if DEBUG && VERBOSE {
		LoggerStdErr.Debug().Interface(key, value).Msg("")
	}
}

func LogDebugResponse(res *http.Response, permission string) {
	if VERBOSE && DEBUG {
		b, err := ioutil.ReadAll(res.Body)
		if err == nil {
			LoggerStdErr.Debug().Str("permission", permission).Str("body", string(b)).Msg("")
		}
	}
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}
