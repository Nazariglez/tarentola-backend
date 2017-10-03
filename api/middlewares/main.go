// Created by nazarigonzalez on 3/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
	"strings"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var list = map[string][]Middleware{
	"isLogged": {
		isLogged,
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
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			controllers.SendForbidden(w, "Invalid auth token.")
			return
		}

		_, err := controllers.ValidateToken(auth[1])
		if err != nil {
			controllers.SendForbidden(w, err.Error())
			return
		}

		next(w, r)
	}
}
