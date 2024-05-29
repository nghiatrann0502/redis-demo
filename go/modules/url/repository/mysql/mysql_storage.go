package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const dbTimeout = 3 * time.Second

type mySqlStorage struct {
	db *sql.DB
}

func NewMySqlStorage() (*mySqlStorage, error) {
	db, err := sql.Open("mysql", "url_shorten_dev:my_secret@tcp(localhost:3306)/url_shorten")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	store := &mySqlStorage{
		db,
	}

	return store, nil
}
