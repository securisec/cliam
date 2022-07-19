package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var (
	MaxThreads     int
	RequestTimeout int
	CLIVerbose     bool
	SaveOutput     bool
	successCounter = 0
	failureCounter = 0
	maybeCounter   = 0
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cliam",
	Short: "Cliam is a tool to enumerate cloud credentials for their permissions.",
	Long:  `Set DEBUG=* and/or VERBOSE=* to see debug output.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	PostRun: PostRunStatsFunc,
}

func init() {
	RootCmd.PersistentFlags().IntVar(&MaxThreads, "max-threads", 5, "Maximum number of threads to use.")
	RootCmd.PersistentFlags().IntVar(&RequestTimeout, "request-timeout", 5, "Timeout for each request in seconds.")
	RootCmd.PersistentFlags().BoolVarP(&CLIVerbose, "verbose", "v", false, "Enable verbose output.")
	RootCmd.Flags().BoolVar(&SaveOutput, "save-output", false, "Save output to file on success")
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func PostRunStatsFunc(cmd *cobra.Command, _ []string) {
	cloudProvider := ""
	cp := strings.Fields(cmd.CommandPath())
	if len(cp) >= 2 {
		cloudProvider = cp[1]
	}

	if CLIVerbose {
		logger.LoggerStdErr.Debug().Msgf(
			"%s: %s, %s, %s",
			cloudProvider,
			shared.Green(fmt.Sprintf("%d success", successCounter)),
			shared.Yellow(fmt.Sprintf("%d maybe", maybeCounter)),
			shared.Red(fmt.Sprintf("%d failure", failureCounter)),
		)
	}
}
