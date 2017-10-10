// Created by nazarigonzalez on 2/10/17.

package controllers

import (
	"github.com/nazariglez/tarentola-backend/database/helpers"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
	"github.com/nazariglez/tarentola-backend/utils"
	"net/http"
	"strconv"
	"strings"
)

func TestToken(w http.ResponseWriter, r *http.Request) {
	SendOk(w, r)
	return
}

type publicUserInfo struct {
	ID     uint       `json:"id"`
	Name   string     `json:"name"`
	Role   roleInfo   `json:"role"`
	Avatar avatarInfo `json:"avatar"`
}

type ownUserInfo struct {
	publicUserInfo
	Email string `json:"email"`
}

//todo update avatar

//public info
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.Form.Get("id"))
	if id == "" {
		SendBadRequest(w, r, "Invalid user id.")
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	user, err := usermodel.GetByID(uint(uid))
	if err != nil {
		if helpers.IsNotFoundErr(err) {
			SendBadRequest(w, r, "Invalid user id.")
			return
		}

		SendServerError(w, r, err)
		return
	}

	data := publicUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Role: roleInfo{user.Role.ID, user.Role.Name},
	}
	SendOk(w, r, data)
	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	user, err := usermodel.GetByID(claims.ID)
	if err != nil {
		if helpers.IsNotFoundErr(err) {
			SendBadRequest(w, r, "Invalid user.")
			return
		}

		SendServerError(w, r, err)
		return
	}

	public := publicUserInfo{
		ID:   user.ID,
		Name: user.Name,
		Role: roleInfo{user.Role.ID, user.Role.Name},
	}

	data := ownUserInfo{
		publicUserInfo: public,
		Email:          user.Email,
	}
	SendOk(w, r, data)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.Form.Get("email"))
	name := strings.TrimSpace(r.Form.Get("name"))
	pass := strings.TrimSpace(r.Form.Get("password"))
	if email == "" || pass == "" || name == "" {
		SendBadRequest(w, r, "All fields are required.")
		return
	}

	if err := utils.ValidateEmailFormat(email); err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	exists, err := usermodel.ExistsEmail(email)
	if err != nil {
		SendServerError(w, r, err)
		return
	}

	if exists {
		SendBadRequest(w, r, "Email already exists.")
		return
	}

	userModel := usermodel.User{
		Email:    email,
		Password: pass,
		Name:     name,
	}

	if err := usermodel.Create(&userModel); err != nil {
		SendServerError(w, r, err)
		return
	}

	SendOk(w, r, "User created.")
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.Form.Get("email"))
	name := strings.TrimSpace(r.Form.Get("name"))
	pass := strings.TrimSpace(r.Form.Get("password"))
	if !(email != "" || name != "" || pass != "") {
		SendBadRequest(w, r, "Invalid field to update.")
		return
	}

	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	data := map[string]interface{}{}
	if email != "" {
		data["email"] = email
	}

	if pass != "" {
		data["password"] = pass
	}

	if name != "" {
		data["name"] = name
	}

	err = usermodel.UpdateFields(claims.ID, data)
	if err != nil {
		if helpers.IsNotFoundErr(err) {
			SendBadRequest(w, r, "Invalid user id.")
			return
		}

		SendServerError(w, r, err)
		return
	}

	SendOk(w, r, "User updated.")
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	claims, err := ValidateToken(GetToken(r))
	if err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	err = usermodel.DeleteByID(claims.ID)
	if err != nil {
		if helpers.IsNotFoundErr(err) {
			SendBadRequest(w, r, "Invalid user id.")
			return
		}

		SendServerError(w, r, err)
		return
	}

	SendOk(w, r, "User deleted.")
	return
}

func GetToken(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Bearer" {
		return ""
	}

	return auth[1]
}
