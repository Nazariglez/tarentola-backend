// Created by nazarigonzalez on 2/10/17.

package controllers

import (
	"fmt"
	"github.com/nazariglez/tarentola-backend/database"
	"net/http"
	"strconv"
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
	id := strings.TrimSpace(r.Form.Get("id"))
	if id == "" {
		SendBadRequest(w, "Invalid user id.")
		return
	}

	//todo isAdmin can delete

	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	um := database.UserModel{Email: claims.Email}
	if err := database.UserModelFindOne(&um); err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	//check if it's the same user
	uid, err := strconv.Atoi(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	if um.ID != uint(uid) {
		SendForbidden(w)
		return
	}

	SendOk(w, fmt.Sprintf("User deleted. %d - %s", um.ID, um.Email))
	return
}

func GetToken(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Bearer" {
		return ""
	}

	return auth[1]
}
