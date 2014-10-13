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

func NewMySQLFromDB(db *sql.DB) *MySQL {
	return &MySQL{DBO{DB: db}}
}

func NewMySQL(dsn string) (*MySQL, error) {

	var (
		m   = new(MySQL)
		err error
	)

	m.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = m.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s (%s)\n", err, dsn)
	}

	return m, nil

}
