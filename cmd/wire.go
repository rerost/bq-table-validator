//+build wireinject

package cmd

import (
	"context"

	"github.com/google/wire"
	"github.com/spf13/cobra"
)

func InitializeCmd(ctx context.Context, cfg Config) (*cobra.Command, error) {
	wire.Build(
		NewCmdRoot,
	)
	return nil, nil
}
