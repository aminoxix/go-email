package services

import (
	"main/configs"
	"net/smtp"
)

func SendEmail(from string, fromPassword string, to []string, subject string, htmlContent string) error {
	fromEmailSMTP := configs.GetEnv("FROM_EMAIL_SMTP")
	smtpAddress := configs.GetEnv("SMTP_ADDRESS")

	auth := smtp.PlainAuth(
		"",
		from,
		fromPassword,
		fromEmailSMTP,
	)

	// email headers
	headers := "MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"From: " + from + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n\r\n"

	message := headers + htmlContent

	return smtp.SendMail(
		smtpAddress,
		auth,
		from,
		to,
		[]byte(message),
	)
}
