package main

import "time"

type Cfg struct {
	BotToken string
	visited  map[string]bool
}
type Event struct {
	title   string
	status  string
	tickets bool
	link    string
}
type Work struct {
	link   string
	chatID int64
}

var cfg = &Cfg{}

func init() {
	cfg.BotToken = "7291370458:AAHgNvMBp47RiQmO4BzeB2l3sM2rY_DCR8E"
}
func main() {

	work := make(chan Work)
	go StartBot(work)
	for job := range work {
		go func() {
			for {
				event := scrap(job.link)
				if event.tickets {
					sendMessage(job.chatID, event.title+": Есть билеты!!!")
					break
				} else {
					time.Sleep(10 * time.Second)
				}
			}
		}()
	}
}
