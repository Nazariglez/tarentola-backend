// Created by nazarigonzalez on 30/9/17.

package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nazariglez/tarentola-backend/config"
)

var db *gorm.DB

func Open(prod, debug bool) error {
	opts := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		config.Data.Database.Host,
		config.Data.Database.User,
		config.Data.Database.Name,
		config.Data.Database.SSLMode,
		config.Data.Database.Password,
	)
	_db, err := gorm.Open("postgres", opts)
	if err != nil {
		return err
	}

	_db.LogMode(config.Data.Database.Debug)
	_db.AutoMigrate(modelList...)

	//todo prod credentials

	db = _db
	return nil
}

func Close() error {
	if db == nil {
		return errors.New("Can not close a database without open first.")
	}

	return db.Close()
}

func GetDB() *gorm.DB {
	return db
}
