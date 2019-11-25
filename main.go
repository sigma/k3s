//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go
//go:generate go fmt pkg/deploy/zz_generated_bindata.go
//go:generate go fmt pkg/static/zz_generated_bindata.go

package main

import (
	"os"

	"github.com/rancher/k3s/pkg/cli/cmds"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var commands []cli.Command

func main() {
	app := cmds.NewApp()
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
