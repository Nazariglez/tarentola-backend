// Created by nazarigonzalez on 30/9/17.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Base struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func sendOk(w http.ResponseWriter, args ...interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	base := Base{
		Success: true,
	}

	if len(args) != 0 {
		base.Data = args[0]
	}

	if err := json.NewEncoder(w).Encode(base); err != nil {
		fmt.Errorf("%+v", err)
		return
	}
}

func sendServerError(w http.ResponseWriter, args ...error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(500)

	msg := "Server error."
	if len(args) != 0 {
		msg = args[0].Error()
	}

	err := json.NewEncoder(w).Encode(Base{
		Success: false,
		Message: msg,
	})

	if err != nil {
		fmt.Errorf("%+v", err)
		return
	}
}

func sendBadRequest(w http.ResponseWriter, args ...string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	msg := "Bad request."
	if len(args) != 0 {
		msg = args[0]
	}

	err := json.NewEncoder(w).Encode(Base{
		Success: false,
		Message: msg,
	})

	if err != nil {
		fmt.Errorf("%+v", err)
		return
	}
}
