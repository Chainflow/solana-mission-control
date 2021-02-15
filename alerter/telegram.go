package alerter

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// SendTelegramMessage to send alert to telegram bot
func (t telegramAlert) SendTelegramMessage(msgText, botToken string, chatID int64) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}
	bot.Debug = true
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = msgText
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
