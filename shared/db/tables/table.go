package tables

import (
	"database/sql"

	"github.com/drmarduk/mvd2/shared/db"
)

// Table is a template for a
type Table struct {
	ctx     *db.DBContext
	Name    string
	_Create string
	_Delete string
	_Insert string
	_Select string
	_Update string
}

// Create creates the table
func (t *Table) _create(ctx *db.DBContext) error {
	_, err := ctx.C.Exec(t._Create)
	return err
}

// Delete deletes the table
func (t *Table) _delete(ctx *db.DBContext, args ...interface{}) error {
	_, err := ctx.C.Exec(t._Delete, args...)
	return err
}

// Insert inserts an object in the table
func (t *Table) _insert(ctx *db.DBContext, args ...interface{}) (int64, error) {
	result, err := ctx.C.Exec(t._Insert, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// SelectOne selects the object from the table
func (t *Table) _selectone(ctx *db.DBContext, args ...interface{}) *sql.Row {
	row := ctx.C.QueryRow(t._Select, args...)
	return row
}

// Select selects the object from the table
func (t *Table) _select(ctx *db.DBContext, args ...interface{}) (*sql.Rows, error) {
	rows, err := ctx.C.Query(t._Select, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Update updates a single row
func (t *Table) _update(ctx *db.DBContext, args ...interface{}) error {
	_, err := ctx.C.Exec(t._Update, args...)
	return err
}
