// Created by nazarigonzalez on 9/10/17.

package controllers

import (
	"bytes"
	"github.com/nazariglez/tarentola-backend/utils"
	"io"
	"net/http"
)

type avatarInfo struct {
	ID  uint `json:"id"`
	URL uint `json:"url"`
}

func GetList(w http.ResponseWriter, r *http.Request) {
	//todo get avatar list
	SendOk(w, r)
}

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	//todo this must be a admin method, not user or public
	file, header, err := r.FormFile("avatar")
	if err != nil {
		SendServerError(w, r, err)
		return
	}

	defer file.Close()
	var b bytes.Buffer
	io.Copy(&b, file)

	//todo, add a random name to the file, and addit to the database
	if err := utils.CreateStaticFile("avatars", header.Filename, b.Bytes()); err != nil {
		SendServerError(w, r, err)
		return
	}

	b.Reset()

	//todo return the relative path to the file to use in client
	SendOk(w, r)
	return
}
