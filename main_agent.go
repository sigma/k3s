// +build !serveronly,!darwin

package main

import (
	"github.com/rancher/k3s/pkg/cli/agent"
	"github.com/rancher/k3s/pkg/cli/cmds"
	"github.com/rancher/k3s/pkg/cli/crictl"
)

func init() {
	commands = append(commands, cmds.NewAgentCommand(agent.Run))
	commands = append(commands, cmds.NewCRICTL(crictl.Run))
}
