// Created by nazarigonzalez on 7/10/17.

package utils

import (
  "regexp"
  "errors"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


func ValidateEmailFormat(email string) error {
  if !emailRegexp.MatchString(email) {
    return errors.New("Invalid email format.")
  }
  return nil
}