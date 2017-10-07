// Created by nazarigonzalez on 7/10/17.

package rolemodel

import (
  "github.com/jinzhu/gorm"
  "github.com/nazariglez/tarentola-backend/database/helpers"
)

type Role struct {
	gorm.Model

	Name  string `gorm:"index, unique"`
	Value int
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Role{}, initModel
}

func initModel() error {
  err := db.Where("name = ?", "User").First(&Role{}).Error
  if helpers.IsNotFoundErr(err) {
    err := db.Create(&Role{Name: "User", Value: 0}).Error
    if err != nil {
      return err
    }

    return nil
  }

  if err != nil {
    return err
  }

  return nil
}

