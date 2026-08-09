package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "scalerctl",
		Short: "CLI for NodeFitter",
		Long: `Scalerctl
		Tool for controlling the NodeFitter daemon`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

// Init function for later. TODO: start the connection
func init() {
	cobra.OnInitialize()
}
