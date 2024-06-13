package db

import (
	"database/sql"
	"fmt"
)

type Review struct {
	Id      int
	Content string
}

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, fmt.Errorf("[DB] opening db: %s", err)
	}
	defer db.Close()

	return db, nil
}
