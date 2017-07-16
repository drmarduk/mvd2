package model

import (
	"time"

	"github.com/drmarduk/mvd2/shared/db"
	"github.com/drmarduk/mvd2/shared/db/tables"
)

// OpenQueue returns all files in the queue
func OpenQueue(ctx *db.DBContext) ([]tables.QueueRow, error) {
	qt, err := tables.NewQueueTable(ctx)
	if err != nil {
		return nil, err
	}

	files, err := qt.Select()
	if err != nil {
		return nil, err
	}

	return files, nil
}

// AddQueue adds a new element in the qaueue
func AddQueue(ctx *db.DBContext, fname, hash string, size int64) error {
	qt, err := tables.NewQueueTable(ctx)
	if err != nil {
		return err
	}

	nq := &tables.QueueRow{
		Filename:  fname,
		Hash:      hash,
		Size:      size,
		DateAdded: time.Now().UTC(),
	}

	_, err = qt.Insert(nq)
	if err != nil {
		return err
	}

	return nil
}
