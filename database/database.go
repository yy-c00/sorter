//Package database provides all implementations of model package that have access to data
package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	db  *sql.DB
	getError func() error
)

func init() {
	dsn := os.Getenv("DSN")
	engine := os.Getenv("ENGINE")

	var err error
	db, err = sql.Open(engine, dsn)
	getError = func() error {
		return err
	}
}

//Connection returns a pool connection with the database
func Connection() *sql.DB {
	return db
}

//Error returns an error if there was an error with database connection
func Error() error {
	return getError()
}

//Close close the connection with the database
func Close() error {
	return db.Close()
}