// Created by nazarigonzalez on 10/10/17.

package scoremodel

import "github.com/jinzhu/gorm"

type Score struct {
	gorm.Model

	Value   int  `sql:"default:'1'"`
	BoardID uint `gorm:"required"`
	UserID  uint `gorm:"required"`
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Score{}, nil
}
