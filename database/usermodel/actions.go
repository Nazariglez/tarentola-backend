// Created by nazarigonzalez on 7/10/17.

package usermodel

import (
	"errors"
	"github.com/nazariglez/tarentola-backend/database/helpers"
	"github.com/nazariglez/tarentola-backend/database/rolemodel"
	"golang.org/x/crypto/bcrypt"
)

func DeleteByID(id uint) error {
	err := db.Select("id").Where("id = ?", id).First(&User{}).Error
	if helpers.IsNotFoundErr(err) {
		return errors.New("Invalid user id.")
	}

	if err != nil {
		return err
	}

	//delete permanently
	return db.Unscoped().Where("id = ?", id).Delete(&User{}).Error
}

func GetByID(id uint) (*User, error) {
	um := &User{}
	err := db.
		Preload("Role").
		Preload("Avatar").
		Where("id = ?", id).
		Find(um).Error
	if err != nil {
		return nil, err
	}
	return um, nil
}

func CreateAdmin(username, email, password string) error {
	u := User{
		Email:     email,
		Name:      username,
		Password:  password,
		RoleRefer: rolemodel.GetID("Admin"),
	}

	return Create(&u)
}

func ExistsAdmin() (bool, error) {
	var c int
	err := db.Model(&User{}).Where("role_refer = ?", rolemodel.GetID("Admin")).Count(&c).Error
	if err != nil {
		return false, err
	}

	if c == 0 {
		return false, nil
	}

	return true, nil
}

func FindOne(u *User) error {
	return db.Where(*u).First(&u).Error
}

func FindToLogin(email, password string) (*User, error) {
	u := User{}

	err := db.
		Preload("Role").
		Preload("Avatar").
		Select("id, email, hashed_password, role_refer").
		Where(map[string]interface{}{
			"email": email,
		}).First(&u).Error

	if err != nil {
		if helpers.IsNotFoundErr(err) {
			return nil, nil
		}

		return nil, err
	}

	if !matchPassword(password, u.HashedPassword) {
		return nil, errors.New("Invalid email or password.")
	}

	return &u, nil
}

func UpdateFields(id uint, fields map[string]interface{}) error {
	return db.Model(&User{}).Where("id = ?", id).Updates(fields).Error
}

func Create(u *User) error {
	return db.Create(u).Error
}

func ExistsEmail(email string) (bool, error) {
	u := User{Email: email}
	err := db.Select("id").Where(u).First(&u).Error
	if err != nil {
		if helpers.IsNotFoundErr(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

//--
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
