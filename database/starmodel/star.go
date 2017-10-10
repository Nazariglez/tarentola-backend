// Created by nazarigonzalez on 10/10/17.

package starmodel

import "github.com/jinzhu/gorm"

type Star struct {
	gorm.Model

	Value  int  `sql:"default:'1'"` //1,2,3,4,5 stars
	GameID uint `gorm:"required"`
	UserID uint `gorm:"required"`
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Star{}, nil
}
