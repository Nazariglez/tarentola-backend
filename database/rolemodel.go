// Created by nazarigonzalez on 4/10/17.

package database

import "github.com/jinzhu/gorm"

type RoleModel struct {
	gorm.Model

	Name  string `gorm:"index"`
	Value int
}
