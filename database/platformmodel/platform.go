// Created by nazarigonzalez on 9/10/17.

package platformmodel

import "github.com/jinzhu/gorm"

type Platform struct {
	gorm.Model

	Name    string //web, android, ios, etc...
	URL     string
	Enabled bool //disable platform with links broken or similar
	GameID  uint
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Platform{}, nil
}
