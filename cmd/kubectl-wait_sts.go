package main

import (
	"os"

	"github.com/oke-py/kubectl-wait-sts-plugin/pkg/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	command := cmd.NewCmdWaitSts()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
