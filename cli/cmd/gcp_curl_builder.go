package cmd

import "github.com/spf13/cobra"

// TODO: implement this
var gcpRestCurlBuilderCmd = &cobra.Command{
	Use:     "curl-builder",
	Short:   "Build the curl command to test a gcp policy.",
	Example: "TODO",
	Run:     gcpRestCurlBuilderFunc,
}

func init() {
	gcpRestCmd.AddCommand(gcpRestCurlBuilderCmd)
	gcpRestCurlBuilderCmd.Flags().StringP("policy", "p", "", "The policy to build")
	gcpRestCurlBuilderCmd.Flags().StringP("operation", "o", "", "The operation to build.")
	gcpRestCurlBuilderCmd.Flags().StringSliceP("values", "n", []string{}, "The values to use for known values.")

	gcpRestCurlBuilderCmd.MarkFlagRequired("policy")
	gcpRestCurlBuilderCmd.MarkFlagRequired("operation")
}

func gcpRestCurlBuilderFunc(cmd *cobra.Command, args []string) {
	// TODO
}
