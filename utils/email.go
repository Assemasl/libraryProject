package utils

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendEmail(to string, subject string, body string) error {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(m)
}
