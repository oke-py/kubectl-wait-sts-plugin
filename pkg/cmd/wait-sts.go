package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	example = `
	# wait for statefulset
	%[1]s wait-sts <statefulset>

	# wait for statefulset in different namespace
	%[1]s wait-sts <statefulset> -n/--namespace <ns>
	`
)

// NewCmdWaitSts creates the cobra command to be executed
func NewCmdWaitSts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "wait-sts [statefulset-name]",
		Short:   "Wait until Statefulset gets ready",
		Example: fmt.Sprintf(example, "kubectl"),
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
