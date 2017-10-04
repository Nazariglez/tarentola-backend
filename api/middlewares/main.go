// Created by nazarigonzalez on 3/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var list = map[string][]Middleware{
	"isLogged": {
		isLogged,
	},
	"isNotLogged": {
		isNotLogged,
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

func ParseForm(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			controllers.SendServerError(w, err)
			return
		}

		next(w, r)
	}
}
