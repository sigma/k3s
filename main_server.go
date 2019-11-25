// +build !agentonly

package main

import (
	"github.com/rancher/k3s/pkg/cli/cmds"
	"github.com/rancher/k3s/pkg/cli/server"
)

func init() {
	commands = append(commands, cmds.NewServerCommand(server.Run))
}
