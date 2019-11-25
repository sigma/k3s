// +build !agentonly,!serveronly

package main

import (
	"github.com/rancher/k3s/pkg/cli/cmds"
	"github.com/rancher/k3s/pkg/cli/kubectl"
)

func init() {
	commands = append(commands, cmds.NewKubectlCommand(kubectl.Run))
}
