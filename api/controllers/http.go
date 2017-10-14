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

func SendOk(w http.ResponseWriter, r *http.Request, args ...interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	base := Base{
		Success: true,
	}

	if len(args) != 0 {
		base.Data = args[0]
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'OK' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'OK' (%s) - %s %s", user, r.RemoteAddr, rid, r.Method, r.URL)
	}
}

func SendServerError(w http.ResponseWriter, r *http.Request, args ...error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(500)

	msg := "Server error."
	if len(args) != 0 {
		msg = args[0].Error()
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'SERVER ERROR' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Errorf("[User:%d - %s] Response 'SERVER ERROR' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, msg)
	}
}

func SendBadRequest(w http.ResponseWriter, r *http.Request, args ...string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	msg := "Bad request."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'BAD REQUEST' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'BAD REQUEST' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, msg)
	}
}

func SendNotFound(w http.ResponseWriter, r *http.Request, args ...string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	msg := "Not found."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'NOT FOUND' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'NOT FOUND' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, msg)
	}
}

func SendForbidden(w http.ResponseWriter, r *http.Request, args ...string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)

	msg := "Forbidden."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'FORBIDDEN' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'FORBIDDEN' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, msg)
	}
}

func SendUnauthorized(w http.ResponseWriter, r *http.Request, args ...string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	msg := "Authentication fail."
	if len(args) != 0 {
		msg = args[0]
	}

	base := Base{
		Success: false,
		Message: msg,
	}

	rid, user := getRequestIdAndUser(r)
	if err := json.NewEncoder(w).Encode(base); err != nil {
		logger.Log.Errorf("[User:%d - %s] ERROR 'UNAUTHORIZED' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, err.Error())
		return
	} else {
		logger.Log.Tracef("[User:%d - %s] Response 'UNAUTHORIZED' (%s) - %s %s '%s'", user, r.RemoteAddr, rid, r.Method, r.URL, msg)
	}
}
