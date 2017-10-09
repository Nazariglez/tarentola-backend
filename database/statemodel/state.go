// Created by nazarigonzalez on 9/10/17.

package statemodel

import (
	"github.com/jinzhu/gorm"
	"github.com/nazariglez/tarentola-backend/database/helpers"
)

var stateList = []string{
	"Unpublished", //working on it
	"Pending",     //pending for a moderator
	"Published",   //public!!!
	"Rejected",    //Pending was rejected
	"Banned",      //banned for life!
}

var currentStates = []State{}

type State struct {
	gorm.Model

	Name string `gorm:"index, unique"`
}

var db *gorm.DB

func Init(database *gorm.DB) (interface{}, func() error) {
	db = database
	return &State{}, initModel
}

func initModel() error {
	states := []State{}
	err := db.Order("id ASC").Find(&states).Error
	if helpers.IsNotFoundErr(err) {
		return createStates(stateList)
	}

	if err != nil {
		return err
	}

	if len(states) != len(stateList) {
		list := []string{}
		for _, role := range stateList {
			exists := false
			for _, dbState := range states {
				if dbState.Name == role {
					exists = true
					break
				}
			}

			if !exists {
				list = append(list, role)
			}
		}

		return createStates(list)
	}

	return fillCurrentStates()
}

func createStates(list []string) error {
	for _, s := range list {
		if err := db.Create(&State{Name: s}).Error; err != nil {
			return err
		}
	}

	return fillCurrentStates()
}

func fillCurrentStates() error {
	states := []State{}
	if err := db.Select("id, name").Find(&states).Error; err != nil {
		return err
	}

	currentStates = states
	return nil
}
