//+build wireinject

package cmd

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/google/wire"
	"github.com/googleapis/google-cloud-go-testing/bigquery/bqiface"
	"github.com/pkg/errors"
	"github.com/rerost/bq-table-validator/domain/bqquery"
	"github.com/rerost/bq-table-validator/domain/validator"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewBQClient(ctx context.Context, cfg Config) (bqiface.Client, error) {
	zap.L().Debug("Create BQ Client", zap.String("ProjectID", cfg.ProjectID))
	bqClient, err := bigquery.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bqiface.AdaptClient(bqClient), nil
}

func NewBQMiddleware(bqClient bqiface.Client) validator.Middleware {
	return bqquery.NewBQQuery(bqClient)
}

func InitializeCmd(ctx context.Context, cfg Config) (*cobra.Command, error) {
	wire.Build(
		NewCmdRoot,
		validator.NewValidator,
		NewBQClient,
		NewBQMiddleware,
	)
	return nil, nil
}
