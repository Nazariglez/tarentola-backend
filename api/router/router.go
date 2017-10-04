// Created by nazarigonzalez on 30/9/17.

package router

import (
	"github.com/gorilla/mux"
	"github.com/nazariglez/tarentola-backend/api/middlewares"
	"net/http"
)

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
)

var router *mux.Router

type route struct {
	method     HttpMethod
	url        string
	handler    http.HandlerFunc
	middleware string
}

func GetRouter() *mux.Router {
	if router != nil {
		return router
	}

	router = mux.NewRouter()
	for _, r := range routeList {
		handler := middlewares.ParseForm(ParseURL(r.handler))
		if r.middleware != "" {
			handler = middlewares.Apply(r.middleware, handler)
		}

		switch r.method {
		case GET:
			router.HandleFunc(r.url, handler).Methods("GET")
		case PUT:
			router.HandleFunc(r.url, handler).Methods("PUT")
		case POST:
			router.HandleFunc(r.url, handler).Methods("POST")
		case DELETE:
			router.HandleFunc(r.url, handler).Methods("DELETE")
		}
	}

	return router
}

func ParseURL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		for k, v := range vars {
			r.Form.Set(k, v)
		}

		next(w, r)
	}
}
