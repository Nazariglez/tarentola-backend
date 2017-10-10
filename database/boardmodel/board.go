// Created by nazarigonzalez on 10/10/17.

package boardmodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/scoremodel"
)

type Board struct {
	gorm.Model

	Name   string `gorm:"required"`
	GameID uint   `gorm:"required"`

	Scores []scoremodel.Score
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Board{}, nil
}
