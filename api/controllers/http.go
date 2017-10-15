// Created by nazarigonzalez on 30/9/17.

package controllers

import (
	"encoding/json"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
)

type Base struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func getRequestIdAndUser(r *http.Request) (string, uint) {
	rid := GetRequestID(r)
	user := GetRequestUserID(r)
	return rid, user
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func SendOk(w http.ResponseWriter, r *http.Request, args ...interface{}) {
	base := Base{
		Success: true,
	}

	if len(args) != 0 {
		base.Data = args[0]
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusOK, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'OK' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'OK' (%s) - %s %s", user, ip, rid, r.Method, r.URL)
	}
}

func SendServerError(w http.ResponseWriter, r *http.Request, args ...error) {
	msg := "Server error."
	if len(args) != 0 {
		msg = args[0].Error()
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusInternalServerError, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'SERVER ERROR' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Errorf("[User:%d - %s] Response 'SERVER ERROR' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, msg)
	}
}

func SendBadRequest(w http.ResponseWriter, r *http.Request, args ...string) {
	msg := "Bad request."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusBadRequest, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'BAD REQUEST' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'BAD REQUEST' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, msg)
	}
}

func SendNotFound(w http.ResponseWriter, r *http.Request, args ...string) {
	msg := "Not found."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusNotFound, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'NOT FOUND' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'NOT FOUND' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, msg)
	}
}

func SendForbidden(w http.ResponseWriter, r *http.Request, args ...string) {
	msg := "Forbidden."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusForbidden, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'FORBIDDEN' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'FORBIDDEN' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, msg)
	}
}

func SendUnauthorized(w http.ResponseWriter, r *http.Request, args ...string) {
	msg := "Authentication fail."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	ip := GetRequestIPAddr(r)
	if err := sendJSON(w, http.StatusUnauthorized, base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'UNAUTHORIZED' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'UNAUTHORIZED' (%s) - %s %s '%s'", user, ip, rid, r.Method, r.URL, msg)
	}
}
