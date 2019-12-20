package main

import (
	"github.com/oke-py/kubectl-wait-sts-plugin/pkg/cmd"
	"os"
)

func main() {
	command := cmd.NewCmdWaitSts()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
