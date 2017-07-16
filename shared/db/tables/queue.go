package tables

import (
	"fmt"
	"time"

	"github.com/drmarduk/mvd2/shared/db"
)

// NewQueueTable returns a new queu instance
func NewQueueTable(ctx *db.DBContext) (*QueueTable, error) {
	queue := &QueueTable{}

	queue.ctx = ctx
	queue.Name = "queue"
	return queue, nil
}

// QueueTable database structure
type QueueTable struct {
	Table
}

// QueueRow is one line in the table
type QueueRow struct {
	ID        int64
	Filename  string
	Hash      string
	Size      int64
	DateAdded time.Time
}

// Create function of queue
func (q *QueueTable) Create() error {
	q._Create = fmt.Sprintf(`
		Create table %s (
			id integer not null,
			filename text not null,
			hash text,
			size integer,
			date_added timestamp
		);
	`, q.Name)

	return q._create(q.ctx)
}

// Delete removes one row from the table
func (q *QueueTable) Delete(id int) error {
	q._Delete = fmt.Sprintf("delete from %s where id = ?;", q.Name)

	return q._delete(q.ctx, id)
}

// Insert adds a new item in the table
func (q *QueueTable) Insert(row *QueueRow) (QueueRow, error) {
	result := *row
	q._Insert = fmt.Sprintf(`
		insert into %s(filename, hash, size, date_added) values(?, ?, ?, ?);`, q.Name)

	id, err := q._insert(q.ctx, row.Filename, row.Hash, row.Size, time.Now().UTC())
	if err != nil {
		return result, err
	}
	result.ID = id
	return result, nil
}

// SelectOne returns a slice or QueueRows
func (q *QueueTable) SelectOne(id int64) (*QueueRow, error) {
	q._Select = fmt.Sprintf(`
		Select id, filename, hash, size, date_added from %s where id = ?`, q.Name)
	tmp := q._selectone(q.ctx, id)
	x := &QueueRow{}
	err := tmp.Scan(&x.ID, &x.Filename, &x.Hash, &x.Size, &x.DateAdded)
	if err != nil {
		return nil, err
	}

	return x, nil
}

// Select returns a slice of QueueRows
func (q *QueueTable) Select() ([]QueueRow, error) {
	q._Select = fmt.Sprintf(`
		Select id, filename, hash, size, date_added from %s;`, q.Name)
	tmp, err := q._select(q.ctx)
	if err != nil {
		return nil, err
	}
	var result []QueueRow

	for tmp.Next() {
		x := QueueRow{}
		err = tmp.Scan(&x.ID, &x.Filename, &x.Hash, &x.Size, &x.DateAdded)
		if err != nil {
			fmt.Printf("could not parse queryrow: %v\n", err)
			continue
		}
		result = append(result, x)
	}
	return result, nil
}
