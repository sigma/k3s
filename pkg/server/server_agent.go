// +build agentonly !serveronly,!darwin

package server

import (
	"context"

	"github.com/rancher/k3s/pkg/rootlessports"
)

func registerRootlessPorts(ctx context.Context, sc *Context, config *Config) error {
	if !config.DisableServiceLB && config.Rootless {
		return rootlessports.Register(ctx, sc.Core.Core().V1().Service(), config.ControlConfig.HTTPSPort)
	}
	return nil
}
