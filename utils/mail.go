package utils

import (
	"go-api/config"
	"net/smtp"
	"strings"
	"fmt"
)

func SendMail(subject string, body string, recipients []string) error {
	mailConfig := config.GetConfig().Mail
	host := mailConfig.Host
	port := mailConfig.Port
	user := mailConfig.User
	pass := mailConfig.Pass
	from := mailConfig.From

	// Create email headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(recipients, ", ")
	headers["Subject"] = subject
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// Build the full email message
	var msgBuilder strings.Builder
	for key, value := range headers {
		msgBuilder.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	msgBuilder.WriteString("\r\n") // Empty line between headers and body
	msgBuilder.WriteString(body)

	msg := []byte(msgBuilder.String())

	auth := smtp.PlainAuth("", user, pass, host)

	err := smtp.SendMail(host+":"+port, auth, from, recipients, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
