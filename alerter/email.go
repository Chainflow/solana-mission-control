package alerter

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
)

// SendEmail to send mail alert
func (e emailAlert) SendEmail(msg string, cfg *config.Config) error {
	accountName := cfg.SendGrid.SendgridName
	fromEmail := cfg.SendGrid.SendgridEmail
	toEmail := cfg.SendGrid.ReceiverEmailAddress
	token := cfg.SendGrid.Token

	from := mail.NewEmail(accountName, fromEmail) //mail.NewEmail("Matic Tool", "matic@vitwit.com")
	subject := msg
	to := mail.NewEmail(accountName, toEmail) //mail.NewEmail("Matic Tool", toEmail)
	plainTextContent := msg
	htmlContent := msg
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(token)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
