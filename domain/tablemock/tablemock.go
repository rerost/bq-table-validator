package tablemock

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rerost/bq-table-validator/types"
)

type TableMock interface {
	// Mock Return SQL and error
	Mock(ctx context.Context, sql string, mocks []types.Mock) (string, error)
}

type tableMockImpl struct {
	currentTime time.Time
}

func NewTableMock(currentTime time.Time) TableMock {
	return tableMockImpl{
		currentTime: currentTime,
	}
}

func (tm tableMockImpl) Mock(ctx context.Context, sql string, mocks []types.Mock) (string, error) {
	if len(mocks) == 0 {
		return sql, nil
	}
	mockedTables := make([]string, len(mocks))
	for i, mock := range mocks {
		mockedTableName := tm.mockedTable(mock.Table)
		sql = strings.ReplaceAll(sql, mock.Table, mockedTableName)
		mockedTable := fmt.Sprintf(`%s AS (
%s
)`, mockedTableName, mock.SQL)
		mockedTables[i] = mockedTable
	}

	return "WITH " + strings.Join(mockedTables, " , ") + "\n" + sql, nil
}

func (tm tableMockImpl) mockedTable(table string) string {
	return fmt.Sprintf("mocked_%s_%s", strings.ReplaceAll(table, ".", "_"), tm.currentTime.Format("20060102150405"))
}
