package cmd

import (
	"github.com/NodeFitter/scalerctl/context"
	"github.com/spf13/cobra"
)

var (
	app = &context.App{}

	rootCmd = &cobra.Command{
		Use:   "scalerctl",
		Short: "CLI for NodeFitter",
		Long: `Scalerctl
		Tool for controlling the NodeFitter daemon`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Connect()
		},

		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return app.Close()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

// Init function for later. TODO: start the connection
func init() {
	cobra.OnInitialize()
}
