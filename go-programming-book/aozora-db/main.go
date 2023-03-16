package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := []string{
		`create table if not exists authors (
			author_id text,
			author text,
			primary key (author_id)
		)`,
		`create table if not exists contents (
			author_id text,
			title_id text,
			title text,
			content text,
			primary key (author_id, title_id)
		)`,
		`create virtual table if not exists contents_fts using fts4 (words)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
