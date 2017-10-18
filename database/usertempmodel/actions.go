// Created by nazarigonzalez on 18/10/17.

package usertempmodel

import (
	"errors"
	"github.com/nazariglez/tarentola-backend/database/helpers"
)

func Create(userID uint) (*UserTemp, error) {
	ut := &UserTemp{UserID: userID}
	err := db.Create(ut).Error
	return ut, err
}

func FindByToken(token string) (*UserTemp, error) {
	u := UserTemp{}
	err := db.Where("token = ?", token).First(&u).Error
	return &u, err
}

func DeleteByID(id uint) error {
	err := db.Select("id").Where("id = ?", id).First(&UserTemp{}).Error
	if helpers.IsNotFoundErr(err) {
		return errors.New("Invalid user_temp id.")
	}

	if err != nil {
		return err
	}

	//delete permanently
	return db.Unscoped().Where("id = ?", id).Delete(&UserTemp{}).Error
}
