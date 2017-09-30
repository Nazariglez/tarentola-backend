// Created by nazarigonzalez on 30/9/17.

package api

import (
	"github.com/gorilla/pat"
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
	method  HttpMethod
	url     string
	handler http.HandlerFunc
}

func GetRouter() *pat.Router {
	if router != nil {
		return router
	}

	router = pat.New()
	for _, r := range routeList {
		switch r.method {
		case GET:
			router.Get(r.url, r.handler)
		case PUT:
			router.Put(r.url, r.handler)
		case POST:
			router.Post(r.url, r.handler)
		case DELETE:
			router.Delete(r.url, r.handler)
		}
	}

	return router
}
