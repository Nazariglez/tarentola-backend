// Created by nazarigonzalez on 2/10/17.

package controllers

import (
	"github.com/nazariglez/tarentola-backend/database"
	"net/http"
	"strings"
)

func TestToken(w http.ResponseWriter, r *http.Request) {
	SendOk(w)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.Form.Get("email"))
	pass := strings.TrimSpace(r.Form.Get("password"))
	if email == "" || pass == "" {
		SendBadRequest(w, "All fields are required.")
		return
	}

	exists, err := database.UserModelExistsEmail(email)
	if err != nil {
		SendServerError(w, err)
		return
	}

	if exists {
		SendBadRequest(w, "Email already exists.")
		return
	}

	userModel := database.UserModel{
		Email:    email,
		Password: pass,
	}

	if err := database.UserModelCreate(&userModel); err != nil {
		SendServerError(w, err)
		return
	}

	SendOk(w, "User created.")
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	SendOk(w, "User updated. "+r.Form.Get("id"))
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	SendOk(w, "User deleted. "+r.Form.Get("id"))
	return
}
