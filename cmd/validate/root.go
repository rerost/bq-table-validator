package validate

import (
	"context"

	"github.com/spf13/cobra"
)

func NewCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "validate",
	}

	return cmd
}
