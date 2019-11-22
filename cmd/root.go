package cmd

import (
	"context"

	"github.com/rerost/bq-table-validator/cmd/validate"
	"github.com/spf13/cobra"
)

func NewCmdRoot(
	ctx context.Context,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bq-table-validator",
		Short: "Manage BigQuery view",
	}

	cmd.AddCommand(
		validate.NewCmd(ctx),
	)

	return cmd
}
