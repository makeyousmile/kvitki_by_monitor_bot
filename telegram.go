package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot(work chan Work) {
	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(c tele.Context) error {
		text := "kvitki_by_monitor_bot starts"
		return c.Send(text)
	})
	b.Handle("/link", func(c tele.Context) error {
		log.Println(c.Chat().ID)
		if len(c.Args()) == 0 {
			return c.Send("Введите ссылку", tele.ModeHTML)
		} else {

			job := Work{
				link:   c.Args()[0],
				ChatID: c.Chat().ID,
			}
			work <- job
		}

		return c.Send("Если есть или появятся билеты в продаже Вам придет сообщение", tele.ModeHTML)
	})

	b.Start()

}

func sendMessage(chatID int64, message string) {

	botToken := cfg.BotToken

	// Создаем настройки для бота
	pref := tele.Settings{
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Инициализируем бота
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// ID чата или пользователя, которому отправляем сообщение
	recipient := &tele.User{ID: chatID}

	// Отправляем сообщение

	_, err = bot.Send(recipient, message, tele.ModeHTML)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Message sent successfully!")
}
