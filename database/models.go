// Created by nazarigonzalez on 30/9/17.

package database

import (
	"github.com/jinzhu/gorm"

	"github.com/nazariglez/tarentola-backend/database/achievementmodel"
	"github.com/nazariglez/tarentola-backend/database/avatarmodel"
	"github.com/nazariglez/tarentola-backend/database/boardmodel"
	"github.com/nazariglez/tarentola-backend/database/categorymodel"
	"github.com/nazariglez/tarentola-backend/database/gamemodel"
	"github.com/nazariglez/tarentola-backend/database/imagemodel"
	"github.com/nazariglez/tarentola-backend/database/platformmodel"
	"github.com/nazariglez/tarentola-backend/database/rolemodel"
	"github.com/nazariglez/tarentola-backend/database/scoremodel"
	"github.com/nazariglez/tarentola-backend/database/starmodel"
	"github.com/nazariglez/tarentola-backend/database/statemodel"
	"github.com/nazariglez/tarentola-backend/database/tagmodel"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
	"github.com/nazariglez/tarentola-backend/database/videomodel"
)

type InitFunc func() error
type ModelInit func(database *gorm.DB) (interface{}, func() error)

var modelInitList = []ModelInit{
	achievementmodel.Init,
	avatarmodel.Init,
	boardmodel.Init,
	categorymodel.Init,
	achievementmodel.Init,
	gamemodel.Init,
	imagemodel.Init,
	platformmodel.Init,
	rolemodel.Init,
	scoremodel.Init,
	starmodel.Init,
	statemodel.Init,
	tagmodel.Init,
	usermodel.Init,
	videomodel.Init,
}
