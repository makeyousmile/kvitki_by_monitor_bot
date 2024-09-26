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
	var dates []string
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
	c.OnHTML(".concert_details_spec_content", func(e *colly.HTMLElement) {
		if e.Text != "" {
			log.Println(e.Text)
			dates = append(dates, strings.TrimSpace(e.Text))
		}
	})

	c.OnHTML(".event_short", func(e *colly.HTMLElement) {

		title := e.ChildText(".event_short_title")
		if strings.TrimSpace(title) == event.title {
			text := e.ChildText(".buy_button_text")
			text = strings.Trim(text, " ")

			if text == "Купить" {
				date := e.ChildText(".mobile_layout_list_mobile_hidden")
				dates = append(dates, date)
				event.tickets = true
			}
		}

	})
	c.OnScraped(func(_ *colly.Response) {

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(link)
	if err != nil {
		log.Print(err)
	}
	wg.Wait()
	event.dates = dates
	return event
}
