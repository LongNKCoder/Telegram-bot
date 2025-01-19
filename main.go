package main

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	token := os.Getenv("TOKEN")
	setting := tele.Settings{
		URL:         "",
		Token:       token,
		Updates:     0,
		Poller:      &tele.LongPoller{Timeout: 10 * time.Second},
		Synchronous: false,
		Verbose:     false,
		ParseMode:   "",
		OnError:     nil,
		Client:      nil,
		Offline:     false,
	}
	bot, err := tele.NewBot(setting)
	if err != nil {
		panic(err)
	}
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello, World!")
	})
	bot.Start()
}
