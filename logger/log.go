package logger

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/securisec/cliam/shared"
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

func LogSuccess(service, permission string, additionalVerbose ...map[string]string) {
	l := Logger.Info().Str(service, ToSnakeCase(permission))
	if len(additionalVerbose) > 0 && VERBOSE {
		for k, v := range additionalVerbose[0] {
			l.Str(k, v)
		}
	}
	l.Msg(shared.GetMessageColor("success"))
}

func LogMaybe(service, permission string) {
	if DEBUG {
		Logger.Debug().Str(service, ToSnakeCase(permission)).Msg(shared.GetMessageColor("warning"))
	}
}

func LogDenied(status int, service, permission string) {
	if DEBUG {
		LoggerStdErr.Error().Int("status", status).Str(service, ToSnakeCase(permission)).Msg(shared.GetMessageColor("error"))
	}
}

func LogError(err error, operation ...string) {
	if err != nil {
		if DEBUG {
			l := LoggerStdErr.Error().Err(err)
			if len(operation) > 0 {
				l.Str("operation", operation[0])
			}
			l.Msg(shared.GetMessageColor("error"))
		}
	}
}

func LogPanic(err error) {
	if err != nil {
		LoggerStdErr.Panic().Err(err).Msg("")
	}
}

func LogWarning(msg string) {
	Logger.Warn().Str("msg", msg).Msg(shared.GetMessageColor("warning"))
}

func LogDebug(key string, value interface{}) {
	if DEBUG {
		LoggerStdErr.Debug().Interface(key, value).Msg(shared.GetMessageColor("debug"))
	}
}

func LogDebugVerbose(key string, value interface{}) {
	if DEBUG && VERBOSE {
		LoggerStdErr.Debug().Interface(key, value).Msg(shared.GetMessageColor("debug"))
	}
}

func LogDebugResponse(res *http.Response, permission string) {
	if VERBOSE && DEBUG {
		b, err := ioutil.ReadAll(res.Body)
		if err == nil {
			LoggerStdErr.Debug().Str("permission", permission).Str("body", string(b)).Msg(shared.GetMessageColor("debug"))
		}
	}
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}
