// Created by nazarigonzalez on 30/9/17.

package api

import (
	"errors"
	"github.com/nazariglez/tarentola-backend/database"
	"net/http"
)

func homeController(w http.ResponseWriter, r *http.Request) {
	role := database.RoleModel{Name: "user", Value: 0}
	err := database.GetDB().Attrs(role).FirstOrCreate(&role).Error
	if err != nil {
		sendServerError(w, err)
		return
	}

	sendOk(w, []int{1, 2, 3})
	return
}

func errorController(w http.ResponseWriter, r *http.Request) {
	sendServerError(w, errors.New("server error!"))
	return
}
