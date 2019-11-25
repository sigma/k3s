// +build serveronly darwin
// +build !agentonly

package server

import (
	"context"

	"github.com/rancher/k3s/pkg/cli/cmds"
	"github.com/rancher/k3s/pkg/server"
	"github.com/urfave/cli"
)

func checkRootless(*cmds.Server) error {
	return nil
}

func runAgent(context.Context, *cli.Context, *cmds.Server, *server.Config) error {
	return nil
}
