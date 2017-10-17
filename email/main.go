// Created by nazarigonzalez on 17/10/17.

package email

import (
	"github.com/gobuffalo/plush"
	"github.com/nazariglez/tarentola-backend/config"
	"gopkg.in/gomail.v2"
)

var dial = func() *gomail.Dialer {
	return gomail.NewDialer(
		config.Data.Email.SMTP,
		config.Data.Email.Port,
		config.Data.Email.User,
		config.Data.Email.Password,
	)
}()

func compileTemplate(tpl string, data map[string]interface{}) (string, error) {
	ctx := plush.NewContextWith(data)
	tpl, err := plush.Render(tpl, ctx)
	return tpl, err
}

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return dial.DialAndSend(m)
}

func SendEmailBCC(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.Email.User)
	m.SetHeader("Bcc", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return dial.DialAndSend(m)
}
