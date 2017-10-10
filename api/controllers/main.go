// Created by nazarigonzalez on 2/10/17.

package controllers

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	SendNotFound(w, r)
	return
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	SendForbidden(w, r)
	return
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	SendUnauthorized(w, r)
	return
}
