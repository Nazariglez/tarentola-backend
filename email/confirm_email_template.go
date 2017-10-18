// Created by nazarigonzalez on 17/10/17.

package email

import (
	"fmt"
	"github.com/nazariglez/tarentola-backend/config"
)

var ConfirmEmailTemplate = `
<html>
  <body>
     <h4>Hello <%= name %>!</h4>
     <p>
      To active your account just click in the next link: <%= confirmation_link %>.
     </p>
  </body>
</html>
`

func SendUserConfirmationEmail(name, email, token string) error {
	e := &Email{
		Body: ConfirmEmailTemplate,
		Data: map[string]interface{}{
			"confirmation_link": fmt.Sprintf("http://%s/user/active/%s", config.Data.FrontURL, token),
			"email":             email,
			"name":              name,
		},
		To:      email,
		Subject: "Confirm your account",
	}

	return e.Send()
}
