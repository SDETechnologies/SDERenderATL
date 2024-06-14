package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[DB] pinging db: %s", err)
	}

	return db, nil
}

func InsertReview(ctx context.Context, db *sql.DB, content string) (Review, error) {
	sqlStmt := "insert into reviews (content) values (?)"

	tx, err := db.Begin()
	if err != nil {
		return Review{}, fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		return Review{}, fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, content)

	if err != nil {
		return Review{}, fmt.Errorf("[DB] executing statement: %s", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return Review{}, fmt.Errorf("[DB] last returned id: %s", err)
	}

	return Review{
		Id:      int(id),
		Content: content,
	}, nil
}

func GetReviews(ctx context.Context, db *sql.DB) ([]Review, error) {
	sqlStmt := "select id, content from reviews"

	tx, err := db.Begin()
	if err != nil {
		return []Review{}, fmt.Errorf("[DB] beginning transaction: %s", err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		return []Review{}, fmt.Errorf("[DB] preparing statement: %s", err)
	}

	defer stmt.Close()

	res, err := stmt.QueryContext(ctx)

	if err != nil {
		return []Review{}, fmt.Errorf("[DB] executing statement: %s", err)
	}

	reviews := []Review{}

	for res.Next() {
		var content string
		var id int64
		err = res.Scan(&id, &content)

		if err != nil {
			return []Review{}, fmt.Errorf("[DB] scanning thread")
		}

		reviews = append(reviews, Review{
			Id:      int(id),
			Content: content,
		})
	}

	return reviews, nil
}
