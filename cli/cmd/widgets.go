package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
)

// promptInput is a helper function to prompt the user for input.
func promptInput(msg string) string {
	prompt := promptui.Prompt{
		Label: msg,
	}
	p, err := prompt.Run()
	if err != nil {
		logger.LoggerStdErr.Err(err).Msg("")
		os.Exit(1)
	}
	// trim whitespace
	return strings.Trim(p, " ")
}

func cliErrorLogger(s scanner.ServiceMap, err error) {
	if errors.Is(err, context.DeadlineExceeded) {
		logger.LoggerStdErr.Error().Str(s.Resource, s.Policy.Permission).Msg(shared.GetMessageColor("timeout"))
	} else {
		if logger.DEBUG {
			logger.LoggerStdErr.Err(err).Msg(shared.GetMessageColor("error"))
		}
	}
}

func cliResponseLoggerAWS(ser scanner.ServiceMap, status int) {
	l := logger.Logger
	cf := ser.Policy.ExtraCommandLineFlag
	flag, ok := ModifyExtraMap(awsKnownResourceMap)[cf]
	if status == http.StatusOK {
		sl := l.Info().Str(ser.Resource, logger.ToSnakeCase(ser.Policy.Permission))
		if ok && flag != "" {
			sl.Str(strings.ReplaceAll(cf, "_", "-"), flag)
		}
		successCounter++
		sl.Msg(shared.GetMessageColor("success"))
		return
	}
	// it was not a successful request
	failureCounter++
	if status != 200 && CLIVerbose {
		dl := l.Error().Str(ser.Resource, logger.ToSnakeCase(ser.Policy.Permission))
		dl.Int("status", status)
		if ok {
			dl.Str(strings.ReplaceAll(cf, "_", "-"), flag)
		}
		dl.Msg(shared.GetMessageColor("error"))
	}
}

func cliLogRegion(r string) {
	logger.LoggerStdErr.Debug().Str("region", r).Msg(shared.GetMessageColor("info"))
}

func cliGcpLogRegion(data map[string]string) {
	l := logger.LoggerStdErr.Debug()
	for k, v := range data {
		l.Str(k, v)
	}
	l.Msg(shared.GetMessageColor("info"))
}
