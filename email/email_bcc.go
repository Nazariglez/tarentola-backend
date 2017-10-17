// Created by nazarigonzalez on 17/10/17.

package email

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
