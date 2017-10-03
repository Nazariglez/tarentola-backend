// Created by nazarigonzalez on 30/9/17.

package router

import (
	"github.com/gorilla/pat"
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

var router *pat.Router

type route struct {
	method     HttpMethod
	url        string
	handler    http.HandlerFunc
	middleware string
}

func GetRouter() *pat.Router {
	if router != nil {
		return router
	}

	router = pat.New()
	for _, r := range routeList {
		handler := middlewares.ParseForm(r.handler)
		if r.middleware != "" {
			handler = middlewares.Apply(r.middleware, handler)
		}

		switch r.method {
		case GET:
			router.Get(r.url, handler)
		case PUT:
			router.Put(r.url, handler)
		case POST:
			router.Post(r.url, handler)
		case DELETE:
			router.Delete(r.url, handler)
		}
	}

	return router
}
