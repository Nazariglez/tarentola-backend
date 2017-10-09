// Created by nazarigonzalez on 7/10/17.

package rolemodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/helpers"
)

var roleList = []string{
	"User",      //a normal player
	"Creator",   //a user who can add new games
	"Moderator", //a user who can accept and publish pending games
	"Admin",     //total overpowered!!!
}

var currentRoles = []Role{}

type Role struct {
	gorm.Model

	Name string `gorm:"index, unique"`
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Role{}, initModel
}

func initModel() error {
	roles := []Role{}
	err := db.Order("id ASC").Find(&roles).Error
	if helpers.IsNotFoundErr(err) {
		return createRoles(roleList)
	}

	if err != nil {
		return err
	}

	if len(roles) != len(roleList) {
		list := []string{}
		for _, role := range roleList {
			exists := false
			for _, dbRole := range roles {
				if dbRole.Name == role {
					exists = true
					break
				}
			}

			if !exists {
				list = append(list, role)
			}
		}

		return createRoles(list)
	}

	return fillCurrentRoles()
}

func createRoles(list []string) error {
	for _, role := range list {
		if err := db.Create(&Role{Name: role}).Error; err != nil {
			return err
		}
	}

	return fillCurrentRoles()
}

func fillCurrentRoles() error {
	roles := []Role{}
	if err := db.Select("id, name").Find(&roles).Error; err != nil {
		return err
	}

	currentRoles = roles
	return nil
}
