// Created by nazarigonzalez on 15/10/17.

package email

import (
	"github.com/nazariglez/tarentola-backend/config"
	"gopkg.in/gomail.v2"
)

func SendTestEmail(to []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.Email.User)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello fff<b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(
		config.Data.Email.SMTP,
		config.Data.Email.Port,
		config.Data.Email.User,
		config.Data.Email.Password,
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
