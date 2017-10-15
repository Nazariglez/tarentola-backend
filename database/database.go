// Created by nazarigonzalez on 30/9/17.

package database

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
	"github.com/nazariglez/tarentola-backend/logger"
	"github.com/nazariglez/tarentola-backend/utils"
	"os"
	"strings"
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

	existsAdmin, err := usermodel.ExistsAdmin()
	if err != nil {
		return err
	}

	if !existsAdmin {
		if err := createAdminAccount(); err != nil {
			return err
		}
	}

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

func createAdminAccount() error {
	logger.Log.Warn("An admin account it's required. Please fill the next fields:")

	email := inputAdminEmail()
	pass := inputAdminPass()
	name := inputAdminName()

	if err := usermodel.CreateAdmin(name, email, pass); err != nil {
		return err
	}

	fmt.Println("")
	logger.Log.Infof("Created an admin account for '%s'", email)
	return nil
}

func inputAdminEmail() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for {
		fmt.Print("\nEmail: ")
		scanner.Scan()
		text = strings.TrimSpace(scanner.Text())

		if err := utils.ValidateEmailFormat(text); err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	return text
}

func inputAdminPass() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for {
		fmt.Print("\nPassword: ")
		scanner.Scan()
		text = strings.TrimSpace(scanner.Text())

		if len(text) < 6 {
			fmt.Println("Minimum 6 characters.")
		} else {
			break
		}
	}

	return text
}

func inputAdminName() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for {
		fmt.Print("\nPublic Name: ")
		scanner.Scan()
		text = strings.TrimSpace(scanner.Text())

		if len(text) < 3 {
			fmt.Println("Minimum 3 characters.")
		} else {
			break
		}
	}

	return text
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
