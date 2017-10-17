// Created by nazarigonzalez on 10/10/17.

package gameachievementmodal

import "github.com/jinzhu/gorm"

type GameAchievement struct {
	gorm.Model

	Name   string
	GameID uint
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &GameAchievement{}, nil
}
