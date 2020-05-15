package database

type SqlHandler interface {
	Execute(string, ...interface{}) (SqlResult, error)
	Query(string, ...interface{}) (SqlRow, error)
}

type SqlResult interface {
	LastInsertedId() (int64, error)
	RowsAffected() (int64, error)
}

type SqlRow interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
