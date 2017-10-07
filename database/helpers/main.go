// Created by nazarigonzalez on 7/10/17.

package helpers

import "github.com/jinzhu/gorm"

func IsNotFoundErr(err error) bool {
	return gorm.ErrRecordNotFound == err
}
