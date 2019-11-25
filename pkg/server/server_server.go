// +build serveronly darwin
// +build !agentonly

package server

import (
	"context"
)

func registerRootlessPorts(ctx context.Context, sc *Context, config *Config) error {
	return nil
}
