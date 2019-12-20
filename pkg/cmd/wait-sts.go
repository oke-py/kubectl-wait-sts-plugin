package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	example = `
	# wait for statefulset
	%[1]s wait-sts <statefulset>

	# wait for statefulset in different namespace
	%[1]s wait-sts <statefulset> -n/--namespace <ns>
	`
)

// WaitStsOptions is the struct holding common properties
type WaitStsOptions struct {
	args      []string
	namespace string
	name      string
}

// NewCmdWaitSts creates the cobra command to be executed
func NewCmdWaitSts() *cobra.Command {
	o := &WaitStsOptions{}

	cmd := &cobra.Command{
		Use:     "wait-sts [statefulset-name]",
		Short:   "Wait until Statefulset gets ready",
		Example: fmt.Sprintf(example, "kubectl"),
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&o.namespace, "namespace", "n", o.namespace, "override the namespace defined in the current context")

	return cmd
}

// Complete sets all information required for waiting statefulset
func (o *WaitStsOptions) Complete(args []string) error {
	o.args = args
	return nil
}

// Validate ensures proper command usage
func (o *WaitStsOptions) Validate() error {
	if len(o.args) != 1 {
		return fmt.Errorf("incorrect number of arguments, see --help for usage instructions")
	}
	o.name = o.args[0]
	return nil
}

// Run waits until statefulset gets ready
func (o *WaitStsOptions) Run() error {
	options := genericclioptions.NewConfigFlags(true)
	kubeConfig := options.ToRawKubeConfigLoader()

	if o.namespace == "" {
		namespace, _, err := kubeConfig.Namespace()

		if err != nil {
			return err
		}
		o.namespace = namespace
	}

	fmt.Printf("%s in namespace/%s", o.name, o.namespace)
	return nil
}
