// Created by nazarigonzalez on 4/10/17.

package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	gorm.Model

	Name           string `gorm:"index"`
	Email          string `gorm:"unique, index, required"`
	Password       string `sql:"-"`
	HashedPassword string `gorm:"required"`
	Role           RoleModel
}

func (um *UserModel) BeforeCreate(scope *gorm.Scope) error {
	pw, err := hashPassword(um.Password)
	if err != nil {
		return err
	}

	scope.SetColumn("HashedPassword", pw)
	return nil
}

func (um *UserModel) BeforeUpdate(scope *gorm.Scope) error {
	pw, err := hashPassword(um.Password)
	if err != nil {
		return err
	}

	scope.SetColumn("HashedPassword", pw)
	return nil
}

func UserModelFindToLogin(email, password string) (*UserModel, error) {
	um := UserModel{}

	if err := db.Where(map[string]interface{}{
		"email": email,
	}).First(&um).Error; err != nil {
		if IsNotFoundErr(err) {
			return nil, nil
		}

		return nil, err
	}

	if !matchPassword(password, um.HashedPassword) {
		return nil, errors.New("Invalid email or password.")
	}

	return &um, nil
}

func UserModelCreate(um *UserModel) error {
	return db.Create(um).Error
}

func hashPassword(pass string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(pass), 11)
	if err != nil {
		return "", err
	}

	return string(p), nil
}

func matchPassword(password string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func UserModelExistsEmail(email string) (bool, error) {
	um := UserModel{Email: email}
	err := db.Select("ID").Where(um).First(&um).Error
	if err != nil {
		if IsNotFoundErr(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
