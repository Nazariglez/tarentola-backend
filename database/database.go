// Created by nazarigonzalez on 30/9/17.

package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/logger"
)

var db *gorm.DB

func Open() error {
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

	db = _db

	db.LogMode(config.Data.Database.Debug)
	initModels(db)
	return nil
}

func initModels(db *gorm.DB) {
	modelList := []interface{}{}
	initList := []InitFunc{}
	for _, init := range modelInitList {
		m, f := init(db)
		modelList = append(modelList, m)
		initList = append(initList, f)
	}

	//init tables
	db.AutoMigrate(modelList...)

	//exec custom init functions
	for _, f := range initList {
		if f != nil {
			if err := f(); err != nil {
				logger.Log.Fatal(err)
			}
		}
	}
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
