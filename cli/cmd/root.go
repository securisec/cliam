package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	MaxThreads     int
	RequestTimeout int
	CLIVerbose     bool
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
}

func init() {
	RootCmd.PersistentFlags().IntVar(&MaxThreads, "max-threads", 5, "Maximum number of threads to use.")
	RootCmd.PersistentFlags().IntVar(&RequestTimeout, "request-timeout", 10, "Timeout for each request in seconds.")
	RootCmd.PersistentFlags().BoolVarP(&CLIVerbose, "verbose", "v", false, "Enable verbose output.")
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
