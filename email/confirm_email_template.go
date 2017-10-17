// Created by nazarigonzalez on 17/10/17.

package email

var ConfirmEmailTemplate = `
<html>
  <body>
     Hello <%= name %> click in this link: <%= confirmation_link %>!
  </body>
</html>
`
