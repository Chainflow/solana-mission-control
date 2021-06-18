package alerter

import (
	"log"
	"strconv"
	"strings"

	"github.com/Chainflow/solana-mission-control/config"
)

// SendTelegramAlert sends the alert to telegram account
// check's alert setting before sending the alert
func SendTelegramAlert(msg string, cfg *config.Config) error {
	if strings.ToUpper(strconv.FormatBool(cfg.EnableAlerts.EnableTelegramAlerts)) == "TRUE" {
		if err := NewTelegramAlerter().SendTelegramMessage(msg, cfg.Telegram.BotToken, cfg.Telegram.ChatID); err != nil {
			log.Printf("Failed to send tg alert : %v of msg : %s", err, msg)
			return err
		}
	}
	return nil
}

// SendEmailAlert sends alert to email account
// by checking user's choice
func SendEmailAlert(msg string, cfg *config.Config) error {
	if strings.ToUpper(strconv.FormatBool(cfg.EnableAlerts.EnableEmailAlerts)) == "TRUE" {
		if err := NewEmailAlerter().SendEmail(msg, cfg); err != nil {
			log.Printf("failed to send email alert: %v", err)
			return err
		}
	}
	return nil
}
