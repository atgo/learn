package learn_sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// example: NewMySQL("user:pass@tcp(host:3306)/db")
func NewMySQL(dsn string) *sql.DB, error {
	return sql.Open("mysql", dsn)
}

// example: NewSQLite("/tmp/db.sqlite")
// memory:  NewSQLite(":memory:")
func NewSQLite(dsn string)  *sql.DB, error {
	return sql.Open("sqlite3", dns)
}
