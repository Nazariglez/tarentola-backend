// Created by nazarigonzalez on 30/9/17.

package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Open(prod, debug bool) error {
	_db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=tarentola sslmode=disable password=postgres")
	if err != nil {
		return err
	}

	_db.LogMode(debug)
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
