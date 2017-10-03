// Created by nazarigonzalez on 3/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var list = map[string][]Middleware{
	"isAdmin": {
		isLogged,
		isAdmin,
	},
}

func Apply(name string, controller http.HandlerFunc) http.HandlerFunc {
	mws, ok := list[name]
	if !ok {
		logger.Log.Fatalf("Invalid '%s' middleware name.", name)
		return nil
	}

	for i := len(mws) - 1; i >= 0; i-- {
		controller = mws[i](controller)
	}

	return controller
}

func isLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Debug("Middleware: isLogged")
		notLogged := true
		if !notLogged {
			controllers.Forbidden(w, r)
			return
		}

		next(w, r)
	}
}

func isAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Debug("Middleware: isAdmin")
		notAdmin := true
		if !notAdmin {
			controllers.Forbidden(w, r)
			return
		}

		next(w, r)
	}
}
