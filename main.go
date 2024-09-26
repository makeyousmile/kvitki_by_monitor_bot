package main

import (
	"log"
	"time"
)

type Cfg struct {
	BotToken string
}
type Event struct {
	title   string
	tickets bool
	dates   []string
}
type Work struct {
	link   string `json:"link"`
	ChatID int64  `json:"chat_id"`
}

var cfg = &Cfg{}

func init() {
	cfg.BotToken = "7291370458:AAHgNvMBp47RiQmO4BzeB2l3sM2rY_DCR8E"
}
func main() {
	db := NewDB()
	work := make(chan Work)
	go startWorkFromDB(db, work)
	go StartBot(work)
	for job := range work {
		err := db.InsertWork(job)
		if err != nil {
			log.Print(err)
		}
		go func() {
			for {
				event := scrap(job.link)
				if event.tickets {
					text := ""
					for _, t := range event.dates {
						text += t + ", "
					}
					sendMessage(job.ChatID, "<b>"+event.title+"</b>"+". Есть билеты на сдедующие даты: <b>"+text+"</b> ")
					db.RemoveWork(job)
					break
				} else {
					time.Sleep(10 * time.Second)
				}
			}
		}()
	}
}
func startWorkFromDB(db DB, work chan Work) {
	works := db.GetWork()
	for _, w := range works {
		work <- w
	}
}
