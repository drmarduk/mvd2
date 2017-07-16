package db

import (
	"database/sql"
	"fmt"

	// Driver import for mysql
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

/*
	List of common sql drivers in go

	SQL Server: https://github.com/denisenkom/go-mssqldb
	MySQL: https://github.com/go-sql-driver/mysql/
	Sqlite3: https://github.com/mattn/go-sqlite3
	Postgres: https://github.com/lib/pq
*/

// DBContext holds the database info for
// accessing the database
type DBContext struct {
	Driver   string
	Host     string
	Database string
	C        *sql.DB
	User     string
	Password string
}

// NewDBContext returns a new db instance based on the driver
// if the driver is sqlite3, then host must be the db-filename on disk
func NewDBContext(host, user, pass, database, driver string) (*DBContext, error) {
	ctx := &DBContext{
		Host:     host,
		Database: database,
		User:     user,
		Password: pass,
		Driver:   driver,
	}
	var err error
	dsn := generateDSN(driver, host, user, pass, database)
	ctx.C, err = sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	err = ctx.C.Ping()
	return ctx, err
}

func generateDSN(driver, host, user, pass, database string) string {
	switch driver {
	case "sqlite3":
		return host
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, database)
	case "sqlserver":
		return fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;encrypt=disable", host, database, user, pass)
	case "mssql":
		return fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;encrypt=disable", host, database, user, pass)
	case "psql":
		return fmt.Sprintf("User ID=%s;Password=%s;Host=%s;Database=%s", user, pass, host, database)
	}
	panic("uncreable")
}
