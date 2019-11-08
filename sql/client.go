package learn_sql

// example:
//   user:pass@tcp(host:3306)/db
func NewMySQL(dsn string) *sql.DB, error {
	return sql.Open("mysql", dsn)
}
