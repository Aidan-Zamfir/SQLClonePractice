package core

import (
	"fmt"

	"github.com/marianogappa/sqlparser/query"
)

type Table struct {
	numRow int
	pages []Page
}

type TableResult int 

const (
	TableFull TableResult = iota
	ExecuteSuccess
	ExecuteFailed
)

func (t *Table) GetTableSize() string {
	return fmt.Sprintf("table size: %d", t.numRow)
}

func (t *Table) ExecuteInsert(q query.Query) TableResult {
	if t.numRow >=MaxRows {
		return TableFull
	}

	row := RowFromQuery(q)

	pageNum := t.numRow / PageSize
	if len(t.pages) <= pageNum {
		t.pages = append(t.pages, Page{})
	}
	page := &t.pages[pageNum]
	if page.PageSize() < PageSize {
		page.AddRow(row)
		t.numRow += 1
		return ExecuteSuccess
	}
	return ExecuteFailed
}