package sqlite

import (
	_ "embed"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	conn sqlx.DB
}

//go:embed journal.sql
var schema string

func NewSqlite(connectionString string) (*Sqlite, error) {
	conn, err := sqlx.Connect("sqlite3", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite connection: %w", err)
	}

	conn.MustExec(schema)

	return &Sqlite{
		conn: *conn,
	}, nil
}
