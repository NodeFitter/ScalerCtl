package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the scheduling process",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Stop()
	},
}
