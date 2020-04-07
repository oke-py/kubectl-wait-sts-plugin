package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
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
	timeout   time.Duration
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
	cmd.Flags().DurationVar(&o.timeout, "timeout", 30*time.Second, "The length of time to wait before giving up.  Zero means check once and don't wait, negative means wait for a week.")
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

	// negative means wait for a week
	if o.timeout < 0 {
		o.timeout = 168 * time.Hour
	}

	return nil
}

// Run waits until statefulset gets ready
func (o *WaitStsOptions) Run() error {
	endTime := time.Now().Add(o.timeout)

	options := genericclioptions.NewConfigFlags(true)
	kubeConfig := options.ToRawKubeConfigLoader()

	if o.namespace == "" {
		namespace, _, err := kubeConfig.Namespace()

		if err != nil {
			return err
		}
		o.namespace = namespace
	}

	restConfig, err := options.ToRESTConfig()
	if err != nil {
		return err
	}

	client := kubernetes.NewForConfigOrDie(restConfig)

	watch, err := client.AppsV1().StatefulSets(o.namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", o.name),
	})
	if err != nil {
		return err
	}

	for {
		sts, err := client.AppsV1().StatefulSets(o.namespace).Get(context.TODO(), o.name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		if *sts.Spec.Replicas == sts.Status.ReadyReplicas {
			watch.Stop()
			return nil
		}

		timeout := time.Until(endTime)
		if timeout < 0 {
			watch.Stop()
			return fmt.Errorf("timed out waiting for the condition on statefulsets/%s", o.name)
		}
	}
}
