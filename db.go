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

func NewDB() DB {
	db := DB{}
	con, err := sql.Open("sqlite3", "work.db")
	db.sql = con
	if err != nil {

	}
	return db
}
func (db *DB) GetWork() []Work {
	rows, err := db.sql.Query("SELECT chat_id, link FROM work")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	var works []Work
	for rows.Next() {
		var work Work

		rows.Scan(&work.ChatID, &work.link)
		works = append(works, work)
	}
	return works
}
func (db *DB) InsertWork(work Work) error {
	stmt, err := db.sql.Prepare("INSERT INTO work(chat_id, link ) values(?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(work.ChatID, work.link)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	log.Println(id)
	return err
}
func (db *DB) RemoveWork(work Work) error {
	stmt, err := db.sql.Prepare("DELETE FROM work WHERE chat_id=? AND link=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(work.ChatID, work.link)
	if err != nil {
		return err
	}
	return nil
}
