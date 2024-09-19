package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func scrap(status chan Event, url string) {

	event := Event{}
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".event_short_title", func(e *colly.HTMLElement) {
		log.Println(e.Text)
		event.title = e.Text
	})
	c.OnHTML(".buy_button_text", func(e *colly.HTMLElement) {
		log.Println(e.Text)
		text := strings.Trim(e.Text, " ")
		if text == "Купить" {
			event.tickets = true
		}
		event.status = event.status + e.Text + "\n"
	})
	c.OnScraped(func(_ *colly.Response) {
		status <- event
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		event.link = r.URL.String()
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
}
