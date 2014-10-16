package pdo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type (
	MySQL struct {
		DBO
	}
)

func NewMysql(db *sql.DB) *MySQL {
	return &MySQL{DBO{DB: db}}
}

func NewMySQLFromDSN(dsn string) (*MySQL, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s (%s)\n", err, dsn)
	}

	return NewMysql(db), nil

}
