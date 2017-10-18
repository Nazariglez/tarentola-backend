// Created by nazarigonzalez on 15/10/17.

package email

import (
	"github.com/nazariglez/tarentola-backend/config"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Body    string
	Data    map[string]interface{}
	To      string
	Subject string
}

func (e *Email) Send() error {
	tpl, err := compileTemplate(e.Body, e.Data)
	if err != nil {
		return err
	}

	return SendEmail(e.To, e.Subject, tpl)
}

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return dial.DialAndSend(m)
}
