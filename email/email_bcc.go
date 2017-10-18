// Created by nazarigonzalez on 17/10/17.

package email

import (
	"github.com/nazariglez/tarentola-backend/config"
	"gopkg.in/gomail.v2"
)

//blind carbon copy emails
type EmailBCC struct {
	Body    string
	Data    map[string]interface{}
	To      []string
	Subject string
}

func (e *EmailBCC) Send() error {
	tpl, err := compileTemplate(e.Body, e.Data)
	if err != nil {
		return err
	}

	return SendEmailBCC(e.To, e.Subject, tpl)
}

func SendEmailBCC(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.Email.User)
	m.SetHeader("Bcc", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return dial.DialAndSend(m)
}
