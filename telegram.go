package main

import (
	"log"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot(link chan string, status chan Event) *tele.Bot {
	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(c tele.Context) error {
		text := "kvitkiby_monitor_bot starts"
		return c.Send(text)
	})
	b.Handle("/link", func(c tele.Context) error {
		link := c.Args()[0]
		c.Send("Если есть или появятся билеты в продаже Вам придет сообщение", tele.ModeHTML)
		go func() {
			for event := range status {
				message := "Мероприятие: <b>" + strings.Trim(event.title, " ") + "</b>\n"
				if event.tickets {
					message = message + "есть билеты на продажу"
					c.Send(message, tele.ModeHTML)
					return
				}

			}

		}()
		go func() {
			for {
				scrap(status, link)
				time.Sleep(10 * time.Second)
			}
		}()

		return nil
	})

	b.Start()
	return b
}
