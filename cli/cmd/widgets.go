package cmd

import (
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/securisec/cliam/logger"
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
