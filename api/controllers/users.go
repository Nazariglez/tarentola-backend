// Created by nazarigonzalez on 2/10/17.

package controllers

import (
	"github.com/nazariglez/tarentola-backend/database"
	"net/http"
	"strconv"
	"strings"
	"github.com/nazariglez/tarentola-backend/utils"
)

func TestToken(w http.ResponseWriter, r *http.Request) {
	SendOk(w)
	return
}

type publicUserInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type ownUserInfo struct {
	publicUserInfo
	Email string `json:"email"`
}

//public info
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.Form.Get("id"))
	if id == "" {
		SendBadRequest(w, "Invalid user id.")
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	user, err := database.UserModelGetByID(uint(uid))
	if err != nil {
		if database.IsNotFoundErr(err) {
			SendBadRequest(w, "Invalid user id.")
			return
		}

		SendServerError(w, err)
		return
	}

	data := publicUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role.Name,
	}
	SendOk(w, data)
	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	user, err := database.UserModelGetByID(claims.ID)
	if err != nil {
		if database.IsNotFoundErr(err) {
			SendBadRequest(w, "Invalid user.")
			return
		}

		SendServerError(w, err)
		return
	}

	public := publicUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role.Name,
	}

	data := ownUserInfo{
		publicUserInfo: public,
		Email:          user.Email,
	}
	SendOk(w, data)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.Form.Get("email"))
	pass := strings.TrimSpace(r.Form.Get("password"))
	if email == "" || pass == "" {
		SendBadRequest(w, "All fields are required.")
		return
	}

	if err := utils.ValidateEmailFormat(email); err != nil {
		SendBadRequest(w, err.Error())
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
	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	err = database.UserModelDeleteByID(claims.ID)
	if err != nil {
		if database.IsNotFoundErr(err) {
			SendBadRequest(w, "Invalid user id.")
			return
		}

		SendServerError(w, err)
		return
	}

	SendOk(w, "User deleted.")
	return
}

func GetToken(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Bearer" {
		return ""
	}

	return auth[1]
}
