package context

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/NodeFitter/scalerctl/comms"
)

type App struct {
	RPC *rpc.Client
}

func (a *App) Connect() error {
	socket := os.Getenv("NODEFITTER_SOCKET")
	if socket == "" {
		socket = "/run/nodefitter/NodeFitter.sock"
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return fmt.Errorf("connect to daemon: %w", err)
	}

	a.RPC = rpc.NewClient(conn)
	return nil
}

func (a *App) Close() error {
	if a.RPC != nil {
		return a.RPC.Close()
	}
	return nil
}

func (a *App) UpdateMemThreshold(threshold float64) error {
	args := &comms.UpdateMemThresholdArgs{
		NewThreshold: threshold,
	}
	reply := &comms.UpdateMemThresholdReply{}

	if err := a.RPC.Call("Controller.UpdateMemThreshold", args, reply); err != nil {
		return err
	}

	if !reply.Success {
		return fmt.Errorf("failed to update memory threshold")
	}

	fmt.Println("Threshold updated successfully")
	return nil
}

func (a *App) UpdateCPUThreshold(threshold float32) error {
	args := &comms.UpdateCPUThresholdArgs{
		NewThreshold: threshold,
	}
	reply := &comms.UpdateCPUThresholdReply{}

	if err := a.RPC.Call("Controller.UpdateCPUThreshold", args, reply); err != nil {
		return err
	}

	if !reply.Success {
		return fmt.Errorf("failed to update CPU threshold")
	}

	fmt.Println("Threshold updated successfully")
	return nil
}

func (a *App) VMs() error {
	args := &comms.EmptyArgs{}
	reply := &comms.PrintVMReply{}

	if err := a.RPC.Call("Controller.PrintVM", args, reply); err != nil {
		return err
	}

	if len(reply.VMs) == 0 {
		fmt.Println("No VMs found.")
		return nil
	}

	w := tabwriter.NewWriter(
		os.Stdout,
		0,   // min width
		4,   // tab width
		2,   // padding
		' ', // padding character
		0,   // flags
	)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tMEMORY\tCPU\tVM GROUP\tTEMPLATE ID\tCREATED")
	fmt.Fprintln(w, "--\t------\t---\t--------\t-----------\t-------")

	for _, vm := range reply.VMs {

		mem := strconv.FormatFloat(vm.AvailableMem, 'f', -1, 64)
		cpu := strconv.FormatFloat(float64(vm.AvailableCPU), 'f', -1, 64)

		if vm.AvailableMem == math.MaxFloat64 {
			mem = "Not Available"
		}

		if vm.AvailableCPU == math.MaxFloat32 {
			cpu = "Not Available"
		}

		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%d\t%s\n",
			vm.Id,
			mem,
			cpu,
			vm.VMGroupName,
			vm.VMTemplateId,
			vm.InstantiationTimestamp.Format("2006-01-02 15:04:05"),
		)
	}

	return nil
}

func (a *App) Start() error {
	args := &comms.EmptyArgs{}
	reply := &comms.EmptyReply{}

	if err := a.RPC.Call("Controller.Start", args, reply); err != nil {
		return err
	}

	fmt.Println("Successfully started the scheduling process")
	return nil
}

func (a *App) Stop() error {
	args := &comms.EmptyArgs{}
	reply := &comms.EmptyReply{}

	if err := a.RPC.Call("Controller.Stop", args, reply); err != nil {
		return err
	}

	fmt.Println("Successfully stopped the scheduling process")
	return nil

}
