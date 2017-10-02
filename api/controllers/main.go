// Created by nazarigonzalez on 2/10/17.

package controllers

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	sendNotFound(w)
	return
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	sendForbidden(w)
	return
}
