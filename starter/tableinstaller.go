package main

import "github.com/drmarduk/mvd2/shared/db"

// Install creates the tables in the database
func Install(ctx *db.DBContext, table TableInstaller) error {
	_, err := ctx.C.Exec(table.CreateTable())
	if err != nil {
		return err
	}
	return nil
}

// TableInstaller blabla
type TableInstaller interface {
	Name() string
	CreateTable() string
}
