// Created by nazarigonzalez on 9/10/17.

package videomodel

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model

	URL    string
	GameID uint
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Video{}, nil
}
