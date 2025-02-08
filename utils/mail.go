package utils

import (
	"go-api/config"
	"net/smtp"
	"fmt"
)

func SendMail(msg []byte, recipients []string) error {
	mailConfig := config.GetConfig().Mail
	host := mailConfig.Host
	port := mailConfig.Port
	user := mailConfig.User
	pass := mailConfig.Pass
	from := mailConfig.From

	auth := smtp.PlainAuth("", user, pass, host)

	err := smtp.SendMail(host+":"+port, auth, from, recipients, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
