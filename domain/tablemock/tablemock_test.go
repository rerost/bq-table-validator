package tablemock_test

import (
	"context"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/rerost/bq-table-validator/domain/tablemock"
	"github.com/rerost/bq-table-validator/types"
)

func TestTableMockMock(t *testing.T) {
	currentTime, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	if err != nil {
		t.Error(err)
	}
	tableMock := tablemock.NewTableMock(currentTime)

	ctx := context.Background()

	type In struct {
		SQL   string
		Mocks []types.Mock
	}

	inOutPairs := []struct {
		name string
		in   In
	}{
		{
			name: "simplecase.sql",
			in: In{
				SQL: "SELECT id FROM users",
				Mocks: []types.Mock{
					{
						Table: "users",
						SQL:   "SELECT 1 AS id",
					},
				},
			},
		},
		{
			name: "nilmock.sql",
			in: In{
				SQL:   "SELECT id FROM users",
				Mocks: nil,
			},
		},
		{
			name: "emptymock.sql",
			in: In{
				SQL:   "SELECT id FROM users",
				Mocks: []types.Mock{},
			},
		},
		{
			name: "twomock.sql",
			in: In{
				SQL: "SELECT id, COUNT(1) FROM users INNER JOIN item ON users.id = item.user_id GROUP BY users.id",
				Mocks: []types.Mock{
					{
						Table: "users",
						SQL:   "SELECT 1 AS id",
					},
					{
						Table: "item",
						SQL:   "SELECT 1 AS user_id",
					},
				},
			},
		},
	}

	for _, inOutPair := range inOutPairs {
		inOutPair := inOutPair
		t.Run(inOutPair.name, func(t *testing.T) {
			sql, err := tableMock.Mock(ctx, inOutPair.in.SQL, inOutPair.in.Mocks)
			if err != nil {
				t.Error(err)
				return
			}

			cupaloy.SnapshotT(t, sql)
		})
	}
}
