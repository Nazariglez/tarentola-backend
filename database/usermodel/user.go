// Created by nazarigonzalez on 7/10/17.

package usermodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/rolemodel"
)

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &User{}, nil
}

type User struct {
	gorm.Model

	Name           string         `gorm:"index; unique; required"`
	Email          string         `gorm:"index; unique; required"`
	Password       string         `sql:"-"`
	HashedPassword string         `gorm:"required"`
	Role           rolemodel.Role `gorm:"ForeignKey:RoleRefer"` // use RoleRefer as foreign key
	RoleRefer      uint           `sql:"default:'1'"`
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
	pw, err := hashPassword(u.Password)
	if err != nil {
		return err
	}

	scope.SetColumn("HashedPassword", pw)
	return nil
}