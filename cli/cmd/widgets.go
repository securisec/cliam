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
		Mask:  '*',
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
		logger.LoggerStdErr.Err(err).Msg(shared.GetMessageColor("error"))
	}
}

func cliResponseLogger(ser scanner.ServiceMap, status int) {
	if status == http.StatusOK {
		logger.LogSuccess(ser.Resource, ser.Policy.Permission)
	}
	if status != 200 && logger.DEBUG {
		logger.LogDenied(status, ser.Resource, ser.Policy.Permission)
	}
}

func cliLogRegion(r string) {
	logger.LoggerStdErr.Debug().Str("region", r).Msg(shared.GetMessageColor("info"))
}

func cliGcpLogRegion(r, z string) {
	logger.LoggerStdErr.Debug().Str("region", r).Str("zone", z).Msg(shared.GetMessageColor("info"))
}
