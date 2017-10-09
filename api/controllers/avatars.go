// Created by nazarigonzalez on 9/10/17.

package controllers

import "net/http"

type avatarInfo struct {
	ID  uint `json:"id"`
	URL uint `json:"url"`
}

func GetList(w http.ResponseWriter, r *http.Request) {
	//todo get avatar list
	SendOk(w)
}
