// Created by nazarigonzalez on 7/10/17.

package usermodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/achievementmodel"
	"github.com/nazariglez/tarentola-backend/database/avatarmodel"
	"github.com/nazariglez/tarentola-backend/database/rolemodel"
	"github.com/nazariglez/tarentola-backend/logger"
	"time"
)

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &User{}, nil
}

type User struct {
	gorm.Model

	Name           string `gorm:"index; unique; required"`
	Email          string `gorm:"index; unique; required"`
	Password       string `sql:"-"`
	HashedPassword string `gorm:"required"`

	Active  bool `sql:"default:'false'"`
	Banned  bool `sql:"default:'false'"`
	BanTime time.Time

	Coins int `sql:"default:'0'"`

	Role      rolemodel.Role `gorm:"ForeignKey:RoleRefer"` // use RoleRefer as foreign key
	RoleRefer uint           `sql:"default:'1'"`

	Avatar      avatarmodel.Avatar `gorm:"ForeignKey:AvatarRefer"` // use RoleRefer as foreign key
	AvatarRefer uint

	//todo guilds?

	Achievements []achievementmodel.Achievement
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	pw, err := hashPassword(u.Password)
	if err != nil {
		return err
	}

	scope.SetColumn("HashedPassword", pw)
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	if u.Password != "" {
		pw, err := hashPassword(u.Password)
		if err != nil {
			return err
		}

		scope.SetColumn("HashedPassword", pw)
	}
	return nil
}

func (u *User) IsActive() bool {
	if u.Banned && u.BanTime != (time.Time{}) {
		diff := time.Now().Sub(u.BanTime)
		if diff >= 0 {

			if err := BanByID(u.ID, false, time.Time{}); err != nil {
				logger.Log.Error(err)
			} else {
				u.Banned = false
				u.BanTime = time.Time{}
			}

		}
	}

	return u.Active && !u.Banned
}
