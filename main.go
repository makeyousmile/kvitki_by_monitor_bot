package main

import "time"

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
	chatID int64  `json:"chat_id"`
}

var cfg = &Cfg{}

func init() {
	cfg.BotToken = "7291370458:AAHgNvMBp47RiQmO4BzeB2l3sM2rY_DCR8E"
}
func main() {
	db := NewDB()

	work := make(chan Work)
	go StartBot(work)
	for job := range work {
		db.InsertWork(job)
		go func() {
			for {
				event := scrap(job.link)
				if event.tickets {
					text := ""
					for _, t := range event.dates {
						text += t + ", "
					}
					sendMessage(job.chatID, "<b>"+event.title+"</b>"+". Есть билеты на сдедующие даты: <b>"+text+"</b> ")
					break
				} else {
					time.Sleep(10 * time.Second)
				}
			}
		}()
	}
}
