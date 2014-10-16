// +build cgo

package pdo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Sqlite struct {
		DBO
	}
)

func NewSqlite(db *sql.DB) *Sqlite {
	return &Sqlite{DBO{DB: db}}
}

func NewSqliteFromDSN(dsn string) (*Sqlite, error) {

	_, err := os.Stat(dsn)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s (%s)\n", err, dsn)
	}

	return NewSqlite(db), nil

}
