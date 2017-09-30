// Created by nazarigonzalez on 30/9/17.

package database

import "github.com/jinzhu/gorm"

var modelList = []interface{}{
	&UserModel{},
	&RoleModel{},
}

type UserModel struct {
	gorm.Model

	Name     string `gorm:"index, required"`
	Email    string `gorm:"unique, index, required"`
	Password string `gorm:"required"`
	Role     RoleModel
}

type RoleModel struct {
	gorm.Model

	Name  string `gorm:"index"`
	Value int
}
