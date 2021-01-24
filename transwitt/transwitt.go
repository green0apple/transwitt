package transwitt

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Run(opconf OperateConfig) error {
	var telegram *tgbotapi.BotAPI
	var err error
	if opconf.Messanger.Telegram != (TelegramConfig{}) {
		telegram, err = tgbotapi.NewBotAPI(opconf.Messanger.Telegram.Token)
		if err != nil {
			return err
		}
		log.Printf("Telegram bot [%s] is authorized", telegram.Self.UserName)

		ucUpdates := tgbotapi.NewUpdate(0)
		ucUpdates.Timeout = 60
		chanUpdate, err := telegram.GetUpdatesChan(ucUpdates)
		if err != nil {
			return err
		}
		// telegram command listener
		go func() {
			for {
				select {
				case u := <-chanUpdate:
					log.Println("updates : ", chanUpdate)
					log.Println("From :", u.Message.From.String())
					log.Println("Date :", u.Message.Date)
					log.Println("Chat.ID :", u.Message.Chat.ID)
					log.Println("Chat.UserName :", u.Message.Chat.UserName)
					log.Println("Text :", u.Message.Text)
				}
			}
		}()
	}

	return nil
}
