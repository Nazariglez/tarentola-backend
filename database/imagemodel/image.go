// Created by nazarigonzalez on 9/10/17.

package imagemodel

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model

	URL string
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Image{}, nil
}
