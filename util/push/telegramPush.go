package push

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"pushbot/config"
)

func PushTelegram(msg tgbotapi.MessageConfig) error {
	token := config.HandleYaml().Push.Telegram.Token
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}
