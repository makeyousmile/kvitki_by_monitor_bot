package main

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

var cfg = &Cfg{}

func init() {
	cfg.visited = make(map[string]bool)
	cfg.BotToken = "7291370458:AAHgNvMBp47RiQmO4BzeB2l3sM2rY_DCR8E"
}
func main() {

	status := make(chan Event)
	link := make(chan string)
	go StartBot(link, status)
	scrap(status, link)

}
