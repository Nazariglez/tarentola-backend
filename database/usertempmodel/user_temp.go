// Created by nazarigonzalez on 18/10/17.

package usertempmodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/helpers"
	"github.com/nazariglez/tarentola-backend/utils"
)

type UserTemp struct {
	gorm.Model

	UserID uint   `gorm:"required"`
	Token  string `gorm:"unique, index"`
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &UserTemp{}, nil
}

func (u *UserTemp) BeforeCreate(scope *gorm.Scope) error {
	//find a valid token
	var (
		err   error
		token string
	)

	for {
		token = utils.GetRandomID(32)
		err = db.Where("token = ?", token).Find(&UserTemp{}).Error

		if helpers.IsNotFoundErr(err) {
			err = nil
			break
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		return err
	}

	scope.SetColumn("Token", token)
	return nil
}
