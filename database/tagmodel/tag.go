// Created by nazarigonzalez on 9/10/17.

package tagmodel

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	Name string
	//todo many ref to game
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Tag{}, nil
}
