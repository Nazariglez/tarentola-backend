// Created by nazarigonzalez on 30/9/17.

package database

import (
	"github.com/jinzhu/gorm"

	"github.com/nazariglez/tarentola-backend/database/rolemodel"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
)

type InitFunc func() error
type ModelInit func(database *gorm.DB) (interface{}, func() error)

var modelInitList = []ModelInit{
	usermodel.Init,
	rolemodel.Init,
}
