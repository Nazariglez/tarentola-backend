// Created by nazarigonzalez on 9/10/17.

package controllers

import (
	"github.com/nazariglez/tarentola-backend/database/rolemodel"
	"net/http"
)

type roleInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	list, err := rolemodel.GetList()
	if err != nil {
		SendServerError(w, err)
		return
	}

	roles := []roleInfo{}
	for _, r := range list {
		roles = append(roles, roleInfo{r.ID, r.Name})
	}

	SendOk(w, roles)
	return
}
