// Created by nazarigonzalez on 9/10/17.

package avatarmodel

import "github.com/jinzhu/gorm"

type Avatar struct {
	gorm.Model

	URL string
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Avatar{}, nil
}
