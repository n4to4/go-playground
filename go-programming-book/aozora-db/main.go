package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
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

	_, err = db.Exec(
		`insert into authors (author_id, author) values (?, ?)`,
		"000879",
		"芥川竜之介",
	)
	if err != nil {
		log.Fatal(err)
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

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatal(err)
	}

	seg := t.Wakati(content)
	_, err = db.Exec(`
		insert into contents_fts (docid, words) values (?, ?)
	`,
		docID,
		strings.Join(seg, " "),
	)
	if err != nil {
		log.Fatal(err)
	}

	query := "虫 AND ココア"
	rows, err := db.Query(`
		select a.author, c.title
		from contents c
		inner join authors a using (author_id)
		inner join contents_fts f
		  on c.rowid = f.docid
		  and words match ?
	`, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var author, title string
		err = rows.Scan(&author, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(author, title)
	}
}
