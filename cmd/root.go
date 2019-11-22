package cmd

import (
	"context"

	"github.com/rerost/bq-table-validator/cmd/validate"
	"github.com/rerost/bq-table-validator/domain/validator"
	"github.com/spf13/cobra"
)

func NewCmdRoot(
	ctx context.Context,
	v validator.Validator,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bq-table-validator",
		Short: "Manage BigQuery view",
	}

	cmd.AddCommand(
		validate.NewCmd(ctx, v),
	)

	return cmd
}
