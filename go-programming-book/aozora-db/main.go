package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/encoding/japanese"
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

	b, err := os.ReadFile("ababababa.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)

	res, err := db.Exec(
		`insert into contents (author_id, title_id, title, content) values (?, ?, ?, ?)`,
		"000879",
		"14",
		"あばばばば",
		content,
	)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := res.LastInsertId()
	_ = docID
}
