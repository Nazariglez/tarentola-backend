// Created by nazarigonzalez on 9/10/17.

package gamemodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/categorymodel"
	"github.com/nazariglez/tarentola-backend/database/imagemodel"
	"github.com/nazariglez/tarentola-backend/database/statemodel"
	"github.com/nazariglez/tarentola-backend/database/tagmodel"
	"github.com/nazariglez/tarentola-backend/database/videomodel"
)

type Game struct {
	gorm.Model

	Title       string `gorm:"index, unique"`
	URL         string `gorm:"unique"`
	Description string

	MainImage      imagemodel.Image `gorm:"ForeignKey:MainImageRefer"` //thumbnail
	MainImageRefer uint
	MainVideo      videomodel.Video `gorm:"ForeignKey:MainVideoRefer"`
	MainVideoRefer uint

	Images     []imagemodel.Image       //images to rotate or show
	Videos     []videomodel.Video       //videos to show
	Categories []categorymodel.Category //multiplayer, shooter, etc..
	Tags       []tagmodel.Tag           //five tags for game max

	State      statemodel.State `gorm:"ForeignKey:StateRefer"` //unpublished, toModerate, published, rejected, banned
	StateRefer uint

	//todo add points or stars

	//builtin games
	//Achievements int //list of achievements to display in the landing
	//Boards       int //boards or rankings
	//Chat bool //enable chat or not
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Game{}, nil
}
