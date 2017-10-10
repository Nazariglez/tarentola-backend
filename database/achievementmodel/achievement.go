// Created by nazarigonzalez on 10/10/17.

package achievementmodel

import "github.com/jinzhu/gorm"

type Achievement struct {
	gorm.Model

	GameAchievementID uint
	UserID            uint
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Achievement{}, nil
}
