// Created by nazarigonzalez on 30/9/17.

package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nazariglez/tarentola-backend/api/middlewares"
	"github.com/nazariglez/tarentola-backend/api/middlewares/policies"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/logger"
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
	method  HttpMethod
	url     string
	handler http.HandlerFunc
	police  string
}

func GetRouter() *mux.Router {
	if router != nil {
		return router
	}

	router = mux.NewRouter()
	router.StrictSlash(true)
	for _, r := range routeList {
		var handler http.HandlerFunc
		if r.police != "" {
			handler = policies.Apply(r.police, r.handler)
		}

		handler = middlewares.ParseForm(ParseURL(handler))

		if config.Data.Middlewares.Logger {
			handler = middlewares.Logger(handler)
		}

		if config.Data.Middlewares.GZIP {
			handler = middlewares.Gzip(handler)
		}

		handler = middlewares.InitRequest(handler)

		if config.Data.Middlewares.RateLimitRPS > 0 {
			handler = middlewares.RateLimit(handler)
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

		next.ServeHTTP(w, r)
	}
}

func AllowCORS(h http.Handler) http.Handler {
	if !config.Data.CORS.Enabled {
		return h
	}

	logger.Log.Debug("HTTP CORS enabled.")

	allowed := []handlers.CORSOption{
		handlers.AllowedOrigins(config.Data.CORS.Origins),
		handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization", "X-Forwarded-For", "X-Real-Ip", "X-User-Final-IP"}),
	}
	return handlers.CORS(allowed...)(h)
}
