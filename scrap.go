package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"sync"
)

func scrap(link string) Event {
	var wg sync.WaitGroup
	event := Event{}
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".concert_details_title", func(e *colly.HTMLElement) {
		if e.Text != "" {
			event.title = strings.TrimSpace(e.Text)
		}
	})
	c.OnHTML(".show_details_title", func(e *colly.HTMLElement) {
		if e.Text != "" {
			event.title = strings.TrimSpace(e.Text)
		}
	})

	c.OnHTML(".event_short_top_bottom", func(e *colly.HTMLElement) {

		title := e.ChildText(".event_short_title")
		if strings.TrimSpace(title) == event.title {
			text := e.ChildText(".buy_button_text")
			text = strings.Trim(text, " ")

			if text == "Купить" {
				event.tickets = true
			}
		}

	})
	c.OnScraped(func(_ *colly.Response) {

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		event.link = r.URL.String()
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(link)
	if err != nil {
		log.Print(err)
	}
	wg.Wait()
	return event
}
