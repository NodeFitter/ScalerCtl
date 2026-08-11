package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the scheduling process",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start()
	},
}
