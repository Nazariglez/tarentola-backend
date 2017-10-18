// Created by nazarigonzalez on 18/10/17.

package usertempmodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/database/helpers"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
	"github.com/nazariglez/tarentola-backend/logger"
	"github.com/nazariglez/tarentola-backend/utils"
	"time"
)

type UserTemp struct {
	gorm.Model

	UserID uint   `gorm:"required"`
	Token  string `gorm:"unique, index"`
}

var db *gorm.DB
var ticker *time.Ticker
var expireTime time.Duration

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &UserTemp{}, initTicker
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

func initTicker() error {
	expireTime = time.Duration(config.Data.Auth.ConfirmTokenExpire) * time.Hour
	ticker = time.NewTicker(time.Second * 3600)

	go runTicker()
	return nil
}

func runTicker() {
	for _ = range ticker.C {
		cleanTempUsers()
	}
}

func cleanTempUsers() {
	logger.Log.Debug("Cleaning temporal users...")
	users := []UserTemp{}
	t := time.Now().Add(-expireTime)

	err := db.Where("created_at < ?", t).Find(&users).Error
	if err != nil && !helpers.IsNotFoundErr(err) {
		logger.Log.Errorf("%s -> Cleaning temporal users.", err.Error())
		return
	}

	if len(users) != 0 {
		for _, u := range users {
			if err := DeleteByID(u.ID); err != nil {
				logger.Log.Errorf("%s -> Cleaning temporal user with id '%d'.", err.Error(), u.ID)
				continue
			}

			if err := usermodel.DeleteByID(u.UserID); err != nil {
				logger.Log.Errorf("%s -> Cleaning tempoaral user with id '%d', and user with id '%d'.", err.Error(), u.ID, u.UserID)
				continue
			}
		}
	}

	logger.Log.Debugf("Cleaned '%d' inactive temporal users.", len(users))
}
