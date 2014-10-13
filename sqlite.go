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

func NewSqliteFromDB(db *sql.DB) *Sqlite {
	return &Sqlite{DBO{DB: db}}
}

func NewSqlite(dsn string) (*Sqlite, error) {

	_, err := os.Stat(dsn)
	if err != nil {
		return nil, err
	}

	s := new(Sqlite)

	s.DB, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = s.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s (%s)\n", err, dsn)
	}

	return s, nil

}
