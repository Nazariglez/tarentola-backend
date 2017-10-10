// Created by nazarigonzalez on 9/10/17.

package gamemodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/boardmodel"
	"github.com/nazariglez/tarentola-backend/database/categorymodel"
	"github.com/nazariglez/tarentola-backend/database/gameachievementmodal"
	"github.com/nazariglez/tarentola-backend/database/imagemodel"
	"github.com/nazariglez/tarentola-backend/database/platformmodel"
	"github.com/nazariglez/tarentola-backend/database/starmodel"
	"github.com/nazariglez/tarentola-backend/database/statemodel"
	"github.com/nazariglez/tarentola-backend/database/tagmodel"
	"github.com/nazariglez/tarentola-backend/database/videomodel"
)

type Game struct {
	gorm.Model

	Title       string `gorm:"index, unique"`
	Platforms   []platformmodel.Platform
	Description string
	UserID      uint

	MainImage      imagemodel.Image `gorm:"ForeignKey:MainImageRefer"` //thumbnail
	MainImageRefer uint
	MainVideo      videomodel.Video `gorm:"ForeignKey:MainVideoRefer"`
	MainVideoRefer uint

	Images     []imagemodel.Image       //images to rotate or show
	Videos     []videomodel.Video       //videos to show
	Categories []categorymodel.Category //multiplayer, shooter, etc..
	Tags       []tagmodel.Tag           //five tags for game max
	Stars      []starmodel.Star

	State      statemodel.State `gorm:"ForeignKey:StateRefer"` //unpublished, toModerate, published, rejected, banned
	StateRefer uint

	Achievements []gameachievementmodal.GameAchievement
	Boards       []boardmodel.Board
	Chat         bool
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &Game{}, nil
}
