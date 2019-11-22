package bqquery

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/googleapis/google-cloud-go-testing/bigquery/bqiface"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type bqQuery struct {
	bqClient bqiface.Client
}

func NewBQQuery(bqClient bqiface.Client) bqQuery {
	return bqQuery{
		bqClient: bqClient,
	}
}

type Result = map[string]bigquery.Value

func (b bqQuery) Query(ctx context.Context, sql string) ([]map[string]interface{}, error) {
	query := b.bqClient.Query(sql)
	rows, err := query.Read(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := []Result{}

	for {
		var row Result
		zap.L().Debug("Load", zap.String("Load type", fmt.Sprintf("%T", row)))
		err := rows.Next(&row)

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, errors.WithStack(err)
		}

		res = append(res, row)
	}

	result := make([]map[string]interface{}, len(res))
	for i, r := range res {
		m := map[string]interface{}{}
		for k, v := range r {
			m[k] = v
		}
		result[i] = m
	}

	return result, nil
}
