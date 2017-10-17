// Created by nazarigonzalez on 15/10/17.

package email

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
