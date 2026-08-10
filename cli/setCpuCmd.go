package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	setCmd.AddCommand(setCpuCmd)
}

var setCpuCmd = &cobra.Command{
	Use:   "cpu <percentage>",
	Short: "Set the CPU threshold for VMs",
	Long:  "Set the maximum CPU usage after which new VMs need to be scheduled",
	RunE: func(cmd *cobra.Command, args []string) error {
		cpu, err := strconv.Atoi(args[0])
		if err != nil || cpu <= 0 || cpu > 100 {
			return fmt.Errorf("CPU threshold must be between 1 and 100")
		}

		return app.UpdateCPUThreshold(float32(cpu))
	},
}
