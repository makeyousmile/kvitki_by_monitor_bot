package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Phone struct {
	number string
	date   time.Time
}
type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
	buffer []Phone
}
type Work1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewDB() DB {
	db := DB{}
	con, err := sql.Open("sqlite3", "work.db")
	db.sql = con
	if err != nil {

	}
	return db
}

func (db *DB) InsertWork(work Work) error {
	stmt, err := db.sql.Prepare("INSERT INTO work(link, chat_id) values(?,?)")
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(work.chatID, work.link)
	if err != nil {
		log.Print(err)
	}
	id, err := res.LastInsertId()
	log.Println(id)
	return nil
}
